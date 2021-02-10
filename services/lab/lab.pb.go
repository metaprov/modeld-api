// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/lab/lab.proto

package lab

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/infra/v1alpha1"
	common "github.com/metaprov/modeldapi/services/common"
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

type ListLabsRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListLabsRequest) Reset()         { *m = ListLabsRequest{} }
func (m *ListLabsRequest) String() string { return proto.CompactTextString(m) }
func (*ListLabsRequest) ProtoMessage()    {}
func (*ListLabsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{0}
}
func (m *ListLabsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabsRequest.Unmarshal(m, b)
}
func (m *ListLabsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabsRequest.Marshal(b, m, deterministic)
}
func (m *ListLabsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabsRequest.Merge(m, src)
}
func (m *ListLabsRequest) XXX_Size() int {
	return xxx_messageInfo_ListLabsRequest.Size(m)
}
func (m *ListLabsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabsRequest proto.InternalMessageInfo

func (m *ListLabsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListLabsRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListLabsResponse struct {
	Items                *v1alpha1.LabList `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListLabsResponse) Reset()         { *m = ListLabsResponse{} }
func (m *ListLabsResponse) String() string { return proto.CompactTextString(m) }
func (*ListLabsResponse) ProtoMessage()    {}
func (*ListLabsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{1}
}
func (m *ListLabsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabsResponse.Unmarshal(m, b)
}
func (m *ListLabsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabsResponse.Marshal(b, m, deterministic)
}
func (m *ListLabsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabsResponse.Merge(m, src)
}
func (m *ListLabsResponse) XXX_Size() int {
	return xxx_messageInfo_ListLabsResponse.Size(m)
}
func (m *ListLabsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabsResponse proto.InternalMessageInfo

func (m *ListLabsResponse) GetItems() *v1alpha1.LabList {
	if m != nil {
		return m.Items
	}
	return nil
}

type LabResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabResponse) Reset()         { *m = LabResponse{} }
func (m *LabResponse) String() string { return proto.CompactTextString(m) }
func (*LabResponse) ProtoMessage()    {}
func (*LabResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{2}
}
func (m *LabResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabResponse.Unmarshal(m, b)
}
func (m *LabResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabResponse.Marshal(b, m, deterministic)
}
func (m *LabResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabResponse.Merge(m, src)
}
func (m *LabResponse) XXX_Size() int {
	return xxx_messageInfo_LabResponse.Size(m)
}
func (m *LabResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LabResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LabResponse proto.InternalMessageInfo

type CreateLabRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.LabSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateLabRequest) Reset()         { *m = CreateLabRequest{} }
func (m *CreateLabRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLabRequest) ProtoMessage()    {}
func (*CreateLabRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{3}
}
func (m *CreateLabRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLabRequest.Unmarshal(m, b)
}
func (m *CreateLabRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLabRequest.Marshal(b, m, deterministic)
}
func (m *CreateLabRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLabRequest.Merge(m, src)
}
func (m *CreateLabRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLabRequest.Size(m)
}
func (m *CreateLabRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLabRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLabRequest proto.InternalMessageInfo

func (m *CreateLabRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateLabRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLabRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateLabRequest) GetSpec() *v1alpha1.LabSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type CreateLabResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLabResponse) Reset()         { *m = CreateLabResponse{} }
func (m *CreateLabResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLabResponse) ProtoMessage()    {}
func (*CreateLabResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{4}
}
func (m *CreateLabResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLabResponse.Unmarshal(m, b)
}
func (m *CreateLabResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLabResponse.Marshal(b, m, deterministic)
}
func (m *CreateLabResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLabResponse.Merge(m, src)
}
func (m *CreateLabResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLabResponse.Size(m)
}
func (m *CreateLabResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLabResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLabResponse proto.InternalMessageInfo

type UpdateLabRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.LabSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateLabRequest) Reset()         { *m = UpdateLabRequest{} }
func (m *UpdateLabRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateLabRequest) ProtoMessage()    {}
func (*UpdateLabRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{5}
}
func (m *UpdateLabRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLabRequest.Unmarshal(m, b)
}
func (m *UpdateLabRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLabRequest.Marshal(b, m, deterministic)
}
func (m *UpdateLabRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLabRequest.Merge(m, src)
}
func (m *UpdateLabRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateLabRequest.Size(m)
}
func (m *UpdateLabRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLabRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLabRequest proto.InternalMessageInfo

func (m *UpdateLabRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateLabRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateLabRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateLabRequest) GetSpec() *v1alpha1.LabSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateLabResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateLabResponse) Reset()         { *m = UpdateLabResponse{} }
func (m *UpdateLabResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateLabResponse) ProtoMessage()    {}
func (*UpdateLabResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{6}
}
func (m *UpdateLabResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLabResponse.Unmarshal(m, b)
}
func (m *UpdateLabResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLabResponse.Marshal(b, m, deterministic)
}
func (m *UpdateLabResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLabResponse.Merge(m, src)
}
func (m *UpdateLabResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateLabResponse.Size(m)
}
func (m *UpdateLabResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLabResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLabResponse proto.InternalMessageInfo

type GetLabRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLabRequest) Reset()         { *m = GetLabRequest{} }
func (m *GetLabRequest) String() string { return proto.CompactTextString(m) }
func (*GetLabRequest) ProtoMessage()    {}
func (*GetLabRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{7}
}
func (m *GetLabRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabRequest.Unmarshal(m, b)
}
func (m *GetLabRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabRequest.Marshal(b, m, deterministic)
}
func (m *GetLabRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabRequest.Merge(m, src)
}
func (m *GetLabRequest) XXX_Size() int {
	return xxx_messageInfo_GetLabRequest.Size(m)
}
func (m *GetLabRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabRequest proto.InternalMessageInfo

func (m *GetLabRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetLabRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetLabResponse struct {
	Item                 *v1alpha1.Lab `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string        `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetLabResponse) Reset()         { *m = GetLabResponse{} }
func (m *GetLabResponse) String() string { return proto.CompactTextString(m) }
func (*GetLabResponse) ProtoMessage()    {}
func (*GetLabResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{8}
}
func (m *GetLabResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabResponse.Unmarshal(m, b)
}
func (m *GetLabResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabResponse.Marshal(b, m, deterministic)
}
func (m *GetLabResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabResponse.Merge(m, src)
}
func (m *GetLabResponse) XXX_Size() int {
	return xxx_messageInfo_GetLabResponse.Size(m)
}
func (m *GetLabResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabResponse proto.InternalMessageInfo

func (m *GetLabResponse) GetItem() *v1alpha1.Lab {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetLabResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type GetLabNamespacesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLabNamespacesRequest) Reset()         { *m = GetLabNamespacesRequest{} }
func (m *GetLabNamespacesRequest) String() string { return proto.CompactTextString(m) }
func (*GetLabNamespacesRequest) ProtoMessage()    {}
func (*GetLabNamespacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{9}
}
func (m *GetLabNamespacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabNamespacesRequest.Unmarshal(m, b)
}
func (m *GetLabNamespacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabNamespacesRequest.Marshal(b, m, deterministic)
}
func (m *GetLabNamespacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabNamespacesRequest.Merge(m, src)
}
func (m *GetLabNamespacesRequest) XXX_Size() int {
	return xxx_messageInfo_GetLabNamespacesRequest.Size(m)
}
func (m *GetLabNamespacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabNamespacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabNamespacesRequest proto.InternalMessageInfo

func (m *GetLabNamespacesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetLabNamespacesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetLabNamespacesResponse struct {
	Namespaces           []*common.NamespaceInfo `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetLabNamespacesResponse) Reset()         { *m = GetLabNamespacesResponse{} }
func (m *GetLabNamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*GetLabNamespacesResponse) ProtoMessage()    {}
func (*GetLabNamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{10}
}
func (m *GetLabNamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabNamespacesResponse.Unmarshal(m, b)
}
func (m *GetLabNamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabNamespacesResponse.Marshal(b, m, deterministic)
}
func (m *GetLabNamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabNamespacesResponse.Merge(m, src)
}
func (m *GetLabNamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_GetLabNamespacesResponse.Size(m)
}
func (m *GetLabNamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabNamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabNamespacesResponse proto.InternalMessageInfo

func (m *GetLabNamespacesResponse) GetNamespaces() []*common.NamespaceInfo {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type DeleteLabRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLabRequest) Reset()         { *m = DeleteLabRequest{} }
func (m *DeleteLabRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLabRequest) ProtoMessage()    {}
func (*DeleteLabRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{11}
}
func (m *DeleteLabRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLabRequest.Unmarshal(m, b)
}
func (m *DeleteLabRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLabRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLabRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLabRequest.Merge(m, src)
}
func (m *DeleteLabRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLabRequest.Size(m)
}
func (m *DeleteLabRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLabRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLabRequest proto.InternalMessageInfo

func (m *DeleteLabRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteLabRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteLabResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLabResponse) Reset()         { *m = DeleteLabResponse{} }
func (m *DeleteLabResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteLabResponse) ProtoMessage()    {}
func (*DeleteLabResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd47e739fd646341, []int{12}
}
func (m *DeleteLabResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLabResponse.Unmarshal(m, b)
}
func (m *DeleteLabResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLabResponse.Marshal(b, m, deterministic)
}
func (m *DeleteLabResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLabResponse.Merge(m, src)
}
func (m *DeleteLabResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteLabResponse.Size(m)
}
func (m *DeleteLabResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLabResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLabResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListLabsRequest)(nil), "github.com.metaprov.modeld.services.lab.ListLabsRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.lab.ListLabsRequest.LabelsEntry")
	proto.RegisterType((*ListLabsResponse)(nil), "github.com.metaprov.modeld.services.lab.ListLabsResponse")
	proto.RegisterType((*LabResponse)(nil), "github.com.metaprov.modeld.services.lab.LabResponse")
	proto.RegisterType((*CreateLabRequest)(nil), "github.com.metaprov.modeld.services.lab.CreateLabRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.lab.CreateLabRequest.LabelsEntry")
	proto.RegisterType((*CreateLabResponse)(nil), "github.com.metaprov.modeld.services.lab.CreateLabResponse")
	proto.RegisterType((*UpdateLabRequest)(nil), "github.com.metaprov.modeld.services.lab.UpdateLabRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.lab.UpdateLabRequest.LabelsEntry")
	proto.RegisterType((*UpdateLabResponse)(nil), "github.com.metaprov.modeld.services.lab.UpdateLabResponse")
	proto.RegisterType((*GetLabRequest)(nil), "github.com.metaprov.modeld.services.lab.GetLabRequest")
	proto.RegisterType((*GetLabResponse)(nil), "github.com.metaprov.modeld.services.lab.GetLabResponse")
	proto.RegisterType((*GetLabNamespacesRequest)(nil), "github.com.metaprov.modeld.services.lab.GetLabNamespacesRequest")
	proto.RegisterType((*GetLabNamespacesResponse)(nil), "github.com.metaprov.modeld.services.lab.GetLabNamespacesResponse")
	proto.RegisterType((*DeleteLabRequest)(nil), "github.com.metaprov.modeld.services.lab.DeleteLabRequest")
	proto.RegisterType((*DeleteLabResponse)(nil), "github.com.metaprov.modeld.services.lab.DeleteLabResponse")
}

func init() { proto.RegisterFile("services/lab/lab.proto", fileDescriptor_dd47e739fd646341) }

var fileDescriptor_dd47e739fd646341 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe5, 0x24, 0x8d, 0x9a, 0xe9, 0xaf, 0x6d, 0x7e, 0xdb, 0x16, 0x22, 0xd3, 0x43, 0xe4,
	0x0b, 0x55, 0x85, 0xd6, 0x6a, 0x91, 0xa0, 0xc9, 0x01, 0x09, 0x68, 0x85, 0x10, 0x55, 0x0f, 0x06,
	0x0e, 0x20, 0x10, 0x5a, 0x27, 0xd3, 0xd4, 0xaa, 0xed, 0x35, 0xde, 0x4d, 0xa4, 0x0a, 0x71, 0xe1,
	0x15, 0x10, 0x07, 0x4e, 0x1c, 0xe0, 0x8e, 0xc4, 0x53, 0x70, 0xe7, 0x15, 0x78, 0x10, 0xb4, 0xeb,
	0x75, 0xea, 0xa6, 0x55, 0x94, 0x3f, 0x07, 0x38, 0x54, 0x5d, 0xaf, 0x32, 0xdf, 0xf9, 0xec, 0x77,
	0x67, 0xc6, 0x86, 0x6b, 0x02, 0xd3, 0x41, 0xd0, 0x41, 0xe1, 0x86, 0xcc, 0x57, 0x7f, 0x34, 0x49,
	0xb9, 0xe4, 0xe4, 0x66, 0x2f, 0x90, 0x27, 0x7d, 0x9f, 0x76, 0x78, 0x44, 0x23, 0x94, 0x2c, 0x49,
	0xf9, 0x80, 0x46, 0xbc, 0x8b, 0x61, 0x97, 0xe6, 0x21, 0x34, 0x64, 0xbe, 0xbd, 0xd9, 0xe3, 0xbc,
	0x17, 0xa2, 0xcb, 0x92, 0xc0, 0x65, 0x71, 0xcc, 0x25, 0x93, 0x01, 0x8f, 0x45, 0x26, 0x63, 0x1f,
	0x9c, 0xcb, 0xb8, 0xb9, 0x8c, 0x9b, 0xc9, 0xa8, 0x80, 0xe4, 0xb4, 0xa7, 0x02, 0x85, 0x1b, 0xc4,
	0xc7, 0x29, 0x73, 0x07, 0x3b, 0x2c, 0x4c, 0x4e, 0xd8, 0x8e, 0xdb, 0xc3, 0x18, 0x53, 0x26, 0xb1,
	0x6b, 0x64, 0xda, 0xe3, 0x65, 0x86, 0x67, 0xe8, 0xf0, 0x28, 0xe2, 0xb1, 0xf9, 0x97, 0xc5, 0x3a,
	0x3f, 0x2d, 0x58, 0x3d, 0x0c, 0x84, 0x3c, 0x64, 0xbe, 0xf0, 0xf0, 0x6d, 0x1f, 0x85, 0x24, 0x9b,
	0x50, 0x8b, 0x59, 0x84, 0x22, 0x61, 0x1d, 0x6c, 0x58, 0x4d, 0x6b, 0xab, 0xe6, 0x9d, 0x6f, 0x90,
	0x57, 0x50, 0x0d, 0x99, 0x8f, 0xa1, 0x68, 0x94, 0x9b, 0xe5, 0xad, 0xa5, 0xdd, 0x7d, 0x3a, 0xa1,
	0x19, 0x74, 0x24, 0x0f, 0x3d, 0xd4, 0x32, 0x07, 0xb1, 0x4c, 0xcf, 0x3c, 0xa3, 0x69, 0xb7, 0x60,
	0xa9, 0xb0, 0x4d, 0xea, 0x50, 0x3e, 0xc5, 0x33, 0x03, 0xa1, 0x96, 0x64, 0x1d, 0x16, 0x06, 0x2c,
	0xec, 0x63, 0xa3, 0xa4, 0xf7, 0xb2, 0x87, 0x76, 0x69, 0xcf, 0x72, 0x4e, 0xa0, 0x7e, 0x9e, 0x41,
	0x24, 0x3c, 0x16, 0x48, 0x9e, 0xc1, 0x42, 0x20, 0x31, 0x12, 0x5a, 0x61, 0x69, 0xf7, 0xde, 0x18,
	0x56, 0x96, 0x04, 0x34, 0x39, 0xed, 0x51, 0xe5, 0x38, 0xd5, 0x8e, 0xd3, 0xdc, 0x71, 0x45, 0xaa,
	0xa4, 0xbd, 0x4c, 0xcc, 0x59, 0xd6, 0x90, 0x79, 0x12, 0xe7, 0x47, 0x09, 0xea, 0x0f, 0x53, 0x64,
	0x12, 0xf5, 0xee, 0x24, 0x26, 0x12, 0xa8, 0xa8, 0x07, 0x73, 0x08, 0xbd, 0x26, 0xaf, 0x47, 0x8c,
	0x3d, 0x98, 0xd8, 0xd8, 0xd1, 0xe4, 0x57, 0x39, 0x4b, 0x3c, 0xa8, 0x88, 0x04, 0x3b, 0x8d, 0xca,
	0xbc, 0x4e, 0x3c, 0x4d, 0xb0, 0xe3, 0x69, 0xad, 0x79, 0x6e, 0x6b, 0x0d, 0xfe, 0x2f, 0x60, 0x17,
	0x9c, 0x7c, 0x9e, 0x74, 0xff, 0x9e, 0x93, 0xa3, 0xc9, 0xc7, 0x3a, 0xb9, 0xf0, 0xcf, 0x38, 0x59,
	0xc0, 0x36, 0x4e, 0xde, 0x87, 0xe5, 0x47, 0x28, 0xe7, 0x71, 0xd1, 0x91, 0xb0, 0x92, 0x4b, 0x98,
	0x6e, 0x3a, 0x82, 0x8a, 0x6a, 0x00, 0xd3, 0x4c, 0xed, 0xd9, 0x0f, 0xee, 0x69, 0x1d, 0x95, 0xf5,
	0x8c, 0x45, 0x61, 0x9e, 0x55, 0xad, 0x9d, 0x27, 0x70, 0x3d, 0xcb, 0x7a, 0x94, 0xc3, 0x89, 0xd9,
	0x8f, 0xd0, 0x87, 0xc6, 0x65, 0x31, 0x73, 0x98, 0x17, 0x00, 0xc3, 0x60, 0x35, 0x1f, 0x54, 0xa1,
	0xb4, 0x26, 0x2a, 0x14, 0x33, 0x40, 0x87, 0x9a, 0x8f, 0xe3, 0x63, 0xee, 0x15, 0xc4, 0x9c, 0x7d,
	0xa8, 0xef, 0x63, 0x88, 0xf3, 0x55, 0xb1, 0xba, 0xd7, 0x82, 0x4a, 0x46, 0xbd, 0xfb, 0xbd, 0x0a,
	0xa0, 0x2a, 0x27, 0x63, 0x21, 0x9f, 0x2c, 0x58, 0xcc, 0x87, 0x1e, 0xd9, 0x9b, 0x75, 0x12, 0xdb,
	0xad, 0x19, 0x22, 0x4d, 0xa1, 0xad, 0x7f, 0xf8, 0xf5, 0xfb, 0x63, 0x69, 0x85, 0xfc, 0xa7, 0xdf,
	0x71, 0x83, 0x1d, 0xf5, 0x96, 0x14, 0xe4, 0xb3, 0x05, 0xb5, 0x61, 0x7b, 0x93, 0xd6, 0xcc, 0x93,
	0xcc, 0x6e, 0xcf, 0x12, 0x6a, 0xd0, 0x36, 0x34, 0xda, 0xaa, 0xb3, 0x98, 0x63, 0xb5, 0xcb, 0x21,
	0xf3, 0xc9, 0x57, 0x0b, 0xaa, 0x59, 0x55, 0x90, 0x3b, 0x13, 0xab, 0x5f, 0x68, 0x26, 0xfb, 0xee,
	0xd4, 0x71, 0x06, 0x69, 0x4b, 0x23, 0x39, 0xa4, 0x59, 0x74, 0xcb, 0x7d, 0xa7, 0x7e, 0xa9, 0x64,
	0xba, 0x4c, 0x32, 0xaa, 0x2e, 0xff, 0x3d, 0xf9, 0x66, 0x41, 0x6d, 0xd8, 0xd6, 0x53, 0x38, 0x38,
	0x3a, 0xc1, 0xa6, 0x70, 0xf0, 0xf2, 0x14, 0x69, 0x6a, 0x5c, 0xdb, 0x5e, 0xbb, 0x80, 0xfa, 0x46,
	0x13, 0x66, 0x66, 0x7e, 0xb1, 0xa0, 0x36, 0xac, 0xd2, 0x29, 0x30, 0x47, 0xfb, 0x63, 0x0a, 0xcc,
	0x4b, 0x4d, 0xe1, 0xdc, 0xd0, 0x98, 0x1b, 0xdb, 0x57, 0x61, 0x3e, 0xb8, 0xf5, 0x72, 0x7b, 0xc2,
	0xef, 0xa3, 0x90, 0xf9, 0x7e, 0x55, 0x7f, 0x16, 0xdd, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdc,
	0x3f, 0x95, 0x0e, 0xfa, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LabServiceClient is the client API for LabService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabServiceClient interface {
	ListLabs(ctx context.Context, in *ListLabsRequest, opts ...grpc.CallOption) (*ListLabsResponse, error)
	CreateLab(ctx context.Context, in *CreateLabRequest, opts ...grpc.CallOption) (*CreateLabResponse, error)
	GetLab(ctx context.Context, in *GetLabRequest, opts ...grpc.CallOption) (*GetLabResponse, error)
	UpdateLab(ctx context.Context, in *UpdateLabRequest, opts ...grpc.CallOption) (*UpdateLabResponse, error)
	DeleteLab(ctx context.Context, in *DeleteLabRequest, opts ...grpc.CallOption) (*DeleteLabResponse, error)
}

type labServiceClient struct {
	cc *grpc.ClientConn
}

func NewLabServiceClient(cc *grpc.ClientConn) LabServiceClient {
	return &labServiceClient{cc}
}

func (c *labServiceClient) ListLabs(ctx context.Context, in *ListLabsRequest, opts ...grpc.CallOption) (*ListLabsResponse, error) {
	out := new(ListLabsResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.lab.LabService/ListLabs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labServiceClient) CreateLab(ctx context.Context, in *CreateLabRequest, opts ...grpc.CallOption) (*CreateLabResponse, error) {
	out := new(CreateLabResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.lab.LabService/CreateLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labServiceClient) GetLab(ctx context.Context, in *GetLabRequest, opts ...grpc.CallOption) (*GetLabResponse, error) {
	out := new(GetLabResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.lab.LabService/GetLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labServiceClient) UpdateLab(ctx context.Context, in *UpdateLabRequest, opts ...grpc.CallOption) (*UpdateLabResponse, error) {
	out := new(UpdateLabResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.lab.LabService/UpdateLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labServiceClient) DeleteLab(ctx context.Context, in *DeleteLabRequest, opts ...grpc.CallOption) (*DeleteLabResponse, error) {
	out := new(DeleteLabResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.lab.LabService/DeleteLab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabServiceServer is the server API for LabService service.
type LabServiceServer interface {
	ListLabs(context.Context, *ListLabsRequest) (*ListLabsResponse, error)
	CreateLab(context.Context, *CreateLabRequest) (*CreateLabResponse, error)
	GetLab(context.Context, *GetLabRequest) (*GetLabResponse, error)
	UpdateLab(context.Context, *UpdateLabRequest) (*UpdateLabResponse, error)
	DeleteLab(context.Context, *DeleteLabRequest) (*DeleteLabResponse, error)
}

// UnimplementedLabServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLabServiceServer struct {
}

func (*UnimplementedLabServiceServer) ListLabs(ctx context.Context, req *ListLabsRequest) (*ListLabsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabs not implemented")
}
func (*UnimplementedLabServiceServer) CreateLab(ctx context.Context, req *CreateLabRequest) (*CreateLabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLab not implemented")
}
func (*UnimplementedLabServiceServer) GetLab(ctx context.Context, req *GetLabRequest) (*GetLabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLab not implemented")
}
func (*UnimplementedLabServiceServer) UpdateLab(ctx context.Context, req *UpdateLabRequest) (*UpdateLabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLab not implemented")
}
func (*UnimplementedLabServiceServer) DeleteLab(ctx context.Context, req *DeleteLabRequest) (*DeleteLabResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLab not implemented")
}

func RegisterLabServiceServer(s *grpc.Server, srv LabServiceServer) {
	s.RegisterService(&_LabService_serviceDesc, srv)
}

func _LabService_ListLabs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabServiceServer).ListLabs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.lab.LabService/ListLabs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabServiceServer).ListLabs(ctx, req.(*ListLabsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabService_CreateLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabServiceServer).CreateLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.lab.LabService/CreateLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabServiceServer).CreateLab(ctx, req.(*CreateLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabService_GetLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabServiceServer).GetLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.lab.LabService/GetLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabServiceServer).GetLab(ctx, req.(*GetLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabService_UpdateLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabServiceServer).UpdateLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.lab.LabService/UpdateLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabServiceServer).UpdateLab(ctx, req.(*UpdateLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabService_DeleteLab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLabRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabServiceServer).DeleteLab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.lab.LabService/DeleteLab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabServiceServer).DeleteLab(ctx, req.(*DeleteLabRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.lab.LabService",
	HandlerType: (*LabServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLabs",
			Handler:    _LabService_ListLabs_Handler,
		},
		{
			MethodName: "CreateLab",
			Handler:    _LabService_CreateLab_Handler,
		},
		{
			MethodName: "GetLab",
			Handler:    _LabService_GetLab_Handler,
		},
		{
			MethodName: "UpdateLab",
			Handler:    _LabService_UpdateLab_Handler,
		},
		{
			MethodName: "DeleteLab",
			Handler:    _LabService_DeleteLab_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/lab/lab.proto",
}
