/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/catalog/v1alpha1"
	"github.com/metaprov/modeldapi/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CatalogV1alpha1Interface interface {
	RESTClient() rest.Interface
	AlgorithmsGetter
	CloudsGetter
	MLFrameworksGetter
	PretrainedModelsGetter
	UserRoleClassesGetter
	WorkloadClassesGetter
}

// CatalogV1alpha1Client is used to interact with features provided by the catalog.modeld.io group.
type CatalogV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CatalogV1alpha1Client) Algorithms(namespace string) AlgorithmInterface {
	return newAlgorithms(c, namespace)
}

func (c *CatalogV1alpha1Client) Clouds(namespace string) CloudInterface {
	return newClouds(c, namespace)
}

func (c *CatalogV1alpha1Client) MLFrameworks(namespace string) MLFrameworkInterface {
	return newMLFrameworks(c, namespace)
}

func (c *CatalogV1alpha1Client) PretrainedModels(namespace string) PretrainedModelInterface {
	return newPretrainedModels(c, namespace)
}

func (c *CatalogV1alpha1Client) UserRoleClasses(namespace string) UserRoleClassInterface {
	return newUserRoleClasses(c, namespace)
}

func (c *CatalogV1alpha1Client) WorkloadClasses(namespace string) WorkloadClassInterface {
	return newWorkloadClasses(c, namespace)
}

// NewForConfig creates a new CatalogV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CatalogV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CatalogV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CatalogV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CatalogV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CatalogV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CatalogV1alpha1Client {
	return &CatalogV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CatalogV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}