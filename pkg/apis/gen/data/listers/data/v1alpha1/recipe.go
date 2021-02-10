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

// RecipeLister helps list Recipes.
type RecipeLister interface {
	// List lists all Recipes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Recipe, err error)
	// Recipes returns an object that can list and get Recipes.
	Recipes(namespace string) RecipeNamespaceLister
	RecipeListerExpansion
}

// recipeLister implements the RecipeLister interface.
type recipeLister struct {
	indexer cache.Indexer
}

// NewRecipeLister returns a new RecipeLister.
func NewRecipeLister(indexer cache.Indexer) RecipeLister {
	return &recipeLister{indexer: indexer}
}

// List lists all Recipes in the indexer.
func (s *recipeLister) List(selector labels.Selector) (ret []*v1alpha1.Recipe, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Recipe))
	})
	return ret, err
}

// Recipes returns an object that can list and get Recipes.
func (s *recipeLister) Recipes(namespace string) RecipeNamespaceLister {
	return recipeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RecipeNamespaceLister helps list and get Recipes.
type RecipeNamespaceLister interface {
	// List lists all Recipes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Recipe, err error)
	// Get retrieves the Recipe from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Recipe, error)
	RecipeNamespaceListerExpansion
}

// recipeNamespaceLister implements the RecipeNamespaceLister
// interface.
type recipeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Recipes in the indexer for a given namespace.
func (s recipeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Recipe, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Recipe))
	})
	return ret, err
}

// Get retrieves the Recipe from the indexer for a given namespace and name.
func (s recipeNamespaceLister) Get(name string) (*v1alpha1.Recipe, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("recipe"), name)
	}
	return obj.(*v1alpha1.Recipe), nil
}
