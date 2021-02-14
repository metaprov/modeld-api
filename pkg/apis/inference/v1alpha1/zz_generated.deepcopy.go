// +build !ignore_autogenerated

/**
*
* Copyright (C) 2017 modeld authors
* For license information, see LICENSE.txt
 */
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	catalogv1alpha1 "github.com/metaprov/modeldapi/pkg/apis/catalog/v1alpha1"
	datav1alpha1 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BotChannelSpec) DeepCopyInto(out *BotChannelSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BotChannelSpec.
func (in *BotChannelSpec) DeepCopy() *BotChannelSpec {
	if in == nil {
		return nil
	}
	out := new(BotChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketChannelSpec) DeepCopyInto(out *BucketChannelSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketChannelSpec.
func (in *BucketChannelSpec) DeepCopy() *BucketChannelSpec {
	if in == nil {
		return nil
	}
	out := new(BucketChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelStatus) DeepCopyInto(out *ChannelStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelStatus.
func (in *ChannelStatus) DeepCopy() *ChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Curtain) DeepCopyInto(out *Curtain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Curtain.
func (in *Curtain) DeepCopy() *Curtain {
	if in == nil {
		return nil
	}
	out := new(Curtain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurtainCondition) DeepCopyInto(out *CurtainCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurtainCondition.
func (in *CurtainCondition) DeepCopy() *CurtainCondition {
	if in == nil {
		return nil
	}
	out := new(CurtainCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurtainList) DeepCopyInto(out *CurtainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Curtain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurtainList.
func (in *CurtainList) DeepCopy() *CurtainList {
	if in == nil {
		return nil
	}
	out := new(CurtainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CurtainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurtainSpec) DeepCopyInto(out *CurtainSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Wizards != nil {
		in, out := &in.Wizards, &out.Wizards
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurtainSpec.
func (in *CurtainSpec) DeepCopy() *CurtainSpec {
	if in == nil {
		return nil
	}
	out := new(CurtainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurtainStatus) DeepCopyInto(out *CurtainStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CurtainCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurtainStatus.
func (in *CurtainStatus) DeepCopy() *CurtainStatus {
	if in == nil {
		return nil
	}
	out := new(CurtainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurtainTemplateSpec) DeepCopyInto(out *CurtainTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurtainTemplateSpec.
func (in *CurtainTemplateSpec) DeepCopy() *CurtainTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CurtainTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriftCheckSpec) DeepCopyInto(out *DriftCheckSpec) {
	*out = *in
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.NotifierName != nil {
		in, out := &in.NotifierName, &out.NotifierName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriftCheckSpec.
func (in *DriftCheckSpec) DeepCopy() *DriftCheckSpec {
	if in == nil {
		return nil
	}
	out := new(DriftCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpec) DeepCopyInto(out *ModelDeploymentSpec) {
	*out = *in
	if in.ModelName != nil {
		in, out := &in.ModelName, &out.ModelName
		*out = new(string)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.AutoScale != nil {
		in, out := &in.AutoScale, &out.AutoScale
		*out = new(bool)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(bool)
		**out = **in
	}
	if in.Shadow != nil {
		in, out := &in.Shadow, &out.Shadow
		*out = new(bool)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.CanaryMetrics != nil {
		in, out := &in.CanaryMetrics, &out.CanaryMetrics
		*out = make([]CanaryMetric, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpec.
func (in *ModelDeploymentSpec) DeepCopy() *ModelDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentStatus) DeepCopyInto(out *ModelDeploymentStatus) {
	*out = *in
	if in.LastPrediction != nil {
		in, out := &in.LastPrediction, &out.LastPrediction
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentStatus.
func (in *ModelDeploymentStatus) DeepCopy() *ModelDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionChannel) DeepCopyInto(out *PredictionChannel) {
	*out = *in
	if in.Table != nil {
		in, out := &in.Table, &out.Table
		*out = new(TableChannelSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Bot != nil {
		in, out := &in.Bot, &out.Bot
		*out = new(BotChannelSpec)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(BucketChannelSpec)
		**out = **in
	}
	if in.Streaming != nil {
		in, out := &in.Streaming, &out.Streaming
		*out = new(StreamingChannelSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionChannel.
func (in *PredictionChannel) DeepCopy() *PredictionChannel {
	if in == nil {
		return nil
	}
	out := new(PredictionChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipeline) DeepCopyInto(out *PredictionPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipeline.
func (in *PredictionPipeline) DeepCopy() *PredictionPipeline {
	if in == nil {
		return nil
	}
	out := new(PredictionPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineCondition) DeepCopyInto(out *PredictionPipelineCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineCondition.
func (in *PredictionPipelineCondition) DeepCopy() *PredictionPipelineCondition {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineList) DeepCopyInto(out *PredictionPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PredictionPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineList.
func (in *PredictionPipelineList) DeepCopy() *PredictionPipelineList {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineRun) DeepCopyInto(out *PredictionPipelineRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineRun.
func (in *PredictionPipelineRun) DeepCopy() *PredictionPipelineRun {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineRunCondition) DeepCopyInto(out *PredictionPipelineRunCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineRunCondition.
func (in *PredictionPipelineRunCondition) DeepCopy() *PredictionPipelineRunCondition {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineRunCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineRunList) DeepCopyInto(out *PredictionPipelineRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PredictionPipelineRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineRunList.
func (in *PredictionPipelineRunList) DeepCopy() *PredictionPipelineRunList {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineRunSpec) DeepCopyInto(out *PredictionPipelineRunSpec) {
	*out = *in
	if in.Objective != nil {
		in, out := &in.Objective, &out.Objective
		*out = new(catalogv1alpha1.Metric)
		**out = **in
	}
	if in.DatasetName != nil {
		in, out := &in.DatasetName, &out.DatasetName
		*out = new(string)
		**out = **in
	}
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = new(datav1alpha1.DataLocation)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(datav1alpha1.DataLocation)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineRunSpec.
func (in *PredictionPipelineRunSpec) DeepCopy() *PredictionPipelineRunSpec {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineRunStatus) DeepCopyInto(out *PredictionPipelineRunStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PredictionPipelineRunCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineRunStatus.
func (in *PredictionPipelineRunStatus) DeepCopy() *PredictionPipelineRunStatus {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineSpec) DeepCopyInto(out *PredictionPipelineSpec) {
	*out = *in
	if in.Objective != nil {
		in, out := &in.Objective, &out.Objective
		*out = new(catalogv1alpha1.Metric)
		**out = **in
	}
	if in.DatasetName != nil {
		in, out := &in.DatasetName, &out.DatasetName
		*out = new(string)
		**out = **in
	}
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = new(datav1alpha1.DataLocation)
		**out = **in
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(datav1alpha1.DataLocation)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineSpec.
func (in *PredictionPipelineSpec) DeepCopy() *PredictionPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionPipelineStatus) DeepCopyInto(out *PredictionPipelineStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PredictionPipelineCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionPipelineStatus.
func (in *PredictionPipelineStatus) DeepCopy() *PredictionPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PredictionPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Predictor) DeepCopyInto(out *Predictor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Predictor.
func (in *Predictor) DeepCopy() *Predictor {
	if in == nil {
		return nil
	}
	out := new(Predictor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorCondition) DeepCopyInto(out *PredictorCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorCondition.
func (in *PredictorCondition) DeepCopy() *PredictorCondition {
	if in == nil {
		return nil
	}
	out := new(PredictorCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorHealth) DeepCopyInto(out *PredictorHealth) {
	*out = *in
	if in.LastDailyPredictions != nil {
		in, out := &in.LastDailyPredictions, &out.LastDailyPredictions
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorHealth.
func (in *PredictorHealth) DeepCopy() *PredictorHealth {
	if in == nil {
		return nil
	}
	out := new(PredictorHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorList) DeepCopyInto(out *PredictorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Predictor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorList.
func (in *PredictorList) DeepCopy() *PredictorList {
	if in == nil {
		return nil
	}
	out := new(PredictorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PredictorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorSpec) DeepCopyInto(out *PredictorSpec) {
	*out = *in
	if in.OwnerName != nil {
		in, out := &in.OwnerName, &out.OwnerName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ProductRef != nil {
		in, out := &in.ProductRef, &out.ProductRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ServingSiteRef != nil {
		in, out := &in.ServingSiteRef, &out.ServingSiteRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(AccessType)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(v1.PodTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Models != nil {
		in, out := &in.Models, &out.Models
		*out = make([]ModelDeploymentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DriftCheck != nil {
		in, out := &in.DriftCheck, &out.DriftCheck
		*out = new(DriftCheckSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Progressive != nil {
		in, out := &in.Progressive, &out.Progressive
		*out = new(ProgressiveSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ArtifactsFolder != nil {
		in, out := &in.ArtifactsFolder, &out.ArtifactsFolder
		*out = new(string)
		**out = **in
	}
	if in.InputChannels != nil {
		in, out := &in.InputChannels, &out.InputChannels
		*out = make([]PredictionChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OutputChannels != nil {
		in, out := &in.OutputChannels, &out.OutputChannels
		*out = make([]PredictionChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorSpec.
func (in *PredictorSpec) DeepCopy() *PredictorSpec {
	if in == nil {
		return nil
	}
	out := new(PredictorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorStatus) DeepCopyInto(out *PredictorStatus) {
	*out = *in
	if in.ModelStatuses != nil {
		in, out := &in.ModelStatuses, &out.ModelStatuses
		*out = make([]ModelDeploymentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PredictorCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.MonitorLastAttemptAt = in.MonitorLastAttemptAt
	in.Health.DeepCopyInto(&out.Health)
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]ChannelStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorStatus.
func (in *PredictorStatus) DeepCopy() *PredictorStatus {
	if in == nil {
		return nil
	}
	out := new(PredictorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorTemplate) DeepCopyInto(out *PredictorTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorTemplate.
func (in *PredictorTemplate) DeepCopy() *PredictorTemplate {
	if in == nil {
		return nil
	}
	out := new(PredictorTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PredictorTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorTemplateList) DeepCopyInto(out *PredictorTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PredictorTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorTemplateList.
func (in *PredictorTemplateList) DeepCopy() *PredictorTemplateList {
	if in == nil {
		return nil
	}
	out := new(PredictorTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PredictorTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorTemplateSpec) DeepCopyInto(out *PredictorTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorTemplateSpec.
func (in *PredictorTemplateSpec) DeepCopy() *PredictorTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(PredictorTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveSpec) DeepCopyInto(out *ProgressiveSpec) {
	*out = *in
	if in.Warmup != nil {
		in, out := &in.Warmup, &out.Warmup
		*out = new(int32)
		**out = **in
	}
	if in.TrafficIncrement != nil {
		in, out := &in.TrafficIncrement, &out.TrafficIncrement
		*out = new(int32)
		**out = **in
	}
	if in.CanaryMetrics != nil {
		in, out := &in.CanaryMetrics, &out.CanaryMetrics
		*out = make([]CanaryMetric, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveSpec.
func (in *ProgressiveSpec) DeepCopy() *ProgressiveSpec {
	if in == nil {
		return nil
	}
	out := new(ProgressiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamingChannelSpec) DeepCopyInto(out *StreamingChannelSpec) {
	*out = *in
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamingChannelSpec.
func (in *StreamingChannelSpec) DeepCopy() *StreamingChannelSpec {
	if in == nil {
		return nil
	}
	out := new(StreamingChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableChannelSpec) DeepCopyInto(out *TableChannelSpec) {
	*out = *in
	if in.DataSourceName != nil {
		in, out := &in.DataSourceName, &out.DataSourceName
		*out = new(string)
		**out = **in
	}
	if in.TableName != nil {
		in, out := &in.TableName, &out.TableName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableChannelSpec.
func (in *TableChannelSpec) DeepCopy() *TableChannelSpec {
	if in == nil {
		return nil
	}
	out := new(TableChannelSpec)
	in.DeepCopyInto(out)
	return out
}
