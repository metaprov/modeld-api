// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/publisherd/publisherd.proto

package publisherd

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1alpha11 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	v1alpha12 "github.com/metaprov/modeldapi/pkg/apis/infra/v1alpha1"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/training/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PublishNotebookRequest struct {
	NotebookName         string                 `protobuf:"bytes,1,opt,name=notebookName,proto3" json:"notebookName,omitempty"`
	NotebookNamespace    string                 `protobuf:"bytes,2,opt,name=notebookNamespace,proto3" json:"notebookNamespace,omitempty"`
	NotebookSpec         *v1alpha1.NotebookSpec `protobuf:"bytes,3,opt,name=notebookSpec,proto3" json:"notebookSpec,omitempty"`
	NotebookContent      string                 `protobuf:"bytes,4,opt,name=NotebookContent,proto3" json:"NotebookContent,omitempty"`
	Dockerfile           string                 `protobuf:"bytes,5,opt,name=Dockerfile,proto3" json:"Dockerfile,omitempty"`
	Provider             string                 `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
	Secret               map[string][]byte      `protobuf:"bytes,7,rep,name=secret,proto3" json:"secret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PublishNotebookRequest) Reset()         { *m = PublishNotebookRequest{} }
func (m *PublishNotebookRequest) String() string { return proto.CompactTextString(m) }
func (*PublishNotebookRequest) ProtoMessage()    {}
func (*PublishNotebookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a081c30c8c3f3f62, []int{0}
}
func (m *PublishNotebookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishNotebookRequest.Unmarshal(m, b)
}
func (m *PublishNotebookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishNotebookRequest.Marshal(b, m, deterministic)
}
func (m *PublishNotebookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishNotebookRequest.Merge(m, src)
}
func (m *PublishNotebookRequest) XXX_Size() int {
	return xxx_messageInfo_PublishNotebookRequest.Size(m)
}
func (m *PublishNotebookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishNotebookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishNotebookRequest proto.InternalMessageInfo

func (m *PublishNotebookRequest) GetNotebookName() string {
	if m != nil {
		return m.NotebookName
	}
	return ""
}

func (m *PublishNotebookRequest) GetNotebookNamespace() string {
	if m != nil {
		return m.NotebookNamespace
	}
	return ""
}

func (m *PublishNotebookRequest) GetNotebookSpec() *v1alpha1.NotebookSpec {
	if m != nil {
		return m.NotebookSpec
	}
	return nil
}

func (m *PublishNotebookRequest) GetNotebookContent() string {
	if m != nil {
		return m.NotebookContent
	}
	return ""
}

func (m *PublishNotebookRequest) GetDockerfile() string {
	if m != nil {
		return m.Dockerfile
	}
	return ""
}

func (m *PublishNotebookRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *PublishNotebookRequest) GetSecret() map[string][]byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

type PublishNotebookResponse struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishNotebookResponse) Reset()         { *m = PublishNotebookResponse{} }
func (m *PublishNotebookResponse) String() string { return proto.CompactTextString(m) }
func (*PublishNotebookResponse) ProtoMessage()    {}
func (*PublishNotebookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a081c30c8c3f3f62, []int{1}
}
func (m *PublishNotebookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishNotebookResponse.Unmarshal(m, b)
}
func (m *PublishNotebookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishNotebookResponse.Marshal(b, m, deterministic)
}
func (m *PublishNotebookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishNotebookResponse.Merge(m, src)
}
func (m *PublishNotebookResponse) XXX_Size() int {
	return xxx_messageInfo_PublishNotebookResponse.Size(m)
}
func (m *PublishNotebookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishNotebookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishNotebookResponse proto.InternalMessageInfo

func (m *PublishNotebookResponse) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

type PublishModelRequest struct {
	DataproductName             string                            `protobuf:"bytes,1,opt,name=dataproductName,proto3" json:"dataproductName,omitempty"`
	DataproductNamespace        string                            `protobuf:"bytes,2,opt,name=dataproductNamespace,proto3" json:"dataproductNamespace,omitempty"`
	DataproductSpec             *v1alpha11.DataProductSpec        `protobuf:"bytes,3,opt,name=dataproductSpec,proto3" json:"dataproductSpec,omitempty"`
	DataproductversionName      string                            `protobuf:"bytes,4,opt,name=dataproductversionName,proto3" json:"dataproductversionName,omitempty"`
	DataproductversionNamespace string                            `protobuf:"bytes,5,opt,name=dataproductversionNamespace,proto3" json:"dataproductversionNamespace,omitempty"`
	DataproductversionSpec      *v1alpha11.DataProductVersionSpec `protobuf:"bytes,6,opt,name=dataproductversionSpec,proto3" json:"dataproductversionSpec,omitempty"`
	ModelName                   string                            `protobuf:"bytes,7,opt,name=modelName,proto3" json:"modelName,omitempty"`
	ModelNamespace              string                            `protobuf:"bytes,8,opt,name=modelNamespace,proto3" json:"modelNamespace,omitempty"`
	ModelSpec                   *v1alpha1.ModelSpec               `protobuf:"bytes,9,opt,name=modelSpec,proto3" json:"modelSpec,omitempty"`
	StudyName                   string                            `protobuf:"bytes,13,opt,name=studyName,proto3" json:"studyName,omitempty"`
	StudyNamespace              string                            `protobuf:"bytes,14,opt,name=studyNamespace,proto3" json:"studyNamespace,omitempty"`
	StudySpec                   *v1alpha1.StudySpec               `protobuf:"bytes,15,opt,name=studySpec,proto3" json:"studySpec,omitempty"`
	DatasourceName              string                            `protobuf:"bytes,16,opt,name=datasourceName,proto3" json:"datasourceName,omitempty"`
	DatasourceNamespace         string                            `protobuf:"bytes,17,opt,name=datasourceNamespace,proto3" json:"datasourceNamespace,omitempty"`
	DatasourceSpec              *v1alpha11.DataSourceSpec         `protobuf:"bytes,18,opt,name=datasourceSpec,proto3" json:"datasourceSpec,omitempty"`
	DatasetName                 string                            `protobuf:"bytes,19,opt,name=datasetName,proto3" json:"datasetName,omitempty"`
	DatasetNamespace            string                            `protobuf:"bytes,20,opt,name=datasetNamespace,proto3" json:"datasetNamespace,omitempty"`
	DatasetSpec                 *v1alpha11.DatasetSpec            `protobuf:"bytes,21,opt,name=datasetSpec,proto3" json:"datasetSpec,omitempty"`
	Provider                    string                            `protobuf:"bytes,22,opt,name=provider,proto3" json:"provider,omitempty"`
	Imagename                   string                            `protobuf:"bytes,23,opt,name=imagename,proto3" json:"imagename,omitempty"`
	Push                        bool                              `protobuf:"varint,24,opt,name=push,proto3" json:"push,omitempty"`
	BucketName                  string                            `protobuf:"bytes,25,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	BucketNamespace             string                            `protobuf:"bytes,26,opt,name=bucketNamespace,proto3" json:"bucketNamespace,omitempty"`
	BucketSpec                  *v1alpha12.VirtualBucketSpec      `protobuf:"bytes,27,opt,name=bucketSpec,proto3" json:"bucketSpec,omitempty"`
	CloudConnectionName         string                            `protobuf:"bytes,28,opt,name=cloudConnectionName,proto3" json:"cloudConnectionName,omitempty"`
	CloudsConnectionNamespace   string                            `protobuf:"bytes,29,opt,name=cloudsConnectionNamespace,proto3" json:"cloudsConnectionNamespace,omitempty"`
	CloudConnectionSpec         *v1alpha12.ConnectionSpec         `protobuf:"bytes,30,opt,name=cloudConnectionSpec,proto3" json:"cloudConnectionSpec,omitempty"`
	CloudSecret                 map[string][]byte                 `protobuf:"bytes,31,rep,name=cloudSecret,proto3" json:"cloudSecret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DockerConnectionName        string                            `protobuf:"bytes,32,opt,name=dockerConnectionName,proto3" json:"dockerConnectionName,omitempty"`
	DockerConnectionNamespace   string                            `protobuf:"bytes,33,opt,name=dockerConnectionNamespace,proto3" json:"dockerConnectionNamespace,omitempty"`
	DockerConnectionSpec        *v1alpha12.ConnectionSpec         `protobuf:"bytes,34,opt,name=dockerConnectionSpec,proto3" json:"dockerConnectionSpec,omitempty"`
	DockerRegistrySecret        map[string][]byte                 `protobuf:"bytes,35,rep,name=dockerRegistrySecret,proto3" json:"dockerRegistrySecret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ModelWeightsUri             string                            `protobuf:"bytes,36,opt,name=modelWeightsUri,proto3" json:"modelWeightsUri,omitempty"`
	PreprocessorWeightsUri      string                            `protobuf:"bytes,37,opt,name=preprocessorWeightsUri,proto3" json:"preprocessorWeightsUri,omitempty"`
	LabelEncoderUri             string                            `protobuf:"bytes,38,opt,name=labelEncoderUri,proto3" json:"labelEncoderUri,omitempty"`
	Kaniko                      bool                              `protobuf:"varint,39,opt,name=kaniko,proto3" json:"kaniko,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                          `json:"-"`
	XXX_unrecognized            []byte                            `json:"-"`
	XXX_sizecache               int32                             `json:"-"`
}

func (m *PublishModelRequest) Reset()         { *m = PublishModelRequest{} }
func (m *PublishModelRequest) String() string { return proto.CompactTextString(m) }
func (*PublishModelRequest) ProtoMessage()    {}
func (*PublishModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a081c30c8c3f3f62, []int{2}
}
func (m *PublishModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishModelRequest.Unmarshal(m, b)
}
func (m *PublishModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishModelRequest.Marshal(b, m, deterministic)
}
func (m *PublishModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishModelRequest.Merge(m, src)
}
func (m *PublishModelRequest) XXX_Size() int {
	return xxx_messageInfo_PublishModelRequest.Size(m)
}
func (m *PublishModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishModelRequest proto.InternalMessageInfo

func (m *PublishModelRequest) GetDataproductName() string {
	if m != nil {
		return m.DataproductName
	}
	return ""
}

func (m *PublishModelRequest) GetDataproductNamespace() string {
	if m != nil {
		return m.DataproductNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetDataproductSpec() *v1alpha11.DataProductSpec {
	if m != nil {
		return m.DataproductSpec
	}
	return nil
}

func (m *PublishModelRequest) GetDataproductversionName() string {
	if m != nil {
		return m.DataproductversionName
	}
	return ""
}

func (m *PublishModelRequest) GetDataproductversionNamespace() string {
	if m != nil {
		return m.DataproductversionNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetDataproductversionSpec() *v1alpha11.DataProductVersionSpec {
	if m != nil {
		return m.DataproductversionSpec
	}
	return nil
}

func (m *PublishModelRequest) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *PublishModelRequest) GetModelNamespace() string {
	if m != nil {
		return m.ModelNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetModelSpec() *v1alpha1.ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PublishModelRequest) GetStudyName() string {
	if m != nil {
		return m.StudyName
	}
	return ""
}

func (m *PublishModelRequest) GetStudyNamespace() string {
	if m != nil {
		return m.StudyNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetStudySpec() *v1alpha1.StudySpec {
	if m != nil {
		return m.StudySpec
	}
	return nil
}

func (m *PublishModelRequest) GetDatasourceName() string {
	if m != nil {
		return m.DatasourceName
	}
	return ""
}

func (m *PublishModelRequest) GetDatasourceNamespace() string {
	if m != nil {
		return m.DatasourceNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetDatasourceSpec() *v1alpha11.DataSourceSpec {
	if m != nil {
		return m.DatasourceSpec
	}
	return nil
}

func (m *PublishModelRequest) GetDatasetName() string {
	if m != nil {
		return m.DatasetName
	}
	return ""
}

func (m *PublishModelRequest) GetDatasetNamespace() string {
	if m != nil {
		return m.DatasetNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetDatasetSpec() *v1alpha11.DatasetSpec {
	if m != nil {
		return m.DatasetSpec
	}
	return nil
}

func (m *PublishModelRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *PublishModelRequest) GetImagename() string {
	if m != nil {
		return m.Imagename
	}
	return ""
}

func (m *PublishModelRequest) GetPush() bool {
	if m != nil {
		return m.Push
	}
	return false
}

func (m *PublishModelRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *PublishModelRequest) GetBucketNamespace() string {
	if m != nil {
		return m.BucketNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetBucketSpec() *v1alpha12.VirtualBucketSpec {
	if m != nil {
		return m.BucketSpec
	}
	return nil
}

func (m *PublishModelRequest) GetCloudConnectionName() string {
	if m != nil {
		return m.CloudConnectionName
	}
	return ""
}

func (m *PublishModelRequest) GetCloudsConnectionNamespace() string {
	if m != nil {
		return m.CloudsConnectionNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetCloudConnectionSpec() *v1alpha12.ConnectionSpec {
	if m != nil {
		return m.CloudConnectionSpec
	}
	return nil
}

func (m *PublishModelRequest) GetCloudSecret() map[string][]byte {
	if m != nil {
		return m.CloudSecret
	}
	return nil
}

func (m *PublishModelRequest) GetDockerConnectionName() string {
	if m != nil {
		return m.DockerConnectionName
	}
	return ""
}

func (m *PublishModelRequest) GetDockerConnectionNamespace() string {
	if m != nil {
		return m.DockerConnectionNamespace
	}
	return ""
}

func (m *PublishModelRequest) GetDockerConnectionSpec() *v1alpha12.ConnectionSpec {
	if m != nil {
		return m.DockerConnectionSpec
	}
	return nil
}

func (m *PublishModelRequest) GetDockerRegistrySecret() map[string][]byte {
	if m != nil {
		return m.DockerRegistrySecret
	}
	return nil
}

func (m *PublishModelRequest) GetModelWeightsUri() string {
	if m != nil {
		return m.ModelWeightsUri
	}
	return ""
}

func (m *PublishModelRequest) GetPreprocessorWeightsUri() string {
	if m != nil {
		return m.PreprocessorWeightsUri
	}
	return ""
}

func (m *PublishModelRequest) GetLabelEncoderUri() string {
	if m != nil {
		return m.LabelEncoderUri
	}
	return ""
}

func (m *PublishModelRequest) GetKaniko() bool {
	if m != nil {
		return m.Kaniko
	}
	return false
}

type PublishModelResponse struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishModelResponse) Reset()         { *m = PublishModelResponse{} }
func (m *PublishModelResponse) String() string { return proto.CompactTextString(m) }
func (*PublishModelResponse) ProtoMessage()    {}
func (*PublishModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a081c30c8c3f3f62, []int{3}
}
func (m *PublishModelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishModelResponse.Unmarshal(m, b)
}
func (m *PublishModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishModelResponse.Marshal(b, m, deterministic)
}
func (m *PublishModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishModelResponse.Merge(m, src)
}
func (m *PublishModelResponse) XXX_Size() int {
	return xxx_messageInfo_PublishModelResponse.Size(m)
}
func (m *PublishModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishModelResponse proto.InternalMessageInfo

func (m *PublishModelResponse) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

type ShutdownRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownRequest) Reset()         { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()    {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a081c30c8c3f3f62, []int{4}
}
func (m *ShutdownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownRequest.Unmarshal(m, b)
}
func (m *ShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownRequest.Marshal(b, m, deterministic)
}
func (m *ShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownRequest.Merge(m, src)
}
func (m *ShutdownRequest) XXX_Size() int {
	return xxx_messageInfo_ShutdownRequest.Size(m)
}
func (m *ShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownRequest proto.InternalMessageInfo

type ShutdownResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownResponse) Reset()         { *m = ShutdownResponse{} }
func (m *ShutdownResponse) String() string { return proto.CompactTextString(m) }
func (*ShutdownResponse) ProtoMessage()    {}
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a081c30c8c3f3f62, []int{5}
}
func (m *ShutdownResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownResponse.Unmarshal(m, b)
}
func (m *ShutdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownResponse.Marshal(b, m, deterministic)
}
func (m *ShutdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownResponse.Merge(m, src)
}
func (m *ShutdownResponse) XXX_Size() int {
	return xxx_messageInfo_ShutdownResponse.Size(m)
}
func (m *ShutdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PublishNotebookRequest)(nil), "github.com.metaprov.modeld.services.publisherd.PublishNotebookRequest")
	proto.RegisterMapType((map[string][]byte)(nil), "github.com.metaprov.modeld.services.publisherd.PublishNotebookRequest.SecretEntry")
	proto.RegisterType((*PublishNotebookResponse)(nil), "github.com.metaprov.modeld.services.publisherd.PublishNotebookResponse")
	proto.RegisterType((*PublishModelRequest)(nil), "github.com.metaprov.modeld.services.publisherd.PublishModelRequest")
	proto.RegisterMapType((map[string][]byte)(nil), "github.com.metaprov.modeld.services.publisherd.PublishModelRequest.CloudSecretEntry")
	proto.RegisterMapType((map[string][]byte)(nil), "github.com.metaprov.modeld.services.publisherd.PublishModelRequest.DockerRegistrySecretEntry")
	proto.RegisterType((*PublishModelResponse)(nil), "github.com.metaprov.modeld.services.publisherd.PublishModelResponse")
	proto.RegisterType((*ShutdownRequest)(nil), "github.com.metaprov.modeld.services.publisherd.ShutdownRequest")
	proto.RegisterType((*ShutdownResponse)(nil), "github.com.metaprov.modeld.services.publisherd.ShutdownResponse")
}

func init() {
	proto.RegisterFile("services/publisherd/publisherd.proto", fileDescriptor_a081c30c8c3f3f62)
}

var fileDescriptor_a081c30c8c3f3f62 = []byte{
	// 1090 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0x4d, 0x9b, 0x26, 0x93, 0xd2, 0x24, 0x27, 0x21, 0xdd, 0x6c, 0x43, 0x31, 0xa6, 0x14,
	0x0b, 0xa1, 0x35, 0x49, 0x11, 0x7f, 0x42, 0x50, 0x6a, 0xa7, 0x2d, 0x12, 0x8d, 0xa2, 0x35, 0x14,
	0x09, 0x89, 0x8b, 0xf1, 0xee, 0xd4, 0xde, 0x7a, 0xbd, 0xb3, 0xcc, 0xce, 0xba, 0xca, 0x3d, 0xf7,
	0x88, 0x17, 0xe0, 0x8a, 0x37, 0xe1, 0x51, 0xb8, 0xe6, 0x1d, 0xaa, 0x39, 0xb3, 0xff, 0x59, 0x47,
	0xb6, 0x93, 0x3b, 0xcf, 0x37, 0x33, 0xdf, 0xf9, 0xbe, 0x33, 0x67, 0xe7, 0x8c, 0xc9, 0xfd, 0x88,
	0x89, 0xa9, 0xe7, 0xb0, 0xa8, 0x13, 0xc6, 0x03, 0xdf, 0x8b, 0x46, 0x4c, 0xb8, 0x85, 0x9f, 0x56,
	0x28, 0xb8, 0xe4, 0x60, 0x0d, 0x3d, 0x39, 0x8a, 0x07, 0x96, 0xc3, 0x27, 0xd6, 0x84, 0x49, 0x1a,
	0x0a, 0x3e, 0xb5, 0x26, 0xdc, 0x65, 0xbe, 0x6b, 0xa5, 0x04, 0x56, 0xbe, 0xcb, 0x7c, 0x96, 0xaf,
	0xef, 0xa4, 0xeb, 0x3b, 0x7a, 0x3d, 0x0d, 0xbd, 0x4e, 0x38, 0x1e, 0x76, 0x68, 0xe8, 0x45, 0x1d,
	0x29, 0xa8, 0x17, 0x78, 0xc1, 0xb0, 0x33, 0x3d, 0xa4, 0x7e, 0x38, 0xa2, 0x87, 0x9d, 0x21, 0x0b,
	0x98, 0xa0, 0x92, 0x25, 0x91, 0xcd, 0xde, 0x9c, 0x4c, 0x2e, 0x95, 0x74, 0x36, 0xcb, 0xf1, 0x9c,
	0x2c, 0x5e, 0xf0, 0x52, 0xcc, 0xa6, 0x69, 0xfd, 0xb7, 0x42, 0xf6, 0x4e, 0xb5, 0xcb, 0x13, 0x2e,
	0xd9, 0x80, 0xf3, 0xb1, 0xcd, 0x7e, 0x8f, 0x59, 0x24, 0xa1, 0x45, 0x6e, 0x05, 0x09, 0x74, 0x42,
	0x27, 0xcc, 0x68, 0x34, 0x1b, 0xed, 0x75, 0xbb, 0x84, 0xc1, 0x27, 0x64, 0xbb, 0x38, 0x8e, 0x42,
	0xea, 0x30, 0xe3, 0x1a, 0x2e, 0x3c, 0x3f, 0x01, 0xaf, 0x72, 0xc6, 0x7e, 0xc8, 0x1c, 0x63, 0xa5,
	0xd9, 0x68, 0x6f, 0x1c, 0x3d, 0xb9, 0xe0, 0x28, 0x68, 0xe8, 0x59, 0xe1, 0x78, 0x68, 0x29, 0x2b,
	0x56, 0x9a, 0x5a, 0x2b, 0x75, 0x63, 0x9d, 0x14, 0xd8, 0xec, 0x12, 0x37, 0xb4, 0xc9, 0x66, 0x3a,
	0xdb, 0xe5, 0x81, 0x64, 0x81, 0x34, 0xae, 0xa3, 0xae, 0x2a, 0x0c, 0xf7, 0x08, 0xe9, 0x71, 0x67,
	0xcc, 0xc4, 0x4b, 0xcf, 0x67, 0xc6, 0x0d, 0x5c, 0x54, 0x40, 0xc0, 0x24, 0x6b, 0x4a, 0x91, 0xe7,
	0x32, 0x61, 0xac, 0xe2, 0x6c, 0x36, 0x86, 0x57, 0x64, 0x35, 0x62, 0x8e, 0x60, 0xd2, 0xb8, 0xd9,
	0x5c, 0x69, 0x6f, 0x1c, 0xd9, 0x0b, 0x96, 0x95, 0x55, 0x9f, 0x7b, 0xab, 0x8f, 0xa4, 0xc7, 0x81,
	0x14, 0x67, 0x76, 0x12, 0xc1, 0xfc, 0x8a, 0x6c, 0x14, 0x60, 0xd8, 0x22, 0x2b, 0x63, 0x76, 0x96,
	0x9c, 0x8a, 0xfa, 0x09, 0xbb, 0xe4, 0xc6, 0x94, 0xfa, 0xb1, 0x3e, 0x80, 0x5b, 0xb6, 0x1e, 0x7c,
	0x7d, 0xed, 0xcb, 0x46, 0xeb, 0x0b, 0x72, 0xe7, 0x5c, 0xa0, 0x28, 0xe4, 0x41, 0xc4, 0xe0, 0x80,
	0xac, 0xff, 0x30, 0xa1, 0x43, 0x56, 0x38, 0xe2, 0x1c, 0x68, 0xfd, 0xbb, 0x43, 0x76, 0x92, 0x9d,
	0xcf, 0x95, 0x89, 0xb4, 0x36, 0xda, 0x64, 0x53, 0x95, 0x67, 0x28, 0xb8, 0x1b, 0x3b, 0xb2, 0xb0,
	0xb7, 0x0a, 0xc3, 0x11, 0xd9, 0xad, 0x40, 0xc5, 0x22, 0xa9, 0x9d, 0x03, 0x5e, 0x62, 0x2f, 0x94,
	0xca, 0xf1, 0xbc, 0xa5, 0xa2, 0xb6, 0xe7, 0x65, 0xd2, 0xa3, 0x92, 0x9e, 0xe6, 0x64, 0x76, 0x95,
	0x1d, 0x3e, 0x27, 0x7b, 0x05, 0x68, 0xca, 0x44, 0xe4, 0xf1, 0x00, 0x5d, 0xe9, 0x9a, 0x99, 0x31,
	0x0b, 0x8f, 0xc8, 0xdd, 0xfa, 0x19, 0xed, 0x51, 0xd7, 0xd2, 0x45, 0x4b, 0xe0, 0x8f, 0x46, 0x5d,
	0x68, 0xb4, 0xbc, 0x8a, 0x96, 0x7f, 0xbc, 0xb4, 0xe5, 0x17, 0x39, 0xa7, 0x3d, 0x23, 0x96, 0xaa,
	0x02, 0xa4, 0x44, 0xcf, 0x37, 0x75, 0x15, 0x64, 0x00, 0x3c, 0x20, 0xb7, 0xb3, 0x81, 0x76, 0xb6,
	0x86, 0x4b, 0x2a, 0x28, 0xd0, 0x84, 0x05, 0xe5, 0xaf, 0xa3, 0xfc, 0xee, 0xf2, 0x1f, 0xf7, 0xf3,
	0x94, 0xca, 0xce, 0x59, 0x95, 0xd0, 0x48, 0xc6, 0xee, 0x19, 0x0a, 0x7d, 0x5b, 0x0b, 0xcd, 0x00,
	0x25, 0x34, 0x1b, 0x68, 0xa1, 0xb7, 0xb5, 0xd0, 0x32, 0xaa, 0x84, 0x22, 0x82, 0x42, 0x37, 0x2f,
	0x2b, 0xb4, 0x9f, 0x52, 0xd9, 0x39, 0xab, 0x92, 0xa2, 0x72, 0x1d, 0xf1, 0x58, 0x38, 0xfa, 0xe3,
	0xda, 0xd2, 0x52, 0xca, 0x28, 0x7c, 0x4a, 0x76, 0xca, 0x88, 0xd6, 0xbd, 0x8d, 0x8b, 0xeb, 0xa6,
	0xc0, 0x2f, 0x32, 0xa3, 0x03, 0x40, 0x07, 0xbd, 0xe5, 0x2b, 0xa5, 0x9f, 0x71, 0xd9, 0x15, 0x6e,
	0x68, 0x92, 0x0d, 0x44, 0x98, 0xfe, 0xca, 0x77, 0x50, 0x57, 0x11, 0x82, 0x8f, 0xc9, 0x56, 0x61,
	0xa8, 0xe5, 0xef, 0xe2, 0xb2, 0x73, 0x38, 0x38, 0x19, 0x1b, 0x0a, 0x7f, 0x07, 0x85, 0x7f, 0xbf,
	0xbc, 0xf0, 0x84, 0xc8, 0x2e, 0xb2, 0x96, 0x2e, 0xec, 0xbd, 0xca, 0x85, 0x7d, 0x40, 0xd6, 0x3d,
	0x75, 0xbb, 0x05, 0xca, 0xcc, 0x1d, 0x5d, 0x3f, 0x19, 0x00, 0x40, 0xae, 0x87, 0x71, 0x34, 0x32,
	0x8c, 0x66, 0xa3, 0xbd, 0x66, 0xe3, 0x6f, 0xd5, 0x1e, 0x06, 0xb1, 0x33, 0x4e, 0xfc, 0xef, 0xeb,
	0xf6, 0x90, 0x23, 0xea, 0x2a, 0xcc, 0x47, 0xda, 0xbd, 0xa9, 0xaf, 0xc2, 0x0a, 0x0c, 0xa3, 0x94,
	0x09, 0xbd, 0xdf, 0x45, 0xef, 0xcf, 0xe6, 0xf5, 0x8e, 0x7d, 0x3c, 0x37, 0xff, 0xc2, 0x13, 0x32,
	0xa6, 0xfe, 0xe3, 0x8c, 0xcf, 0x2e, 0x70, 0xab, 0xa2, 0x72, 0x7c, 0x1e, 0xbb, 0x5d, 0x1e, 0x04,
	0xcc, 0x91, 0xe9, 0x65, 0x76, 0xa0, 0x8b, 0xaa, 0x66, 0x0a, 0xbe, 0x21, 0xfb, 0x08, 0x47, 0x65,
	0x5c, 0xfb, 0x79, 0x17, 0xf7, 0xcd, 0x5e, 0x00, 0xaf, 0xcf, 0xc5, 0x43, 0x8b, 0xf7, 0x16, 0xbb,
	0xb4, 0x2b, 0x16, 0xcb, 0x64, 0x76, 0x5d, 0x04, 0x98, 0x92, 0x0d, 0x84, 0x75, 0x63, 0x34, 0xde,
	0xc3, 0x26, 0xfc, 0xd3, 0x92, 0x4d, 0xb8, 0xd8, 0xe1, 0xac, 0x6e, 0x4e, 0xab, 0xdb, 0x70, 0x31,
	0x10, 0x76, 0x35, 0x7c, 0x21, 0x54, 0x32, 0xdc, 0x4c, 0xba, 0x5a, 0xcd, 0x9c, 0x4a, 0x71, 0x1d,
	0xae, 0x53, 0xfc, 0xbe, 0x4e, 0xf1, 0xcc, 0x05, 0x70, 0x76, 0x3e, 0x22, 0xe6, 0xb8, 0x75, 0x95,
	0x39, 0xae, 0x0d, 0x01, 0x7f, 0x35, 0xd2, 0xd8, 0x36, 0x1b, 0x7a, 0x91, 0x14, 0x67, 0x49, 0xba,
	0x3f, 0xc0, 0x74, 0xff, 0x76, 0x15, 0xe9, 0xee, 0xd5, 0xf0, 0xeb, 0xbc, 0xd7, 0x86, 0x56, 0x5f,
	0x1d, 0x46, 0xf8, 0x85, 0x79, 0xc3, 0x91, 0x8c, 0x7e, 0x16, 0x9e, 0x71, 0x5f, 0x7f, 0x75, 0x15,
	0x58, 0xf5, 0xf6, 0x50, 0xb0, 0x50, 0x70, 0x87, 0x45, 0x11, 0x17, 0x85, 0x0d, 0x1f, 0xea, 0xde,
	0x5e, 0x3f, 0xab, 0x22, 0xf8, 0x74, 0xc0, 0xfc, 0xe3, 0xc0, 0xe1, 0x2e, 0x13, 0x6a, 0xc3, 0x03,
	0x1d, 0xa1, 0x02, 0xc3, 0x1e, 0x59, 0x1d, 0xd3, 0xc0, 0x1b, 0x73, 0xe3, 0x23, 0xbc, 0x37, 0x92,
	0x91, 0xf9, 0x2d, 0xd9, 0xaa, 0x56, 0xd1, 0x22, 0xaf, 0x36, 0xf3, 0x29, 0xd9, 0x9f, 0x99, 0x96,
	0x85, 0x9e, 0x7f, 0x9f, 0x91, 0xdd, 0x72, 0xce, 0xe7, 0x7a, 0xfb, 0x6d, 0x93, 0xcd, 0xfe, 0x28,
	0x96, 0x2e, 0x7f, 0x1d, 0x24, 0xa7, 0xd4, 0x02, 0xb2, 0x95, 0x43, 0x9a, 0xe4, 0xe8, 0xff, 0x15,
	0xb2, 0x7d, 0x9a, 0x1d, 0x6e, 0x5f, 0x9f, 0x37, 0xfc, 0xdd, 0x20, 0x9b, 0x09, 0xda, 0xe5, 0x93,
	0x81, 0x17, 0x30, 0x17, 0xba, 0x57, 0x50, 0x28, 0x66, 0xef, 0x72, 0x24, 0x5a, 0x74, 0xeb, 0x2d,
	0xf8, 0x27, 0x17, 0x98, 0xbe, 0x89, 0xe1, 0xc9, 0xd5, 0xbc, 0xde, 0xcd, 0xa7, 0x97, 0xe6, 0xc9,
	0x64, 0xfe, 0xd9, 0x20, 0x6b, 0x69, 0xca, 0xe1, 0xbb, 0x45, 0x79, 0x2b, 0xe7, 0x67, 0x3e, 0x5a,
	0x9e, 0x20, 0x55, 0xf4, 0xf8, 0xe1, 0xaf, 0x87, 0x17, 0xff, 0xf5, 0xac, 0xf9, 0xfb, 0x3d, 0x58,
	0xc5, 0x7f, 0x9b, 0x0f, 0xdf, 0x04, 0x00, 0x00, 0xff, 0xff, 0x60, 0xc9, 0xb0, 0x23, 0x9c, 0x0f,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PublisherdServiceClient is the client API for PublisherdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublisherdServiceClient interface {
	PublishCombined(ctx context.Context, in *PublishModelRequest, opts ...grpc.CallOption) (*PublishModelResponse, error)
	// Publish a notebook.
	PublishNotebook(ctx context.Context, in *PublishNotebookRequest, opts ...grpc.CallOption) (*PublishNotebookResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
}

type publisherdServiceClient struct {
	cc *grpc.ClientConn
}

func NewPublisherdServiceClient(cc *grpc.ClientConn) PublisherdServiceClient {
	return &publisherdServiceClient{cc}
}

func (c *publisherdServiceClient) PublishCombined(ctx context.Context, in *PublishModelRequest, opts ...grpc.CallOption) (*PublishModelResponse, error) {
	out := new(PublishModelResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.publisherd.PublisherdService/PublishCombined", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherdServiceClient) PublishNotebook(ctx context.Context, in *PublishNotebookRequest, opts ...grpc.CallOption) (*PublishNotebookResponse, error) {
	out := new(PublishNotebookResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.publisherd.PublisherdService/PublishNotebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherdServiceClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.publisherd.PublisherdService/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublisherdServiceServer is the server API for PublisherdService service.
type PublisherdServiceServer interface {
	PublishCombined(context.Context, *PublishModelRequest) (*PublishModelResponse, error)
	// Publish a notebook.
	PublishNotebook(context.Context, *PublishNotebookRequest) (*PublishNotebookResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
}

// UnimplementedPublisherdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPublisherdServiceServer struct {
}

func (*UnimplementedPublisherdServiceServer) PublishCombined(ctx context.Context, req *PublishModelRequest) (*PublishModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishCombined not implemented")
}
func (*UnimplementedPublisherdServiceServer) PublishNotebook(ctx context.Context, req *PublishNotebookRequest) (*PublishNotebookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishNotebook not implemented")
}
func (*UnimplementedPublisherdServiceServer) Shutdown(ctx context.Context, req *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}

func RegisterPublisherdServiceServer(s *grpc.Server, srv PublisherdServiceServer) {
	s.RegisterService(&_PublisherdService_serviceDesc, srv)
}

func _PublisherdService_PublishCombined_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherdServiceServer).PublishCombined(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.publisherd.PublisherdService/PublishCombined",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherdServiceServer).PublishCombined(ctx, req.(*PublishModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherdService_PublishNotebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishNotebookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherdServiceServer).PublishNotebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.publisherd.PublisherdService/PublishNotebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherdServiceServer).PublishNotebook(ctx, req.(*PublishNotebookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublisherdService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherdServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.publisherd.PublisherdService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherdServiceServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PublisherdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.publisherd.PublisherdService",
	HandlerType: (*PublisherdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishCombined",
			Handler:    _PublisherdService_PublishCombined_Handler,
		},
		{
			MethodName: "PublishNotebook",
			Handler:    _PublisherdService_PublishNotebook_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _PublisherdService_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/publisherd/publisherd.proto",
}