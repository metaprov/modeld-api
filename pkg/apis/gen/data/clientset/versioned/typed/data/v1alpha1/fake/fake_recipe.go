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

// FakeRecipes implements RecipeInterface
type FakeRecipes struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var recipesResource = schema.GroupVersionResource{Group: "data.modeld.io", Version: "v1alpha1", Resource: "recipes"}

var recipesKind = schema.GroupVersionKind{Group: "data.modeld.io", Version: "v1alpha1", Kind: "Recipe"}

// Get takes name of the recipe, and returns the corresponding recipe object, and an error if there is any.
func (c *FakeRecipes) Get(name string, options v1.GetOptions) (result *v1alpha1.Recipe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(recipesResource, c.ns, name), &v1alpha1.Recipe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Recipe), err
}

// List takes label and field selectors, and returns the list of Recipes that match those selectors.
func (c *FakeRecipes) List(opts v1.ListOptions) (result *v1alpha1.RecipeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(recipesResource, recipesKind, c.ns, opts), &v1alpha1.RecipeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RecipeList{ListMeta: obj.(*v1alpha1.RecipeList).ListMeta}
	for _, item := range obj.(*v1alpha1.RecipeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested recipes.
func (c *FakeRecipes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(recipesResource, c.ns, opts))

}

// Create takes the representation of a recipe and creates it.  Returns the server's representation of the recipe, and an error, if there is any.
func (c *FakeRecipes) Create(recipe *v1alpha1.Recipe) (result *v1alpha1.Recipe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(recipesResource, c.ns, recipe), &v1alpha1.Recipe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Recipe), err
}

// Update takes the representation of a recipe and updates it. Returns the server's representation of the recipe, and an error, if there is any.
func (c *FakeRecipes) Update(recipe *v1alpha1.Recipe) (result *v1alpha1.Recipe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(recipesResource, c.ns, recipe), &v1alpha1.Recipe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Recipe), err
}

// Delete takes name of the recipe and deletes it. Returns an error if one occurs.
func (c *FakeRecipes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(recipesResource, c.ns, name), &v1alpha1.Recipe{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecipes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(recipesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RecipeList{})
	return err
}

// Patch applies the patch and returns the patched recipe.
func (c *FakeRecipes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Recipe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(recipesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Recipe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Recipe), err
}
