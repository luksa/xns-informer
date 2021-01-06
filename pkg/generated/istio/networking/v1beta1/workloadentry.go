// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	informers "istio.io/client-go/pkg/informers/externalversions/networking/v1beta1"
	listers "istio.io/client-go/pkg/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type workloadEntryInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.WorkloadEntryInformer = &workloadEntryInformer{}

func NewWorkloadEntryInformer(f xnsinformers.SharedInformerFactory) informers.WorkloadEntryInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("workloadentries")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1beta1.WorkloadEntry{},
		&v1beta1.WorkloadEntryList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &workloadEntryInformer{informer: informer.Informer()}
}

func (i *workloadEntryInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *workloadEntryInformer) Lister() listers.WorkloadEntryLister {
	return listers.NewWorkloadEntryLister(i.informer.GetIndexer())
}
