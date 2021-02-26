/*
 * Copyright (c) 2019.
 *
 * Metaprov.com
 */

package v1alpha1

import (
	catalog "github.com/metaprov/modeldapi/pkg/apis/catalog/v1alpha1"
	"github.com/metaprov/modeldapi/pkg/apis/common"
	"github.com/metaprov/modeldapi/pkg/util"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (model *Model) Default() {

	if model.Spec.OwnerName == nil {
		model.Spec.OwnerName = util.StrPtr("")
	}
	// Objective
	if model.Spec.Training == nil {
		model.Spec.Training = &TrainingSpec{}
	}

	if model.Spec.Training.Priority == nil {
		model.Spec.Training.Priority = util.Int32Ptr(int32(0)) // default prioirty is 5
	}

	if model.Spec.Aborted == nil {
		model.Spec.Aborted = util.BoolPtr(false)
	}

	if model.Spec.Published == nil {
		model.Spec.Published = util.BoolPtr(false)
	}

	if model.Spec.Reported == nil {
		model.Spec.Reported = util.BoolPtr(false)
	}

	if model.Spec.Paused == nil {
		model.Spec.Paused = util.BoolPtr(false)
	}

	if model.Spec.Profiled == nil {
		model.Spec.Profiled = util.BoolPtr(false)
	}

	if model.Spec.Training.CV == nil {
		model.Spec.Training.CV = util.BoolPtr(false)
	}

	if model.Spec.Objective == nil {
		o := model.DefaultObjective()
		model.Spec.Objective = &o
	}
	if model.Spec.Training.Priority == nil {
		model.Spec.Training.Priority = util.Int32Ptr(1)
	}
	if model.Spec.Estimator == nil {
		model.Spec.Estimator = &ClassicalEstimatorSpec{}
	}
	if model.Spec.Preprocessing == nil {
		model.Spec.Preprocessing = &PreprocessingSpec{}
	}

	if model.Spec.Preprocessing.Text == nil {
		tfidf := catalog.TfIdf
		model.Spec.Preprocessing.Text = &TextPipelineSpec{
			Columns:   []string{},
			Encoder:   &tfidf,
			Tokenizer: util.StrPtr(""),
			StopWords: util.BoolPtr(false),
			Pos:       util.BoolPtr(false),
			Lemma:     util.BoolPtr(false),
			Stem:      util.BoolPtr(false),
			Embedding: nil,
		}
	}

	if model.Spec.Ensemble == nil {
		model.Spec.Ensemble = &EnsembleSpec{
			Base: make([]string, 0),
		}
	}

	if model.Spec.Training.Folds == nil {
		model.Spec.Training.Folds = util.Int32Ptr(5)
	}

	if model.Spec.Preprocessing.Imbalanced == nil {
		model.Spec.Preprocessing.Imbalanced = util.BoolPtr(false)
	}
	if model.Spec.Tested == nil {
		model.Spec.Tested = util.BoolPtr(false)
	}
	if model.Spec.Aborted == nil {
		model.Spec.Aborted = util.BoolPtr(false)
	}
	if model.Spec.Published == nil {
		model.Spec.Published = util.BoolPtr(false)
	}
	if model.Spec.Reported == nil {
		model.Spec.Reported = util.BoolPtr(false)
	}
	if model.Spec.Paused == nil {
		model.Spec.Paused = util.BoolPtr(false)
	}
	if model.Spec.Profiled == nil {
		model.Spec.Profiled = util.BoolPtr(false)
	}
	if model.Spec.Archived == nil {
		model.Spec.Archived = util.BoolPtr(false)
	}
	if model.Spec.Published == nil {
		model.Spec.Published = util.BoolPtr(false)
	}

}

// validation
var _ webhook.Validator = &Model{}

//Set up the webhook with the manager.
func (r *Model) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

func (in *Model) ValidateDelete() error {
	return nil
}

//==============================================================================
// Validate
//==============================================================================

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (model *Model) ValidateCreate() error {
	return model.validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (model *Model) ValidateUpdate(old runtime.Object) error {
	return model.validate()
}

func (notifier *Model) validate() error {
	var allErrs field.ErrorList
	allErrs = append(allErrs, notifier.validateMeta(field.NewPath("metadata"))...)
	allErrs = append(allErrs, notifier.validateSpec(field.NewPath("spec"))...)
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "training.modeld.io", Kind: "dModel"},
		notifier.Name, allErrs)
}

func (dataset *Model) validateMeta(fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	allErrs = append(allErrs, dataset.validateName(fldPath.Child("name"))...)
	return allErrs
}

func (dataset *Model) validateName(fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	err := common.ValidateResourceName(dataset.Name)
	if err != nil {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("FileName"), dataset.Name, err.Error()))
	}
	return allErrs
}

func (dataset *Model) validateSpec(fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	return allErrs
}