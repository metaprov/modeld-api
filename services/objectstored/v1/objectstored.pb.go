// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: github.com/metaprov/modeldapi/services/objectstored/v1/objectstored.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/infra/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IngestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LiveKey string `protobuf:"bytes,1,opt,name=liveKey,proto3" json:"liveKey,omitempty"` // the new live key
}

func (x *IngestResponse) Reset() {
	*x = IngestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestResponse) ProtoMessage() {}

func (x *IngestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestResponse.ProtoReflect.Descriptor instead.
func (*IngestResponse) Descriptor() ([]byte, []int) {
	return file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescGZIP(), []int{0}
}

func (x *IngestResponse) GetLiveKey() string {
	if x != nil {
		return x.LiveKey
	}
	return ""
}

type ObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchivePath string `protobuf:"bytes,1,opt,name=archivePath,proto3" json:"archivePath,omitempty"`
}

func (x *ObjectResponse) Reset() {
	*x = ObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectResponse) ProtoMessage() {}

func (x *ObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectResponse.ProtoReflect.Descriptor instead.
func (*ObjectResponse) Descriptor() ([]byte, []int) {
	return file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectResponse) GetArchivePath() string {
	if x != nil {
		return x.ArchivePath
	}
	return ""
}

// A request to recover archived object from the archive into the depot.
type ObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName     string                      `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	BucketSpec     *v1alpha1.VirtualBucketSpec `protobuf:"bytes,2,opt,name=bucketSpec,proto3" json:"bucketSpec,omitempty"`
	CredNamespace  string                      `protobuf:"bytes,3,opt,name=credNamespace,proto3" json:"credNamespace,omitempty"`
	CredName       string                      `protobuf:"bytes,4,opt,name=credName,proto3" json:"credName,omitempty"`
	ConnectionSpec *v1alpha1.ConnectionSpec    `protobuf:"bytes,5,opt,name=connectionSpec,proto3" json:"connectionSpec,omitempty"`
	Secret         map[string][]byte           `protobuf:"bytes,6,rep,name=secret,proto3" json:"secret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Key            string                      `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"` // the depot location of the object
}

func (x *ObjectRequest) Reset() {
	*x = ObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectRequest) ProtoMessage() {}

func (x *ObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectRequest.ProtoReflect.Descriptor instead.
func (*ObjectRequest) Descriptor() ([]byte, []int) {
	return file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *ObjectRequest) GetBucketSpec() *v1alpha1.VirtualBucketSpec {
	if x != nil {
		return x.BucketSpec
	}
	return nil
}

func (x *ObjectRequest) GetCredNamespace() string {
	if x != nil {
		return x.CredNamespace
	}
	return ""
}

func (x *ObjectRequest) GetCredName() string {
	if x != nil {
		return x.CredName
	}
	return ""
}

func (x *ObjectRequest) GetConnectionSpec() *v1alpha1.ConnectionSpec {
	if x != nil {
		return x.ConnectionSpec
	}
	return nil
}

func (x *ObjectRequest) GetSecret() map[string][]byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *ObjectRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ExistInVirtualBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *ExistInVirtualBucketResponse) Reset() {
	*x = ExistInVirtualBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistInVirtualBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistInVirtualBucketResponse) ProtoMessage() {}

func (x *ExistInVirtualBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistInVirtualBucketResponse.ProtoReflect.Descriptor instead.
func (*ExistInVirtualBucketResponse) Descriptor() ([]byte, []int) {
	return file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescGZIP(), []int{3}
}

func (x *ExistInVirtualBucketResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto protoreflect.FileDescriptor

var file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x76, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0e, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x69, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x69, 0x76, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x32, 0x0a, 0x0e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x82, 0x04, 0x0a, 0x0d, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x68, 0x0a, 0x0a,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x69, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x34, 0x0a, 0x1c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x32, 0xa2, 0x05, 0x0a, 0x12, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a,
	0x06, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x12, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x12, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x45, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0xb5, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x45, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x54, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x76, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x64, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescOnce sync.Once
	file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescData = file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDesc
)

func file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescGZIP() []byte {
	file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescOnce.Do(func() {
		file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescData)
	})
	return file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDescData
}

var file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_goTypes = []interface{}{
	(*IngestResponse)(nil),               // 0: github.com.metaprov.modeldapi.services.objectstored.v1.IngestResponse
	(*ObjectResponse)(nil),               // 1: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectResponse
	(*ObjectRequest)(nil),                // 2: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest
	(*ExistInVirtualBucketResponse)(nil), // 3: github.com.metaprov.modeldapi.services.objectstored.v1.ExistInVirtualBucketResponse
	nil,                                  // 4: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest.SecretEntry
	(*v1alpha1.VirtualBucketSpec)(nil),   // 5: github.com.metaprov.modeldapi.pkg.apis.infra.v1alpha1.VirtualBucketSpec
	(*v1alpha1.ConnectionSpec)(nil),      // 6: github.com.metaprov.modeldapi.pkg.apis.infra.v1alpha1.ConnectionSpec
}
var file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_depIdxs = []int32{
	5, // 0: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest.bucketSpec:type_name -> github.com.metaprov.modeldapi.pkg.apis.infra.v1alpha1.VirtualBucketSpec
	6, // 1: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest.connectionSpec:type_name -> github.com.metaprov.modeldapi.pkg.apis.infra.v1alpha1.ConnectionSpec
	4, // 2: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest.secret:type_name -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest.SecretEntry
	2, // 3: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.Ingest:input_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest
	2, // 4: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.Archive:input_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest
	2, // 5: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.Recover:input_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest
	2, // 6: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.ExistInVirtualBucket:input_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectRequest
	1, // 7: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.Ingest:output_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectResponse
	1, // 8: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.Archive:output_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectResponse
	1, // 9: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.Recover:output_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ObjectResponse
	3, // 10: github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService.ExistInVirtualBucket:output_type -> github.com.metaprov.modeldapi.services.objectstored.v1.ExistInVirtualBucketResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_init() }
func file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_init() {
	if File_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistInVirtualBucketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_goTypes,
		DependencyIndexes: file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_depIdxs,
		MessageInfos:      file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_msgTypes,
	}.Build()
	File_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto = out.File
	file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_rawDesc = nil
	file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_goTypes = nil
	file_github_com_metaprov_modeldapi_services_objectstored_v1_objectstored_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ObjectstoreServiceClient is the client API for ObjectstoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ObjectstoreServiceClient interface {
	// Ingest a new dataset to the store, the store creates a new layouts and set of keys
	// for the new dataset
	Ingest(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error)
	// Delete a dataset from the store. Datasets are not deleted but moved to archive mode
	Archive(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error)
	// Recover an item from the archive into the depot
	Recover(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error)
	// Check if a dataset exist in bucket
	ExistInVirtualBucket(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ExistInVirtualBucketResponse, error)
}

type objectstoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectstoreServiceClient(cc grpc.ClientConnInterface) ObjectstoreServiceClient {
	return &objectstoreServiceClient{cc}
}

func (c *objectstoreServiceClient) Ingest(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/Ingest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectstoreServiceClient) Archive(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/Archive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectstoreServiceClient) Recover(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectstoreServiceClient) ExistInVirtualBucket(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ExistInVirtualBucketResponse, error) {
	out := new(ExistInVirtualBucketResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/ExistInVirtualBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectstoreServiceServer is the server API for ObjectstoreService service.
type ObjectstoreServiceServer interface {
	// Ingest a new dataset to the store, the store creates a new layouts and set of keys
	// for the new dataset
	Ingest(context.Context, *ObjectRequest) (*ObjectResponse, error)
	// Delete a dataset from the store. Datasets are not deleted but moved to archive mode
	Archive(context.Context, *ObjectRequest) (*ObjectResponse, error)
	// Recover an item from the archive into the depot
	Recover(context.Context, *ObjectRequest) (*ObjectResponse, error)
	// Check if a dataset exist in bucket
	ExistInVirtualBucket(context.Context, *ObjectRequest) (*ExistInVirtualBucketResponse, error)
}

// UnimplementedObjectstoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedObjectstoreServiceServer struct {
}

func (*UnimplementedObjectstoreServiceServer) Ingest(context.Context, *ObjectRequest) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}
func (*UnimplementedObjectstoreServiceServer) Archive(context.Context, *ObjectRequest) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Archive not implemented")
}
func (*UnimplementedObjectstoreServiceServer) Recover(context.Context, *ObjectRequest) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (*UnimplementedObjectstoreServiceServer) ExistInVirtualBucket(context.Context, *ObjectRequest) (*ExistInVirtualBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistInVirtualBucket not implemented")
}

func RegisterObjectstoreServiceServer(s *grpc.Server, srv ObjectstoreServiceServer) {
	s.RegisterService(&_ObjectstoreService_serviceDesc, srv)
}

func _ObjectstoreService_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectstoreServiceServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectstoreServiceServer).Ingest(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectstoreService_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectstoreServiceServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/Archive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectstoreServiceServer).Archive(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectstoreService_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectstoreServiceServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectstoreServiceServer).Recover(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectstoreService_ExistInVirtualBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectstoreServiceServer).ExistInVirtualBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService/ExistInVirtualBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectstoreServiceServer).ExistInVirtualBucket(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectstoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeldapi.services.objectstored.v1.ObjectstoreService",
	HandlerType: (*ObjectstoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _ObjectstoreService_Ingest_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _ObjectstoreService_Archive_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _ObjectstoreService_Recover_Handler,
		},
		{
			MethodName: "ExistInVirtualBucket",
			Handler:    _ObjectstoreService_ExistInVirtualBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/metaprov/modeldapi/services/objectstored/v1/objectstored.proto",
}
