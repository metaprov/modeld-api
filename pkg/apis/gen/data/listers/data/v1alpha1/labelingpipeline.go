/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LabelingPipelineLister helps list LabelingPipelines.
type LabelingPipelineLister interface {
	// List lists all LabelingPipelines in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LabelingPipeline, err error)
	// LabelingPipelines returns an object that can list and get LabelingPipelines.
	LabelingPipelines(namespace string) LabelingPipelineNamespaceLister
	LabelingPipelineListerExpansion
}

// labelingPipelineLister implements the LabelingPipelineLister interface.
type labelingPipelineLister struct {
	indexer cache.Indexer
}

// NewLabelingPipelineLister returns a new LabelingPipelineLister.
func NewLabelingPipelineLister(indexer cache.Indexer) LabelingPipelineLister {
	return &labelingPipelineLister{indexer: indexer}
}

// List lists all LabelingPipelines in the indexer.
func (s *labelingPipelineLister) List(selector labels.Selector) (ret []*v1alpha1.LabelingPipeline, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LabelingPipeline))
	})
	return ret, err
}

// LabelingPipelines returns an object that can list and get LabelingPipelines.
func (s *labelingPipelineLister) LabelingPipelines(namespace string) LabelingPipelineNamespaceLister {
	return labelingPipelineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LabelingPipelineNamespaceLister helps list and get LabelingPipelines.
type LabelingPipelineNamespaceLister interface {
	// List lists all LabelingPipelines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LabelingPipeline, err error)
	// Get retrieves the LabelingPipeline from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LabelingPipeline, error)
	LabelingPipelineNamespaceListerExpansion
}

// labelingPipelineNamespaceLister implements the LabelingPipelineNamespaceLister
// interface.
type labelingPipelineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LabelingPipelines in the indexer for a given namespace.
func (s labelingPipelineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LabelingPipeline, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LabelingPipeline))
	})
	return ret, err
}

// Get retrieves the LabelingPipeline from the indexer for a given namespace and name.
func (s labelingPipelineNamespaceLister) Get(name string) (*v1alpha1.LabelingPipeline, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("labelingpipeline"), name)
	}
	return obj.(*v1alpha1.LabelingPipeline), nil
}
