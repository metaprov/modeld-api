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

	trainingv1alpha1 "github.com/metaprov/modeldapi/pkg/apis/training/v1alpha1"
	versioned "github.com/metaprov/modeldapi/pkg/client/clientset/versioned"
	internalinterfaces "github.com/metaprov/modeldapi/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/client/listers/training/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NotebookRunInformer provides access to a shared informer and lister for
// NotebookRuns.
type NotebookRunInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NotebookRunLister
}

type notebookRunInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNotebookRunInformer constructs a new informer for NotebookRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNotebookRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNotebookRunInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNotebookRunInformer constructs a new informer for NotebookRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNotebookRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1alpha1().NotebookRuns(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TrainingV1alpha1().NotebookRuns(namespace).Watch(context.TODO(), options)
			},
		},
		&trainingv1alpha1.NotebookRun{},
		resyncPeriod,
		indexers,
	)
}

func (f *notebookRunInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNotebookRunInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *notebookRunInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&trainingv1alpha1.NotebookRun{}, f.defaultInformer)
}

func (f *notebookRunInformer) Lister() v1alpha1.NotebookRunLister {
	return v1alpha1.NewNotebookRunLister(f.Informer().GetIndexer())
}