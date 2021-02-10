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

// FakeAccounts implements AccountInterface
type FakeAccounts struct {
	Fake *FakeInfraV1alpha1
	ns   string
}

var accountsResource = schema.GroupVersionResource{Group: "infra.modeld.io", Version: "v1alpha1", Resource: "accounts"}

var accountsKind = schema.GroupVersionKind{Group: "infra.modeld.io", Version: "v1alpha1", Kind: "Account"}

// Get takes name of the account, and returns the corresponding account object, and an error if there is any.
func (c *FakeAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountsResource, c.ns, name), &v1alpha1.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// List takes label and field selectors, and returns the list of Accounts that match those selectors.
func (c *FakeAccounts) List(opts v1.ListOptions) (result *v1alpha1.AccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountsResource, accountsKind, c.ns, opts), &v1alpha1.AccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountList{ListMeta: obj.(*v1alpha1.AccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accounts.
func (c *FakeAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountsResource, c.ns, opts))

}

// Create takes the representation of a account and creates it.  Returns the server's representation of the account, and an error, if there is any.
func (c *FakeAccounts) Create(account *v1alpha1.Account) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountsResource, c.ns, account), &v1alpha1.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// Update takes the representation of a account and updates it. Returns the server's representation of the account, and an error, if there is any.
func (c *FakeAccounts) Update(account *v1alpha1.Account) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountsResource, c.ns, account), &v1alpha1.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccounts) UpdateStatus(account *v1alpha1.Account) (*v1alpha1.Account, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountsResource, "status", c.ns, account), &v1alpha1.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// Delete takes name of the account and deletes it. Returns an error if one occurs.
func (c *FakeAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accountsResource, c.ns, name), &v1alpha1.Account{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountList{})
	return err
}

// Patch applies the patch and returns the patched account.
func (c *FakeAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Account{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}
