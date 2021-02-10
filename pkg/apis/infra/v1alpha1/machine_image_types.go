package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//==============================================================================
// MachineImage
//==============================================================================

// MachineImage condition
type MachineImageConditionType string

/// MachineImage Condition
const (
	MachineImageReady MachineImageConditionType = "Ready"
)

// MachineImageCondition describes the state of the MachineImage at a certain point.
type MachineImageCondition struct {
	// Type of account condition.
	Type MachineImageConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=MachineImageConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=machineimages,singular=machineimage,categories={infra,modeld}
// MachineImage is an abstraction of virtual machine image that is used when
// adding new nodes to the cluster
type MachineImage struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Spec holds the desired state of the machine image.
	Spec MachineImageSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	// Status holds the desired state of the machine image.
	//+optional
	Status MachineImageStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
type MachineImageList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []MachineImage `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// MachineImageSpec defines the desired state of MachineImage
type MachineImageSpec struct {
	// The packer file definition.
	Content string `json:"content" protobuf:"bytes,2,opt,name=content"`
}

type MachineImageStatus struct {
	//+optional
	Conditions []MachineImageCondition `json:"conditions,omitempty" protobuf:"bytes,1,rep,name=conditions"`
}
