// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type endpointsInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.EndpointsInformer = &endpointsInformer{}

func (f *endpointsInformer) resource() schema.GroupVersionResource {
	return v1.SchemeGroupVersion.WithResource("endpoints")
}

func (f *endpointsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *endpointsInformer) Lister() listers.EndpointsLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1.Endpoints{})
	return listers.NewEndpointsLister(idx)
}