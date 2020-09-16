// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	"k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/rbac/v1beta1"
	listers "k8s.io/client-go/listers/rbac/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type roleInformer struct {
	factory xnsinformers.SharedInformerFactory
}

var _ informers.RoleInformer = &roleInformer{}

func (f *roleInformer) resource() schema.GroupVersionResource {
	return v1beta1.SchemeGroupVersion.WithResource("roles")
}

func (f *roleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *roleInformer) Lister() listers.RoleLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1beta1.Role{})
	return listers.NewRoleLister(idx)
}
