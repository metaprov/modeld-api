/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRecipeRuns implements RecipeRunInterface
type FakeRecipeRuns struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var reciperunsResource = schema.GroupVersionResource{Group: "data.modeld.io", Version: "v1alpha1", Resource: "reciperuns"}

var reciperunsKind = schema.GroupVersionKind{Group: "data.modeld.io", Version: "v1alpha1", Kind: "RecipeRun"}

// Get takes name of the recipeRun, and returns the corresponding recipeRun object, and an error if there is any.
func (c *FakeRecipeRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RecipeRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reciperunsResource, c.ns, name), &v1alpha1.RecipeRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecipeRun), err
}

// List takes label and field selectors, and returns the list of RecipeRuns that match those selectors.
func (c *FakeRecipeRuns) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RecipeRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reciperunsResource, reciperunsKind, c.ns, opts), &v1alpha1.RecipeRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RecipeRunList{ListMeta: obj.(*v1alpha1.RecipeRunList).ListMeta}
	for _, item := range obj.(*v1alpha1.RecipeRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested recipeRuns.
func (c *FakeRecipeRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reciperunsResource, c.ns, opts))

}

// Create takes the representation of a recipeRun and creates it.  Returns the server's representation of the recipeRun, and an error, if there is any.
func (c *FakeRecipeRuns) Create(ctx context.Context, recipeRun *v1alpha1.RecipeRun, opts v1.CreateOptions) (result *v1alpha1.RecipeRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reciperunsResource, c.ns, recipeRun), &v1alpha1.RecipeRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecipeRun), err
}

// Update takes the representation of a recipeRun and updates it. Returns the server's representation of the recipeRun, and an error, if there is any.
func (c *FakeRecipeRuns) Update(ctx context.Context, recipeRun *v1alpha1.RecipeRun, opts v1.UpdateOptions) (result *v1alpha1.RecipeRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reciperunsResource, c.ns, recipeRun), &v1alpha1.RecipeRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecipeRun), err
}

// Delete takes name of the recipeRun and deletes it. Returns an error if one occurs.
func (c *FakeRecipeRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reciperunsResource, c.ns, name), &v1alpha1.RecipeRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecipeRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reciperunsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RecipeRunList{})
	return err
}

// Patch applies the patch and returns the patched recipeRun.
func (c *FakeRecipeRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RecipeRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reciperunsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RecipeRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RecipeRun), err
}