// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/apps/v1beta1"
	listers "k8s.io/client-go/listers/apps/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type controllerRevisionInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.ControllerRevisionInformer = &controllerRevisionInformer{}

func (f *controllerRevisionInformer) resource() schema.GroupVersionResource {
	return v1beta1.SchemeGroupVersion.WithResource("controllerrevisions")
}

func (f *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *controllerRevisionInformer) Lister() listers.ControllerRevisionLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1beta1.ControllerRevision{})
	return listers.NewControllerRevisionLister(idx)
}
