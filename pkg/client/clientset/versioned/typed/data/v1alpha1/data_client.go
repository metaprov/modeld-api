/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	"github.com/metaprov/modeldapi/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type DataV1alpha1Interface interface {
	RESTClient() rest.Interface
	DataPipelinesGetter
	DataPipelineRunsGetter
	DataProductsGetter
	DataProductVersionsGetter
	DataSourcesGetter
	DatasetsGetter
	EntitiesGetter
	FeaturesGetter
	FeaturePipelinesGetter
	FeaturePipelineRunsGetter
	FeaturesetsGetter
	LabelingPipelinesGetter
	LabelingPipelineRunsGetter
	RecipesGetter
	RecipeRunsGetter
}

// DataV1alpha1Client is used to interact with features provided by the data.modeld.io group.
type DataV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DataV1alpha1Client) DataPipelines(namespace string) DataPipelineInterface {
	return newDataPipelines(c, namespace)
}

func (c *DataV1alpha1Client) DataPipelineRuns(namespace string) DataPipelineRunInterface {
	return newDataPipelineRuns(c, namespace)
}

func (c *DataV1alpha1Client) DataProducts(namespace string) DataProductInterface {
	return newDataProducts(c, namespace)
}

func (c *DataV1alpha1Client) DataProductVersions(namespace string) DataProductVersionInterface {
	return newDataProductVersions(c, namespace)
}

func (c *DataV1alpha1Client) DataSources(namespace string) DataSourceInterface {
	return newDataSources(c, namespace)
}

func (c *DataV1alpha1Client) Datasets(namespace string) DatasetInterface {
	return newDatasets(c, namespace)
}

func (c *DataV1alpha1Client) Entities(namespace string) EntityInterface {
	return newEntities(c, namespace)
}

func (c *DataV1alpha1Client) Features(namespace string) FeatureInterface {
	return newFeatures(c, namespace)
}

func (c *DataV1alpha1Client) FeaturePipelines(namespace string) FeaturePipelineInterface {
	return newFeaturePipelines(c, namespace)
}

func (c *DataV1alpha1Client) FeaturePipelineRuns(namespace string) FeaturePipelineRunInterface {
	return newFeaturePipelineRuns(c, namespace)
}

func (c *DataV1alpha1Client) Featuresets(namespace string) FeaturesetInterface {
	return newFeaturesets(c, namespace)
}

func (c *DataV1alpha1Client) LabelingPipelines(namespace string) LabelingPipelineInterface {
	return newLabelingPipelines(c, namespace)
}

func (c *DataV1alpha1Client) LabelingPipelineRuns(namespace string) LabelingPipelineRunInterface {
	return newLabelingPipelineRuns(c, namespace)
}

func (c *DataV1alpha1Client) Recipes(namespace string) RecipeInterface {
	return newRecipes(c, namespace)
}

func (c *DataV1alpha1Client) RecipeRuns(namespace string) RecipeRunInterface {
	return newRecipeRuns(c, namespace)
}

// NewForConfig creates a new DataV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*DataV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DataV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new DataV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DataV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DataV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *DataV1alpha1Client {
	return &DataV1alpha1Client{c}
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
func (c *DataV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}