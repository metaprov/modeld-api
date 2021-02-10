/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataProducts implements DataProductInterface
type FakeDataProducts struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var dataproductsResource = schema.GroupVersionResource{Group: "data.modeld.io", Version: "v1alpha1", Resource: "dataproducts"}

var dataproductsKind = schema.GroupVersionKind{Group: "data.modeld.io", Version: "v1alpha1", Kind: "DataProduct"}

// Get takes name of the dataProduct, and returns the corresponding dataProduct object, and an error if there is any.
func (c *FakeDataProducts) Get(name string, options v1.GetOptions) (result *v1alpha1.DataProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataproductsResource, c.ns, name), &v1alpha1.DataProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataProduct), err
}

// List takes label and field selectors, and returns the list of DataProducts that match those selectors.
func (c *FakeDataProducts) List(opts v1.ListOptions) (result *v1alpha1.DataProductList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataproductsResource, dataproductsKind, c.ns, opts), &v1alpha1.DataProductList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataProductList{ListMeta: obj.(*v1alpha1.DataProductList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataProductList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataProducts.
func (c *FakeDataProducts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataproductsResource, c.ns, opts))

}

// Create takes the representation of a dataProduct and creates it.  Returns the server's representation of the dataProduct, and an error, if there is any.
func (c *FakeDataProducts) Create(dataProduct *v1alpha1.DataProduct) (result *v1alpha1.DataProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataproductsResource, c.ns, dataProduct), &v1alpha1.DataProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataProduct), err
}

// Update takes the representation of a dataProduct and updates it. Returns the server's representation of the dataProduct, and an error, if there is any.
func (c *FakeDataProducts) Update(dataProduct *v1alpha1.DataProduct) (result *v1alpha1.DataProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataproductsResource, c.ns, dataProduct), &v1alpha1.DataProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataProduct), err
}

// Delete takes name of the dataProduct and deletes it. Returns an error if one occurs.
func (c *FakeDataProducts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataproductsResource, c.ns, name), &v1alpha1.DataProduct{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataProducts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataproductsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataProductList{})
	return err
}

// Patch applies the patch and returns the patched dataProduct.
func (c *FakeDataProducts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataProduct, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataproductsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataProduct{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataProduct), err
}
