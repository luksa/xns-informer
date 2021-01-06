// Code generated by xns-informer-gen. DO NOT EDIT.

package istio

import (
	"fmt"

	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}
type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=networking, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("destinationrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().DestinationRules().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("envoyfilters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().EnvoyFilters().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().Gateways().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("serviceentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().ServiceEntries().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("sidecars"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().Sidecars().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().VirtualServices().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("workloadentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().WorkloadEntries().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("workloadgroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1alpha3().WorkloadGroups().Informer()}, nil

		// Group=networking, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("destinationrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1beta1().DestinationRules().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1beta1().Gateways().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("serviceentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1beta1().ServiceEntries().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("sidecars"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1beta1().Sidecars().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("virtualservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1beta1().VirtualServices().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("workloadentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().V1beta1().WorkloadEntries().Informer()}, nil

		// Group=security, Version=v1beta1
	case securityv1beta1.SchemeGroupVersion.WithResource("authorizationpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Security().V1beta1().AuthorizationPolicies().Informer()}, nil
	case securityv1beta1.SchemeGroupVersion.WithResource("peerauthentications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Security().V1beta1().PeerAuthentications().Informer()}, nil
	case securityv1beta1.SchemeGroupVersion.WithResource("requestauthentications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Security().V1beta1().RequestAuthentications().Informer()}, nil

	}
	return nil, fmt.Errorf("no informer found for %v", resource)
}
