module stash.appscode.dev/stash

go 1.15

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/kubernetes-csi/external-snapshotter/client/v4 v4.1.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	go.bytebuilders.dev/audit v0.0.10
	go.bytebuilders.dev/license-verifier v0.9.3
	go.bytebuilders.dev/license-verifier/kubernetes v0.9.3
	gomodules.xyz/blobfs v0.1.7
	gomodules.xyz/cert v1.2.0
	gomodules.xyz/encoding v0.0.2
	gomodules.xyz/envsubst v0.1.0
	gomodules.xyz/flags v0.1.2
	gomodules.xyz/go-sh v0.1.0
	gomodules.xyz/logs v0.0.4
	gomodules.xyz/pointer v0.1.0
	gomodules.xyz/runtime v0.2.0
	gomodules.xyz/stow v0.2.4
	gomodules.xyz/x v0.0.8
	gopkg.in/ini.v1 v1.51.0
	k8s.io/api v0.21.1
	k8s.io/apiextensions-apiserver v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/apiserver v0.21.1
	k8s.io/client-go v0.21.1
	k8s.io/klog/v2 v2.8.0
	k8s.io/kube-aggregator v0.21.1
	k8s.io/kubernetes v1.21.1
	kmodules.xyz/client-go v0.0.0-20210928133955-8df5bb467db6
	kmodules.xyz/constants v0.0.0-20210218100002-2c304bfda278
	kmodules.xyz/custom-resources v0.0.0-20211007080833-72bd9e8cae6e
	kmodules.xyz/objectstore-api v0.0.0-20210928135706-fdf68f88ea6e
	kmodules.xyz/offshoot-api v0.0.0-20210829122105-6f4d481b0c61
	kmodules.xyz/openshift v0.0.0-20210618001443-f2507caa512f
	kmodules.xyz/prober v0.0.0-20210618020259-5836fb959027
	kmodules.xyz/webhook-runtime v0.0.0-20210928141616-7f73c2ab318a
	stash.appscode.dev/apimachinery v0.16.0
)

replace bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d

replace cloud.google.com/go => cloud.google.com/go v0.54.0

replace cloud.google.com/go/bigquery => cloud.google.com/go/bigquery v1.4.0

replace cloud.google.com/go/datastore => cloud.google.com/go/datastore v1.1.0

replace cloud.google.com/go/firestore => cloud.google.com/go/firestore v1.1.0

replace cloud.google.com/go/pubsub => cloud.google.com/go/pubsub v1.2.0

replace cloud.google.com/go/storage => cloud.google.com/go/storage v1.6.0

replace github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v43.0.0+incompatible

replace github.com/Azure/go-ansiterm => github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.0+incompatible

replace github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.11.12

replace github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.9.5

replace github.com/Azure/go-autorest/autorest/date => github.com/Azure/go-autorest/autorest/date v0.3.0

replace github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.4.1

replace github.com/Azure/go-autorest/autorest/to => github.com/Azure/go-autorest/autorest/to v0.2.0

replace github.com/Azure/go-autorest/autorest/validation => github.com/Azure/go-autorest/autorest/validation v0.1.0

replace github.com/Azure/go-autorest/logger => github.com/Azure/go-autorest/logger v0.2.0

replace github.com/Azure/go-autorest/tracing => github.com/Azure/go-autorest/tracing v0.6.0

replace github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d

replace github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible

replace github.com/go-openapi/analysis => github.com/go-openapi/analysis v0.19.5

replace github.com/go-openapi/errors => github.com/go-openapi/errors v0.19.2

replace github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.3

replace github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.3

replace github.com/go-openapi/loads => github.com/go-openapi/loads v0.19.4

replace github.com/go-openapi/runtime => github.com/go-openapi/runtime v0.19.4

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.5

replace github.com/go-openapi/strfmt => github.com/go-openapi/strfmt v0.19.5

replace github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.5

replace github.com/go-openapi/validate => github.com/gomodules/validate v0.19.8-1.16

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/golang/protobuf => github.com/golang/protobuf v1.4.3

replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

replace github.com/prometheus-operator/prometheus-operator => github.com/prometheus-operator/prometheus-operator v0.47.0

replace github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring => github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.47.0

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.10.0

replace go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200910180754-dd1b699fc489

replace google.golang.org/api => google.golang.org/api v0.20.0

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a

replace google.golang.org/grpc => google.golang.org/grpc v1.27.1

replace helm.sh/helm/v3 => github.com/kubepack/helm/v3 v3.6.1-0.20210518225915-c3e0ce48dd1b

replace k8s.io/api => k8s.io/api v0.21.1

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.21.2-rc.0.0.20210617231004-332981b97d2d

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.21.2-0.20210716212718-83e5493ac170

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.21.1

replace k8s.io/client-go => k8s.io/client-go v0.21.1

replace k8s.io/component-base => k8s.io/component-base v0.21.1

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.22.0-alpha.0.0.20210617232219-a432af45d932

replace k8s.io/utils => k8s.io/utils v0.0.0-20201110183641-67b214c5f920

replace sigs.k8s.io/application => github.com/kmodules/application v0.8.4-0.20210427030912-90eeee3bc4ad

replace github.com/satori/go.uuid => github.com/gofrs/uuid v4.0.0+incompatible

replace github.com/dgrijalva/jwt-go => github.com/gomodules/jwt v3.2.2+incompatible

replace github.com/golang-jwt/jwt => github.com/golang-jwt/jwt v3.2.2+incompatible

replace github.com/form3tech-oss/jwt-go => github.com/form3tech-oss/jwt-go v3.2.5+incompatible
