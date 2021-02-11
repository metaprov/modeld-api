/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	catalogv1alpha1 "github.com/metaprov/modeldapi/pkg/apis/catalog/v1alpha1"
	versioned "github.com/metaprov/modeldapi/pkg/client/clientset/versioned"
	internalinterfaces "github.com/metaprov/modeldapi/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/client/listers/catalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CloudInformer provides access to a shared informer and lister for
// Clouds.
type CloudInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CloudLister
}

type cloudInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCloudInformer constructs a new informer for Cloud type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCloudInformer constructs a new informer for Cloud type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CatalogV1alpha1().Clouds(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CatalogV1alpha1().Clouds(namespace).Watch(context.TODO(), options)
			},
		},
		&catalogv1alpha1.Cloud{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&catalogv1alpha1.Cloud{}, f.defaultInformer)
}

func (f *cloudInformer) Lister() v1alpha1.CloudLister {
	return v1alpha1.NewCloudLister(f.Informer().GetIndexer())
}