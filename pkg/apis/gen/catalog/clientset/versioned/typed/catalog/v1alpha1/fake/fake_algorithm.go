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

// FakeAlgorithms implements AlgorithmInterface
type FakeAlgorithms struct {
	Fake *FakeCatalogV1alpha1
	ns   string
}

var algorithmsResource = schema.GroupVersionResource{Group: "catalog.modeld.io", Version: "v1alpha1", Resource: "algorithms"}

var algorithmsKind = schema.GroupVersionKind{Group: "catalog.modeld.io", Version: "v1alpha1", Kind: "Algorithm"}

// Get takes name of the algorithm, and returns the corresponding algorithm object, and an error if there is any.
func (c *FakeAlgorithms) Get(name string, options v1.GetOptions) (result *v1alpha1.Algorithm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(algorithmsResource, c.ns, name), &v1alpha1.Algorithm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Algorithm), err
}

// List takes label and field selectors, and returns the list of Algorithms that match those selectors.
func (c *FakeAlgorithms) List(opts v1.ListOptions) (result *v1alpha1.AlgorithmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(algorithmsResource, algorithmsKind, c.ns, opts), &v1alpha1.AlgorithmList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlgorithmList{ListMeta: obj.(*v1alpha1.AlgorithmList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlgorithmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested algorithms.
func (c *FakeAlgorithms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(algorithmsResource, c.ns, opts))

}

// Create takes the representation of a algorithm and creates it.  Returns the server's representation of the algorithm, and an error, if there is any.
func (c *FakeAlgorithms) Create(algorithm *v1alpha1.Algorithm) (result *v1alpha1.Algorithm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(algorithmsResource, c.ns, algorithm), &v1alpha1.Algorithm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Algorithm), err
}

// Update takes the representation of a algorithm and updates it. Returns the server's representation of the algorithm, and an error, if there is any.
func (c *FakeAlgorithms) Update(algorithm *v1alpha1.Algorithm) (result *v1alpha1.Algorithm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(algorithmsResource, c.ns, algorithm), &v1alpha1.Algorithm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Algorithm), err
}

// Delete takes name of the algorithm and deletes it. Returns an error if one occurs.
func (c *FakeAlgorithms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(algorithmsResource, c.ns, name), &v1alpha1.Algorithm{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlgorithms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(algorithmsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlgorithmList{})
	return err
}

// Patch applies the patch and returns the patched algorithm.
func (c *FakeAlgorithms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Algorithm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(algorithmsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Algorithm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Algorithm), err
}
