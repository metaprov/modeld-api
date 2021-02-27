/*
 * Copyright (c) 2019.
 *
 * Metaprov.com
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//============
// BucketName
// ===========

type PostMortemConditionType string

/// Post Moretem Condition
const (
	PostMortemReady PostMortemConditionType = "Ready"
)

// Post Morem Condition describes the state of a post mortem at a certain point.
type PostMortemCondition struct {
	// Type of bucket condition.
	Type PostMortemConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PostMortemConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description=""
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=postmortems,shortName=pt,singular=postmortem,categories={team,modeld,all}
// PostMortem represent a post mortem
type PostMortem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              PostMortemSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            PostMortemStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// PostMortemList is a list of PostMortems
type PostMortemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []PostMortem `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// TimeLineEvent
type TimeLineEvent struct {
	// +optional
	// +kubebuilder:default = ""
	Content *string `json:"content" protobuf:"bytes,1,opt,name=content"`
	// +optional
	At *metav1.Time `json:"at" protobuf:"bytes,2,opt,name=at"`
}

// Alert is an alert that was raised during the incident
type Alert struct {
	// +optional
	Name *string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// +optional
	At *metav1.Time `json:"at" protobuf:"bytes,2,opt,name=at"`
}

// PostMortemSpec defines the desired state of PostMorem
type PostMortemSpec struct {
	// The target object of the post mortem for example a predictor.
	// +optional
	EntityRef *v1.ObjectReference `json:"entityRef" protobuf:"bytes,1,opt,name=entityRef"`
	// The place of the events, can be lab or serving site.
	// +optional
	LocationRef *v1.ObjectReference `json:"locationRef" protobuf:"bytes,2,opt,name=locationRef"`
	// Reference to accounts that handled the post-mortem
	// +optional
	HandledBy []v1.ObjectReference `json:"handledBy" protobuf:"bytes,3,rep,name=handledBy"`
	// Alerts
	// +optional
	Alerts []Alert `json:"alerts" protobuf:"bytes,4,rep,name=alert"`
	// Timeline is the list of events that happed during the events
	// +optional
	TimeLine []TimeLineEvent `json:"timeline" protobuf:"bytes,5,rep,name=timeline"`
	// Root cause
	// +kubebuilder:default = ""
	// +optional
	RootCause *string `json:"rootCause" protobuf:"bytes,6,opt,name=rootCause"`
	// Summary of the post mortem
	// +kubebuilder:default = ""
	// +optional
	Summary *string `json:"summary" protobuf:"bytes,7,opt,name=summary"`
	// The owner account name
	// +kubebuilder:default = ""
	// +optional
	Owner *string `json:"owner,omitempty" protobuf:"bytes,8,opt,name=owner"`
}

// PostMortemStatus defines the actual state of a PostMortem
type PostMortemStatus struct {
	//+optional
	Conditions []PostMortemCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
}
