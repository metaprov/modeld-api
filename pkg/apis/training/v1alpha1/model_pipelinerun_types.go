package v1alpha1

import (
	"github.com/metaprov/modeldapi/pkg/apis/training"
	"github.com/metaprov/modeldapi/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PipelinePhase string

const (
	PipelinePhase_PreprocessingData               PipelinePhase = "PreprocessingData"
	PipelinePhase_PreprocessingCompleted          PipelinePhase = "PreprocessingCompleted"
	PipelinePhase_PreprocessingFailed             PipelinePhase = "PreprocessingFailed"
	PipelinePhase_NotebookRunning                 PipelinePhase = "NotebookRunning"
	PipelinePhase_NotebookReady                   PipelinePhase = "NotebookReady"
	PipelinePhase_NotebookFailed                  PipelinePhase = "NotebookFailed"
	PipelinePhase_StudyRunning                    PipelinePhase = "StudyRunning"
	PipelinePhase_StudyFailed                     PipelinePhase = "StudyFailed"
	PipelinePhase_StudyCompleted                  PipelinePhase = "StudyCompleted"
	PipelinePhase_ModelPublished                  PipelinePhase = "ModelPublished"
	PipelinePhase_TrainingApproved                PipelinePhase = "TrainingApproved"
	PipelinePhase_AcceptancePredictorCreating     PipelinePhase = "AcceptancePredictorCreating"
	PipelinePhase_AcceptancePredictorReady        PipelinePhase = "AcceptancePredictorReady"
	PipelinePhase_AcceptancePrecitionRunCreating  PipelinePhase = "AcceptancePrecitionRunCreating"
	PipelinePhase_AcceptancePrecitionRunCompleted PipelinePhase = "AcceptancePrecitionRunCompleted"
	PipelinePhase_AcceptancePrecitionRunFailed    PipelinePhase = "AcceptancePrecitionRunFailed"
	PipelinePhase_AcceptanceApproved              PipelinePhase = "AcceptanceApproved"
	PipelinePhase_CapacityPredictorCreating       PipelinePhase = "CapacityPredictorCreating"
	PipelinePhase_CapacityPredictorReady          PipelinePhase = "CapacityPredictorReady"
	PipelinePhase_CapacityPrecitionCreating       PipelinePhase = "CapacityPrecitionCreating"
	PipelinePhase_CapacityPrecitionRunCompleted   PipelinePhase = "CapacityPrecitionRunCompleted"
	PipelinePhase_CapacityPrecitionRunFailed      PipelinePhase = "CapacityPrecitionRunFailed"
	PipelinePhase_CapcityApproved                 PipelinePhase = "CapcityApproved"
	PipelinePhase_ReleaseModelDeployed            PipelinePhase = "ReleaseModelDeployed"
	PipelinePhase_ReleaseModelLive                PipelinePhase = "ReleaseModelLive"
)

// +kubebuilder:validation:Enum="data";"train";"test";"prod"
type EnvName string

const (
	DataEnvName  EnvName = "data"
	TrainEnvName EnvName = "training"
	TestEnvName  EnvName = "staging"
	ProdEnvName  EnvName = "prod"
)

// Condition on the dataset
type PipelineRunConditionType string

/// ModelPipelineRun Condition
const (
	PipelineRunAvailable PipelineRunConditionType = "Available"
	PipelineRunInDb      PipelineRunConditionType = "InDb"

	// Paths stage
	PipelineRunDataReady PipelineRunConditionType = "DataReady"

	// TrainingSpec stage
	PipelineRunTrainingNotebookReady PipelineRunConditionType = "TrainingNotebookReady"
	PipelineRunTrainingStudyReady    PipelineRunConditionType = "TrainingStudyReady"
	PipelineRunTrainingApproved      PipelineRunConditionType = "TrainingApproved"

	// Acceptance stage
	PipelineRunAcceptancePredictorReady  PipelineRunConditionType = "AcceptancePredictorReady"
	PipelineRunAcceptancePredictionReady PipelineRunConditionType = "AcceptancePredictionReady"
	PipelineRunAcceptanceApproved        PipelineRunConditionType = "AcceptanceApproved"

	//Capacity
	PipelineRunCapacityPredictorReady  PipelineRunConditionType = "CapacityPredictorReady"
	PipelineRunCapacityPredictionReady PipelineRunConditionType = "CapacityPredictionReady"
	PipelineRunCapacityApproved        PipelineRunConditionType = "CapacityApproved"

	// Deployed
	PipelineModelDeployed PipelineRunConditionType = "ModelDeployed"

	PipelineRunReady  PipelineRunConditionType = "Ready"
	PipelineRunFailed PipelineRunConditionType = "Failed"
)

// ModelPipelineRunCondition describes the state of a ModelPipelineRun at a certain point.
type ModelPipelineRunCondition struct {
	// Type of  condition.
	Type PipelineRunConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=PipelineRunConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="ModelPipeline",type="string",JSONPath=".status.pipelineName"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".status.versionName"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:path=modelpipelineruns,shortName=pr,singular=modelpipelinerun,categories={train,modeld,all}
// ModelPipelineRun represent a execution of a pipeline
type ModelPipelineRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ModelPipelineRunSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	// +kubebuilder:validation:Optional
	Status ModelPipelineRunStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

func (run *ModelPipelineRun) ValidateDelete() error {
	return nil
}

// +kubebuilder:object:root=true
// ModelPipelineRunList represent list of pipelineruns
type ModelPipelineRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ModelPipelineRun `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// ModelPipelineRunSpec is the desired state of the ModelPipelineRun resource
type ModelPipelineRunSpec struct {
	// User provided description
	// +kubebuilder:default:=""
	// +kubebuilder:validation:Optional
	VersionName *string `json:"versionName,omitempty" protobuf:"bytes,1,opt,name=versionName"`
	// User provided description
	// +kubebuilder:default:=""
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	// TriggerName denote the trigger that fired this pipeline run
	// +kubebuilder:default:=""
	// +kubebuilder:validation:Optional
	TriggerName *string `json:"triggerName,omitempty" protobuf:"bytes,3,opt,name=triggerName"`
	// PipelineName is the name of the pipeline specification for this pipeline
	// +kubebuilder:default:=""
	// +kubebuilder:validation:Optional
	PipelineName *string `json:"pipelineName,omitempty" protobuf:"bytes,4,opt,name=pipelineName"`
	// The owner account name
	// +kubebuilder:default:=""
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" protobuf:"bytes,5,opt,name=owner"`
}

// ModelPipelineRunStageStatus is the observed state of the PipelineRunStage.
type ModelPipelineRunStageStatus struct {
	// +kubebuilder:validation:Optional
	Approved bool `json:"approved,omitempty" protobuf:"bytes,1,opt,name=approved"`
	// +kubebuilder:validation:Optional
	ApprovedBy string `json:"approvedBy,omitempty" protobuf:"bytes,2,opt,name=approvedBy"`
	// +kubebuilder:validation:Optional
	Score float64 `json:"score,omitempty" protobuf:"bytes,3,opt,name=score"`
	// +kubebuilder:validation:Optional
	StartTime *metav1.Time `json:"startTime,omitempty" protobuf:"bytes,4,opt,name=startTime"`
	// +kubebuilder:validation:Optional
	EndTime *metav1.Time `json:"endTime,omitempty" protobuf:"bytes,5,opt,name=endTime"`
}

// ModelPipelineRunStatus is the observed state of the ModelPipelineRun resource .
type ModelPipelineRunStatus struct {
	// +kubebuilder:validation:Optional
	StudyName string `json:"studyName" protobuf:"bytes,1,opt,name=studyName"`
	// +kubebuilder:validation:Optional
	// data is the status for the data stage.
	DataStatus ModelPipelineRunStageStatus `json:"data,omitempty" protobuf:"bytes,3,opt,name=data"`
	// +kubebuilder:validation:Optional
	// dev is the status for the dev stage.
	TrainingStatus ModelPipelineRunStageStatus `json:"training,omitempty" protobuf:"bytes,4,opt,name=training"`
	// Staging is status for each qa stage.
	// +kubebuilder:validation:Optional
	AcceptanceStatus ModelPipelineRunStageStatus `json:"acceptance,omitempty" protobuf:"bytes,5,rep,name=acceptance"`
	// ReleaseStage is the status for the production stage.
	// +kubebuilder:validation:Optional
	CapacityStatus ModelPipelineRunStageStatus `json:"capacity,omitempty" protobuf:"bytes,6,rep,name=capacity"`
	// ReleaseStage is the status for the production stage.
	// +kubebuilder:validation:Optional
	ReleaseState ModelPipelineRunStageStatus `json:"prod,omitempty" protobuf:"bytes,7,rep,name=prod"`
	// +kubebuilder:validation:Optional
	StartTime *metav1.Time `json:"startTime,omitempty" protobuf:"bytes,8,opt,name=startTime"`
	// +kubebuilder:validation:Optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty" protobuf:"bytes,9,opt,name=completionTime"`
	// The phase of the pipeline run
	// +kubebuilder:validation:Optional
	Phase string `json:"phase" protobuf:"bytes,10,opt,name=phase"`
	//+optional
	Conditions []ModelPipelineRunCondition `json:"conditions,omitempty" protobuf:"bytes,11,rep,name=conditions"`
	// Folder for pipeline run artifacts. This is assigned by the system
	// The folder contains all the pipeline artifacts - metadata, logs
	// +kubebuilder:validation:Optional
	Folder string `json:"folder,omitempty" protobuf:"bytes,13,opt,name=evalMetrics"`
}

//==============================================================================
// Finalizer
//==============================================================================

func (run *ModelPipelineRun) HasFinalizer() bool {
	return util.HasFin(&run.ObjectMeta, training.GroupName)
}
func (run *ModelPipelineRun) AddFinalizer()    { util.AddFin(&run.ObjectMeta, training.GroupName) }
func (run *ModelPipelineRun) RemoveFinalizer() { util.RemoveFin(&run.ObjectMeta, training.GroupName) }

//==============================================================================
// PIpeline stage status
//==============================================================================

func (this *ModelPipelineRunStageStatus) RecordApprove(name string) {
	this.Approved = true
	this.ApprovedBy = name
}

func (this *ModelPipelineRunStageStatus) RecordDeny(name string) {
	this.Approved = false
	this.ApprovedBy = name
}
