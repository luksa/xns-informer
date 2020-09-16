// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/certificates/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/certificates/v1beta1"
	listers "k8s.io/client-go/listers/certificates/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type certificateSigningRequestInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.CertificateSigningRequestInformer = &certificateSigningRequestInformer{}

func (f *certificateSigningRequestInformer) resource() schema.GroupVersionResource {
	return v1beta1.SchemeGroupVersion.WithResource("certificatesigningrequests")
}

func (f *certificateSigningRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.ClusterResource(f.resource()).Informer()
}

func (f *certificateSigningRequestInformer) Lister() listers.CertificateSigningRequestLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1beta1.CertificateSigningRequest{})
	return listers.NewCertificateSigningRequestLister(idx)
}