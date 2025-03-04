/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package backup

import (
	"context"
	"fmt"
	"time"

	"stash.appscode.dev/apimachinery/apis"
	api "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
	"stash.appscode.dev/apimachinery/client/clientset/versioned/scheme"
	"stash.appscode.dev/stash/pkg/eventer"
	"stash.appscode.dev/stash/pkg/util"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/reference"
	"k8s.io/klog/v2"
)

func (c *Controller) BackupScheduler() error {
	stopBackup := make(chan struct{})
	defer close(stopBackup)

	// split code from here for leader election
	switch c.opt.Workload.Kind {
	case apis.KindDeployment, apis.KindReplicaSet, apis.KindReplicationController:
		if err := c.electLeader(); err != nil {
			return err
		}
	default:
		if err := c.setupAndRunScheduler(stopBackup); err != nil {
			return err
		}
	}
	select {} // no error, wait forever
}

func (c *Controller) setupAndRunScheduler(stopBackup <-chan struct{}) error {
	if restic, _, err := c.setup(); err != nil {
		err = fmt.Errorf("failed to setup backup. Error: %v", err)
		if restic != nil {
			ref, rerr := reference.GetReference(scheme.Scheme, restic)
			if rerr == nil {
				eventer.CreateEventWithLog(
					c.k8sClient,
					BackupEventComponent,
					ref,
					core.EventTypeWarning,
					eventer.EventReasonFailedSetup,
					err.Error(),
				)
			} else {
				klog.Errorf("Failed to write event on %s %s. Reason: %s", restic.Kind, restic.Name, rerr)
			}
		}
		return err
	}
	c.initResticWatcher() // setup restic watcher, not required for offline backup
	go c.runScheduler(stopBackup)
	return nil
}

func (c *Controller) electLeader() error {
	rlc := resourcelock.ResourceLockConfig{
		Identity:      c.opt.PodName,
		EventRecorder: c.recorder,
	}
	resLock, err := resourcelock.New(
		resourcelock.ConfigMapsResourceLock,
		c.opt.Namespace,
		util.GetConfigmapLockName(c.opt.Workload),
		c.k8sClient.CoreV1(),
		c.k8sClient.CoordinationV1(),
		rlc,
	)
	if err != nil {
		return fmt.Errorf("error during leader election: %s", err)
	}
	go func() {
		leaderelection.RunOrDie(context.Background(), leaderelection.LeaderElectionConfig{
			Lock:          resLock,
			LeaseDuration: 15 * time.Second,
			RenewDeadline: 10 * time.Second,
			RetryPeriod:   2 * time.Second,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStartedLeading: func(ctx context.Context) {
					klog.Infoln("Got leadership, preparing backup backup")
					err := c.setupAndRunScheduler(ctx.Done())
					if err != nil {
						klog.Errorln(err)
					}
				},
				OnStoppedLeading: func() {
					klog.Infoln("Lost leadership, stopping backup backup")
				},
			},
		})
	}()
	return nil
}

func (c *Controller) runScheduler(stopCh <-chan struct{}) {
	c.cron.Start()
	c.locked <- struct{}{}

	c.stashInformerFactory.Start(stopCh)

	for _, v := range c.stashInformerFactory.WaitForCacheSync(stopCh) {
		if !v {
			runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
			return
		}
	}

	c.rQueue.Run(stopCh)

	<-stopCh
	klog.Info("Stopping Stash backup")
}

func (c *Controller) configureScheduler(r *api.Restic) error {
	// Remove previous jobs
	for _, v := range c.cron.Entries() {
		c.cron.Remove(v.ID)
	}
	_, err := c.cron.AddFunc(r.Spec.Schedule, func() {
		if err := c.runOnceForScheduler(); err != nil {
			ref, rerr := reference.GetReference(scheme.Scheme, r)
			if rerr == nil {
				c.recorder.Event(ref, core.EventTypeWarning, eventer.EventReasonFailedCronJob, err.Error())
			} else {
				klog.Errorf("Failed to write event on %s %s. Reason: %s", r.Kind, r.Name, rerr)
			}
			klog.Errorln(err)
		}
	})
	if err != nil {
		return err
	}
	_, err = c.cron.AddFunc("0 0 */3 * *", func() {
		err2 := c.checkOnceForScheduler()
		if err2 != nil {
			klog.Errorln(err2)
		}
	})
	return err
}

func (c *Controller) runOnceForScheduler() error {
	select {
	case <-c.locked:
		klog.Infof("Acquired lock for Restic %s/%s", c.opt.Namespace, c.opt.ResticName)
		defer func() {
			c.locked <- struct{}{}
		}()
	default:
		klog.Warningf("Skipping backup schedule for Restic %s/%s", c.opt.Namespace, c.opt.ResticName)
		return nil
	}

	// check restic again, previously done in setup()
	restic, err := c.rLister.Restics(c.opt.Namespace).Get(c.opt.ResticName)
	if kerr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}
	if restic.Spec.Backend.StorageSecretName == "" {
		return errors.New("missing repository secret name")
	}
	secret, err := c.k8sClient.CoreV1().Secrets(restic.Namespace).Get(context.TODO(), restic.Spec.Backend.StorageSecretName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// setup restic again, previously done in setup()
	prefix := ""
	if prefix, err = c.resticCLI.SetupEnv(restic.Spec.Backend, secret, c.opt.SmartPrefix); err != nil {
		return err
	}
	if err = c.resticCLI.InitRepositoryIfAbsent(); err != nil {
		return err
	}
	repository, err := c.createRepositoryCrdIfNotExist(restic, prefix)
	if err != nil {
		return err
	}

	// run final restic backup command
	return c.runResticBackup(restic, repository)
}

func (c *Controller) checkOnceForScheduler() (err error) {

	var repository *api.Repository
	repository, err = c.stashClient.StashV1alpha1().Repositories(c.opt.Namespace).Get(context.TODO(), c.opt.Workload.GetRepositoryCRDName(c.opt.PodName, c.opt.NodeName), metav1.GetOptions{})
	if kerr.IsNotFound(err) {
		err = nil
		return
	} else if err != nil {
		return
	}

	select {
	case <-c.locked:
		klog.Infof("Acquired lock for Repository %s/%s", repository.Namespace, repository.Name)
		defer func() {
			c.locked <- struct{}{}
		}()
	default:
		klog.Warningf("Skipping checkup schedule for Repository %s/%s", repository.Namespace, repository.Name)
		return
	}

	err = c.resticCLI.Check()
	if err != nil {
		ref, rerr := reference.GetReference(scheme.Scheme, repository)
		if rerr == nil {
			c.recorder.Eventf(
				ref,
				core.EventTypeWarning,
				eventer.EventReasonFailedToCheck,
				"Repository check failed for workload %s %s/%s. Reason: %v",
				c.opt.Workload.Kind, c.opt.Namespace, c.opt.Workload.Name, err)
		} else {
			klog.Errorf("Failed to write event on %s %s. Reason: %s", repository.Kind, repository.Name, rerr)
		}
	}
	return
}
