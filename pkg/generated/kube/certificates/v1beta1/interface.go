// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/certificates/v1beta1"
)

type Interface interface {
	CertificateSigningRequests() informers.CertificateSigningRequestInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) CertificateSigningRequests() informers.CertificateSigningRequestInformer {
	return &certificateSigningRequestInformer{factory: v.factory}
}