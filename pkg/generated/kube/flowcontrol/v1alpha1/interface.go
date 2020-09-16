// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	informers "k8s.io/client-go/informers/flowcontrol/v1alpha1"
)

type Interface interface {
	FlowSchemas() informers.FlowSchemaInformer
	PriorityLevelConfigurations() informers.PriorityLevelConfigurationInformer
}

type version struct {
	factory xnsinformers.SharedInformerFactory
}

func New(factory xnsinformers.SharedInformerFactory) Interface {
	return &version{factory: factory}
}
func (v *version) FlowSchemas() informers.FlowSchemaInformer {
	return &flowSchemaInformer{factory: v.factory}
}
func (v *version) PriorityLevelConfigurations() informers.PriorityLevelConfigurationInformer {
	return &priorityLevelConfigurationInformer{factory: v.factory}
}
