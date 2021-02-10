/*
 * Copyright (c) 2019.
 *
 * Metaprov.com
 */

package v1alpha1

import (
	"testing"

	"github.com/metaprov/modeldapi/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDataPipelineRunValidate(t *testing.T) {
	dfr := DataPipelineRun{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "w",
			Namespace: "modeld-system",
		},
		Spec: DataPipelineRunSpec{
			DataPipelineName: "",
		},
	}
	dfr.Default()
	err := dfr.validate()
	util.OK(t, err)
}
