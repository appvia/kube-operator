module github.com/appvia/kube-operator

require (
	github.com/appvia/gke-operator v0.0.0-20191105162906-92844b1efca9 // indirect
	github.com/appvia/hub-apis v0.0.0-20191122123803-5128583f5c8b
	github.com/davecgh/go-spew v1.1.1
	github.com/gambol99/hub-utils v0.0.0-20191108161735-f3ffc063b42e
	github.com/go-openapi/spec v0.17.2
	github.com/operator-framework/operator-sdk v0.12.1-0.20191106084303-5006ea85501e
	github.com/spf13/pflag v1.0.3
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20191108065633-c18f71bf2947
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
	sigs.k8s.io/controller-runtime v0.3.0
)

// Pinned to kubernetes-1.14.1
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190409023720-1bc0c81fa51d
)

replace (
	// Indirect operator-sdk dependencies use git.apache.org, which is frequently
	// down. The github mirror should be used instead.
	// Locking to a specific version (from 'go mod graph'):
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/coreos/prometheus-operator => github.com/coreos/prometheus-operator v0.31.1
	// Pinned to v2.10.0 (kubernetes-1.14.1) so https://proxy.golang.org can
	// resolve it correctly.
	github.com/prometheus/prometheus => github.com/prometheus/prometheus v1.8.2-0.20190525122359-d20e84d0fb64
)

replace github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.11.0

go 1.13
