/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/infra/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualBuckets implements VirtualBucketInterface
type FakeVirtualBuckets struct {
	Fake *FakeInfraV1alpha1
	ns   string
}

var virtualbucketsResource = schema.GroupVersionResource{Group: "infra.modeld.io", Version: "v1alpha1", Resource: "virtualbuckets"}

var virtualbucketsKind = schema.GroupVersionKind{Group: "infra.modeld.io", Version: "v1alpha1", Kind: "VirtualBucket"}

// Get takes name of the virtualBucket, and returns the corresponding virtualBucket object, and an error if there is any.
func (c *FakeVirtualBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualbucketsResource, c.ns, name), &v1alpha1.VirtualBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualBucket), err
}

// List takes label and field selectors, and returns the list of VirtualBuckets that match those selectors.
func (c *FakeVirtualBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualbucketsResource, virtualbucketsKind, c.ns, opts), &v1alpha1.VirtualBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualBucketList{ListMeta: obj.(*v1alpha1.VirtualBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualBuckets.
func (c *FakeVirtualBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualbucketsResource, c.ns, opts))

}

// Create takes the representation of a virtualBucket and creates it.  Returns the server's representation of the virtualBucket, and an error, if there is any.
func (c *FakeVirtualBuckets) Create(ctx context.Context, virtualBucket *v1alpha1.VirtualBucket, opts v1.CreateOptions) (result *v1alpha1.VirtualBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualbucketsResource, c.ns, virtualBucket), &v1alpha1.VirtualBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualBucket), err
}

// Update takes the representation of a virtualBucket and updates it. Returns the server's representation of the virtualBucket, and an error, if there is any.
func (c *FakeVirtualBuckets) Update(ctx context.Context, virtualBucket *v1alpha1.VirtualBucket, opts v1.UpdateOptions) (result *v1alpha1.VirtualBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualbucketsResource, c.ns, virtualBucket), &v1alpha1.VirtualBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualBuckets) UpdateStatus(ctx context.Context, virtualBucket *v1alpha1.VirtualBucket, opts v1.UpdateOptions) (*v1alpha1.VirtualBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualbucketsResource, "status", c.ns, virtualBucket), &v1alpha1.VirtualBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualBucket), err
}

// Delete takes name of the virtualBucket and deletes it. Returns an error if one occurs.
func (c *FakeVirtualBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualbucketsResource, c.ns, name), &v1alpha1.VirtualBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualbucketsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualBucketList{})
	return err
}

// Patch applies the patch and returns the patched virtualBucket.
func (c *FakeVirtualBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualbucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualBucket), err
}