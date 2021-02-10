/*
 * Copyright (c) 2020.
 *
 * Metaprov.com
 */

package v1alpha1

import (
	"fmt"
	"github.com/metaprov/modeld-api/pkg/apis/data"
	"github.com/metaprov/modeld-api/pkg/util"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

func NewDataPipeline(ns string, name string, schemaName string) *DataPipeline {
	result := &DataPipeline{}
	result.ObjectMeta.Name = name
	result.ObjectMeta.Namespace = ns
	result.Kind = "DataPipeline"
	result.Default()
	return result
}

func (wr *DataPipeline) HasFinalizer() bool {
	return util.HasFin(&wr.ObjectMeta, data.GroupName)
}
func (wr *DataPipeline) AddFinalizer()    { util.AddFin(&wr.ObjectMeta, data.GroupName) }
func (wr *DataPipeline) RemoveFinalizer() { util.RemoveFin(&wr.ObjectMeta, data.GroupName) }

//==============================================================================
// Validate
//==============================================================================

// Merge or update condition
func (wr *DataPipeline) CreateOrUpdateCond(cond DataPipelineCondition) {
	i := wr.GetCondIdx(cond.Type)
	now := metav1.Now()
	if i == -1 { // not found
		cond.LastTransitionTime = &now
		wr.Status.Conditions = append(wr.Status.Conditions, cond)
		return
	}
	// else we already have the condition, update it
	current := wr.Status.Conditions[i]
	current.Message = cond.Message
	current.Reason = cond.Reason
	current.LastTransitionTime = &now
	if current.Status != cond.Status {
		current.Status = cond.Status
	}
	wr.Status.Conditions[i] = current
}

func (wr *DataPipeline) GetCondIdx(t DataPipelineConditionType) int {
	for i, v := range wr.Status.Conditions {
		if v.Type == t {
			return i
		}
	}
	return -1
}

func (wr *DataPipeline) GetCond(t DataPipelineConditionType) DataPipelineCondition {
	for _, v := range wr.Status.Conditions {
		if v.Type == t {
			return v
		}
	}
	// if we did not find the condition, we return an unknown object
	return DataPipelineCondition{
		Type:    t,
		Status:  v1.ConditionUnknown,
		Reason:  "",
		Message: "",
	}

}

func (wr *DataPipeline) RootUri() string {
	return fmt.Sprintf("dataproducts/%s/versions/%s/wranglers/%s", wr.Namespace, wr.Spec.VersionName, wr.Name)
}

func (wr *DataPipeline) ManifestUri() string {
	return fmt.Sprintf("%s/%s-wrangler.yaml", wr.RootUri(), wr.Name)
}

func ParseDataPipelineYaml(content []byte) (*DataPipeline, error) {
	requiredObj, err := runtime.Decode(scheme.Codecs.UniversalDecoder(SchemeGroupVersion), content)
	if err != nil {
		return nil, err
	}
	r := requiredObj.(*DataPipeline)
	return r, nil
}

func (wr *DataPipeline) ToYamlFile() ([]byte, error) {
	return yaml.Marshal(wr)
}
