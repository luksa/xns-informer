// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/scheduling/v1beta1"
)

type Interface interface {
	PriorityClasses() informers.PriorityClassInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) PriorityClasses() informers.PriorityClassInformer {
	return &priorityClassInformer{factory: v.factory}
}
