/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/metaprov/modeldapi/pkg/apis/gen/infra/clientset/versioned"
	internalinterfaces "github.com/metaprov/modeldapi/pkg/apis/gen/infra/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/gen/infra/listers/infra/v1alpha1"
	infrav1alpha1 "github.com/metaprov/modeldapi/pkg/apis/infra/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TenantInformer provides access to a shared informer and lister for
// Tenants.
type TenantInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TenantLister
}

type tenantInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTenantInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTenantInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTenantInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InfraV1alpha1().Tenants(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InfraV1alpha1().Tenants(namespace).Watch(options)
			},
		},
		&infrav1alpha1.Tenant{},
		resyncPeriod,
		indexers,
	)
}

func (f *tenantInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTenantInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tenantInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&infrav1alpha1.Tenant{}, f.defaultInformer)
}

func (f *tenantInformer) Lister() v1alpha1.TenantLister {
	return v1alpha1.NewTenantLister(f.Informer().GetIndexer())
}
