/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/team/v1alpha1"
	"github.com/metaprov/modeldapi/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type TeamV1alpha1Interface interface {
	RESTClient() rest.Interface
	ConversationsGetter
	PostMortemsGetter
	RunBooksGetter
}

// TeamV1alpha1Client is used to interact with features provided by the team.modeld.io group.
type TeamV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TeamV1alpha1Client) Conversations(namespace string) ConversationInterface {
	return newConversations(c, namespace)
}

func (c *TeamV1alpha1Client) PostMortems(namespace string) PostMortemInterface {
	return newPostMortems(c, namespace)
}

func (c *TeamV1alpha1Client) RunBooks(namespace string) RunBookInterface {
	return newRunBooks(c, namespace)
}

// NewForConfig creates a new TeamV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*TeamV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TeamV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TeamV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TeamV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TeamV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TeamV1alpha1Client {
	return &TeamV1alpha1Client{c}
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
func (c *TeamV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
