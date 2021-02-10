/*
 * Copyright (c) 2020.
 *
 * Metaprov.com
 */

package v1alpha1

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/metaprov/modeld-api/pkg/apis/common"
	"github.com/metaprov/modeld-api/pkg/apis/data"
	"github.com/metaprov/modeld-api/pkg/util"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

//==============================================================================
// EntityRef
//==============================================================================

func (lpr *LabelingPipelineRun) AddConfiditions() {
	lpr.Status.Conditions = make([]LabelingPipelineRunCondition, 1)
	lpr.Status.Conditions[0] = LabelingPipelineRunCondition{
		Type:   LabelingPipelineRunReady,
		Status: v1.ConditionUnknown,
	}
}

func (lpr *LabelingPipelineRun) HasFinalizer() bool {
	return util.HasFin(&lpr.ObjectMeta, data.GroupName)
}
func (lpr *LabelingPipelineRun) AddFinalizer() { util.AddFin(&lpr.ObjectMeta, data.GroupName) }
func (lpr *LabelingPipelineRun) RemoveFinalizer() {
	util.RemoveFin(&lpr.ObjectMeta, data.GroupName)
}

//==============================================================================
// Trackable
//==============================================================================

// Return the on disk rep location

func (lpr *LabelingPipelineRun) ToYamlFile() ([]byte, error) {
	return yaml.Marshal(lpr)
}

func (lpr *LabelingPipelineRun) Age() string {
	return humanize.Time(lpr.CreationTimestamp.Time)
}

//==============================================================================
// Factory method
//==============================================================================

// Parse an data
func ParseLabelPipelineRun(content string, user string, commit string) (*LabelingPipelineRun, error) {
	this := &LabelingPipelineRun{}
	err := yaml.Unmarshal([]byte(content), this)
	if err != nil {
		return nil, err
	}
	return this, nil
}

//==============================================================================
// Assign commit and id
//==============================================================================

func (lpr *LabelingPipelineRun) LabelWithCommit(commit string, uname string, branch string) {
	lpr.ObjectMeta.Labels[common.CommitLabelKey] = commit
	lpr.ObjectMeta.Labels[common.UnameLabelKey] = uname
	lpr.ObjectMeta.Labels[common.BranchLabelKey] = branch
}

func (lpr *LabelingPipelineRun) IsGitObj() bool {
	label, ok := lpr.ObjectMeta.Labels[common.CommitLabelKey]
	if !ok {
		return false
	}
	return label != ""
}

func (lpr *LabelingPipelineRun) SetChanged() {
	lpr.ObjectMeta.Labels[common.ChangedLabelKey] = "true"

}

// Merge or update condition
// Merge or update condition
func (lpr *LabelingPipelineRun) CreateOrUpdateCond(cond LabelingPipelineRunCondition) {
	i := lpr.GetCondIdx(cond.Type)
	now := metav1.Now()
	if i == -1 { // not found
		cond.LastTransitionTime = &now
		lpr.Status.Conditions = append(lpr.Status.Conditions, cond)
		return
	}
	// else we already have the condition, update it
	current := lpr.Status.Conditions[i]
	current.Message = cond.Message
	current.Reason = cond.Reason
	current.LastTransitionTime = &now
	if current.Status != cond.Status {
		current.Status = cond.Status
	}
	lpr.Status.Conditions[i] = current
}

func (lpr *LabelingPipelineRun) GetCondIdx(t LabelingPipelineRunConditionType) int {
	for i, v := range lpr.Status.Conditions {
		if v.Type == t {
			return i
		}
	}
	return -1
}

func (lpr *LabelingPipelineRun) GetCond(t LabelingPipelineRunConditionType) LabelingPipelineRunCondition {
	for _, v := range lpr.Status.Conditions {
		if v.Type == t {
			return v
		}
	}
	// if we did not find the condition, we return an unknown object
	return LabelingPipelineRunCondition{
		Type:    t,
		Status:  v1.ConditionUnknown,
		Reason:  "",
		Message: "",
	}

}

func (lpr *LabelingPipelineRun) IsReady() bool {
	return lpr.GetCond(LabelingPipelineRunReady).Status == v1.ConditionTrue
}

func (lpr *LabelingPipelineRun) Key() string {
	return fmt.Sprintf("%s/%s/%s", "features", lpr.Namespace, lpr.Name)
}

func ParseLabelPipelineRunYaml(content []byte) (*LabelingPipelineRun, error) {
	requiredObj, err := runtime.Decode(scheme.Codecs.UniversalDecoder(SchemeGroupVersion), content)
	if err != nil {
		return nil, err
	}
	r := requiredObj.(*LabelingPipelineRun)
	return r, nil
}

func (lpr *LabelingPipelineRun) MarkReady() {
	// update the lab state to ready
	lpr.CreateOrUpdateCond(LabelingPipelineRunCondition{
		Type:   LabelingPipelineRunReady,
		Status: v1.ConditionTrue,
	})

}
