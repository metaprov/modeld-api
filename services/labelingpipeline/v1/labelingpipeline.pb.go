// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/labelingpipeline/v1/labelingpipeline.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/data/v1alpha1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListLabelingPipelineRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListLabelingPipelineRequest) Reset()         { *m = ListLabelingPipelineRequest{} }
func (m *ListLabelingPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*ListLabelingPipelineRequest) ProtoMessage()    {}
func (*ListLabelingPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{0}
}
func (m *ListLabelingPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabelingPipelineRequest.Unmarshal(m, b)
}
func (m *ListLabelingPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabelingPipelineRequest.Marshal(b, m, deterministic)
}
func (m *ListLabelingPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabelingPipelineRequest.Merge(m, src)
}
func (m *ListLabelingPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_ListLabelingPipelineRequest.Size(m)
}
func (m *ListLabelingPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabelingPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabelingPipelineRequest proto.InternalMessageInfo

func (m *ListLabelingPipelineRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListLabelingPipelineRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListLabelingPipelineResponse struct {
	Items                *v1alpha1.LabelingPipelineList `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ListLabelingPipelineResponse) Reset()         { *m = ListLabelingPipelineResponse{} }
func (m *ListLabelingPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*ListLabelingPipelineResponse) ProtoMessage()    {}
func (*ListLabelingPipelineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{1}
}
func (m *ListLabelingPipelineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabelingPipelineResponse.Unmarshal(m, b)
}
func (m *ListLabelingPipelineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabelingPipelineResponse.Marshal(b, m, deterministic)
}
func (m *ListLabelingPipelineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabelingPipelineResponse.Merge(m, src)
}
func (m *ListLabelingPipelineResponse) XXX_Size() int {
	return xxx_messageInfo_ListLabelingPipelineResponse.Size(m)
}
func (m *ListLabelingPipelineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabelingPipelineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabelingPipelineResponse proto.InternalMessageInfo

func (m *ListLabelingPipelineResponse) GetItems() *v1alpha1.LabelingPipelineList {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateLabelingPipelineRequest struct {
	Namespace            string                         `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string              `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.LabelingPipelineSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CreateLabelingPipelineRequest) Reset()         { *m = CreateLabelingPipelineRequest{} }
func (m *CreateLabelingPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLabelingPipelineRequest) ProtoMessage()    {}
func (*CreateLabelingPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{2}
}
func (m *CreateLabelingPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLabelingPipelineRequest.Unmarshal(m, b)
}
func (m *CreateLabelingPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLabelingPipelineRequest.Marshal(b, m, deterministic)
}
func (m *CreateLabelingPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLabelingPipelineRequest.Merge(m, src)
}
func (m *CreateLabelingPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLabelingPipelineRequest.Size(m)
}
func (m *CreateLabelingPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLabelingPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLabelingPipelineRequest proto.InternalMessageInfo

func (m *CreateLabelingPipelineRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateLabelingPipelineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLabelingPipelineRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateLabelingPipelineRequest) GetSpec() *v1alpha1.LabelingPipelineSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type CreateLabelingPipelineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLabelingPipelineResponse) Reset()         { *m = CreateLabelingPipelineResponse{} }
func (m *CreateLabelingPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLabelingPipelineResponse) ProtoMessage()    {}
func (*CreateLabelingPipelineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{3}
}
func (m *CreateLabelingPipelineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLabelingPipelineResponse.Unmarshal(m, b)
}
func (m *CreateLabelingPipelineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLabelingPipelineResponse.Marshal(b, m, deterministic)
}
func (m *CreateLabelingPipelineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLabelingPipelineResponse.Merge(m, src)
}
func (m *CreateLabelingPipelineResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLabelingPipelineResponse.Size(m)
}
func (m *CreateLabelingPipelineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLabelingPipelineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLabelingPipelineResponse proto.InternalMessageInfo

type UpdateLabelingPipelineRequest struct {
	Namespace            string                         `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string              `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.LabelingPipelineSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UpdateLabelingPipelineRequest) Reset()         { *m = UpdateLabelingPipelineRequest{} }
func (m *UpdateLabelingPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateLabelingPipelineRequest) ProtoMessage()    {}
func (*UpdateLabelingPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{4}
}
func (m *UpdateLabelingPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLabelingPipelineRequest.Unmarshal(m, b)
}
func (m *UpdateLabelingPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLabelingPipelineRequest.Marshal(b, m, deterministic)
}
func (m *UpdateLabelingPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLabelingPipelineRequest.Merge(m, src)
}
func (m *UpdateLabelingPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateLabelingPipelineRequest.Size(m)
}
func (m *UpdateLabelingPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLabelingPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLabelingPipelineRequest proto.InternalMessageInfo

func (m *UpdateLabelingPipelineRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateLabelingPipelineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateLabelingPipelineRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateLabelingPipelineRequest) GetSpec() *v1alpha1.LabelingPipelineSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateLabelingPipelineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateLabelingPipelineResponse) Reset()         { *m = UpdateLabelingPipelineResponse{} }
func (m *UpdateLabelingPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateLabelingPipelineResponse) ProtoMessage()    {}
func (*UpdateLabelingPipelineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{5}
}
func (m *UpdateLabelingPipelineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLabelingPipelineResponse.Unmarshal(m, b)
}
func (m *UpdateLabelingPipelineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLabelingPipelineResponse.Marshal(b, m, deterministic)
}
func (m *UpdateLabelingPipelineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLabelingPipelineResponse.Merge(m, src)
}
func (m *UpdateLabelingPipelineResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateLabelingPipelineResponse.Size(m)
}
func (m *UpdateLabelingPipelineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLabelingPipelineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLabelingPipelineResponse proto.InternalMessageInfo

type GetLabelingPipelineRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLabelingPipelineRequest) Reset()         { *m = GetLabelingPipelineRequest{} }
func (m *GetLabelingPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*GetLabelingPipelineRequest) ProtoMessage()    {}
func (*GetLabelingPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{6}
}
func (m *GetLabelingPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabelingPipelineRequest.Unmarshal(m, b)
}
func (m *GetLabelingPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabelingPipelineRequest.Marshal(b, m, deterministic)
}
func (m *GetLabelingPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabelingPipelineRequest.Merge(m, src)
}
func (m *GetLabelingPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_GetLabelingPipelineRequest.Size(m)
}
func (m *GetLabelingPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabelingPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabelingPipelineRequest proto.InternalMessageInfo

func (m *GetLabelingPipelineRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetLabelingPipelineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetLabelingPipelineResponse struct {
	Item                 *v1alpha1.LabelingPipeline `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string                     `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetLabelingPipelineResponse) Reset()         { *m = GetLabelingPipelineResponse{} }
func (m *GetLabelingPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*GetLabelingPipelineResponse) ProtoMessage()    {}
func (*GetLabelingPipelineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{7}
}
func (m *GetLabelingPipelineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabelingPipelineResponse.Unmarshal(m, b)
}
func (m *GetLabelingPipelineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabelingPipelineResponse.Marshal(b, m, deterministic)
}
func (m *GetLabelingPipelineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabelingPipelineResponse.Merge(m, src)
}
func (m *GetLabelingPipelineResponse) XXX_Size() int {
	return xxx_messageInfo_GetLabelingPipelineResponse.Size(m)
}
func (m *GetLabelingPipelineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabelingPipelineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabelingPipelineResponse proto.InternalMessageInfo

func (m *GetLabelingPipelineResponse) GetItem() *v1alpha1.LabelingPipeline {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetLabelingPipelineResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type DeleteLabelingPipelineRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLabelingPipelineRequest) Reset()         { *m = DeleteLabelingPipelineRequest{} }
func (m *DeleteLabelingPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLabelingPipelineRequest) ProtoMessage()    {}
func (*DeleteLabelingPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{8}
}
func (m *DeleteLabelingPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLabelingPipelineRequest.Unmarshal(m, b)
}
func (m *DeleteLabelingPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLabelingPipelineRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLabelingPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLabelingPipelineRequest.Merge(m, src)
}
func (m *DeleteLabelingPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLabelingPipelineRequest.Size(m)
}
func (m *DeleteLabelingPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLabelingPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLabelingPipelineRequest proto.InternalMessageInfo

func (m *DeleteLabelingPipelineRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteLabelingPipelineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteLabelingPipelineResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLabelingPipelineResponse) Reset()         { *m = DeleteLabelingPipelineResponse{} }
func (m *DeleteLabelingPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteLabelingPipelineResponse) ProtoMessage()    {}
func (*DeleteLabelingPipelineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d347002ed9b2beb, []int{9}
}
func (m *DeleteLabelingPipelineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLabelingPipelineResponse.Unmarshal(m, b)
}
func (m *DeleteLabelingPipelineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLabelingPipelineResponse.Marshal(b, m, deterministic)
}
func (m *DeleteLabelingPipelineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLabelingPipelineResponse.Merge(m, src)
}
func (m *DeleteLabelingPipelineResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteLabelingPipelineResponse.Size(m)
}
func (m *DeleteLabelingPipelineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLabelingPipelineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLabelingPipelineResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListLabelingPipelineRequest)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.ListLabelingPipelineRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.ListLabelingPipelineRequest.LabelsEntry")
	proto.RegisterType((*ListLabelingPipelineResponse)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.ListLabelingPipelineResponse")
	proto.RegisterType((*CreateLabelingPipelineRequest)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.CreateLabelingPipelineRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.CreateLabelingPipelineRequest.LabelsEntry")
	proto.RegisterType((*CreateLabelingPipelineResponse)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.CreateLabelingPipelineResponse")
	proto.RegisterType((*UpdateLabelingPipelineRequest)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.UpdateLabelingPipelineRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.UpdateLabelingPipelineRequest.LabelsEntry")
	proto.RegisterType((*UpdateLabelingPipelineResponse)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.UpdateLabelingPipelineResponse")
	proto.RegisterType((*GetLabelingPipelineRequest)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.GetLabelingPipelineRequest")
	proto.RegisterType((*GetLabelingPipelineResponse)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.GetLabelingPipelineResponse")
	proto.RegisterType((*DeleteLabelingPipelineRequest)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.DeleteLabelingPipelineRequest")
	proto.RegisterType((*DeleteLabelingPipelineResponse)(nil), "github.com.metaprov.modeld.services.labelingpipeline.v1.DeleteLabelingPipelineResponse")
}

func init() {
	proto.RegisterFile("services/labelingpipeline/v1/labelingpipeline.proto", fileDescriptor_5d347002ed9b2beb)
}

var fileDescriptor_5d347002ed9b2beb = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x77, 0xb7, 0x69, 0x21, 0x13, 0x04, 0x19, 0x6b, 0x8c, 0x9b, 0x54, 0xe2, 0x7a, 0x89, 0x97,
	0x59, 0x92, 0x1e, 0xb4, 0x41, 0x10, 0xb4, 0x2a, 0x48, 0x11, 0xdd, 0xb6, 0x0a, 0x3d, 0x48, 0x27,
	0xc9, 0x63, 0xbb, 0x64, 0xff, 0x8c, 0x3b, 0x93, 0xc5, 0x28, 0x82, 0x78, 0x14, 0x3c, 0xf9, 0x55,
	0xfc, 0x16, 0x1e, 0xfd, 0x02, 0x82, 0x1e, 0x3c, 0x7a, 0xf0, 0x03, 0xc8, 0xcc, 0x6e, 0x52, 0x6d,
	0x36, 0x2b, 0x24, 0x5b, 0xe8, 0x69, 0x67, 0x1e, 0xcc, 0xef, 0xbd, 0xdf, 0xfb, 0xbd, 0xf9, 0xed,
	0xa0, 0x4d, 0x0e, 0x51, 0xec, 0xf6, 0x81, 0x5b, 0x1e, 0xed, 0x81, 0xe7, 0x06, 0x0e, 0x73, 0x99,
	0xfc, 0x82, 0x15, 0xb7, 0x67, 0x62, 0x84, 0x45, 0xa1, 0x08, 0xf1, 0x4d, 0xc7, 0x15, 0x47, 0xa3,
	0x1e, 0xe9, 0x87, 0x3e, 0xf1, 0x41, 0x50, 0x16, 0x85, 0x31, 0xf1, 0xc3, 0x01, 0x78, 0x03, 0x32,
	0xc1, 0x23, 0x33, 0x67, 0xe3, 0xb6, 0xd1, 0x70, 0xc2, 0xd0, 0xf1, 0xc0, 0xa2, 0xcc, 0xb5, 0x68,
	0x10, 0x84, 0x82, 0x0a, 0x37, 0x0c, 0x78, 0x02, 0x6b, 0x6c, 0x1f, 0xc3, 0x5a, 0x13, 0x58, 0x2b,
	0x81, 0x95, 0x07, 0xd8, 0xd0, 0x91, 0x07, 0xb9, 0x35, 0xa0, 0x82, 0x5a, 0x71, 0x9b, 0x7a, 0xec,
	0x88, 0xb6, 0x2d, 0x07, 0x02, 0x88, 0xa8, 0x80, 0x41, 0x82, 0x62, 0xfe, 0xd2, 0x50, 0x7d, 0xc7,
	0xe5, 0x62, 0x27, 0xcd, 0xff, 0x24, 0xcd, 0x6f, 0xc3, 0xcb, 0x11, 0x70, 0x81, 0x1b, 0xa8, 0x1c,
	0x50, 0x1f, 0x38, 0xa3, 0x7d, 0xa8, 0x69, 0x4d, 0xad, 0x55, 0xb6, 0x8f, 0x03, 0xf8, 0x15, 0x5a,
	0x53, 0x85, 0xf3, 0x9a, 0xde, 0x5c, 0x69, 0x55, 0x3a, 0x87, 0x64, 0x41, 0xae, 0x24, 0xa7, 0x06,
	0xa2, 0xe2, 0xfc, 0x7e, 0x20, 0xa2, 0xb1, 0x9d, 0xe6, 0x33, 0xb6, 0x50, 0xe5, 0xaf, 0x30, 0xbe,
	0x80, 0x56, 0x86, 0x30, 0x4e, 0x0b, 0x94, 0x4b, 0xbc, 0x8e, 0x56, 0x63, 0xea, 0x8d, 0xa0, 0xa6,
	0xab, 0x58, 0xb2, 0xe9, 0xea, 0xb7, 0x34, 0xf3, 0x9d, 0x86, 0x1a, 0xd9, 0xe9, 0x38, 0x0b, 0x03,
	0x0e, 0xf8, 0x10, 0xad, 0xba, 0x02, 0x7c, 0xae, 0xe0, 0x2a, 0x9d, 0x47, 0x39, 0xa4, 0x28, 0x73,
	0x09, 0x1b, 0x3a, 0x44, 0x76, 0x9a, 0xc8, 0x4e, 0x93, 0x49, 0xa7, 0xc9, 0x49, 0x78, 0x99, 0xd2,
	0x4e, 0x80, 0xcd, 0x9f, 0x3a, 0xda, 0xb8, 0x17, 0x01, 0x15, 0xb0, 0x58, 0xdf, 0x31, 0x2a, 0xc9,
	0x4d, 0xca, 0x4d, 0xad, 0xf1, 0xeb, 0xa9, 0x16, 0x2b, 0x4a, 0x8b, 0xde, 0xc2, 0x5a, 0xe4, 0x56,
	0x96, 0xa5, 0x06, 0x7e, 0x81, 0x4a, 0x9c, 0x41, 0xbf, 0x56, 0x2a, 0xb2, 0x61, 0xbb, 0x0c, 0xfa,
	0xb6, 0xc2, 0x5d, 0x46, 0xed, 0x26, 0xba, 0x3a, 0x8f, 0x4f, 0x22, 0xb7, 0x12, 0x63, 0x9f, 0x0d,
	0xce, 0xa8, 0x18, 0xb9, 0x95, 0xe5, 0x8a, 0xb1, 0x7a, 0x26, 0xc5, 0x98, 0xc7, 0x27, 0x15, 0xe3,
	0x31, 0x32, 0x1e, 0x82, 0x28, 0x4c, 0x08, 0xf3, 0xa3, 0x86, 0xea, 0x99, 0x80, 0xe9, 0x5d, 0x3f,
	0x40, 0x25, 0x79, 0x25, 0xd3, 0xab, 0xfe, 0xa0, 0x98, 0x66, 0xd9, 0x0a, 0x53, 0xd6, 0x33, 0xa6,
	0xbe, 0x37, 0xa9, 0x47, 0xae, 0xcd, 0xa7, 0x68, 0x63, 0x1b, 0x3c, 0x28, 0x70, 0xd6, 0x64, 0x53,
	0xe7, 0x41, 0x26, 0x24, 0x3b, 0x9f, 0xcb, 0xe8, 0xf2, 0x8c, 0xa0, 0xc9, 0xd0, 0xe1, 0x6f, 0x1a,
	0xba, 0x94, 0xe5, 0x86, 0x1c, 0xef, 0x9d, 0x86, 0x99, 0x1b, 0xfb, 0x05, 0xa3, 0xa6, 0x73, 0x63,
	0xbc, 0xff, 0xfa, 0xe3, 0x93, 0xbe, 0x8e, 0xf1, 0xf4, 0x47, 0xcc, 0xa6, 0x44, 0x7e, 0x6b, 0xa8,
	0x9a, 0xed, 0x01, 0xf8, 0xd9, 0xe9, 0x98, 0xa4, 0xf1, 0xbc, 0x70, 0xdc, 0x94, 0xe7, 0x0d, 0xc5,
	0xf3, 0xba, 0x99, 0xc1, 0xb3, 0x7b, 0xfe, 0x9f, 0x3d, 0xfe, 0xae, 0xa1, 0x8b, 0x19, 0xa3, 0x8f,
	0x77, 0x17, 0xae, 0x6d, 0xfe, 0xcd, 0x34, 0xf6, 0x8a, 0x05, 0x4d, 0xd9, 0x5e, 0x53, 0x6c, 0xeb,
	0xf8, 0xca, 0x2c, 0x5b, 0xeb, 0x8d, 0x1c, 0xfe, 0xb7, 0xf8, 0x83, 0x8e, 0xaa, 0xd9, 0x9e, 0xb2,
	0x84, 0xb8, 0xb9, 0xa6, 0xbb, 0x84, 0xb8, 0xff, 0x31, 0xbf, 0x3b, 0x8a, 0xee, 0x96, 0xd1, 0xca,
	0xa2, 0xab, 0xf6, 0x2a, 0x9d, 0x32, 0x1f, 0xc5, 0xfe, 0xa4, 0xe4, 0x5f, 0x34, 0x54, 0xcd, 0xf6,
	0x82, 0x25, 0x9a, 0x91, 0xeb, 0x57, 0x4b, 0x34, 0x23, 0xdf, 0xb4, 0xcc, 0x73, 0x77, 0x6f, 0x1f,
	0x74, 0xf3, 0xdf, 0xb8, 0x79, 0xaf, 0xf1, 0xde, 0x9a, 0x7a, 0xe0, 0x6e, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0x62, 0x62, 0x5d, 0xb4, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LabelingPipelineServiceClient is the client API for LabelingPipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabelingPipelineServiceClient interface {
	ListLabelingPipelines(ctx context.Context, in *ListLabelingPipelineRequest, opts ...grpc.CallOption) (*ListLabelingPipelineResponse, error)
	CreateLabelingPipeline(ctx context.Context, in *CreateLabelingPipelineRequest, opts ...grpc.CallOption) (*CreateLabelingPipelineResponse, error)
	GetLabelingPipeline(ctx context.Context, in *GetLabelingPipelineRequest, opts ...grpc.CallOption) (*GetLabelingPipelineResponse, error)
	UpdateLabelingPipeline(ctx context.Context, in *UpdateLabelingPipelineRequest, opts ...grpc.CallOption) (*UpdateLabelingPipelineResponse, error)
	DeleteLabelingPipeline(ctx context.Context, in *DeleteLabelingPipelineRequest, opts ...grpc.CallOption) (*DeleteLabelingPipelineResponse, error)
}

type labelingPipelineServiceClient struct {
	cc *grpc.ClientConn
}

func NewLabelingPipelineServiceClient(cc *grpc.ClientConn) LabelingPipelineServiceClient {
	return &labelingPipelineServiceClient{cc}
}

func (c *labelingPipelineServiceClient) ListLabelingPipelines(ctx context.Context, in *ListLabelingPipelineRequest, opts ...grpc.CallOption) (*ListLabelingPipelineResponse, error) {
	out := new(ListLabelingPipelineResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/ListLabelingPipelines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelingPipelineServiceClient) CreateLabelingPipeline(ctx context.Context, in *CreateLabelingPipelineRequest, opts ...grpc.CallOption) (*CreateLabelingPipelineResponse, error) {
	out := new(CreateLabelingPipelineResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/CreateLabelingPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelingPipelineServiceClient) GetLabelingPipeline(ctx context.Context, in *GetLabelingPipelineRequest, opts ...grpc.CallOption) (*GetLabelingPipelineResponse, error) {
	out := new(GetLabelingPipelineResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/GetLabelingPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelingPipelineServiceClient) UpdateLabelingPipeline(ctx context.Context, in *UpdateLabelingPipelineRequest, opts ...grpc.CallOption) (*UpdateLabelingPipelineResponse, error) {
	out := new(UpdateLabelingPipelineResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/UpdateLabelingPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelingPipelineServiceClient) DeleteLabelingPipeline(ctx context.Context, in *DeleteLabelingPipelineRequest, opts ...grpc.CallOption) (*DeleteLabelingPipelineResponse, error) {
	out := new(DeleteLabelingPipelineResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/DeleteLabelingPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelingPipelineServiceServer is the server API for LabelingPipelineService service.
type LabelingPipelineServiceServer interface {
	ListLabelingPipelines(context.Context, *ListLabelingPipelineRequest) (*ListLabelingPipelineResponse, error)
	CreateLabelingPipeline(context.Context, *CreateLabelingPipelineRequest) (*CreateLabelingPipelineResponse, error)
	GetLabelingPipeline(context.Context, *GetLabelingPipelineRequest) (*GetLabelingPipelineResponse, error)
	UpdateLabelingPipeline(context.Context, *UpdateLabelingPipelineRequest) (*UpdateLabelingPipelineResponse, error)
	DeleteLabelingPipeline(context.Context, *DeleteLabelingPipelineRequest) (*DeleteLabelingPipelineResponse, error)
}

// UnimplementedLabelingPipelineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLabelingPipelineServiceServer struct {
}

func (*UnimplementedLabelingPipelineServiceServer) ListLabelingPipelines(ctx context.Context, req *ListLabelingPipelineRequest) (*ListLabelingPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabelingPipelines not implemented")
}
func (*UnimplementedLabelingPipelineServiceServer) CreateLabelingPipeline(ctx context.Context, req *CreateLabelingPipelineRequest) (*CreateLabelingPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabelingPipeline not implemented")
}
func (*UnimplementedLabelingPipelineServiceServer) GetLabelingPipeline(ctx context.Context, req *GetLabelingPipelineRequest) (*GetLabelingPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabelingPipeline not implemented")
}
func (*UnimplementedLabelingPipelineServiceServer) UpdateLabelingPipeline(ctx context.Context, req *UpdateLabelingPipelineRequest) (*UpdateLabelingPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLabelingPipeline not implemented")
}
func (*UnimplementedLabelingPipelineServiceServer) DeleteLabelingPipeline(ctx context.Context, req *DeleteLabelingPipelineRequest) (*DeleteLabelingPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLabelingPipeline not implemented")
}

func RegisterLabelingPipelineServiceServer(s *grpc.Server, srv LabelingPipelineServiceServer) {
	s.RegisterService(&_LabelingPipelineService_serviceDesc, srv)
}

func _LabelingPipelineService_ListLabelingPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabelingPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelingPipelineServiceServer).ListLabelingPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/ListLabelingPipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelingPipelineServiceServer).ListLabelingPipelines(ctx, req.(*ListLabelingPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelingPipelineService_CreateLabelingPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabelingPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelingPipelineServiceServer).CreateLabelingPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/CreateLabelingPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelingPipelineServiceServer).CreateLabelingPipeline(ctx, req.(*CreateLabelingPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelingPipelineService_GetLabelingPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelingPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelingPipelineServiceServer).GetLabelingPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/GetLabelingPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelingPipelineServiceServer).GetLabelingPipeline(ctx, req.(*GetLabelingPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelingPipelineService_UpdateLabelingPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabelingPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelingPipelineServiceServer).UpdateLabelingPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/UpdateLabelingPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelingPipelineServiceServer).UpdateLabelingPipeline(ctx, req.(*UpdateLabelingPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelingPipelineService_DeleteLabelingPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLabelingPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelingPipelineServiceServer).DeleteLabelingPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService/DeleteLabelingPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelingPipelineServiceServer).DeleteLabelingPipeline(ctx, req.(*DeleteLabelingPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabelingPipelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.labelingpipeline.v1.LabelingPipelineService",
	HandlerType: (*LabelingPipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLabelingPipelines",
			Handler:    _LabelingPipelineService_ListLabelingPipelines_Handler,
		},
		{
			MethodName: "CreateLabelingPipeline",
			Handler:    _LabelingPipelineService_CreateLabelingPipeline_Handler,
		},
		{
			MethodName: "GetLabelingPipeline",
			Handler:    _LabelingPipelineService_GetLabelingPipeline_Handler,
		},
		{
			MethodName: "UpdateLabelingPipeline",
			Handler:    _LabelingPipelineService_UpdateLabelingPipeline_Handler,
		},
		{
			MethodName: "DeleteLabelingPipeline",
			Handler:    _LabelingPipelineService_DeleteLabelingPipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/labelingpipeline/v1/labelingpipeline.proto",
}