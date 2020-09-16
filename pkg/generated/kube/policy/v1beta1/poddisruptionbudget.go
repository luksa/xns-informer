// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/policy/v1beta1"
	listers "k8s.io/client-go/listers/policy/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type podDisruptionBudgetInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.PodDisruptionBudgetInformer = &podDisruptionBudgetInformer{}

func (f *podDisruptionBudgetInformer) resource() schema.GroupVersionResource {
	return v1beta1.SchemeGroupVersion.WithResource("poddisruptionbudgets")
}

func (f *podDisruptionBudgetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *podDisruptionBudgetInformer) Lister() listers.PodDisruptionBudgetLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1beta1.PodDisruptionBudget{})
	return listers.NewPodDisruptionBudgetLister(idx)
}
