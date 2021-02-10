/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/infra/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLabs implements LabInterface
type FakeLabs struct {
	Fake *FakeInfraV1alpha1
	ns   string
}

var labsResource = schema.GroupVersionResource{Group: "infra.modeld.io", Version: "v1alpha1", Resource: "labs"}

var labsKind = schema.GroupVersionKind{Group: "infra.modeld.io", Version: "v1alpha1", Kind: "Lab"}

// Get takes name of the lab, and returns the corresponding lab object, and an error if there is any.
func (c *FakeLabs) Get(name string, options v1.GetOptions) (result *v1alpha1.Lab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(labsResource, c.ns, name), &v1alpha1.Lab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lab), err
}

// List takes label and field selectors, and returns the list of Labs that match those selectors.
func (c *FakeLabs) List(opts v1.ListOptions) (result *v1alpha1.LabList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(labsResource, labsKind, c.ns, opts), &v1alpha1.LabList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LabList{ListMeta: obj.(*v1alpha1.LabList).ListMeta}
	for _, item := range obj.(*v1alpha1.LabList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested labs.
func (c *FakeLabs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(labsResource, c.ns, opts))

}

// Create takes the representation of a lab and creates it.  Returns the server's representation of the lab, and an error, if there is any.
func (c *FakeLabs) Create(lab *v1alpha1.Lab) (result *v1alpha1.Lab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(labsResource, c.ns, lab), &v1alpha1.Lab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lab), err
}

// Update takes the representation of a lab and updates it. Returns the server's representation of the lab, and an error, if there is any.
func (c *FakeLabs) Update(lab *v1alpha1.Lab) (result *v1alpha1.Lab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(labsResource, c.ns, lab), &v1alpha1.Lab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lab), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLabs) UpdateStatus(lab *v1alpha1.Lab) (*v1alpha1.Lab, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(labsResource, "status", c.ns, lab), &v1alpha1.Lab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lab), err
}

// Delete takes name of the lab and deletes it. Returns an error if one occurs.
func (c *FakeLabs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(labsResource, c.ns, name), &v1alpha1.Lab{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLabs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(labsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LabList{})
	return err
}

// Patch applies the patch and returns the patched lab.
func (c *FakeLabs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Lab, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(labsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Lab{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lab), err
}
