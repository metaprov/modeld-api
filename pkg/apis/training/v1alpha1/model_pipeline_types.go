package v1alpha1

import (
	catalog "github.com/metaprov/modeldapi/pkg/apis/catalog/v1alpha1"
	data "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ModelPipelineConditionType string

const (
	ModelPipelineReady ModelPipelineConditionType = "Ready"
	ModelPipelineSaved ModelPipelineConditionType = "Saved"
)

// ModelPipelineCondition describes the state of a pipeline at a certain point.
type ModelPipelineCondition struct {
	// Type of account condition.
	Type ModelPipelineConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ModelPipelineConditionType"`
	// Status of the condition, one of True, False, AutoScaler.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
// +kubebuilder:printcolumn:name="Schedule",type="string",JSONPath=".spec.schedule",description=""
// +kubebuilder:printcolumn:name="Last Run",type="date",JSONPath=".status.lastRun",description=""
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:path=modelpipelines,singular=modelpipeline,shortName=pipe,categories={training,modeld,all}
// ModelPipeline represent a CI/CD machine learning pipeline definition
type ModelPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ModelPipelineSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            ModelPipelineStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +kubebuilder:object:root=true
// ModelPipelineList represent list of pipelines
type ModelPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ModelPipeline `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// ModelPipelineSpec define the desired state of the ModelPipeline resource.
type ModelPipelineSpec struct {
	// The product version of the resource
	// +kubebuilder:default ="latest"
	// +kubebuilder:validation:Optional
	VersionName *string `json:"versionName,omitempty" protobuf:"bytes,1,opt,name=versionName"`
	// User provided description
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" protobuf:"bytes,3,opt,name=description"`
	// DatasetSelector is used to select dataset for training
	// +kubebuilder:validation:Optional
	DatasetSelector map[string]string `json:"datasetSelector,omitempty" protobuf:"bytes,4,rep,name=datasetSelector"`
	// Datastage build new dataset from the data sources.
	// +kubebuilder:validation:Optional
	Data *DataStageSpec `json:"data,omitempty" protobuf:"bytes,5,opt,name=data"`
	// TrainingSpec stage
	// +kubebuilder:validation:Optional
	Training *TrainingStageSpec `json:"training,omitempty" protobuf:"bytes,6,opt,name=training"`
	// Acceptance stage is used for further testing
	// +kubebuilder:validation:Optional
	UAT *UATStageSpec `json:"uat,omitempty" protobuf:"bytes,7,opt,name=uat"`
	// Capacity stage for capacity
	// +kubebuilder:validation:Optional
	Capacity *CapacityStageSpec `json:"capacity,omitempty" protobuf:"bytes,8,opt,name=capacity"`
	// Deployment stage define how to place the model into production.
	// +kubebuilder:validation:Optional
	Deployment *DeploymentStageSpec `json:"deployment,omitempty" protobuf:"bytes,9,opt,name=deployment"`
	// Deployment stage define how to place the model into production.
	// +kubebuilder:validation:Optional
	Release *ReleaseStageSpec `json:"release,omitempty" protobuf:"bytes,10,opt,name=release"`
	// Folder for the pipeline and pipeline run artifacts.
	// The folder contains all the study artifacts - metadata, reports, profile,models
	// +kubebuilder:validation:Optional
	Location *data.DataLocation `json:"location,omitempty" protobuf:"bytes,13,opt,name=location"`
	// Schedule for running the pipeline
	// +kubebuilder:validation:Optional
	Schedule string `json:"schedule,omitempty" protobuf:"bytes,14,opt,name=schedule"`
	// The owner of the run, set to the owner of the pipeline
	// +kubebuilder:default:="no-one"
	// +kubebuilder:validation:Pattern="[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*"
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" protobuf:"bytes,15,opt,name=owner"`
	// ApproverAccountName is the name of the approver for stages that need approvals.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	ApproverAccountName *string `json:"approverAccountName,omitempty" protobuf:"bytes,16,opt,name=approverAccountName"`
	// NotifierName is the name of the notifier to use in case of pipeline failure
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	NotifierName *string `json:"notifierName,omitempty" protobuf:"bytes,17,opt,name=notifierName"`
	// BaselineModelName is the name of the model which is used to compare with this pipeline results.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	BaselineModelName *string `json:"baselineModelName,omitempty" protobuf:"bytes,18,opt,name=baselineModelName"`
	// The priority of this pipeline. The defualt is medium.
	// +kubebuilder:default:=medium
	// +kubebuilder:validation:Optional
	Priority *catalog.PriorityLevel `json:"priority,omitempty" protobuf:"bytes,30,opt,name=priority"`
	// Set to true to pause the model pipeline
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	Paused *bool `json:"paused,omitempty" protobuf:"bytes,31,opt,name=paused"`
}

//DataStageSpec is the desired state of the data preprocesing step of the pipeline.
//Data preprocessing will be done via
type DataStageSpec struct {
	// Enabled indicates that the stage is enabled
	// +kubebuilder:default:=true
	Enabled *bool `json:"enabled,omitempty" protobuf:"varint,1,opt,name=enabled"`
	// LabName is the lab that execute processing of the data pipeline
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	LabName *string `json:"labName,omitempty" protobuf:"varint,2,opt,name=labName"`
	// If not null, run the data pipeline and create a dataset. else, use the data in the data location
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	DataPipelineName *string `json:"dataPipelineName,omitempty" protobuf:"varint,3,opt,name=datapipelineName"`
	// The data source name for the data in the location. The data source will be used to create a new dataset for this pipeline
	// based on the file in the location.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	DatasourceName *string `json:"datasourceName,omitempty" protobuf:"bytes,4,opt,name=datasourceName"`
	// If Not null, run a docker image is used in order to generate the data.
	// The data must reside in location after the container run
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	DockerImage *string `json:"dockerImage,omitempty" protobuf:"bytes,5,opt,name=dockerImage"`
}

// TrainingStageSpec is the desired state of the training step of the pipeline
type TrainingStageSpec struct {
	// Enabled indicates that the stage is enabled
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" protobuf:"varint,1,opt,name=enabled"`
	// NotebookName specify the notebook to run before training.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	NotebookName *string `json:"notebookName,omitempty" protobuf:"bytes,2,opt,name=notebookName"`
	// LabName is the name of the lab used for training. If empty, the system will use the default lab assigned to the data product
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	LabName *string `json:"labName,omitempty" protobuf:"bytes,3,opt,name=labName"`
	// StudyName is the name of a study template. The actual study will clone the study template and will
	// use the dataset created in the data stage.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Required
	StudyTemplateName *string `json:"studyTemplateName,omitempty" protobuf:"bytes,4,opt,name=studyTemplateName"`
	// Validations defines the machine learning test cases to run against the new trained model.
	// +kubebuilder:validation:Optional
	Validations []ModelValidation `json:"validations,omitempty" protobuf:"bytes,5,rep,name=validations"`
}

//UATStageSpec is the specification of the user acceptance test.
type UATStageSpec struct {
	// Enabled indicates that the stage is enabled
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" protobuf:"varint,1,opt,name=enabled"`
	// The serving site (name space) used for running the uat tests. If the serving site is empty, the system
	// will skip the uat stage
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	ServingSiteName *string `json:"servingSiteName,omitempty" protobuf:"bytes,2,opt,name=servingSiteName"`
	// Validations defines the machine learning test cases to run against the new trained model.
	// +kubebuilder:validation:Optional
	Validations []ModelValidation `json:"validations,omitempty" protobuf:"bytes,3,rep,name=validations"`
	// WorkloadClassName is a reference to the workload class that is used for running the tests in the serving site.
	// +kubebuilder:default:="nano-cpu-250m-mem-256mi"
	// +kubebuilder:validation:Optional
	WorkloadClassName *string `json:"workloadClassName,omitempty" protobuf:"bytes,4,opt,name=workloadClassName"`
}

// CapacityStageSpec is the desired state of the capcity testing.
type CapacityStageSpec struct {
	// Enabled indicates that the stage is enabled
	// +kubebuilder:default:=false
	//+kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" protobuf:"varint,1,opt,name=enabled"`
	// ServingSiteName is the serving site for the testing during the capacity stage
	// If the serving site is empty or null, the system will skip the capacity stage unit tests.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	ServingSiteName *string `json:"servingSiteName,omitempty" protobuf:"bytes,2,opt,name=servingSiteName"`
	// Validations is the specification of tests to run in this stage
	// +kubebuilder:validation:Optional
	Validations []ModelValidation `json:"validations,omitempty" protobuf:"bytes,3,rep,name=validations"`
	// A reference to the workload class that is used for running the prediction
	// +kubebuilder:default:="nano-cpu-250m-mem-256mi"
	// +kubebuilder:validation:Optional
	WorkloadClassName *string `json:"workloadClassName,omitempty" protobuf:"bytes,5,opt,name=workloadClassName"`
}

//DeploymentStageSpec define the testing and relesing the resulting model to production.
type DeploymentStageSpec struct {
	// Enabled indicates that we want to release the model into production
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" protobuf:"varint,1,opt,name=enabled"`
	// ServingSiteName is the serving site for the release, if empty, the system will use the default serving site name
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	ServingSiteName *string `json:"servingSiteName,omitempty" protobuf:"bytes,2,opt,name=servingSiteName"`
	// ManualApproval dentoes if we need manual approval before advancing from deployed to released
	// By default a user is needed to apporve the release to production
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	ManualApproval *bool `json:"manualApproval,omitempty" protobuf:"varint,3,opt,name=manualApproval"`
	// Validations is the specification of tests to run in this stage
	// +kubebuilder:validation:Optional
	Validations []ModelValidation `json:"validations,omitempty" protobuf:"bytes,4,rep,name=validations"`
	// A reference to the workload class that is used for running the test prediction
	// +kubebuilder:default:="nano-cpu-250m-mem-256mi"
	// +kubebuilder:validation:Optional
	WorkloadClassName *string `json:"workloadClassName,omitempty" protobuf:"bytes,5,opt,name=workloadClassName"`
}

type ReleaseStageSpec struct {
	// Enabled indicates that we want to release the model into production
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" protobuf:"varint,1,opt,name=enabled"`
	// ServingSiteName is the serving site for the release, if empty, the system will use the default serving site name
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	ServingSiteName *string `json:"servingSiteName,omitempty" protobuf:"bytes,2,opt,name=servingSiteName"`
	// PredictorName is the release predictor. The predictor will be created if it does not exist.
	// +kubebuilder:default =""
	// +kubebuilder:validation:Optional
	PredictorName *string `json:"predictorName,omitempty" protobuf:"bytes,3,opt,name=predictorName"`
	// Template defines the default model deployment for this model
	// +kubebuilder:validation:Optional
	Template *catalog.ModelDeploymentSpec `json:"template,omitempty" protobuf:"bytes,4,opt,name=template"`
	// ManualApproval dentoes if we need manual approval before advancing from deployed to released
	// By default a user is needed to apporve the release to production
	// +kubebuilder:default:=true
	// +kubebuilder:validation:Optional
	ManualApproval *bool `json:"manualApproval,omitempty" protobuf:"varint,5,opt,name=manualApproval"`
	// Validations is the List of expectation run against the deployed model before moving production traffic to the model
	// +kubebuilder:validation:Optional
	Validations []ModelValidation `json:"validations,omitempty" protobuf:"bytes,6,rep,name=validations"`
	// A reference to the workload class that is used for running the release
	// +kubebuilder:default:="nano-cpu-250m-mem-256mi"
	// +kubebuilder:validation:Optional
	WorkloadClassName *string `json:"workloadClassName,omitempty" protobuf:"bytes,7,opt,name=workloadClassName"`
}

// ModelPipelineStatus define the observed state of the pipeline
type ModelPipelineStatus struct {
	// Last run is the last time a run was created
	//+kubebuilder:validation:Optional
	LastRun *metav1.Time `json:"lastRun,omitempty" protobuf:"bytes,1,opt,name=lastRun"`
	// ObservedGeneration is the Last generation that was acted on
	//+kubebuilder:validation:Optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,2,opt,name=observedGeneration"`
	// Last time the object was updated
	//+kubebuilder:validation:Optional
	LastUpdated *metav1.Time `json:"lastUpdated,omitempty" protobuf:"bytes,3,opt,name=lastUpdated"`
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +kubebuilder:validation:Optional
	Conditions []ModelPipelineCondition `json:"conditions,omitempty" protobuf:"bytes,4,rep,name=conditions"`
}
