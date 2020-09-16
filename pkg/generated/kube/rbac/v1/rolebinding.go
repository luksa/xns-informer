// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/rbac/v1"
	listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/tools/cache"
)

type roleBindingInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.RoleBindingInformer = &roleBindingInformer{}

func (f *roleBindingInformer) resource() schema.GroupVersionResource {
	return v1.SchemeGroupVersion.WithResource("rolebindings")
}

func (f *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *roleBindingInformer) Lister() listers.RoleBindingLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1.RoleBinding{})
	return listers.NewRoleBindingLister(idx)
}
