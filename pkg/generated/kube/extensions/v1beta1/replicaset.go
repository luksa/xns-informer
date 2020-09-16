// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/extensions/v1beta1"
	listers "k8s.io/client-go/listers/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type replicaSetInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.ReplicaSetInformer = &replicaSetInformer{}

func (f *replicaSetInformer) resource() schema.GroupVersionResource {
	return v1beta1.SchemeGroupVersion.WithResource("replicasets")
}

func (f *replicaSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *replicaSetInformer) Lister() listers.ReplicaSetLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1beta1.ReplicaSet{})
	return listers.NewReplicaSetLister(idx)
}
