// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/curtain/curtain.proto

package curtain

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/inference/v1alpha1"
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

type ListCurtainsRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListCurtainsRequest) Reset()         { *m = ListCurtainsRequest{} }
func (m *ListCurtainsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCurtainsRequest) ProtoMessage()    {}
func (*ListCurtainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{0}
}
func (m *ListCurtainsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCurtainsRequest.Unmarshal(m, b)
}
func (m *ListCurtainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCurtainsRequest.Marshal(b, m, deterministic)
}
func (m *ListCurtainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCurtainsRequest.Merge(m, src)
}
func (m *ListCurtainsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCurtainsRequest.Size(m)
}
func (m *ListCurtainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCurtainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCurtainsRequest proto.InternalMessageInfo

func (m *ListCurtainsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListCurtainsRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListCurtainsResponse struct {
	Items                *v1alpha1.CurtainList `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListCurtainsResponse) Reset()         { *m = ListCurtainsResponse{} }
func (m *ListCurtainsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCurtainsResponse) ProtoMessage()    {}
func (*ListCurtainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{1}
}
func (m *ListCurtainsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCurtainsResponse.Unmarshal(m, b)
}
func (m *ListCurtainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCurtainsResponse.Marshal(b, m, deterministic)
}
func (m *ListCurtainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCurtainsResponse.Merge(m, src)
}
func (m *ListCurtainsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCurtainsResponse.Size(m)
}
func (m *ListCurtainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCurtainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCurtainsResponse proto.InternalMessageInfo

func (m *ListCurtainsResponse) GetItems() *v1alpha1.CurtainList {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateCurtainResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCurtainResponse) Reset()         { *m = CreateCurtainResponse{} }
func (m *CreateCurtainResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCurtainResponse) ProtoMessage()    {}
func (*CreateCurtainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{2}
}
func (m *CreateCurtainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCurtainResponse.Unmarshal(m, b)
}
func (m *CreateCurtainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCurtainResponse.Marshal(b, m, deterministic)
}
func (m *CreateCurtainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCurtainResponse.Merge(m, src)
}
func (m *CreateCurtainResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCurtainResponse.Size(m)
}
func (m *CreateCurtainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCurtainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCurtainResponse proto.InternalMessageInfo

type CreateCurtainRequest struct {
	Namespace            string                `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string     `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.CurtainSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCurtainRequest) Reset()         { *m = CreateCurtainRequest{} }
func (m *CreateCurtainRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCurtainRequest) ProtoMessage()    {}
func (*CreateCurtainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{3}
}
func (m *CreateCurtainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCurtainRequest.Unmarshal(m, b)
}
func (m *CreateCurtainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCurtainRequest.Marshal(b, m, deterministic)
}
func (m *CreateCurtainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCurtainRequest.Merge(m, src)
}
func (m *CreateCurtainRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCurtainRequest.Size(m)
}
func (m *CreateCurtainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCurtainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCurtainRequest proto.InternalMessageInfo

func (m *CreateCurtainRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateCurtainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCurtainRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateCurtainRequest) GetSpec() *v1alpha1.CurtainSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateCurtainRequest struct {
	Namespace            string                `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string     `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.CurtainSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCurtainRequest) Reset()         { *m = UpdateCurtainRequest{} }
func (m *UpdateCurtainRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCurtainRequest) ProtoMessage()    {}
func (*UpdateCurtainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{4}
}
func (m *UpdateCurtainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCurtainRequest.Unmarshal(m, b)
}
func (m *UpdateCurtainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCurtainRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCurtainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCurtainRequest.Merge(m, src)
}
func (m *UpdateCurtainRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCurtainRequest.Size(m)
}
func (m *UpdateCurtainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCurtainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCurtainRequest proto.InternalMessageInfo

func (m *UpdateCurtainRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateCurtainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCurtainRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateCurtainRequest) GetSpec() *v1alpha1.CurtainSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateCurtainResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCurtainResponse) Reset()         { *m = UpdateCurtainResponse{} }
func (m *UpdateCurtainResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCurtainResponse) ProtoMessage()    {}
func (*UpdateCurtainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{5}
}
func (m *UpdateCurtainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCurtainResponse.Unmarshal(m, b)
}
func (m *UpdateCurtainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCurtainResponse.Marshal(b, m, deterministic)
}
func (m *UpdateCurtainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCurtainResponse.Merge(m, src)
}
func (m *UpdateCurtainResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCurtainResponse.Size(m)
}
func (m *UpdateCurtainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCurtainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCurtainResponse proto.InternalMessageInfo

type GetCurtainRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurtainRequest) Reset()         { *m = GetCurtainRequest{} }
func (m *GetCurtainRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurtainRequest) ProtoMessage()    {}
func (*GetCurtainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{6}
}
func (m *GetCurtainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurtainRequest.Unmarshal(m, b)
}
func (m *GetCurtainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurtainRequest.Marshal(b, m, deterministic)
}
func (m *GetCurtainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurtainRequest.Merge(m, src)
}
func (m *GetCurtainRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurtainRequest.Size(m)
}
func (m *GetCurtainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurtainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurtainRequest proto.InternalMessageInfo

func (m *GetCurtainRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetCurtainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetCurtainResponse struct {
	Item                 *v1alpha1.Curtain `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string            `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetCurtainResponse) Reset()         { *m = GetCurtainResponse{} }
func (m *GetCurtainResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurtainResponse) ProtoMessage()    {}
func (*GetCurtainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{7}
}
func (m *GetCurtainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurtainResponse.Unmarshal(m, b)
}
func (m *GetCurtainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurtainResponse.Marshal(b, m, deterministic)
}
func (m *GetCurtainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurtainResponse.Merge(m, src)
}
func (m *GetCurtainResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurtainResponse.Size(m)
}
func (m *GetCurtainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurtainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurtainResponse proto.InternalMessageInfo

func (m *GetCurtainResponse) GetItem() *v1alpha1.Curtain {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetCurtainResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type DeleteCurtainRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCurtainRequest) Reset()         { *m = DeleteCurtainRequest{} }
func (m *DeleteCurtainRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCurtainRequest) ProtoMessage()    {}
func (*DeleteCurtainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{8}
}
func (m *DeleteCurtainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCurtainRequest.Unmarshal(m, b)
}
func (m *DeleteCurtainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCurtainRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCurtainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCurtainRequest.Merge(m, src)
}
func (m *DeleteCurtainRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCurtainRequest.Size(m)
}
func (m *DeleteCurtainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCurtainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCurtainRequest proto.InternalMessageInfo

func (m *DeleteCurtainRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteCurtainRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteCurtainResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCurtainResponse) Reset()         { *m = DeleteCurtainResponse{} }
func (m *DeleteCurtainResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCurtainResponse) ProtoMessage()    {}
func (*DeleteCurtainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a75508fc92f4958, []int{9}
}
func (m *DeleteCurtainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCurtainResponse.Unmarshal(m, b)
}
func (m *DeleteCurtainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCurtainResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCurtainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCurtainResponse.Merge(m, src)
}
func (m *DeleteCurtainResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCurtainResponse.Size(m)
}
func (m *DeleteCurtainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCurtainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCurtainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListCurtainsRequest)(nil), "github.com.metaprov.modeld.services.curtain.ListCurtainsRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.curtain.ListCurtainsRequest.LabelsEntry")
	proto.RegisterType((*ListCurtainsResponse)(nil), "github.com.metaprov.modeld.services.curtain.ListCurtainsResponse")
	proto.RegisterType((*CreateCurtainResponse)(nil), "github.com.metaprov.modeld.services.curtain.CreateCurtainResponse")
	proto.RegisterType((*CreateCurtainRequest)(nil), "github.com.metaprov.modeld.services.curtain.CreateCurtainRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.curtain.CreateCurtainRequest.LabelsEntry")
	proto.RegisterType((*UpdateCurtainRequest)(nil), "github.com.metaprov.modeld.services.curtain.UpdateCurtainRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.curtain.UpdateCurtainRequest.LabelsEntry")
	proto.RegisterType((*UpdateCurtainResponse)(nil), "github.com.metaprov.modeld.services.curtain.UpdateCurtainResponse")
	proto.RegisterType((*GetCurtainRequest)(nil), "github.com.metaprov.modeld.services.curtain.GetCurtainRequest")
	proto.RegisterType((*GetCurtainResponse)(nil), "github.com.metaprov.modeld.services.curtain.GetCurtainResponse")
	proto.RegisterType((*DeleteCurtainRequest)(nil), "github.com.metaprov.modeld.services.curtain.DeleteCurtainRequest")
	proto.RegisterType((*DeleteCurtainResponse)(nil), "github.com.metaprov.modeld.services.curtain.DeleteCurtainResponse")
}

func init() { proto.RegisterFile("services/curtain/curtain.proto", fileDescriptor_3a75508fc92f4958) }

var fileDescriptor_3a75508fc92f4958 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0xdb, 0x3f, 0x7a, 0xfa, 0x83, 0x4e, 0xb7, 0x74, 0x89, 0x55, 0x4a, 0xf0, 0xa2, 0x28,
	0x4c, 0x6c, 0x45, 0xd0, 0x5e, 0xa8, 0x6d, 0xad, 0x3f, 0x50, 0x6f, 0x22, 0x7a, 0x51, 0xbc, 0x99,
	0x26, 0xc7, 0x34, 0x34, 0x3f, 0x63, 0x66, 0x76, 0xa1, 0x88, 0x20, 0xbe, 0x82, 0x0f, 0x21, 0x7a,
	0x2b, 0xbe, 0x82, 0x77, 0x5e, 0xe9, 0x23, 0xf8, 0x20, 0x32, 0x93, 0x49, 0x77, 0xb3, 0x86, 0x62,
	0x8c, 0x20, 0x5e, 0xed, 0xc9, 0x0c, 0xe7, 0x3b, 0xdf, 0xf9, 0xce, 0xcf, 0x0e, 0x5c, 0x12, 0x98,
	0x0f, 0x22, 0x1f, 0x85, 0xeb, 0xf7, 0x73, 0xc9, 0xa2, 0xb4, 0xfc, 0xa5, 0x3c, 0xcf, 0x64, 0x46,
	0xae, 0x86, 0x91, 0x3c, 0xea, 0x1f, 0x52, 0x3f, 0x4b, 0x68, 0x82, 0x92, 0xf1, 0x3c, 0x1b, 0xd0,
	0x24, 0x0b, 0x30, 0x0e, 0x68, 0xe9, 0x4a, 0x8d, 0x8b, 0xbd, 0x1a, 0x66, 0x59, 0x18, 0xa3, 0xcb,
	0x78, 0xe4, 0xb2, 0x34, 0xcd, 0x24, 0x93, 0x51, 0x96, 0x8a, 0x02, 0xca, 0x7e, 0x34, 0x84, 0x72,
	0x4b, 0x28, 0xb7, 0x80, 0x52, 0x0e, 0xfc, 0x38, 0x54, 0x8e, 0xc2, 0x8d, 0xd2, 0x17, 0x98, 0x63,
	0xea, 0xa3, 0x3b, 0xd8, 0x60, 0x31, 0x3f, 0x62, 0x1b, 0x6e, 0x88, 0x29, 0xe6, 0x4c, 0x62, 0x50,
	0x40, 0x39, 0xdf, 0x2d, 0x58, 0xda, 0x8f, 0x84, 0xdc, 0x2d, 0x02, 0x0b, 0x0f, 0x5f, 0xf6, 0x51,
	0x48, 0xb2, 0x0a, 0xb3, 0x29, 0x4b, 0x50, 0x70, 0xe6, 0x63, 0xcf, 0x5a, 0xb3, 0xd6, 0x67, 0xbd,
	0xe1, 0x01, 0x09, 0x60, 0x3a, 0x66, 0x87, 0x18, 0x8b, 0x5e, 0x67, 0x6d, 0x62, 0x7d, 0x6e, 0x73,
	0x9f, 0x36, 0x48, 0x8e, 0xd6, 0xc4, 0xa3, 0xfb, 0x1a, 0x6e, 0x2f, 0x95, 0xf9, 0x89, 0x67, 0xb0,
	0xed, 0x5b, 0x30, 0x37, 0x72, 0x4c, 0xce, 0xc1, 0xc4, 0x31, 0x9e, 0x18, 0x32, 0xca, 0x24, 0x5d,
	0x98, 0x1a, 0xb0, 0xb8, 0x8f, 0xbd, 0x8e, 0x3e, 0x2b, 0x3e, 0xb6, 0x3a, 0x37, 0x2d, 0x47, 0x42,
	0xb7, 0x1a, 0x45, 0xf0, 0x2c, 0x15, 0x48, 0x9e, 0xc3, 0x54, 0x24, 0x31, 0x11, 0x1a, 0x65, 0x6e,
	0xf3, 0xfe, 0x19, 0xbc, 0x19, 0x8f, 0x28, 0x3f, 0x0e, 0xa9, 0x52, 0x92, 0x9e, 0x2a, 0x49, 0x4b,
	0x25, 0xa9, 0xc1, 0x56, 0x61, 0xbc, 0x02, 0xd4, 0x59, 0x81, 0xe5, 0xdd, 0x1c, 0x99, 0x44, 0x73,
	0x57, 0x86, 0x75, 0xbe, 0x74, 0xa0, 0x3b, 0x76, 0xf3, 0x3b, 0x32, 0x13, 0x98, 0x54, 0x1f, 0x26,
	0x3d, 0x6d, 0x13, 0x3c, 0x95, 0x7e, 0x42, 0x4b, 0xff, 0xb8, 0x91, 0xf4, 0x75, 0x24, 0xea, 0xb4,
	0x27, 0x07, 0x30, 0x29, 0x38, 0xfa, 0xbd, 0xa9, 0xbf, 0xa5, 0xd3, 0x13, 0x8e, 0xbe, 0xa7, 0x31,
	0xdb, 0xd4, 0x55, 0x09, 0xf9, 0x94, 0x07, 0xff, 0x5e, 0xc8, 0x3a, 0x12, 0xff, 0x93, 0x90, 0x2b,
	0xb0, 0x3c, 0x96, 0x82, 0x69, 0xd5, 0x3d, 0x38, 0xff, 0x00, 0x65, 0x5b, 0x75, 0x9d, 0x37, 0x16,
	0x90, 0x51, 0x1c, 0x33, 0x7f, 0xcf, 0x60, 0x52, 0x8d, 0x8a, 0x19, 0xbf, 0x9d, 0xf6, 0x6a, 0x78,
	0x1a, 0x4f, 0x51, 0x38, 0x61, 0x49, 0x5c, 0x52, 0x50, 0xb6, 0xf3, 0x10, 0xba, 0xf7, 0x30, 0xc6,
	0xf6, 0xad, 0xa2, 0xc4, 0x1a, 0x43, 0x2a, 0xd2, 0xd9, 0x7c, 0x3f, 0x03, 0x8b, 0x65, 0x59, 0x8a,
	0x0e, 0x21, 0x1f, 0x2c, 0x98, 0x1f, 0x5d, 0x3d, 0xe4, 0x6e, 0xdb, 0xdd, 0x68, 0x6f, 0xb7, 0x40,
	0x30, 0x55, 0xed, 0xbe, 0xfd, 0xf6, 0xe3, 0x5d, 0x67, 0x91, 0xcc, 0xbb, 0x83, 0x8d, 0xf2, 0x7f,
	0x49, 0x90, 0xcf, 0x16, 0x2c, 0x54, 0x36, 0x02, 0xd9, 0x6e, 0xbd, 0x4d, 0xec, 0x9d, 0x36, 0x10,
	0x86, 0xee, 0x45, 0x4d, 0x77, 0xc5, 0xa9, 0xd0, 0xdd, 0x9a, 0x31, 0x16, 0xf9, 0x68, 0x01, 0x0c,
	0x9b, 0x8b, 0xdc, 0x6e, 0x14, 0xf1, 0x97, 0xee, 0xb6, 0xef, 0xfc, 0xb1, 0xbf, 0xa1, 0x7b, 0x41,
	0xd3, 0x5d, 0x26, 0x4b, 0xa3, 0x74, 0xdd, 0x57, 0xaa, 0x77, 0x5e, 0x93, 0xaf, 0x16, 0x2c, 0x54,
	0x46, 0xad, 0xa1, 0xc8, 0x75, 0x9b, 0xa6, 0xa1, 0xc8, 0xf5, 0x93, 0x7e, 0x43, 0xb3, 0x76, 0xed,
	0xcb, 0x55, 0xd6, 0xa5, 0x93, 0x42, 0x0d, 0x98, 0x64, 0x54, 0xa7, 0x31, 0x14, 0xff, 0x93, 0x05,
	0x0b, 0x95, 0x69, 0x68, 0x98, 0x4f, 0xdd, 0x4c, 0x36, 0xcc, 0xa7, 0x76, 0x18, 0xcb, 0x2a, 0x5c,
	0xa9, 0xab, 0xc2, 0xce, 0xb5, 0x03, 0x7a, 0xf6, 0xa3, 0x69, 0xfc, 0xf5, 0x76, 0x38, 0xad, 0x1f,
	0x48, 0xd7, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x13, 0xb9, 0x4e, 0x4e, 0xd8, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CurtainServiceClient is the client API for CurtainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurtainServiceClient interface {
	ListCurtains(ctx context.Context, in *ListCurtainsRequest, opts ...grpc.CallOption) (*ListCurtainsResponse, error)
	CreateCurtain(ctx context.Context, in *CreateCurtainRequest, opts ...grpc.CallOption) (*CreateCurtainResponse, error)
	GetCurtain(ctx context.Context, in *GetCurtainRequest, opts ...grpc.CallOption) (*GetCurtainResponse, error)
	UpdateCurtain(ctx context.Context, in *UpdateCurtainRequest, opts ...grpc.CallOption) (*UpdateCurtainResponse, error)
	DeleteCurtain(ctx context.Context, in *DeleteCurtainRequest, opts ...grpc.CallOption) (*DeleteCurtainResponse, error)
}

type curtainServiceClient struct {
	cc *grpc.ClientConn
}

func NewCurtainServiceClient(cc *grpc.ClientConn) CurtainServiceClient {
	return &curtainServiceClient{cc}
}

func (c *curtainServiceClient) ListCurtains(ctx context.Context, in *ListCurtainsRequest, opts ...grpc.CallOption) (*ListCurtainsResponse, error) {
	out := new(ListCurtainsResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.curtain.CurtainService/ListCurtains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *curtainServiceClient) CreateCurtain(ctx context.Context, in *CreateCurtainRequest, opts ...grpc.CallOption) (*CreateCurtainResponse, error) {
	out := new(CreateCurtainResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.curtain.CurtainService/CreateCurtain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *curtainServiceClient) GetCurtain(ctx context.Context, in *GetCurtainRequest, opts ...grpc.CallOption) (*GetCurtainResponse, error) {
	out := new(GetCurtainResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.curtain.CurtainService/GetCurtain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *curtainServiceClient) UpdateCurtain(ctx context.Context, in *UpdateCurtainRequest, opts ...grpc.CallOption) (*UpdateCurtainResponse, error) {
	out := new(UpdateCurtainResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.curtain.CurtainService/UpdateCurtain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *curtainServiceClient) DeleteCurtain(ctx context.Context, in *DeleteCurtainRequest, opts ...grpc.CallOption) (*DeleteCurtainResponse, error) {
	out := new(DeleteCurtainResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.curtain.CurtainService/DeleteCurtain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurtainServiceServer is the server API for CurtainService service.
type CurtainServiceServer interface {
	ListCurtains(context.Context, *ListCurtainsRequest) (*ListCurtainsResponse, error)
	CreateCurtain(context.Context, *CreateCurtainRequest) (*CreateCurtainResponse, error)
	GetCurtain(context.Context, *GetCurtainRequest) (*GetCurtainResponse, error)
	UpdateCurtain(context.Context, *UpdateCurtainRequest) (*UpdateCurtainResponse, error)
	DeleteCurtain(context.Context, *DeleteCurtainRequest) (*DeleteCurtainResponse, error)
}

// UnimplementedCurtainServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCurtainServiceServer struct {
}

func (*UnimplementedCurtainServiceServer) ListCurtains(ctx context.Context, req *ListCurtainsRequest) (*ListCurtainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurtains not implemented")
}
func (*UnimplementedCurtainServiceServer) CreateCurtain(ctx context.Context, req *CreateCurtainRequest) (*CreateCurtainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCurtain not implemented")
}
func (*UnimplementedCurtainServiceServer) GetCurtain(ctx context.Context, req *GetCurtainRequest) (*GetCurtainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurtain not implemented")
}
func (*UnimplementedCurtainServiceServer) UpdateCurtain(ctx context.Context, req *UpdateCurtainRequest) (*UpdateCurtainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurtain not implemented")
}
func (*UnimplementedCurtainServiceServer) DeleteCurtain(ctx context.Context, req *DeleteCurtainRequest) (*DeleteCurtainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurtain not implemented")
}

func RegisterCurtainServiceServer(s *grpc.Server, srv CurtainServiceServer) {
	s.RegisterService(&_CurtainService_serviceDesc, srv)
}

func _CurtainService_ListCurtains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurtainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurtainServiceServer).ListCurtains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.curtain.CurtainService/ListCurtains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurtainServiceServer).ListCurtains(ctx, req.(*ListCurtainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurtainService_CreateCurtain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCurtainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurtainServiceServer).CreateCurtain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.curtain.CurtainService/CreateCurtain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurtainServiceServer).CreateCurtain(ctx, req.(*CreateCurtainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurtainService_GetCurtain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurtainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurtainServiceServer).GetCurtain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.curtain.CurtainService/GetCurtain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurtainServiceServer).GetCurtain(ctx, req.(*GetCurtainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurtainService_UpdateCurtain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurtainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurtainServiceServer).UpdateCurtain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.curtain.CurtainService/UpdateCurtain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurtainServiceServer).UpdateCurtain(ctx, req.(*UpdateCurtainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurtainService_DeleteCurtain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCurtainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurtainServiceServer).DeleteCurtain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.curtain.CurtainService/DeleteCurtain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurtainServiceServer).DeleteCurtain(ctx, req.(*DeleteCurtainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CurtainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.curtain.CurtainService",
	HandlerType: (*CurtainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCurtains",
			Handler:    _CurtainService_ListCurtains_Handler,
		},
		{
			MethodName: "CreateCurtain",
			Handler:    _CurtainService_CreateCurtain_Handler,
		},
		{
			MethodName: "GetCurtain",
			Handler:    _CurtainService_GetCurtain_Handler,
		},
		{
			MethodName: "UpdateCurtain",
			Handler:    _CurtainService_UpdateCurtain_Handler,
		},
		{
			MethodName: "DeleteCurtain",
			Handler:    _CurtainService_DeleteCurtain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/curtain/curtain.proto",
}