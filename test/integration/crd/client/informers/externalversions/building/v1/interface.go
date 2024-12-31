// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/slok/kubewebhook/v2/test/integration/crd/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Houses returns a HouseInformer.
	Houses() HouseInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Houses returns a HouseInformer.
func (v *version) Houses() HouseInformer {
	return &houseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}