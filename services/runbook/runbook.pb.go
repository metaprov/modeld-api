// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/runbook/runbook.proto

package runbook

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/team/v1alpha1"
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

type ListRunBooksRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListRunBooksRequest) Reset()         { *m = ListRunBooksRequest{} }
func (m *ListRunBooksRequest) String() string { return proto.CompactTextString(m) }
func (*ListRunBooksRequest) ProtoMessage()    {}
func (*ListRunBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{0}
}
func (m *ListRunBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunBooksRequest.Unmarshal(m, b)
}
func (m *ListRunBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunBooksRequest.Marshal(b, m, deterministic)
}
func (m *ListRunBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunBooksRequest.Merge(m, src)
}
func (m *ListRunBooksRequest) XXX_Size() int {
	return xxx_messageInfo_ListRunBooksRequest.Size(m)
}
func (m *ListRunBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunBooksRequest proto.InternalMessageInfo

func (m *ListRunBooksRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListRunBooksRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListRunBooksResponse struct {
	Items                *v1alpha1.RunBookList `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRunBooksResponse) Reset()         { *m = ListRunBooksResponse{} }
func (m *ListRunBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListRunBooksResponse) ProtoMessage()    {}
func (*ListRunBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{1}
}
func (m *ListRunBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunBooksResponse.Unmarshal(m, b)
}
func (m *ListRunBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunBooksResponse.Marshal(b, m, deterministic)
}
func (m *ListRunBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunBooksResponse.Merge(m, src)
}
func (m *ListRunBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListRunBooksResponse.Size(m)
}
func (m *ListRunBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunBooksResponse proto.InternalMessageInfo

func (m *ListRunBooksResponse) GetItems() *v1alpha1.RunBookList {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateRunBookRequest struct {
	Namespace            string                `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string     `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.RunBookSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRunBookRequest) Reset()         { *m = CreateRunBookRequest{} }
func (m *CreateRunBookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRunBookRequest) ProtoMessage()    {}
func (*CreateRunBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{2}
}
func (m *CreateRunBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRunBookRequest.Unmarshal(m, b)
}
func (m *CreateRunBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRunBookRequest.Marshal(b, m, deterministic)
}
func (m *CreateRunBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRunBookRequest.Merge(m, src)
}
func (m *CreateRunBookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRunBookRequest.Size(m)
}
func (m *CreateRunBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRunBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRunBookRequest proto.InternalMessageInfo

func (m *CreateRunBookRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateRunBookRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRunBookRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateRunBookRequest) GetSpec() *v1alpha1.RunBookSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type CreateRunBookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRunBookResponse) Reset()         { *m = CreateRunBookResponse{} }
func (m *CreateRunBookResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRunBookResponse) ProtoMessage()    {}
func (*CreateRunBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{3}
}
func (m *CreateRunBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRunBookResponse.Unmarshal(m, b)
}
func (m *CreateRunBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRunBookResponse.Marshal(b, m, deterministic)
}
func (m *CreateRunBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRunBookResponse.Merge(m, src)
}
func (m *CreateRunBookResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRunBookResponse.Size(m)
}
func (m *CreateRunBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRunBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRunBookResponse proto.InternalMessageInfo

type UpdateRunBookRequest struct {
	Namespace            string                `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string     `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.RunBookSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRunBookRequest) Reset()         { *m = UpdateRunBookRequest{} }
func (m *UpdateRunBookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRunBookRequest) ProtoMessage()    {}
func (*UpdateRunBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{4}
}
func (m *UpdateRunBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRunBookRequest.Unmarshal(m, b)
}
func (m *UpdateRunBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRunBookRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRunBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRunBookRequest.Merge(m, src)
}
func (m *UpdateRunBookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRunBookRequest.Size(m)
}
func (m *UpdateRunBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRunBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRunBookRequest proto.InternalMessageInfo

func (m *UpdateRunBookRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateRunBookRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRunBookRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateRunBookRequest) GetSpec() *v1alpha1.RunBookSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateRunBookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRunBookResponse) Reset()         { *m = UpdateRunBookResponse{} }
func (m *UpdateRunBookResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateRunBookResponse) ProtoMessage()    {}
func (*UpdateRunBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{5}
}
func (m *UpdateRunBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRunBookResponse.Unmarshal(m, b)
}
func (m *UpdateRunBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRunBookResponse.Marshal(b, m, deterministic)
}
func (m *UpdateRunBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRunBookResponse.Merge(m, src)
}
func (m *UpdateRunBookResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateRunBookResponse.Size(m)
}
func (m *UpdateRunBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRunBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRunBookResponse proto.InternalMessageInfo

type GetRunBookRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRunBookRequest) Reset()         { *m = GetRunBookRequest{} }
func (m *GetRunBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetRunBookRequest) ProtoMessage()    {}
func (*GetRunBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{6}
}
func (m *GetRunBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunBookRequest.Unmarshal(m, b)
}
func (m *GetRunBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunBookRequest.Marshal(b, m, deterministic)
}
func (m *GetRunBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunBookRequest.Merge(m, src)
}
func (m *GetRunBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetRunBookRequest.Size(m)
}
func (m *GetRunBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRunBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRunBookRequest proto.InternalMessageInfo

func (m *GetRunBookRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetRunBookRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRunBookResponse struct {
	Item                 *v1alpha1.RunBook `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string            `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetRunBookResponse) Reset()         { *m = GetRunBookResponse{} }
func (m *GetRunBookResponse) String() string { return proto.CompactTextString(m) }
func (*GetRunBookResponse) ProtoMessage()    {}
func (*GetRunBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{7}
}
func (m *GetRunBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunBookResponse.Unmarshal(m, b)
}
func (m *GetRunBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunBookResponse.Marshal(b, m, deterministic)
}
func (m *GetRunBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunBookResponse.Merge(m, src)
}
func (m *GetRunBookResponse) XXX_Size() int {
	return xxx_messageInfo_GetRunBookResponse.Size(m)
}
func (m *GetRunBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRunBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRunBookResponse proto.InternalMessageInfo

func (m *GetRunBookResponse) GetItem() *v1alpha1.RunBook {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetRunBookResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type DeleteRunBookRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRunBookRequest) Reset()         { *m = DeleteRunBookRequest{} }
func (m *DeleteRunBookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRunBookRequest) ProtoMessage()    {}
func (*DeleteRunBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{8}
}
func (m *DeleteRunBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRunBookRequest.Unmarshal(m, b)
}
func (m *DeleteRunBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRunBookRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRunBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRunBookRequest.Merge(m, src)
}
func (m *DeleteRunBookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRunBookRequest.Size(m)
}
func (m *DeleteRunBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRunBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRunBookRequest proto.InternalMessageInfo

func (m *DeleteRunBookRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteRunBookRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteRunBookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRunBookResponse) Reset()         { *m = DeleteRunBookResponse{} }
func (m *DeleteRunBookResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteRunBookResponse) ProtoMessage()    {}
func (*DeleteRunBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31954e5121d41113, []int{9}
}
func (m *DeleteRunBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRunBookResponse.Unmarshal(m, b)
}
func (m *DeleteRunBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRunBookResponse.Marshal(b, m, deterministic)
}
func (m *DeleteRunBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRunBookResponse.Merge(m, src)
}
func (m *DeleteRunBookResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteRunBookResponse.Size(m)
}
func (m *DeleteRunBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRunBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRunBookResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListRunBooksRequest)(nil), "github.com.metaprov.modeld.services.runbook.ListRunBooksRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.runbook.ListRunBooksRequest.LabelsEntry")
	proto.RegisterType((*ListRunBooksResponse)(nil), "github.com.metaprov.modeld.services.runbook.ListRunBooksResponse")
	proto.RegisterType((*CreateRunBookRequest)(nil), "github.com.metaprov.modeld.services.runbook.CreateRunBookRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.runbook.CreateRunBookRequest.LabelsEntry")
	proto.RegisterType((*CreateRunBookResponse)(nil), "github.com.metaprov.modeld.services.runbook.CreateRunBookResponse")
	proto.RegisterType((*UpdateRunBookRequest)(nil), "github.com.metaprov.modeld.services.runbook.UpdateRunBookRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.runbook.UpdateRunBookRequest.LabelsEntry")
	proto.RegisterType((*UpdateRunBookResponse)(nil), "github.com.metaprov.modeld.services.runbook.UpdateRunBookResponse")
	proto.RegisterType((*GetRunBookRequest)(nil), "github.com.metaprov.modeld.services.runbook.GetRunBookRequest")
	proto.RegisterType((*GetRunBookResponse)(nil), "github.com.metaprov.modeld.services.runbook.GetRunBookResponse")
	proto.RegisterType((*DeleteRunBookRequest)(nil), "github.com.metaprov.modeld.services.runbook.DeleteRunBookRequest")
	proto.RegisterType((*DeleteRunBookResponse)(nil), "github.com.metaprov.modeld.services.runbook.DeleteRunBookResponse")
}

func init() { proto.RegisterFile("services/runbook/runbook.proto", fileDescriptor_31954e5121d41113) }

var fileDescriptor_31954e5121d41113 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcd, 0x6a, 0x14, 0x41,
	0x10, 0xc7, 0x9d, 0xcd, 0x26, 0x21, 0x95, 0x0f, 0xb4, 0xb3, 0x21, 0xcb, 0x18, 0x25, 0x0c, 0x1e,
	0x02, 0x42, 0x8f, 0x89, 0x08, 0x1a, 0xf0, 0x23, 0x31, 0x41, 0x0f, 0xf1, 0xe0, 0x4a, 0x10, 0xbc,
	0xf5, 0xee, 0x16, 0x93, 0x61, 0x67, 0xa6, 0xdb, 0xe9, 0xde, 0x85, 0x10, 0xbc, 0x78, 0xf6, 0xe6,
	0xc5, 0x57, 0xd0, 0xbb, 0x0f, 0xe0, 0x03, 0x78, 0xd2, 0x47, 0xf0, 0x41, 0xa4, 0x7b, 0x7a, 0xd8,
	0x9d, 0x75, 0x58, 0x1c, 0x47, 0xd0, 0x53, 0x7a, 0x3a, 0xd4, 0xbf, 0xfe, 0xf5, 0xeb, 0xaa, 0xde,
	0x86, 0xeb, 0x12, 0xd3, 0x51, 0xd8, 0x43, 0xe9, 0xa7, 0xc3, 0xa4, 0xcb, 0xf9, 0x20, 0xff, 0x4b,
	0x45, 0xca, 0x15, 0x27, 0x37, 0x83, 0x50, 0x9d, 0x0d, 0xbb, 0xb4, 0xc7, 0x63, 0x1a, 0xa3, 0x62,
	0x22, 0xe5, 0x23, 0x1a, 0xf3, 0x3e, 0x46, 0x7d, 0x9a, 0x87, 0x52, 0x1b, 0xe2, 0x6e, 0x05, 0x9c,
	0x07, 0x11, 0xfa, 0x4c, 0x84, 0x3e, 0x4b, 0x12, 0xae, 0x98, 0x0a, 0x79, 0x22, 0x33, 0x29, 0xf7,
	0x68, 0x2c, 0xe5, 0xe7, 0x52, 0x7e, 0x26, 0xa5, 0x03, 0xc4, 0x20, 0xd0, 0x81, 0xd2, 0x57, 0xc8,
	0x62, 0x7f, 0xb4, 0xcb, 0x22, 0x71, 0xc6, 0x76, 0xfd, 0x00, 0x13, 0x4c, 0x99, 0xc2, 0x7e, 0xa6,
	0xe2, 0x7d, 0x77, 0x60, 0xfd, 0x24, 0x94, 0xaa, 0x33, 0x4c, 0x0e, 0x39, 0x1f, 0xc8, 0x0e, 0xbe,
	0x1e, 0xa2, 0x54, 0x64, 0x0b, 0x96, 0x12, 0x16, 0xa3, 0x14, 0xac, 0x87, 0x6d, 0x67, 0xdb, 0xd9,
	0x59, 0xea, 0x8c, 0x37, 0x48, 0x1f, 0x16, 0x22, 0xd6, 0xc5, 0x48, 0xb6, 0xe7, 0xb6, 0xe7, 0x76,
	0x96, 0xf7, 0x4e, 0x68, 0x85, 0xba, 0x68, 0x49, 0x3e, 0x7a, 0x62, 0xe4, 0x8e, 0x13, 0x95, 0x9e,
	0x77, 0xac, 0xb6, 0x7b, 0x0f, 0x96, 0x27, 0xb6, 0xc9, 0x65, 0x98, 0x1b, 0xe0, 0xb9, 0x35, 0xa3,
	0x97, 0xa4, 0x05, 0xf3, 0x23, 0x16, 0x0d, 0xb1, 0xdd, 0x30, 0x7b, 0xd9, 0xc7, 0x7e, 0xe3, 0xae,
	0xe3, 0x71, 0x68, 0x15, 0xb3, 0x48, 0xc1, 0x13, 0x89, 0xe4, 0x25, 0xcc, 0x87, 0x0a, 0x63, 0x69,
	0x54, 0x96, 0xf7, 0x0e, 0x66, 0xf8, 0x66, 0x22, 0xa4, 0x62, 0x10, 0x50, 0x0d, 0x91, 0x6a, 0x88,
	0x34, 0x87, 0x48, 0xad, 0xac, 0xc9, 0x90, 0xe9, 0x79, 0x5f, 0x1a, 0xd0, 0x7a, 0x9c, 0x22, 0x53,
	0x68, 0xff, 0xf9, 0x7b, 0x20, 0x09, 0x34, 0xf5, 0x87, 0x2d, 0xc0, 0xac, 0x09, 0x4e, 0xc1, 0x7d,
	0x56, 0x09, 0x6e, 0x99, 0x89, 0x32, 0xba, 0xe4, 0x14, 0x9a, 0x52, 0x60, 0xaf, 0xdd, 0xfc, 0x0b,
	0x24, 0x5e, 0x08, 0xec, 0x75, 0x8c, 0x5c, 0x9d, 0x43, 0xdb, 0x84, 0x8d, 0x29, 0xf7, 0xd9, 0xa9,
	0x19, 0xb8, 0xa7, 0xa2, 0xff, 0xef, 0xe1, 0x96, 0x99, 0x98, 0x09, 0x77, 0xfe, 0x7f, 0x82, 0x3b,
	0xe5, 0xde, 0xc2, 0x3d, 0x86, 0x2b, 0x4f, 0x50, 0xd5, 0x05, 0xeb, 0x5d, 0x00, 0x99, 0x94, 0xb1,
	0xf3, 0xf6, 0x1c, 0x9a, 0x7a, 0x3e, 0xec, 0xb8, 0xdd, 0xaf, 0xc5, 0xa1, 0x63, 0xa4, 0x74, 0xf2,
	0x73, 0x16, 0x47, 0x79, 0x72, 0xbd, 0xf6, 0x9e, 0x42, 0xeb, 0x08, 0x23, 0xac, 0xdf, 0x1f, 0x1a,
	0xd3, 0x94, 0x52, 0x56, 0xc9, 0xde, 0xbb, 0x45, 0x58, 0xcb, 0x0f, 0x24, 0x6b, 0x0b, 0xf2, 0xd1,
	0x81, 0x95, 0xc9, 0x5b, 0x86, 0x3c, 0xaa, 0x7b, 0x0d, 0xba, 0x07, 0x35, 0x14, 0xec, 0x79, 0xb6,
	0xde, 0x7e, 0xfb, 0xf1, 0xbe, 0xb1, 0x46, 0x56, 0xfc, 0xd1, 0x6e, 0xfe, 0xeb, 0x23, 0xc9, 0x67,
	0x07, 0x56, 0x0b, 0xc3, 0x45, 0x0e, 0x6a, 0x5f, 0x2b, 0xee, 0x61, 0x1d, 0x09, 0x6b, 0xf7, 0x9a,
	0xb1, 0xbb, 0xe9, 0x15, 0xec, 0xee, 0x2f, 0xda, 0x15, 0xf9, 0xe4, 0x00, 0x8c, 0xfb, 0x8a, 0x3c,
	0xa8, 0x94, 0xf1, 0x97, 0xbe, 0x76, 0x1f, 0xfe, 0x71, 0xbc, 0xb5, 0x7b, 0xd5, 0xd8, 0xdd, 0x20,
	0xeb, 0x93, 0x76, 0xfd, 0x0b, 0xdd, 0x3b, 0x6f, 0xc8, 0x57, 0x07, 0x56, 0x0b, 0x43, 0x56, 0x11,
	0x72, 0xd9, 0xf5, 0x52, 0x11, 0x72, 0xf9, 0x8c, 0xdf, 0x31, 0xae, 0x7d, 0xf7, 0x46, 0xd1, 0x75,
	0x1e, 0xa4, 0x55, 0xfb, 0x4c, 0x31, 0x6a, 0xca, 0x18, 0xc3, 0xff, 0xe0, 0xc0, 0x6a, 0x61, 0x1a,
	0x2a, 0xd6, 0x53, 0x36, 0x93, 0x15, 0xeb, 0x29, 0x1d, 0x46, 0xef, 0xd2, 0xe1, 0xad, 0x57, 0x74,
	0xf6, 0xfb, 0x67, 0xfa, 0x21, 0xd6, 0x5d, 0x30, 0x0f, 0x9e, 0xdb, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0x3f, 0xaf, 0x1b, 0xa3, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RunBookServiceClient is the client API for RunBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunBookServiceClient interface {
	ListRunBooks(ctx context.Context, in *ListRunBooksRequest, opts ...grpc.CallOption) (*ListRunBooksResponse, error)
	CreateRunBook(ctx context.Context, in *CreateRunBookRequest, opts ...grpc.CallOption) (*CreateRunBookResponse, error)
	GetRunBook(ctx context.Context, in *GetRunBookRequest, opts ...grpc.CallOption) (*GetRunBookResponse, error)
	UpdateRunBook(ctx context.Context, in *UpdateRunBookRequest, opts ...grpc.CallOption) (*UpdateRunBookResponse, error)
	DeleteRunBook(ctx context.Context, in *DeleteRunBookRequest, opts ...grpc.CallOption) (*DeleteRunBookResponse, error)
}

type runBookServiceClient struct {
	cc *grpc.ClientConn
}

func NewRunBookServiceClient(cc *grpc.ClientConn) RunBookServiceClient {
	return &runBookServiceClient{cc}
}

func (c *runBookServiceClient) ListRunBooks(ctx context.Context, in *ListRunBooksRequest, opts ...grpc.CallOption) (*ListRunBooksResponse, error) {
	out := new(ListRunBooksResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.runbook.RunBookService/ListRunBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runBookServiceClient) CreateRunBook(ctx context.Context, in *CreateRunBookRequest, opts ...grpc.CallOption) (*CreateRunBookResponse, error) {
	out := new(CreateRunBookResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.runbook.RunBookService/CreateRunBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runBookServiceClient) GetRunBook(ctx context.Context, in *GetRunBookRequest, opts ...grpc.CallOption) (*GetRunBookResponse, error) {
	out := new(GetRunBookResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.runbook.RunBookService/GetRunBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runBookServiceClient) UpdateRunBook(ctx context.Context, in *UpdateRunBookRequest, opts ...grpc.CallOption) (*UpdateRunBookResponse, error) {
	out := new(UpdateRunBookResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.runbook.RunBookService/UpdateRunBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runBookServiceClient) DeleteRunBook(ctx context.Context, in *DeleteRunBookRequest, opts ...grpc.CallOption) (*DeleteRunBookResponse, error) {
	out := new(DeleteRunBookResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.runbook.RunBookService/DeleteRunBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunBookServiceServer is the server API for RunBookService service.
type RunBookServiceServer interface {
	ListRunBooks(context.Context, *ListRunBooksRequest) (*ListRunBooksResponse, error)
	CreateRunBook(context.Context, *CreateRunBookRequest) (*CreateRunBookResponse, error)
	GetRunBook(context.Context, *GetRunBookRequest) (*GetRunBookResponse, error)
	UpdateRunBook(context.Context, *UpdateRunBookRequest) (*UpdateRunBookResponse, error)
	DeleteRunBook(context.Context, *DeleteRunBookRequest) (*DeleteRunBookResponse, error)
}

// UnimplementedRunBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRunBookServiceServer struct {
}

func (*UnimplementedRunBookServiceServer) ListRunBooks(ctx context.Context, req *ListRunBooksRequest) (*ListRunBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRunBooks not implemented")
}
func (*UnimplementedRunBookServiceServer) CreateRunBook(ctx context.Context, req *CreateRunBookRequest) (*CreateRunBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRunBook not implemented")
}
func (*UnimplementedRunBookServiceServer) GetRunBook(ctx context.Context, req *GetRunBookRequest) (*GetRunBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunBook not implemented")
}
func (*UnimplementedRunBookServiceServer) UpdateRunBook(ctx context.Context, req *UpdateRunBookRequest) (*UpdateRunBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRunBook not implemented")
}
func (*UnimplementedRunBookServiceServer) DeleteRunBook(ctx context.Context, req *DeleteRunBookRequest) (*DeleteRunBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRunBook not implemented")
}

func RegisterRunBookServiceServer(s *grpc.Server, srv RunBookServiceServer) {
	s.RegisterService(&_RunBookService_serviceDesc, srv)
}

func _RunBookService_ListRunBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRunBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunBookServiceServer).ListRunBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.runbook.RunBookService/ListRunBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunBookServiceServer).ListRunBooks(ctx, req.(*ListRunBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunBookService_CreateRunBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRunBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunBookServiceServer).CreateRunBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.runbook.RunBookService/CreateRunBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunBookServiceServer).CreateRunBook(ctx, req.(*CreateRunBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunBookService_GetRunBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunBookServiceServer).GetRunBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.runbook.RunBookService/GetRunBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunBookServiceServer).GetRunBook(ctx, req.(*GetRunBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunBookService_UpdateRunBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRunBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunBookServiceServer).UpdateRunBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.runbook.RunBookService/UpdateRunBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunBookServiceServer).UpdateRunBook(ctx, req.(*UpdateRunBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunBookService_DeleteRunBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRunBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunBookServiceServer).DeleteRunBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.runbook.RunBookService/DeleteRunBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunBookServiceServer).DeleteRunBook(ctx, req.(*DeleteRunBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RunBookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.runbook.RunBookService",
	HandlerType: (*RunBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRunBooks",
			Handler:    _RunBookService_ListRunBooks_Handler,
		},
		{
			MethodName: "CreateRunBook",
			Handler:    _RunBookService_CreateRunBook_Handler,
		},
		{
			MethodName: "GetRunBook",
			Handler:    _RunBookService_GetRunBook_Handler,
		},
		{
			MethodName: "UpdateRunBook",
			Handler:    _RunBookService_UpdateRunBook_Handler,
		},
		{
			MethodName: "DeleteRunBook",
			Handler:    _RunBookService_DeleteRunBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/runbook/runbook.proto",
}
