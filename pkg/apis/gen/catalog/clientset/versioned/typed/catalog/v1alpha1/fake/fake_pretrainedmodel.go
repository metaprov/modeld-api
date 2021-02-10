/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/catalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePretrainedModels implements PretrainedModelInterface
type FakePretrainedModels struct {
	Fake *FakeCatalogV1alpha1
	ns   string
}

var pretrainedmodelsResource = schema.GroupVersionResource{Group: "catalog.modeld.io", Version: "v1alpha1", Resource: "pretrainedmodels"}

var pretrainedmodelsKind = schema.GroupVersionKind{Group: "catalog.modeld.io", Version: "v1alpha1", Kind: "PretrainedModel"}

// Get takes name of the pretrainedModel, and returns the corresponding pretrainedModel object, and an error if there is any.
func (c *FakePretrainedModels) Get(name string, options v1.GetOptions) (result *v1alpha1.PretrainedModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pretrainedmodelsResource, c.ns, name), &v1alpha1.PretrainedModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PretrainedModel), err
}

// List takes label and field selectors, and returns the list of PretrainedModels that match those selectors.
func (c *FakePretrainedModels) List(opts v1.ListOptions) (result *v1alpha1.PretrainedModelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pretrainedmodelsResource, pretrainedmodelsKind, c.ns, opts), &v1alpha1.PretrainedModelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PretrainedModelList{ListMeta: obj.(*v1alpha1.PretrainedModelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PretrainedModelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pretrainedModels.
func (c *FakePretrainedModels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pretrainedmodelsResource, c.ns, opts))

}

// Create takes the representation of a pretrainedModel and creates it.  Returns the server's representation of the pretrainedModel, and an error, if there is any.
func (c *FakePretrainedModels) Create(pretrainedModel *v1alpha1.PretrainedModel) (result *v1alpha1.PretrainedModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pretrainedmodelsResource, c.ns, pretrainedModel), &v1alpha1.PretrainedModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PretrainedModel), err
}

// Update takes the representation of a pretrainedModel and updates it. Returns the server's representation of the pretrainedModel, and an error, if there is any.
func (c *FakePretrainedModels) Update(pretrainedModel *v1alpha1.PretrainedModel) (result *v1alpha1.PretrainedModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pretrainedmodelsResource, c.ns, pretrainedModel), &v1alpha1.PretrainedModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PretrainedModel), err
}

// Delete takes name of the pretrainedModel and deletes it. Returns an error if one occurs.
func (c *FakePretrainedModels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pretrainedmodelsResource, c.ns, name), &v1alpha1.PretrainedModel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePretrainedModels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pretrainedmodelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PretrainedModelList{})
	return err
}

// Patch applies the patch and returns the patched pretrainedModel.
func (c *FakePretrainedModels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PretrainedModel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pretrainedmodelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PretrainedModel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PretrainedModel), err
}
