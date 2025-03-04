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

package util

import (
	api_v1alpha1 "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
	api "stash.appscode.dev/apimachinery/apis/stash/v1beta1"
	"stash.appscode.dev/apimachinery/pkg/restic"

	go_str "gomodules.xyz/x/strings"
)

// options that don't come from repository, backup-config, backup-session, restore-session
type ExtraOptions struct {
	Host        string
	SecretDir   string
	CacertFile  string
	ScratchDir  string
	EnableCache bool
}

func BackupOptionsForBackupTarget(backupTarget *api.BackupTarget, retentionPolicy api_v1alpha1.RetentionPolicy, extraOpt ExtraOptions) restic.BackupOptions {
	backupOpt := restic.BackupOptions{
		Host:            extraOpt.Host,
		RetentionPolicy: retentionPolicy,
	}
	if backupTarget != nil {
		backupOpt.BackupPaths = backupTarget.Paths
		backupOpt.Exclude = backupTarget.Exclude
		backupOpt.Args = backupTarget.Args
	}
	return backupOpt
}

// return the matching rule
// if targetHosts is empty for a rule, it will match any hostname
func RestoreOptionsForHost(hostname string, rules []api.Rule) restic.RestoreOptions {
	var matchedRule restic.RestoreOptions
	// first check for rules non-empty targetHost
	for _, rule := range rules {
		// if sourceHost is specified in the rule then use it. otherwise use workload itself as host
		sourceHost := hostname
		if rule.SourceHost != "" {
			sourceHost = rule.SourceHost
		}

		if len(rule.TargetHosts) == 0 || go_str.Contains(rule.TargetHosts, hostname) {
			matchedRule = restic.RestoreOptions{
				Host:         hostname,
				SourceHost:   sourceHost,
				RestorePaths: rule.Paths,
				Snapshots:    rule.Snapshots,
				Include:      rule.Include,
				Exclude:      rule.Exclude,
			}
			// if rule has empty targetHost then check further rules to see if any other rule with non-empty targetHost matches
			if len(rule.TargetHosts) == 0 {
				continue
			} else {
				return matchedRule
			}
		}
	}
	// matchedRule is either empty or contains restore option for the rules with empty targetHost field.
	return matchedRule
}

func SetupOptionsForRepository(repository api_v1alpha1.Repository, extraOpt ExtraOptions) (restic.SetupOptions, error) {
	provider, err := repository.Spec.Backend.Provider()
	if err != nil {
		return restic.SetupOptions{}, err
	}
	bucket, err := repository.Spec.Backend.Container()
	if err != nil {
		return restic.SetupOptions{}, err
	}
	prefix, err := repository.Spec.Backend.Prefix()
	if err != nil {
		return restic.SetupOptions{}, err
	}
	endpoint, _ := repository.Spec.Backend.Endpoint()
	region, _ := repository.Spec.Backend.Region()

	return restic.SetupOptions{
		Provider:       provider,
		Bucket:         bucket,
		Path:           prefix,
		Endpoint:       endpoint,
		Region:         region,
		CacertFile:     extraOpt.CacertFile,
		SecretDir:      extraOpt.SecretDir,
		ScratchDir:     extraOpt.ScratchDir,
		EnableCache:    extraOpt.EnableCache,
		MaxConnections: repository.Spec.Backend.MaxConnections(),
	}, nil
}
