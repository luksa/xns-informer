// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/flowcontrol/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/flowcontrol/v1alpha1"
	listers "k8s.io/client-go/listers/flowcontrol/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

type flowSchemaInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.FlowSchemaInformer = &flowSchemaInformer{}

func (f *flowSchemaInformer) resource() schema.GroupVersionResource {
	return v1alpha1.SchemeGroupVersion.WithResource("flowschemas")
}

func (f *flowSchemaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.ClusterResource(f.resource()).Informer()
}

func (f *flowSchemaInformer) Lister() listers.FlowSchemaLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1alpha1.FlowSchema{})
	return listers.NewFlowSchemaLister(idx)
}
