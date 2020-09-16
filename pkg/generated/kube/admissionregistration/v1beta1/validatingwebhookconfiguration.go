// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/admissionregistration/v1beta1"
	listers "k8s.io/client-go/listers/admissionregistration/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type validatingWebhookConfigurationInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.ValidatingWebhookConfigurationInformer = &validatingWebhookConfigurationInformer{}

func (f *validatingWebhookConfigurationInformer) resource() schema.GroupVersionResource {
	return v1beta1.SchemeGroupVersion.WithResource("validatingwebhookconfigurations")
}

func (f *validatingWebhookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.ClusterResource(f.resource()).Informer()
}

func (f *validatingWebhookConfigurationInformer) Lister() listers.ValidatingWebhookConfigurationLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1beta1.ValidatingWebhookConfiguration{})
	return listers.NewValidatingWebhookConfigurationLister(idx)
}