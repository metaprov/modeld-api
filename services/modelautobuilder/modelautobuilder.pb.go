// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/modelautobuilder/modelautobuilder.proto

package modelautobuilder

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/training/v1alpha1"
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

type ListModelAutobuildersRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListModelAutobuildersRequest) Reset()         { *m = ListModelAutobuildersRequest{} }
func (m *ListModelAutobuildersRequest) String() string { return proto.CompactTextString(m) }
func (*ListModelAutobuildersRequest) ProtoMessage()    {}
func (*ListModelAutobuildersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{0}
}
func (m *ListModelAutobuildersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListModelAutobuildersRequest.Unmarshal(m, b)
}
func (m *ListModelAutobuildersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListModelAutobuildersRequest.Marshal(b, m, deterministic)
}
func (m *ListModelAutobuildersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListModelAutobuildersRequest.Merge(m, src)
}
func (m *ListModelAutobuildersRequest) XXX_Size() int {
	return xxx_messageInfo_ListModelAutobuildersRequest.Size(m)
}
func (m *ListModelAutobuildersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListModelAutobuildersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListModelAutobuildersRequest proto.InternalMessageInfo

func (m *ListModelAutobuildersRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListModelAutobuildersRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListModelAutobuildersResponse struct {
	Items                *v1alpha1.ModelAutobuilderList `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ListModelAutobuildersResponse) Reset()         { *m = ListModelAutobuildersResponse{} }
func (m *ListModelAutobuildersResponse) String() string { return proto.CompactTextString(m) }
func (*ListModelAutobuildersResponse) ProtoMessage()    {}
func (*ListModelAutobuildersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{1}
}
func (m *ListModelAutobuildersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListModelAutobuildersResponse.Unmarshal(m, b)
}
func (m *ListModelAutobuildersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListModelAutobuildersResponse.Marshal(b, m, deterministic)
}
func (m *ListModelAutobuildersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListModelAutobuildersResponse.Merge(m, src)
}
func (m *ListModelAutobuildersResponse) XXX_Size() int {
	return xxx_messageInfo_ListModelAutobuildersResponse.Size(m)
}
func (m *ListModelAutobuildersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListModelAutobuildersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListModelAutobuildersResponse proto.InternalMessageInfo

func (m *ListModelAutobuildersResponse) GetItems() *v1alpha1.ModelAutobuilderList {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateModelAutobuilderRequest struct {
	Namespace            string                         `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string              `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.ModelAutobuilderSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CreateModelAutobuilderRequest) Reset()         { *m = CreateModelAutobuilderRequest{} }
func (m *CreateModelAutobuilderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateModelAutobuilderRequest) ProtoMessage()    {}
func (*CreateModelAutobuilderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{2}
}
func (m *CreateModelAutobuilderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModelAutobuilderRequest.Unmarshal(m, b)
}
func (m *CreateModelAutobuilderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModelAutobuilderRequest.Marshal(b, m, deterministic)
}
func (m *CreateModelAutobuilderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModelAutobuilderRequest.Merge(m, src)
}
func (m *CreateModelAutobuilderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateModelAutobuilderRequest.Size(m)
}
func (m *CreateModelAutobuilderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModelAutobuilderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModelAutobuilderRequest proto.InternalMessageInfo

func (m *CreateModelAutobuilderRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateModelAutobuilderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateModelAutobuilderRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateModelAutobuilderRequest) GetSpec() *v1alpha1.ModelAutobuilderSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type CreateModelAutobuilderResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateModelAutobuilderResponse) Reset()         { *m = CreateModelAutobuilderResponse{} }
func (m *CreateModelAutobuilderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateModelAutobuilderResponse) ProtoMessage()    {}
func (*CreateModelAutobuilderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{3}
}
func (m *CreateModelAutobuilderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModelAutobuilderResponse.Unmarshal(m, b)
}
func (m *CreateModelAutobuilderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModelAutobuilderResponse.Marshal(b, m, deterministic)
}
func (m *CreateModelAutobuilderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModelAutobuilderResponse.Merge(m, src)
}
func (m *CreateModelAutobuilderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateModelAutobuilderResponse.Size(m)
}
func (m *CreateModelAutobuilderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModelAutobuilderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModelAutobuilderResponse proto.InternalMessageInfo

type UpdateModelAutobuilderRequest struct {
	Namespace            string                         `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string              `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.ModelAutobuilderSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UpdateModelAutobuilderRequest) Reset()         { *m = UpdateModelAutobuilderRequest{} }
func (m *UpdateModelAutobuilderRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateModelAutobuilderRequest) ProtoMessage()    {}
func (*UpdateModelAutobuilderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{4}
}
func (m *UpdateModelAutobuilderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateModelAutobuilderRequest.Unmarshal(m, b)
}
func (m *UpdateModelAutobuilderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateModelAutobuilderRequest.Marshal(b, m, deterministic)
}
func (m *UpdateModelAutobuilderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateModelAutobuilderRequest.Merge(m, src)
}
func (m *UpdateModelAutobuilderRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateModelAutobuilderRequest.Size(m)
}
func (m *UpdateModelAutobuilderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateModelAutobuilderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateModelAutobuilderRequest proto.InternalMessageInfo

func (m *UpdateModelAutobuilderRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateModelAutobuilderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateModelAutobuilderRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateModelAutobuilderRequest) GetSpec() *v1alpha1.ModelAutobuilderSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateModelAutobuilderResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateModelAutobuilderResponse) Reset()         { *m = UpdateModelAutobuilderResponse{} }
func (m *UpdateModelAutobuilderResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateModelAutobuilderResponse) ProtoMessage()    {}
func (*UpdateModelAutobuilderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{5}
}
func (m *UpdateModelAutobuilderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateModelAutobuilderResponse.Unmarshal(m, b)
}
func (m *UpdateModelAutobuilderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateModelAutobuilderResponse.Marshal(b, m, deterministic)
}
func (m *UpdateModelAutobuilderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateModelAutobuilderResponse.Merge(m, src)
}
func (m *UpdateModelAutobuilderResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateModelAutobuilderResponse.Size(m)
}
func (m *UpdateModelAutobuilderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateModelAutobuilderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateModelAutobuilderResponse proto.InternalMessageInfo

type GetModelAutobuilderRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetModelAutobuilderRequest) Reset()         { *m = GetModelAutobuilderRequest{} }
func (m *GetModelAutobuilderRequest) String() string { return proto.CompactTextString(m) }
func (*GetModelAutobuilderRequest) ProtoMessage()    {}
func (*GetModelAutobuilderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{6}
}
func (m *GetModelAutobuilderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelAutobuilderRequest.Unmarshal(m, b)
}
func (m *GetModelAutobuilderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelAutobuilderRequest.Marshal(b, m, deterministic)
}
func (m *GetModelAutobuilderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelAutobuilderRequest.Merge(m, src)
}
func (m *GetModelAutobuilderRequest) XXX_Size() int {
	return xxx_messageInfo_GetModelAutobuilderRequest.Size(m)
}
func (m *GetModelAutobuilderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelAutobuilderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelAutobuilderRequest proto.InternalMessageInfo

func (m *GetModelAutobuilderRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetModelAutobuilderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetModelAutobuilderResponse struct {
	Item                 *v1alpha1.ModelAutobuilder `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string                     `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetModelAutobuilderResponse) Reset()         { *m = GetModelAutobuilderResponse{} }
func (m *GetModelAutobuilderResponse) String() string { return proto.CompactTextString(m) }
func (*GetModelAutobuilderResponse) ProtoMessage()    {}
func (*GetModelAutobuilderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{7}
}
func (m *GetModelAutobuilderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelAutobuilderResponse.Unmarshal(m, b)
}
func (m *GetModelAutobuilderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelAutobuilderResponse.Marshal(b, m, deterministic)
}
func (m *GetModelAutobuilderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelAutobuilderResponse.Merge(m, src)
}
func (m *GetModelAutobuilderResponse) XXX_Size() int {
	return xxx_messageInfo_GetModelAutobuilderResponse.Size(m)
}
func (m *GetModelAutobuilderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelAutobuilderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelAutobuilderResponse proto.InternalMessageInfo

func (m *GetModelAutobuilderResponse) GetItem() *v1alpha1.ModelAutobuilder {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetModelAutobuilderResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type DeleteModelAutobuilderRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteModelAutobuilderRequest) Reset()         { *m = DeleteModelAutobuilderRequest{} }
func (m *DeleteModelAutobuilderRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteModelAutobuilderRequest) ProtoMessage()    {}
func (*DeleteModelAutobuilderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{8}
}
func (m *DeleteModelAutobuilderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteModelAutobuilderRequest.Unmarshal(m, b)
}
func (m *DeleteModelAutobuilderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteModelAutobuilderRequest.Marshal(b, m, deterministic)
}
func (m *DeleteModelAutobuilderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteModelAutobuilderRequest.Merge(m, src)
}
func (m *DeleteModelAutobuilderRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteModelAutobuilderRequest.Size(m)
}
func (m *DeleteModelAutobuilderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteModelAutobuilderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteModelAutobuilderRequest proto.InternalMessageInfo

func (m *DeleteModelAutobuilderRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteModelAutobuilderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteModelAutobuilderResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteModelAutobuilderResponse) Reset()         { *m = DeleteModelAutobuilderResponse{} }
func (m *DeleteModelAutobuilderResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteModelAutobuilderResponse) ProtoMessage()    {}
func (*DeleteModelAutobuilderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_615501b4f4fb5bd2, []int{9}
}
func (m *DeleteModelAutobuilderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteModelAutobuilderResponse.Unmarshal(m, b)
}
func (m *DeleteModelAutobuilderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteModelAutobuilderResponse.Marshal(b, m, deterministic)
}
func (m *DeleteModelAutobuilderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteModelAutobuilderResponse.Merge(m, src)
}
func (m *DeleteModelAutobuilderResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteModelAutobuilderResponse.Size(m)
}
func (m *DeleteModelAutobuilderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteModelAutobuilderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteModelAutobuilderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListModelAutobuildersRequest)(nil), "github.com.metaprov.modeld.services.modelautobuilder.ListModelAutobuildersRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.modelautobuilder.ListModelAutobuildersRequest.LabelsEntry")
	proto.RegisterType((*ListModelAutobuildersResponse)(nil), "github.com.metaprov.modeld.services.modelautobuilder.ListModelAutobuildersResponse")
	proto.RegisterType((*CreateModelAutobuilderRequest)(nil), "github.com.metaprov.modeld.services.modelautobuilder.CreateModelAutobuilderRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.modelautobuilder.CreateModelAutobuilderRequest.LabelsEntry")
	proto.RegisterType((*CreateModelAutobuilderResponse)(nil), "github.com.metaprov.modeld.services.modelautobuilder.CreateModelAutobuilderResponse")
	proto.RegisterType((*UpdateModelAutobuilderRequest)(nil), "github.com.metaprov.modeld.services.modelautobuilder.UpdateModelAutobuilderRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.modelautobuilder.UpdateModelAutobuilderRequest.LabelsEntry")
	proto.RegisterType((*UpdateModelAutobuilderResponse)(nil), "github.com.metaprov.modeld.services.modelautobuilder.UpdateModelAutobuilderResponse")
	proto.RegisterType((*GetModelAutobuilderRequest)(nil), "github.com.metaprov.modeld.services.modelautobuilder.GetModelAutobuilderRequest")
	proto.RegisterType((*GetModelAutobuilderResponse)(nil), "github.com.metaprov.modeld.services.modelautobuilder.GetModelAutobuilderResponse")
	proto.RegisterType((*DeleteModelAutobuilderRequest)(nil), "github.com.metaprov.modeld.services.modelautobuilder.DeleteModelAutobuilderRequest")
	proto.RegisterType((*DeleteModelAutobuilderResponse)(nil), "github.com.metaprov.modeld.services.modelautobuilder.DeleteModelAutobuilderResponse")
}

func init() {
	proto.RegisterFile("services/modelautobuilder/modelautobuilder.proto", fileDescriptor_615501b4f4fb5bd2)
}

var fileDescriptor_615501b4f4fb5bd2 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6a, 0xd5, 0x4e,
	0x14, 0xfe, 0x25, 0xbd, 0x2d, 0x74, 0xba, 0x29, 0xf3, 0xb3, 0xed, 0x25, 0xde, 0x5b, 0x4a, 0x70,
	0x51, 0x41, 0x26, 0xb6, 0x0a, 0xda, 0xae, 0xfc, 0xaf, 0x48, 0x2d, 0x36, 0xd5, 0x8d, 0x8b, 0xca,
	0xdc, 0xe4, 0x90, 0x86, 0x26, 0x99, 0x31, 0x33, 0x89, 0x14, 0x71, 0x23, 0xb8, 0x72, 0x25, 0x3e,
	0x86, 0xcf, 0xe1, 0x13, 0xf8, 0x00, 0xba, 0x70, 0xe3, 0xca, 0x8d, 0x0f, 0x20, 0x33, 0xc9, 0xb5,
	0x72, 0x6f, 0x12, 0x25, 0x46, 0xe8, 0x6e, 0x72, 0x98, 0xf9, 0xce, 0xf9, 0xce, 0x77, 0xe6, 0xcb,
	0xa0, 0x8b, 0x02, 0xd2, 0x3c, 0xf4, 0x40, 0x38, 0x31, 0xf3, 0x21, 0xa2, 0x99, 0x64, 0xa3, 0x2c,
	0x8c, 0x7c, 0x48, 0xa7, 0x02, 0x84, 0xa7, 0x4c, 0x32, 0x7c, 0x39, 0x08, 0xe5, 0x61, 0x36, 0x22,
	0x1e, 0x8b, 0x49, 0x0c, 0x92, 0xf2, 0x94, 0xe5, 0x44, 0xef, 0xf5, 0xc9, 0x18, 0x8c, 0x4c, 0x9e,
	0xb5, 0x06, 0x01, 0x63, 0x41, 0x04, 0x0e, 0xe5, 0xa1, 0x43, 0x93, 0x84, 0x49, 0x2a, 0x43, 0x96,
	0x88, 0x02, 0xd3, 0xba, 0x77, 0x82, 0xe9, 0x8c, 0x31, 0x8b, 0xfc, 0xbe, 0x3a, 0xc0, 0x8f, 0x02,
	0x75, 0x50, 0x38, 0x32, 0xa5, 0x61, 0x12, 0x26, 0x81, 0x93, 0x6f, 0xd0, 0x88, 0x1f, 0xd2, 0x0d,
	0x27, 0x80, 0x04, 0x52, 0x2a, 0xc1, 0x2f, 0x90, 0xec, 0x6f, 0x06, 0x1a, 0xec, 0x84, 0x42, 0x3e,
	0x50, 0x87, 0xaf, 0x9f, 0x14, 0x20, 0x5c, 0x78, 0x96, 0x81, 0x90, 0x78, 0x80, 0xe6, 0x13, 0x1a,
	0x83, 0xe0, 0xd4, 0x83, 0xbe, 0xb1, 0x66, 0xac, 0xcf, 0xbb, 0x27, 0x01, 0x9c, 0xa3, 0xb9, 0x88,
	0x8e, 0x20, 0x12, 0x7d, 0x73, 0x6d, 0x66, 0x7d, 0x61, 0xf3, 0x80, 0xb4, 0x61, 0x4b, 0x9a, 0x2a,
	0x20, 0x3b, 0x3a, 0xc1, 0xed, 0x44, 0xa6, 0xc7, 0x6e, 0x99, 0xcd, 0xda, 0x42, 0x0b, 0xbf, 0x84,
	0xf1, 0x22, 0x9a, 0x39, 0x82, 0xe3, 0xb2, 0x3c, 0xb5, 0xc4, 0x67, 0xd0, 0x6c, 0x4e, 0xa3, 0x0c,
	0xfa, 0xa6, 0x8e, 0x15, 0x1f, 0xdb, 0xe6, 0x55, 0xc3, 0x7e, 0x6d, 0xa0, 0x61, 0x4d, 0x3e, 0xc1,
	0x59, 0x22, 0x00, 0xfb, 0x68, 0x36, 0x94, 0x10, 0x0b, 0x8d, 0xb7, 0xb0, 0xb9, 0xdb, 0xc0, 0x89,
	0xf2, 0x90, 0xf0, 0xa3, 0x80, 0xa8, 0x6e, 0x93, 0x71, 0xb7, 0xc9, 0xb8, 0xdb, 0x64, 0x32, 0x87,
	0xca, 0xeb, 0x16, 0xe0, 0xf6, 0x57, 0x13, 0x0d, 0x6f, 0xa6, 0x40, 0x25, 0x4c, 0xee, 0xfa, 0xb3,
	0xd6, 0x63, 0xd4, 0x53, 0x1f, 0x25, 0x41, 0xbd, 0xc6, 0xcf, 0x7f, 0xca, 0x31, 0xa3, 0xe5, 0x78,
	0xda, 0x4e, 0x8e, 0xc6, 0xb2, 0xaa, 0xf4, 0xc0, 0x23, 0xd4, 0x13, 0x1c, 0xbc, 0x7e, 0xaf, 0xeb,
	0x8e, 0xed, 0x73, 0xf0, 0x5c, 0x8d, 0xfd, 0x37, 0x9a, 0xaf, 0xa1, 0xd5, 0x3a, 0x4e, 0x85, 0xe6,
	0x5a, 0x8d, 0xc7, 0xdc, 0x3f, 0x8d, 0x6a, 0x34, 0x96, 0xd5, 0xa8, 0xc6, 0xec, 0xa9, 0x55, 0xa3,
	0x8e, 0x53, 0xa9, 0xc6, 0x2e, 0xb2, 0xee, 0x82, 0xec, 0x4c, 0x09, 0xfb, 0xad, 0x81, 0xce, 0x56,
	0x02, 0x96, 0x37, 0xfe, 0x00, 0xf5, 0xd4, 0xa5, 0x2c, 0x2f, 0xfc, 0xfd, 0xee, 0x1a, 0xe6, 0x6a,
	0x5c, 0x55, 0xd3, 0x31, 0x8d, 0xa3, 0x71, 0x4d, 0x6a, 0x6d, 0xef, 0xa1, 0xe1, 0x2d, 0x88, 0xa0,
	0xc3, 0x81, 0x53, 0x8d, 0xad, 0x83, 0x2c, 0x88, 0x6e, 0xbe, 0x9f, 0x47, 0x2b, 0x53, 0xa2, 0x16,
	0x93, 0x87, 0x3f, 0x19, 0x68, 0xa9, 0xd2, 0x18, 0xb1, 0xdb, 0xbd, 0xab, 0x5b, 0xfb, 0x9d, 0x62,
	0x96, 0x73, 0x33, 0x7c, 0xf5, 0xf1, 0xcb, 0x3b, 0x73, 0x05, 0x2f, 0x39, 0xf9, 0xc6, 0xd4, 0xff,
	0x58, 0xe0, 0xef, 0x06, 0x5a, 0xae, 0xf6, 0x01, 0xbc, 0xff, 0x0f, 0x9c, 0xd2, 0x7a, 0xd4, 0x2d,
	0x68, 0x49, 0xd2, 0xd1, 0x24, 0xcf, 0xdb, 0xd5, 0x24, 0xb7, 0x17, 0x27, 0x43, 0xf8, 0xb3, 0x81,
	0xfe, 0xaf, 0x98, 0x7e, 0xfc, 0xb0, 0x5d, 0x79, 0xf5, 0x37, 0xd3, 0xda, 0xeb, 0x10, 0xb1, 0x64,
	0x7b, 0x4e, 0xb3, 0x5d, 0xc5, 0x83, 0x4a, 0xb6, 0xce, 0x0b, 0x35, 0xf8, 0x2f, 0xf1, 0x1b, 0x13,
	0x2d, 0x57, 0x7b, 0x4a, 0x5b, 0x65, 0x1b, 0x5d, 0xb7, 0xad, 0xb2, 0xbf, 0xb1, 0xbd, 0x3b, 0x9a,
	0xeb, 0x35, 0xeb, 0x42, 0x0d, 0x57, 0x8f, 0xa5, 0x3c, 0x13, 0x3a, 0x9f, 0x4f, 0x25, 0x25, 0x9a,
	0x7b, 0x85, 0xe0, 0x1f, 0x0c, 0xb4, 0x5c, 0x6d, 0x04, 0x6d, 0xbb, 0xd1, 0xe8, 0x54, 0x6d, 0xbb,
	0xd1, 0xec, 0x55, 0xf6, 0x7f, 0x37, 0xb6, 0x9e, 0x5c, 0x69, 0x7e, 0xe8, 0xd6, 0x3e, 0xc6, 0x47,
	0x73, 0xfa, 0x79, 0x7b, 0xe9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xb8, 0xc4, 0x3d, 0xb0,
	0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModelAutobuilderServiceClient is the client API for ModelAutobuilderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModelAutobuilderServiceClient interface {
	ListModelAutobuilders(ctx context.Context, in *ListModelAutobuildersRequest, opts ...grpc.CallOption) (*ListModelAutobuildersResponse, error)
	CreateModelAutobuilder(ctx context.Context, in *CreateModelAutobuilderRequest, opts ...grpc.CallOption) (*CreateModelAutobuilderResponse, error)
	GetModelAutobuilder(ctx context.Context, in *GetModelAutobuilderRequest, opts ...grpc.CallOption) (*GetModelAutobuilderResponse, error)
	UpdateModelAutobuilder(ctx context.Context, in *UpdateModelAutobuilderRequest, opts ...grpc.CallOption) (*UpdateModelAutobuilderResponse, error)
	DeleteModelAutobuilder(ctx context.Context, in *DeleteModelAutobuilderRequest, opts ...grpc.CallOption) (*DeleteModelAutobuilderResponse, error)
}

type modelAutobuilderServiceClient struct {
	cc *grpc.ClientConn
}

func NewModelAutobuilderServiceClient(cc *grpc.ClientConn) ModelAutobuilderServiceClient {
	return &modelAutobuilderServiceClient{cc}
}

func (c *modelAutobuilderServiceClient) ListModelAutobuilders(ctx context.Context, in *ListModelAutobuildersRequest, opts ...grpc.CallOption) (*ListModelAutobuildersResponse, error) {
	out := new(ListModelAutobuildersResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/ListModelAutobuilders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelAutobuilderServiceClient) CreateModelAutobuilder(ctx context.Context, in *CreateModelAutobuilderRequest, opts ...grpc.CallOption) (*CreateModelAutobuilderResponse, error) {
	out := new(CreateModelAutobuilderResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/CreateModelAutobuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelAutobuilderServiceClient) GetModelAutobuilder(ctx context.Context, in *GetModelAutobuilderRequest, opts ...grpc.CallOption) (*GetModelAutobuilderResponse, error) {
	out := new(GetModelAutobuilderResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/GetModelAutobuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelAutobuilderServiceClient) UpdateModelAutobuilder(ctx context.Context, in *UpdateModelAutobuilderRequest, opts ...grpc.CallOption) (*UpdateModelAutobuilderResponse, error) {
	out := new(UpdateModelAutobuilderResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/UpdateModelAutobuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelAutobuilderServiceClient) DeleteModelAutobuilder(ctx context.Context, in *DeleteModelAutobuilderRequest, opts ...grpc.CallOption) (*DeleteModelAutobuilderResponse, error) {
	out := new(DeleteModelAutobuilderResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/DeleteModelAutobuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelAutobuilderServiceServer is the server API for ModelAutobuilderService service.
type ModelAutobuilderServiceServer interface {
	ListModelAutobuilders(context.Context, *ListModelAutobuildersRequest) (*ListModelAutobuildersResponse, error)
	CreateModelAutobuilder(context.Context, *CreateModelAutobuilderRequest) (*CreateModelAutobuilderResponse, error)
	GetModelAutobuilder(context.Context, *GetModelAutobuilderRequest) (*GetModelAutobuilderResponse, error)
	UpdateModelAutobuilder(context.Context, *UpdateModelAutobuilderRequest) (*UpdateModelAutobuilderResponse, error)
	DeleteModelAutobuilder(context.Context, *DeleteModelAutobuilderRequest) (*DeleteModelAutobuilderResponse, error)
}

// UnimplementedModelAutobuilderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedModelAutobuilderServiceServer struct {
}

func (*UnimplementedModelAutobuilderServiceServer) ListModelAutobuilders(ctx context.Context, req *ListModelAutobuildersRequest) (*ListModelAutobuildersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModelAutobuilders not implemented")
}
func (*UnimplementedModelAutobuilderServiceServer) CreateModelAutobuilder(ctx context.Context, req *CreateModelAutobuilderRequest) (*CreateModelAutobuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModelAutobuilder not implemented")
}
func (*UnimplementedModelAutobuilderServiceServer) GetModelAutobuilder(ctx context.Context, req *GetModelAutobuilderRequest) (*GetModelAutobuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelAutobuilder not implemented")
}
func (*UnimplementedModelAutobuilderServiceServer) UpdateModelAutobuilder(ctx context.Context, req *UpdateModelAutobuilderRequest) (*UpdateModelAutobuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModelAutobuilder not implemented")
}
func (*UnimplementedModelAutobuilderServiceServer) DeleteModelAutobuilder(ctx context.Context, req *DeleteModelAutobuilderRequest) (*DeleteModelAutobuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModelAutobuilder not implemented")
}

func RegisterModelAutobuilderServiceServer(s *grpc.Server, srv ModelAutobuilderServiceServer) {
	s.RegisterService(&_ModelAutobuilderService_serviceDesc, srv)
}

func _ModelAutobuilderService_ListModelAutobuilders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelAutobuildersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelAutobuilderServiceServer).ListModelAutobuilders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/ListModelAutobuilders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelAutobuilderServiceServer).ListModelAutobuilders(ctx, req.(*ListModelAutobuildersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelAutobuilderService_CreateModelAutobuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelAutobuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelAutobuilderServiceServer).CreateModelAutobuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/CreateModelAutobuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelAutobuilderServiceServer).CreateModelAutobuilder(ctx, req.(*CreateModelAutobuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelAutobuilderService_GetModelAutobuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelAutobuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelAutobuilderServiceServer).GetModelAutobuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/GetModelAutobuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelAutobuilderServiceServer).GetModelAutobuilder(ctx, req.(*GetModelAutobuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelAutobuilderService_UpdateModelAutobuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModelAutobuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelAutobuilderServiceServer).UpdateModelAutobuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/UpdateModelAutobuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelAutobuilderServiceServer).UpdateModelAutobuilder(ctx, req.(*UpdateModelAutobuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelAutobuilderService_DeleteModelAutobuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelAutobuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelAutobuilderServiceServer).DeleteModelAutobuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService/DeleteModelAutobuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelAutobuilderServiceServer).DeleteModelAutobuilder(ctx, req.(*DeleteModelAutobuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ModelAutobuilderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.modelautobuilder.ModelAutobuilderService",
	HandlerType: (*ModelAutobuilderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListModelAutobuilders",
			Handler:    _ModelAutobuilderService_ListModelAutobuilders_Handler,
		},
		{
			MethodName: "CreateModelAutobuilder",
			Handler:    _ModelAutobuilderService_CreateModelAutobuilder_Handler,
		},
		{
			MethodName: "GetModelAutobuilder",
			Handler:    _ModelAutobuilderService_GetModelAutobuilder_Handler,
		},
		{
			MethodName: "UpdateModelAutobuilder",
			Handler:    _ModelAutobuilderService_UpdateModelAutobuilder_Handler,
		},
		{
			MethodName: "DeleteModelAutobuilder",
			Handler:    _ModelAutobuilderService_DeleteModelAutobuilder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/modelautobuilder/modelautobuilder.proto",
}