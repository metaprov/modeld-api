// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/conversation/conversation.proto

package conversation

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1alpha1 "github.com/metaprov/modeldapi/pkg/apis/team/v1alpha1"
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

type ConversationQuery struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConversationQuery) Reset()         { *m = ConversationQuery{} }
func (m *ConversationQuery) String() string { return proto.CompactTextString(m) }
func (*ConversationQuery) ProtoMessage()    {}
func (*ConversationQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{0}
}
func (m *ConversationQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationQuery.Unmarshal(m, b)
}
func (m *ConversationQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationQuery.Marshal(b, m, deterministic)
}
func (m *ConversationQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationQuery.Merge(m, src)
}
func (m *ConversationQuery) XXX_Size() int {
	return xxx_messageInfo_ConversationQuery.Size(m)
}
func (m *ConversationQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationQuery proto.InternalMessageInfo

func (m *ConversationQuery) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConversationQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConversationQuery) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ConversationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversationResponse) Reset()         { *m = ConversationResponse{} }
func (m *ConversationResponse) String() string { return proto.CompactTextString(m) }
func (*ConversationResponse) ProtoMessage()    {}
func (*ConversationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{1}
}
func (m *ConversationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationResponse.Unmarshal(m, b)
}
func (m *ConversationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationResponse.Marshal(b, m, deterministic)
}
func (m *ConversationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationResponse.Merge(m, src)
}
func (m *ConversationResponse) XXX_Size() int {
	return xxx_messageInfo_ConversationResponse.Size(m)
}
func (m *ConversationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationResponse proto.InternalMessageInfo

type ConversationCreateRequest struct {
	Namespace            string                     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string          `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.ConversationSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	Password             string                     `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Upsert               bool                       `protobuf:"varint,7,opt,name=upsert,proto3" json:"upsert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ConversationCreateRequest) Reset()         { *m = ConversationCreateRequest{} }
func (m *ConversationCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ConversationCreateRequest) ProtoMessage()    {}
func (*ConversationCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{2}
}
func (m *ConversationCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationCreateRequest.Unmarshal(m, b)
}
func (m *ConversationCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationCreateRequest.Marshal(b, m, deterministic)
}
func (m *ConversationCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationCreateRequest.Merge(m, src)
}
func (m *ConversationCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ConversationCreateRequest.Size(m)
}
func (m *ConversationCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationCreateRequest proto.InternalMessageInfo

func (m *ConversationCreateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConversationCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConversationCreateRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ConversationCreateRequest) GetSpec() *v1alpha1.ConversationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ConversationCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ConversationCreateRequest) GetUpsert() bool {
	if m != nil {
		return m.Upsert
	}
	return false
}

type ConversationUpdateRequest struct {
	Namespace            string                     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string          `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.ConversationSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ConversationUpdateRequest) Reset()         { *m = ConversationUpdateRequest{} }
func (m *ConversationUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ConversationUpdateRequest) ProtoMessage()    {}
func (*ConversationUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{3}
}
func (m *ConversationUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationUpdateRequest.Unmarshal(m, b)
}
func (m *ConversationUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ConversationUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationUpdateRequest.Merge(m, src)
}
func (m *ConversationUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ConversationUpdateRequest.Size(m)
}
func (m *ConversationUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationUpdateRequest proto.InternalMessageInfo

func (m *ConversationUpdateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConversationUpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConversationUpdateRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ConversationUpdateRequest) GetSpec() *v1alpha1.ConversationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ConversationGetResponse struct {
	Item                 *v1alpha1.Conversation `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string                 `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConversationGetResponse) Reset()         { *m = ConversationGetResponse{} }
func (m *ConversationGetResponse) String() string { return proto.CompactTextString(m) }
func (*ConversationGetResponse) ProtoMessage()    {}
func (*ConversationGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{4}
}
func (m *ConversationGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationGetResponse.Unmarshal(m, b)
}
func (m *ConversationGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationGetResponse.Marshal(b, m, deterministic)
}
func (m *ConversationGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationGetResponse.Merge(m, src)
}
func (m *ConversationGetResponse) XXX_Size() int {
	return xxx_messageInfo_ConversationGetResponse.Size(m)
}
func (m *ConversationGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationGetResponse proto.InternalMessageInfo

func (m *ConversationGetResponse) GetItem() *v1alpha1.Conversation {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ConversationGetResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type ConversationGetNamespacesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversationGetNamespacesRequest) Reset()         { *m = ConversationGetNamespacesRequest{} }
func (m *ConversationGetNamespacesRequest) String() string { return proto.CompactTextString(m) }
func (*ConversationGetNamespacesRequest) ProtoMessage()    {}
func (*ConversationGetNamespacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{5}
}
func (m *ConversationGetNamespacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationGetNamespacesRequest.Unmarshal(m, b)
}
func (m *ConversationGetNamespacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationGetNamespacesRequest.Marshal(b, m, deterministic)
}
func (m *ConversationGetNamespacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationGetNamespacesRequest.Merge(m, src)
}
func (m *ConversationGetNamespacesRequest) XXX_Size() int {
	return xxx_messageInfo_ConversationGetNamespacesRequest.Size(m)
}
func (m *ConversationGetNamespacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationGetNamespacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationGetNamespacesRequest proto.InternalMessageInfo

func (m *ConversationGetNamespacesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConversationGetNamespacesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ConversationGetNamespacesResponse struct {
	Namespaces           []*common.NamespaceInfo `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConversationGetNamespacesResponse) Reset()         { *m = ConversationGetNamespacesResponse{} }
func (m *ConversationGetNamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*ConversationGetNamespacesResponse) ProtoMessage()    {}
func (*ConversationGetNamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c315746dc759528, []int{6}
}
func (m *ConversationGetNamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversationGetNamespacesResponse.Unmarshal(m, b)
}
func (m *ConversationGetNamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversationGetNamespacesResponse.Marshal(b, m, deterministic)
}
func (m *ConversationGetNamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversationGetNamespacesResponse.Merge(m, src)
}
func (m *ConversationGetNamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_ConversationGetNamespacesResponse.Size(m)
}
func (m *ConversationGetNamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversationGetNamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConversationGetNamespacesResponse proto.InternalMessageInfo

func (m *ConversationGetNamespacesResponse) GetNamespaces() []*common.NamespaceInfo {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*ConversationQuery)(nil), "github.com.metaprov.modeld.services.conversation.ConversationQuery")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.conversation.ConversationQuery.LabelsEntry")
	proto.RegisterType((*ConversationResponse)(nil), "github.com.metaprov.modeld.services.conversation.ConversationResponse")
	proto.RegisterType((*ConversationCreateRequest)(nil), "github.com.metaprov.modeld.services.conversation.ConversationCreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.conversation.ConversationCreateRequest.LabelsEntry")
	proto.RegisterType((*ConversationUpdateRequest)(nil), "github.com.metaprov.modeld.services.conversation.ConversationUpdateRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.conversation.ConversationUpdateRequest.LabelsEntry")
	proto.RegisterType((*ConversationGetResponse)(nil), "github.com.metaprov.modeld.services.conversation.ConversationGetResponse")
	proto.RegisterType((*ConversationGetNamespacesRequest)(nil), "github.com.metaprov.modeld.services.conversation.ConversationGetNamespacesRequest")
	proto.RegisterType((*ConversationGetNamespacesResponse)(nil), "github.com.metaprov.modeld.services.conversation.ConversationGetNamespacesResponse")
}

func init() {
	proto.RegisterFile("services/conversation/conversation.proto", fileDescriptor_6c315746dc759528)
}

var fileDescriptor_6c315746dc759528 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x86, 0x3f, 0x27, 0xa9, 0xbf, 0x76, 0xc2, 0x02, 0x86, 0xaa, 0x18, 0xb7, 0x48, 0xc1, 0xb0,
	0x88, 0x40, 0x1a, 0xd3, 0x54, 0x48, 0x6d, 0x97, 0xfd, 0xa1, 0x54, 0x54, 0x45, 0x98, 0x3f, 0xd1,
	0xdd, 0x24, 0x39, 0x75, 0xad, 0xda, 0x9e, 0xc1, 0x33, 0x09, 0x8a, 0x10, 0x2c, 0x90, 0x58, 0xb0,
	0xe6, 0x06, 0xb8, 0x0c, 0x36, 0x2c, 0xb8, 0x06, 0xc4, 0x9e, 0x05, 0x1b, 0xee, 0x02, 0xcd, 0xd8,
	0x4e, 0xed, 0x42, 0xa3, 0xd2, 0x44, 0x88, 0x55, 0x7c, 0x3c, 0x99, 0x67, 0xde, 0x73, 0xde, 0x93,
	0x33, 0x41, 0x4d, 0x01, 0x49, 0x3f, 0xe8, 0x80, 0x70, 0x3b, 0x2c, 0xee, 0x43, 0x22, 0xa8, 0x0c,
	0x58, 0x5c, 0x0a, 0x08, 0x4f, 0x98, 0x64, 0xf8, 0x96, 0x1f, 0xc8, 0x83, 0x5e, 0x9b, 0x74, 0x58,
	0x44, 0x22, 0x90, 0x94, 0x27, 0xac, 0x4f, 0x22, 0xd6, 0x85, 0xb0, 0x4b, 0x72, 0x08, 0x29, 0xee,
	0xb3, 0xe7, 0x7d, 0xc6, 0xfc, 0x10, 0x5c, 0xbd, 0xbf, 0xdd, 0xdb, 0x77, 0x21, 0xe2, 0x72, 0x90,
	0xe2, 0xec, 0x8d, 0x23, 0x9c, 0x9b, 0xe3, 0xdc, 0x14, 0x47, 0x79, 0xe0, 0xf2, 0x43, 0xdf, 0xa5,
	0x3c, 0x10, 0xae, 0x04, 0x1a, 0xb9, 0xfd, 0x45, 0x1a, 0xf2, 0x03, 0xba, 0xe8, 0xfa, 0x10, 0x43,
	0x42, 0x25, 0x74, 0x33, 0xca, 0xea, 0x68, 0x4a, 0x21, 0xb9, 0x28, 0xd2, 0x69, 0xa9, 0x8f, 0x6c,
	0xef, 0x42, 0x26, 0x4f, 0x7d, 0x91, 0xc6, 0x31, 0x93, 0x5a, 0xb5, 0x48, 0x57, 0x9d, 0x1f, 0x06,
	0xba, 0xb0, 0x5e, 0xc8, 0xe6, 0x41, 0x0f, 0x92, 0x01, 0x5e, 0x40, 0x33, 0x31, 0x8d, 0x40, 0x70,
	0xda, 0x01, 0xcb, 0x68, 0x18, 0xcd, 0x19, 0xef, 0xe8, 0x05, 0xc6, 0xa8, 0xa6, 0x02, 0xab, 0xa2,
	0x17, 0xf4, 0x33, 0xf6, 0x91, 0x19, 0xd2, 0x36, 0x84, 0xc2, 0xaa, 0x36, 0xaa, 0xcd, 0x7a, 0xeb,
	0x3e, 0xf9, 0xd3, 0x3a, 0x92, 0x5f, 0x64, 0x90, 0x1d, 0x4d, 0xdc, 0x8c, 0x65, 0x32, 0xf0, 0x32,
	0xbc, 0xbd, 0x82, 0xea, 0x85, 0xd7, 0xf8, 0x3c, 0xaa, 0x1e, 0xc2, 0x20, 0xd3, 0xa8, 0x1e, 0xf1,
	0x2c, 0x9a, 0xea, 0xd3, 0xb0, 0x97, 0xcb, 0x4b, 0x83, 0xd5, 0xca, 0xb2, 0xe1, 0xcc, 0xa1, 0xd9,
	0xe2, 0x19, 0x1e, 0x08, 0xce, 0x62, 0x01, 0xce, 0xbb, 0x2a, 0xba, 0x5c, 0x5c, 0x58, 0x4f, 0x80,
	0x4a, 0xf0, 0xe0, 0x79, 0x0f, 0x84, 0x3c, 0x43, 0x2d, 0xd8, 0xb1, 0x5a, 0x3c, 0x1d, 0xaf, 0x16,
	0x25, 0x39, 0xbf, 0xab, 0x09, 0xde, 0x43, 0x35, 0xc1, 0xa1, 0x63, 0x4d, 0x35, 0x8c, 0x66, 0xbd,
	0x75, 0x67, 0xc4, 0x71, 0x94, 0x07, 0x84, 0x1f, 0xfa, 0x44, 0xf5, 0x1c, 0x51, 0x3d, 0x47, 0xf2,
	0x9e, 0x2b, 0x1d, 0xf9, 0x90, 0x43, 0xc7, 0xd3, 0x4c, 0x6c, 0xa3, 0x69, 0x4e, 0x85, 0x78, 0xc1,
	0x92, 0xae, 0x65, 0xea, 0x24, 0x87, 0x31, 0x9e, 0x43, 0x66, 0x8f, 0x0b, 0x48, 0xa4, 0xf5, 0x7f,
	0xc3, 0x68, 0x4e, 0x7b, 0x59, 0x34, 0x8e, 0x47, 0x5f, 0x2b, 0x65, 0x2f, 0x1e, 0xf3, 0xee, 0xbf,
	0xe4, 0x45, 0x49, 0xce, 0xdf, 0xf6, 0x62, 0x9c, 0xba, 0xbe, 0x35, 0xd0, 0xa5, 0x22, 0x75, 0x0b,
	0x64, 0xde, 0xff, 0xf8, 0x09, 0xaa, 0x05, 0x12, 0x22, 0x0d, 0xaa, 0xb7, 0xd6, 0xc6, 0x97, 0xec,
	0x69, 0x9e, 0xf2, 0x63, 0x40, 0xa3, 0x30, 0xf7, 0x43, 0x3d, 0x3b, 0x8f, 0x50, 0xe3, 0x98, 0x8c,
	0xdd, 0xdc, 0x3f, 0x71, 0x66, 0x97, 0x9d, 0xd7, 0xe8, 0xea, 0x08, 0x6a, 0x96, 0xe6, 0x33, 0x84,
	0x86, 0x14, 0x61, 0x19, 0xba, 0x1d, 0x56, 0x4e, 0xd9, 0x0e, 0x7a, 0x9e, 0x0e, 0x99, 0xdb, 0xf1,
	0x3e, 0xf3, 0x0a, 0xb0, 0xd6, 0x37, 0x13, 0x5d, 0x2c, 0x79, 0x96, 0xee, 0xc4, 0x1f, 0x0d, 0x54,
	0xdb, 0x09, 0x84, 0xc4, 0xeb, 0x13, 0x18, 0x87, 0xf6, 0x04, 0x9a, 0x49, 0x89, 0x71, 0xae, 0xbd,
	0xf9, 0xf2, 0xfd, 0x7d, 0xe5, 0x0a, 0x9e, 0xd7, 0x77, 0xc2, 0xf0, 0xd2, 0x29, 0x9e, 0x2e, 0xf0,
	0x07, 0x03, 0x99, 0xe9, 0xe4, 0xc1, 0xf7, 0x26, 0x38, 0xbf, 0xec, 0x39, 0x92, 0xde, 0x47, 0x24,
	0xbf, 0x2e, 0xc9, 0xa6, 0xba, 0x2e, 0x9d, 0x45, 0x2d, 0xea, 0xa6, 0x33, 0x4a, 0xd4, 0xea, 0xb9,
	0x62, 0x88, 0x3f, 0x1b, 0xa8, 0xba, 0x05, 0x13, 0x2a, 0xee, 0xf6, 0x78, 0x90, 0xc2, 0xef, 0xc9,
	0xb9, 0xa1, 0x53, 0xb9, 0x8e, 0x9d, 0x11, 0xa9, 0xb8, 0x2f, 0x55, 0xfb, 0xbc, 0xc2, 0x9f, 0x0c,
	0x64, 0xa6, 0x43, 0x65, 0xdc, 0x32, 0x97, 0x46, 0xd3, 0x89, 0x65, 0xde, 0xd5, 0xda, 0xee, 0xda,
	0xcb, 0x23, 0xb5, 0x95, 0x0e, 0x51, 0x52, 0xba, 0x54, 0x52, 0xa2, 0x15, 0x1f, 0xf3, 0xc0, 0x47,
	0xe6, 0x06, 0x84, 0x20, 0x61, 0x32, 0x2e, 0x9c, 0x24, 0xfb, 0xbf, 0xb5, 0xdb, 0x7b, 0x4b, 0xa7,
	0xfe, 0x13, 0x74, 0x44, 0x6d, 0x9b, 0x1a, 0xb4, 0xf4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x59, 0x18,
	0xf4, 0x84, 0x01, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConversationServiceClient is the client API for ConversationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConversationServiceClient interface {
	List(ctx context.Context, in *ConversationQuery, opts ...grpc.CallOption) (*v1alpha1.ConversationList, error)
	Create(ctx context.Context, in *ConversationCreateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *ConversationQuery, opts ...grpc.CallOption) (*ConversationGetResponse, error)
	Update(ctx context.Context, in *ConversationUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *ConversationQuery, opts ...grpc.CallOption) (*empty.Empty, error)
}

type conversationServiceClient struct {
	cc *grpc.ClientConn
}

func NewConversationServiceClient(cc *grpc.ClientConn) ConversationServiceClient {
	return &conversationServiceClient{cc}
}

func (c *conversationServiceClient) List(ctx context.Context, in *ConversationQuery, opts ...grpc.CallOption) (*v1alpha1.ConversationList, error) {
	out := new(v1alpha1.ConversationList)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.conversation.ConversationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) Create(ctx context.Context, in *ConversationCreateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.conversation.ConversationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) Get(ctx context.Context, in *ConversationQuery, opts ...grpc.CallOption) (*ConversationGetResponse, error) {
	out := new(ConversationGetResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.conversation.ConversationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) Update(ctx context.Context, in *ConversationUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.conversation.ConversationService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationServiceClient) Delete(ctx context.Context, in *ConversationQuery, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.conversation.ConversationService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServiceServer is the server API for ConversationService service.
type ConversationServiceServer interface {
	List(context.Context, *ConversationQuery) (*v1alpha1.ConversationList, error)
	Create(context.Context, *ConversationCreateRequest) (*empty.Empty, error)
	Get(context.Context, *ConversationQuery) (*ConversationGetResponse, error)
	Update(context.Context, *ConversationUpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *ConversationQuery) (*empty.Empty, error)
}

// UnimplementedConversationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConversationServiceServer struct {
}

func (*UnimplementedConversationServiceServer) List(ctx context.Context, req *ConversationQuery) (*v1alpha1.ConversationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedConversationServiceServer) Create(ctx context.Context, req *ConversationCreateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedConversationServiceServer) Get(ctx context.Context, req *ConversationQuery) (*ConversationGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedConversationServiceServer) Update(ctx context.Context, req *ConversationUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedConversationServiceServer) Delete(ctx context.Context, req *ConversationQuery) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterConversationServiceServer(s *grpc.Server, srv ConversationServiceServer) {
	s.RegisterService(&_ConversationService_serviceDesc, srv)
}

func _ConversationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.conversation.ConversationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).List(ctx, req.(*ConversationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.conversation.ConversationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).Create(ctx, req.(*ConversationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.conversation.ConversationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).Get(ctx, req.(*ConversationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.conversation.ConversationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).Update(ctx, req.(*ConversationUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.conversation.ConversationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).Delete(ctx, req.(*ConversationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConversationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.conversation.ConversationService",
	HandlerType: (*ConversationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ConversationService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ConversationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ConversationService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ConversationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ConversationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/conversation/conversation.proto",
}