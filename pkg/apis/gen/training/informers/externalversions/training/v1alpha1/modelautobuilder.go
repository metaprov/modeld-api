/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/metaprov/modeldapi/pkg/apis/gen/training/clientset/versioned"
	internalinterfaces "github.com/metaprov/modeldapi/pkg/apis/gen/training/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/gen/training/listers/training/v1alpha1"
	trainingv1alpha1 "github.com/metaprov/modeldapi/pkg/apis/training/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ModelAutobuilderInformer provides access to a shared informer and lister for
// ModelAutobuilders.
type ModelAutobuilderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ModelAutobuilderLister
}

type modelAutobuilderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewModelAutobuilderInformer constructs a new informer for ModelAutobuilder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewModelAutobuilderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredModelAutobuilderInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredModelAutobuilderInformer constructs a new informer for ModelAutobuilder type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredModelAutobuilderInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1alpha1().ModelAutobuilders(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1alpha1().ModelAutobuilders(namespace).Watch(options)
			},
		},
		&trainingv1alpha1.ModelAutobuilder{},
		resyncPeriod,
		indexers,
	)
}

func (f *modelAutobuilderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredModelAutobuilderInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *modelAutobuilderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&trainingv1alpha1.ModelAutobuilder{}, f.defaultInformer)
}

func (f *modelAutobuilderInformer) Lister() v1alpha1.ModelAutobuilderLister {
	return v1alpha1.NewModelAutobuilderLister(f.Informer().GetIndexer())
}
