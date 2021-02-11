/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/team/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePostMortems implements PostMortemInterface
type FakePostMortems struct {
	Fake *FakeTeamV1alpha1
	ns   string
}

var postmortemsResource = schema.GroupVersionResource{Group: "team.modeld.io", Version: "v1alpha1", Resource: "postmortems"}

var postmortemsKind = schema.GroupVersionKind{Group: "team.modeld.io", Version: "v1alpha1", Kind: "PostMortem"}

// Get takes name of the postMortem, and returns the corresponding postMortem object, and an error if there is any.
func (c *FakePostMortems) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PostMortem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postmortemsResource, c.ns, name), &v1alpha1.PostMortem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostMortem), err
}

// List takes label and field selectors, and returns the list of PostMortems that match those selectors.
func (c *FakePostMortems) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PostMortemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postmortemsResource, postmortemsKind, c.ns, opts), &v1alpha1.PostMortemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PostMortemList{ListMeta: obj.(*v1alpha1.PostMortemList).ListMeta}
	for _, item := range obj.(*v1alpha1.PostMortemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postMortems.
func (c *FakePostMortems) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postmortemsResource, c.ns, opts))

}

// Create takes the representation of a postMortem and creates it.  Returns the server's representation of the postMortem, and an error, if there is any.
func (c *FakePostMortems) Create(ctx context.Context, postMortem *v1alpha1.PostMortem, opts v1.CreateOptions) (result *v1alpha1.PostMortem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postmortemsResource, c.ns, postMortem), &v1alpha1.PostMortem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostMortem), err
}

// Update takes the representation of a postMortem and updates it. Returns the server's representation of the postMortem, and an error, if there is any.
func (c *FakePostMortems) Update(ctx context.Context, postMortem *v1alpha1.PostMortem, opts v1.UpdateOptions) (result *v1alpha1.PostMortem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postmortemsResource, c.ns, postMortem), &v1alpha1.PostMortem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostMortem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePostMortems) UpdateStatus(ctx context.Context, postMortem *v1alpha1.PostMortem, opts v1.UpdateOptions) (*v1alpha1.PostMortem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(postmortemsResource, "status", c.ns, postMortem), &v1alpha1.PostMortem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostMortem), err
}

// Delete takes name of the postMortem and deletes it. Returns an error if one occurs.
func (c *FakePostMortems) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(postmortemsResource, c.ns, name), &v1alpha1.PostMortem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostMortems) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postmortemsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PostMortemList{})
	return err
}

// Patch applies the patch and returns the patched postMortem.
func (c *FakePostMortems) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PostMortem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postmortemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PostMortem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostMortem), err
}