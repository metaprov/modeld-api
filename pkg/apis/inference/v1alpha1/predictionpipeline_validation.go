package v1alpha1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// validation
var _ webhook.Validator = &PredictionPipeline{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (prediction *PredictionPipeline) ValidateCreate() error {
	return prediction.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (prediction *PredictionPipeline) ValidateUpdate(old runtime.Object) error {
	return prediction.validate()
}

func (prediction *PredictionPipeline) validate() error {
	var allErrs field.ErrorList
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "inference.modeld.io", Kind: "PredictionPipeline"},
		prediction.Name, allErrs)
}

func (prediction *PredictionPipeline) ValidateDelete() error {
	return nil
}
