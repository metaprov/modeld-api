/*
 * Copyright (c) 2019.
 *
 * Metaprov.com
 */

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// validation
var _ webhook.Validator = &Account{}

func (account *Account) ValidateCreate() error {
	return account.validate()
}

func (account *Account) ValidateUpdate(old runtime.Object) error {
	return account.validate()
}

func (account *Account) validate() error {
	return nil
}

func (account *Account) ValidateDelete() error {
	return nil
}
