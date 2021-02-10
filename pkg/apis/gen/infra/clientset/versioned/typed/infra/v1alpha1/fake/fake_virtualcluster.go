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

// FakeVirtualClusters implements VirtualClusterInterface
type FakeVirtualClusters struct {
	Fake *FakeInfraV1alpha1
	ns   string
}

var virtualclustersResource = schema.GroupVersionResource{Group: "infra.modeld.io", Version: "v1alpha1", Resource: "virtualclusters"}

var virtualclustersKind = schema.GroupVersionKind{Group: "infra.modeld.io", Version: "v1alpha1", Kind: "VirtualCluster"}

// Get takes name of the virtualCluster, and returns the corresponding virtualCluster object, and an error if there is any.
func (c *FakeVirtualClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualclustersResource, c.ns, name), &v1alpha1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualCluster), err
}

// List takes label and field selectors, and returns the list of VirtualClusters that match those selectors.
func (c *FakeVirtualClusters) List(opts v1.ListOptions) (result *v1alpha1.VirtualClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualclustersResource, virtualclustersKind, c.ns, opts), &v1alpha1.VirtualClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualClusterList{ListMeta: obj.(*v1alpha1.VirtualClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualClusters.
func (c *FakeVirtualClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualclustersResource, c.ns, opts))

}

// Create takes the representation of a virtualCluster and creates it.  Returns the server's representation of the virtualCluster, and an error, if there is any.
func (c *FakeVirtualClusters) Create(virtualCluster *v1alpha1.VirtualCluster) (result *v1alpha1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualclustersResource, c.ns, virtualCluster), &v1alpha1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualCluster), err
}

// Update takes the representation of a virtualCluster and updates it. Returns the server's representation of the virtualCluster, and an error, if there is any.
func (c *FakeVirtualClusters) Update(virtualCluster *v1alpha1.VirtualCluster) (result *v1alpha1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualclustersResource, c.ns, virtualCluster), &v1alpha1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualClusters) UpdateStatus(virtualCluster *v1alpha1.VirtualCluster) (*v1alpha1.VirtualCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualclustersResource, "status", c.ns, virtualCluster), &v1alpha1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualCluster), err
}

// Delete takes name of the virtualCluster and deletes it. Returns an error if one occurs.
func (c *FakeVirtualClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualclustersResource, c.ns, name), &v1alpha1.VirtualCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualClusterList{})
	return err
}

// Patch applies the patch and returns the patched virtualCluster.
func (c *FakeVirtualClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualCluster), err
}
