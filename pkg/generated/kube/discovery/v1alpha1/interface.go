// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/discovery/v1alpha1"
)

type Interface interface {
	EndpointSlices() informers.EndpointSliceInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) EndpointSlices() informers.EndpointSliceInformer {
	return &endpointSliceInformer{factory: v.factory}
}
