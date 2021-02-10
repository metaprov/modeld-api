package v1alpha1

import (
	"github.com/metaprov/modeld-api/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func OneModel() Model {
	return Model{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Model",
			APIVersion: "v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "iris-model",
			Namespace: "lab1",
		},
		Spec: ModelSpec{
			VersionName: "",
			StudyName:   "",
			Task:        "",
			Objective:   nil,
			Estimator:   nil,
			Dnn:         nil,
		},
	}
}

func TestModelConditions(t *testing.T) {
	model := OneModel()
	model.Default()

	cond := ModelCondition{
		Type:   ModelTrained,
		Status: v1.ConditionUnknown,
	}

	idx := model.GetCondIdx(ModelTrained)
	util.Assert(t, idx == -1, "not found")

	model.CreateOrUpdateCond(cond)
	util.Assert(t, len(model.Status.Conditions) == 1, "expect 1 conditions")

	after := model.GetCond(ModelTrained)
	util.Assert(t, !after.LastTransitionTime.IsZero(), "expect valid time")

	// do an update
}
