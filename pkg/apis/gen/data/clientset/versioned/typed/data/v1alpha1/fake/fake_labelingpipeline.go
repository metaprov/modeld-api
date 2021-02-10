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

// FakeLabelingPipelines implements LabelingPipelineInterface
type FakeLabelingPipelines struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var labelingpipelinesResource = schema.GroupVersionResource{Group: "data.modeld.io", Version: "v1alpha1", Resource: "labelingpipelines"}

var labelingpipelinesKind = schema.GroupVersionKind{Group: "data.modeld.io", Version: "v1alpha1", Kind: "LabelingPipeline"}

// Get takes name of the labelingPipeline, and returns the corresponding labelingPipeline object, and an error if there is any.
func (c *FakeLabelingPipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.LabelingPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(labelingpipelinesResource, c.ns, name), &v1alpha1.LabelingPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LabelingPipeline), err
}

// List takes label and field selectors, and returns the list of LabelingPipelines that match those selectors.
func (c *FakeLabelingPipelines) List(opts v1.ListOptions) (result *v1alpha1.LabelingPipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(labelingpipelinesResource, labelingpipelinesKind, c.ns, opts), &v1alpha1.LabelingPipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LabelingPipelineList{ListMeta: obj.(*v1alpha1.LabelingPipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.LabelingPipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested labelingPipelines.
func (c *FakeLabelingPipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(labelingpipelinesResource, c.ns, opts))

}

// Create takes the representation of a labelingPipeline and creates it.  Returns the server's representation of the labelingPipeline, and an error, if there is any.
func (c *FakeLabelingPipelines) Create(labelingPipeline *v1alpha1.LabelingPipeline) (result *v1alpha1.LabelingPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(labelingpipelinesResource, c.ns, labelingPipeline), &v1alpha1.LabelingPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LabelingPipeline), err
}

// Update takes the representation of a labelingPipeline and updates it. Returns the server's representation of the labelingPipeline, and an error, if there is any.
func (c *FakeLabelingPipelines) Update(labelingPipeline *v1alpha1.LabelingPipeline) (result *v1alpha1.LabelingPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(labelingpipelinesResource, c.ns, labelingPipeline), &v1alpha1.LabelingPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LabelingPipeline), err
}

// Delete takes name of the labelingPipeline and deletes it. Returns an error if one occurs.
func (c *FakeLabelingPipelines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(labelingpipelinesResource, c.ns, name), &v1alpha1.LabelingPipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLabelingPipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(labelingpipelinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LabelingPipelineList{})
	return err
}

// Patch applies the patch and returns the patched labelingPipeline.
func (c *FakeLabelingPipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LabelingPipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(labelingpipelinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LabelingPipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LabelingPipeline), err
}
