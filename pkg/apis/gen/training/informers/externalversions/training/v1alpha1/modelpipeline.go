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

// ModelPipelineInformer provides access to a shared informer and lister for
// ModelPipelines.
type ModelPipelineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ModelPipelineLister
}

type modelPipelineInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewModelPipelineInformer constructs a new informer for ModelPipeline type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewModelPipelineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredModelPipelineInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredModelPipelineInformer constructs a new informer for ModelPipeline type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredModelPipelineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1alpha1().ModelPipelines(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1alpha1().ModelPipelines(namespace).Watch(options)
			},
		},
		&trainingv1alpha1.ModelPipeline{},
		resyncPeriod,
		indexers,
	)
}

func (f *modelPipelineInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredModelPipelineInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *modelPipelineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&trainingv1alpha1.ModelPipeline{}, f.defaultInformer)
}

func (f *modelPipelineInformer) Lister() v1alpha1.ModelPipelineLister {
	return v1alpha1.NewModelPipelineLister(f.Informer().GetIndexer())
}
