// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/notifier/notifier.proto

package notifier

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

type ListNotifiersRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListNotifiersRequest) Reset()         { *m = ListNotifiersRequest{} }
func (m *ListNotifiersRequest) String() string { return proto.CompactTextString(m) }
func (*ListNotifiersRequest) ProtoMessage()    {}
func (*ListNotifiersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{0}
}
func (m *ListNotifiersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotifiersRequest.Unmarshal(m, b)
}
func (m *ListNotifiersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotifiersRequest.Marshal(b, m, deterministic)
}
func (m *ListNotifiersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotifiersRequest.Merge(m, src)
}
func (m *ListNotifiersRequest) XXX_Size() int {
	return xxx_messageInfo_ListNotifiersRequest.Size(m)
}
func (m *ListNotifiersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotifiersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotifiersRequest proto.InternalMessageInfo

func (m *ListNotifiersRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListNotifiersRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type ListNotifiersResponse struct {
	Items                *v1alpha1.NotifierList `protobuf:"bytes,1,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListNotifiersResponse) Reset()         { *m = ListNotifiersResponse{} }
func (m *ListNotifiersResponse) String() string { return proto.CompactTextString(m) }
func (*ListNotifiersResponse) ProtoMessage()    {}
func (*ListNotifiersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{1}
}
func (m *ListNotifiersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotifiersResponse.Unmarshal(m, b)
}
func (m *ListNotifiersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotifiersResponse.Marshal(b, m, deterministic)
}
func (m *ListNotifiersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotifiersResponse.Merge(m, src)
}
func (m *ListNotifiersResponse) XXX_Size() int {
	return xxx_messageInfo_ListNotifiersResponse.Size(m)
}
func (m *ListNotifiersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotifiersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotifiersResponse proto.InternalMessageInfo

func (m *ListNotifiersResponse) GetItems() *v1alpha1.NotifierList {
	if m != nil {
		return m.Items
	}
	return nil
}

type NotifierResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifierResponse) Reset()         { *m = NotifierResponse{} }
func (m *NotifierResponse) String() string { return proto.CompactTextString(m) }
func (*NotifierResponse) ProtoMessage()    {}
func (*NotifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{2}
}
func (m *NotifierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifierResponse.Unmarshal(m, b)
}
func (m *NotifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifierResponse.Marshal(b, m, deterministic)
}
func (m *NotifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifierResponse.Merge(m, src)
}
func (m *NotifierResponse) XXX_Size() int {
	return xxx_messageInfo_NotifierResponse.Size(m)
}
func (m *NotifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifierResponse proto.InternalMessageInfo

type CreateNotifierRequest struct {
	Namespace            string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string      `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.NotifierSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateNotifierRequest) Reset()         { *m = CreateNotifierRequest{} }
func (m *CreateNotifierRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNotifierRequest) ProtoMessage()    {}
func (*CreateNotifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{3}
}
func (m *CreateNotifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNotifierRequest.Unmarshal(m, b)
}
func (m *CreateNotifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNotifierRequest.Marshal(b, m, deterministic)
}
func (m *CreateNotifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNotifierRequest.Merge(m, src)
}
func (m *CreateNotifierRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNotifierRequest.Size(m)
}
func (m *CreateNotifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNotifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNotifierRequest proto.InternalMessageInfo

func (m *CreateNotifierRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateNotifierRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateNotifierRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateNotifierRequest) GetSpec() *v1alpha1.NotifierSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type CreateNotifierResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNotifierResponse) Reset()         { *m = CreateNotifierResponse{} }
func (m *CreateNotifierResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNotifierResponse) ProtoMessage()    {}
func (*CreateNotifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{4}
}
func (m *CreateNotifierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNotifierResponse.Unmarshal(m, b)
}
func (m *CreateNotifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNotifierResponse.Marshal(b, m, deterministic)
}
func (m *CreateNotifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNotifierResponse.Merge(m, src)
}
func (m *CreateNotifierResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNotifierResponse.Size(m)
}
func (m *CreateNotifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNotifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNotifierResponse proto.InternalMessageInfo

type UpdateNotifierRequest struct {
	Namespace            string                 `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels               map[string]string      `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec                 *v1alpha1.NotifierSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UpdateNotifierRequest) Reset()         { *m = UpdateNotifierRequest{} }
func (m *UpdateNotifierRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNotifierRequest) ProtoMessage()    {}
func (*UpdateNotifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{5}
}
func (m *UpdateNotifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNotifierRequest.Unmarshal(m, b)
}
func (m *UpdateNotifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNotifierRequest.Marshal(b, m, deterministic)
}
func (m *UpdateNotifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNotifierRequest.Merge(m, src)
}
func (m *UpdateNotifierRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNotifierRequest.Size(m)
}
func (m *UpdateNotifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNotifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNotifierRequest proto.InternalMessageInfo

func (m *UpdateNotifierRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateNotifierRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNotifierRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UpdateNotifierRequest) GetSpec() *v1alpha1.NotifierSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type UpdateNotifierResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNotifierResponse) Reset()         { *m = UpdateNotifierResponse{} }
func (m *UpdateNotifierResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateNotifierResponse) ProtoMessage()    {}
func (*UpdateNotifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{6}
}
func (m *UpdateNotifierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNotifierResponse.Unmarshal(m, b)
}
func (m *UpdateNotifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNotifierResponse.Marshal(b, m, deterministic)
}
func (m *UpdateNotifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNotifierResponse.Merge(m, src)
}
func (m *UpdateNotifierResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateNotifierResponse.Size(m)
}
func (m *UpdateNotifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNotifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNotifierResponse proto.InternalMessageInfo

type GetNotifierRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNotifierRequest) Reset()         { *m = GetNotifierRequest{} }
func (m *GetNotifierRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotifierRequest) ProtoMessage()    {}
func (*GetNotifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{7}
}
func (m *GetNotifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotifierRequest.Unmarshal(m, b)
}
func (m *GetNotifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotifierRequest.Marshal(b, m, deterministic)
}
func (m *GetNotifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotifierRequest.Merge(m, src)
}
func (m *GetNotifierRequest) XXX_Size() int {
	return xxx_messageInfo_GetNotifierRequest.Size(m)
}
func (m *GetNotifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotifierRequest proto.InternalMessageInfo

func (m *GetNotifierRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetNotifierRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNotifierResponse struct {
	Item                 *v1alpha1.Notifier `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Yaml                 string             `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetNotifierResponse) Reset()         { *m = GetNotifierResponse{} }
func (m *GetNotifierResponse) String() string { return proto.CompactTextString(m) }
func (*GetNotifierResponse) ProtoMessage()    {}
func (*GetNotifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{8}
}
func (m *GetNotifierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotifierResponse.Unmarshal(m, b)
}
func (m *GetNotifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotifierResponse.Marshal(b, m, deterministic)
}
func (m *GetNotifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotifierResponse.Merge(m, src)
}
func (m *GetNotifierResponse) XXX_Size() int {
	return xxx_messageInfo_GetNotifierResponse.Size(m)
}
func (m *GetNotifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotifierResponse proto.InternalMessageInfo

func (m *GetNotifierResponse) GetItem() *v1alpha1.Notifier {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *GetNotifierResponse) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

type GetNotifierNamespacesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNotifierNamespacesRequest) Reset()         { *m = GetNotifierNamespacesRequest{} }
func (m *GetNotifierNamespacesRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotifierNamespacesRequest) ProtoMessage()    {}
func (*GetNotifierNamespacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{9}
}
func (m *GetNotifierNamespacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotifierNamespacesRequest.Unmarshal(m, b)
}
func (m *GetNotifierNamespacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotifierNamespacesRequest.Marshal(b, m, deterministic)
}
func (m *GetNotifierNamespacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotifierNamespacesRequest.Merge(m, src)
}
func (m *GetNotifierNamespacesRequest) XXX_Size() int {
	return xxx_messageInfo_GetNotifierNamespacesRequest.Size(m)
}
func (m *GetNotifierNamespacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotifierNamespacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotifierNamespacesRequest proto.InternalMessageInfo

func (m *GetNotifierNamespacesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetNotifierNamespacesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNotifierNamespacesResponse struct {
	Namespaces           []*common.NamespaceInfo `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetNotifierNamespacesResponse) Reset()         { *m = GetNotifierNamespacesResponse{} }
func (m *GetNotifierNamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*GetNotifierNamespacesResponse) ProtoMessage()    {}
func (*GetNotifierNamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{10}
}
func (m *GetNotifierNamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotifierNamespacesResponse.Unmarshal(m, b)
}
func (m *GetNotifierNamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotifierNamespacesResponse.Marshal(b, m, deterministic)
}
func (m *GetNotifierNamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotifierNamespacesResponse.Merge(m, src)
}
func (m *GetNotifierNamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_GetNotifierNamespacesResponse.Size(m)
}
func (m *GetNotifierNamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotifierNamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotifierNamespacesResponse proto.InternalMessageInfo

func (m *GetNotifierNamespacesResponse) GetNamespaces() []*common.NamespaceInfo {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type DeleteNotifierRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNotifierRequest) Reset()         { *m = DeleteNotifierRequest{} }
func (m *DeleteNotifierRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNotifierRequest) ProtoMessage()    {}
func (*DeleteNotifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{11}
}
func (m *DeleteNotifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNotifierRequest.Unmarshal(m, b)
}
func (m *DeleteNotifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNotifierRequest.Marshal(b, m, deterministic)
}
func (m *DeleteNotifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNotifierRequest.Merge(m, src)
}
func (m *DeleteNotifierRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNotifierRequest.Size(m)
}
func (m *DeleteNotifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNotifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNotifierRequest proto.InternalMessageInfo

func (m *DeleteNotifierRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteNotifierRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteNotifierResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNotifierResponse) Reset()         { *m = DeleteNotifierResponse{} }
func (m *DeleteNotifierResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteNotifierResponse) ProtoMessage()    {}
func (*DeleteNotifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe5f71c64661c43, []int{12}
}
func (m *DeleteNotifierResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNotifierResponse.Unmarshal(m, b)
}
func (m *DeleteNotifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNotifierResponse.Marshal(b, m, deterministic)
}
func (m *DeleteNotifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNotifierResponse.Merge(m, src)
}
func (m *DeleteNotifierResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteNotifierResponse.Size(m)
}
func (m *DeleteNotifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNotifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNotifierResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListNotifiersRequest)(nil), "github.com.metaprov.modeld.services.notifier.ListNotifiersRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.notifier.ListNotifiersRequest.LabelsEntry")
	proto.RegisterType((*ListNotifiersResponse)(nil), "github.com.metaprov.modeld.services.notifier.ListNotifiersResponse")
	proto.RegisterType((*NotifierResponse)(nil), "github.com.metaprov.modeld.services.notifier.NotifierResponse")
	proto.RegisterType((*CreateNotifierRequest)(nil), "github.com.metaprov.modeld.services.notifier.CreateNotifierRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.notifier.CreateNotifierRequest.LabelsEntry")
	proto.RegisterType((*CreateNotifierResponse)(nil), "github.com.metaprov.modeld.services.notifier.CreateNotifierResponse")
	proto.RegisterType((*UpdateNotifierRequest)(nil), "github.com.metaprov.modeld.services.notifier.UpdateNotifierRequest")
	proto.RegisterMapType((map[string]string)(nil), "github.com.metaprov.modeld.services.notifier.UpdateNotifierRequest.LabelsEntry")
	proto.RegisterType((*UpdateNotifierResponse)(nil), "github.com.metaprov.modeld.services.notifier.UpdateNotifierResponse")
	proto.RegisterType((*GetNotifierRequest)(nil), "github.com.metaprov.modeld.services.notifier.GetNotifierRequest")
	proto.RegisterType((*GetNotifierResponse)(nil), "github.com.metaprov.modeld.services.notifier.GetNotifierResponse")
	proto.RegisterType((*GetNotifierNamespacesRequest)(nil), "github.com.metaprov.modeld.services.notifier.GetNotifierNamespacesRequest")
	proto.RegisterType((*GetNotifierNamespacesResponse)(nil), "github.com.metaprov.modeld.services.notifier.GetNotifierNamespacesResponse")
	proto.RegisterType((*DeleteNotifierRequest)(nil), "github.com.metaprov.modeld.services.notifier.DeleteNotifierRequest")
	proto.RegisterType((*DeleteNotifierResponse)(nil), "github.com.metaprov.modeld.services.notifier.DeleteNotifierResponse")
}

func init() { proto.RegisterFile("services/notifier/notifier.proto", fileDescriptor_ffe5f71c64661c43) }

var fileDescriptor_ffe5f71c64661c43 = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0xc7, 0xd3, 0x65, 0x21, 0xf0, 0x36, 0x22, 0x19, 0x41, 0xd7, 0x06, 0x75, 0x53, 0x13, 0x83,
	0xc6, 0x4c, 0x65, 0xbd, 0x00, 0x31, 0x51, 0x59, 0xd0, 0x90, 0x10, 0x34, 0x25, 0xc6, 0xe8, 0x6d,
	0xd8, 0x9d, 0x2d, 0x0d, 0x6d, 0xa7, 0x76, 0x86, 0x4d, 0xd0, 0xe8, 0xc1, 0xbb, 0x27, 0x3f, 0x8d,
	0x7e, 0x00, 0xbe, 0x80, 0x31, 0x1e, 0xbd, 0xf8, 0x41, 0xcc, 0x4c, 0xa7, 0xa5, 0xbb, 0x76, 0x09,
	0xb5, 0x1c, 0x38, 0xed, 0x74, 0x36, 0xf3, 0x7b, 0xff, 0xf7, 0x7f, 0xf3, 0x5e, 0x0b, 0x2d, 0x4e,
	0xe3, 0x81, 0xd7, 0xa5, 0xdc, 0x0e, 0x99, 0xf0, 0xfa, 0x1e, 0x8d, 0xb3, 0x05, 0x8e, 0x62, 0x26,
	0x18, 0xba, 0xef, 0x7a, 0x62, 0xff, 0x70, 0x0f, 0x77, 0x59, 0x80, 0x03, 0x2a, 0x48, 0x14, 0xb3,
	0x01, 0x0e, 0x58, 0x8f, 0xfa, 0x3d, 0x9c, 0x1e, 0xc6, 0xe9, 0x19, 0x73, 0xd1, 0x65, 0xcc, 0xf5,
	0xa9, 0x4d, 0x22, 0xcf, 0x26, 0x61, 0xc8, 0x04, 0x11, 0x1e, 0x0b, 0x79, 0xc2, 0x32, 0x37, 0x4f,
	0x58, 0x76, 0xca, 0xb2, 0x13, 0x96, 0x3c, 0x10, 0x1d, 0xb8, 0xf2, 0x20, 0xb7, 0xbd, 0xb0, 0x1f,
	0x13, 0x7b, 0xb0, 0x4c, 0xfc, 0x68, 0x9f, 0x2c, 0xdb, 0x2e, 0x0d, 0x69, 0x4c, 0x04, 0xed, 0x69,
	0xcc, 0xda, 0xe9, 0x98, 0x2c, 0xa5, 0x2e, 0x0b, 0x02, 0x16, 0xea, 0x9f, 0xe4, 0xac, 0xf5, 0xcb,
	0x80, 0xf9, 0x6d, 0x8f, 0x8b, 0x1d, 0xad, 0x98, 0x3b, 0xf4, 0xdd, 0x21, 0xe5, 0x02, 0x2d, 0xc2,
	0x4c, 0x48, 0x02, 0xca, 0x23, 0xd2, 0xa5, 0x4d, 0xa3, 0x65, 0x2c, 0xcd, 0x38, 0x27, 0x1b, 0xa8,
	0x0f, 0x53, 0x3e, 0xd9, 0xa3, 0x3e, 0x6f, 0x4e, 0xb4, 0x26, 0x96, 0x1a, 0xed, 0x1d, 0x5c, 0xc6,
	0x16, 0x5c, 0x14, 0x11, 0x6f, 0x2b, 0xe0, 0x66, 0x28, 0xe2, 0x23, 0x47, 0xd3, 0xcd, 0x55, 0x68,
	0xe4, 0xb6, 0xd1, 0x1c, 0x4c, 0x1c, 0xd0, 0x23, 0x2d, 0x47, 0x2e, 0xd1, 0x3c, 0x4c, 0x0e, 0x88,
	0x7f, 0x48, 0x9b, 0x35, 0xb5, 0x97, 0x3c, 0xac, 0xd5, 0x56, 0x0c, 0x2b, 0x86, 0x85, 0x91, 0x30,
	0x3c, 0x62, 0x21, 0xa7, 0xe8, 0x0d, 0x4c, 0x7a, 0x82, 0x06, 0x5c, 0x61, 0x1a, 0xed, 0xce, 0x29,
	0xd2, 0x49, 0xe4, 0xe1, 0xe8, 0xc0, 0xc5, 0xb2, 0x0a, 0x58, 0x55, 0x01, 0xa7, 0x55, 0xc0, 0x29,
	0x58, 0x06, 0x71, 0x12, 0xa2, 0x85, 0x60, 0x2e, 0xdd, 0x4e, 0xc3, 0x59, 0xc7, 0x35, 0x58, 0xe8,
	0xc4, 0x94, 0x08, 0x7a, 0xf2, 0xd7, 0x59, 0x2c, 0x46, 0x50, 0x97, 0x0f, 0x3a, 0x31, 0xb5, 0x46,
	0xee, 0x88, 0xed, 0x2f, 0xca, 0xd9, 0x5e, 0x28, 0xa3, 0xc8, 0x77, 0xf4, 0x1a, 0xea, 0x3c, 0xa2,
	0xdd, 0x66, 0xfd, 0x5c, 0x2c, 0xda, 0x8d, 0x68, 0xd7, 0x51, 0xc0, 0x2a, 0x05, 0x6d, 0xc2, 0xd5,
	0xd1, 0x04, 0x72, 0x16, 0xbf, 0x8a, 0x7a, 0x17, 0xc1, 0xe2, 0x42, 0x19, 0xa7, 0x5a, 0x3c, 0x79,
	0xb1, 0x2c, 0x1e, 0x4d, 0x40, 0x5b, 0xfc, 0x0c, 0xd0, 0x73, 0x2a, 0x2a, 0xdb, 0x6b, 0x7d, 0x82,
	0x2b, 0x43, 0x1c, 0xdd, 0x93, 0xbb, 0x50, 0x97, 0x1d, 0xa4, 0x5b, 0xf2, 0x71, 0x45, 0x33, 0x1c,
	0x05, 0x93, 0xf1, 0x8f, 0x48, 0xe0, 0xa7, 0xf1, 0xe5, 0xda, 0x7a, 0x09, 0x8b, 0xb9, 0xf8, 0x3b,
	0xa9, 0x56, 0xfe, 0xff, 0x19, 0xbd, 0x87, 0x1b, 0x63, 0x88, 0xd9, 0xbc, 0x81, 0x8c, 0x20, 0x87,
	0x8e, 0xbc, 0x55, 0xab, 0x67, 0xba, 0x55, 0x7a, 0x52, 0x67, 0xcc, 0xad, 0xb0, 0xcf, 0x9c, 0x1c,
	0xcc, 0xda, 0x82, 0x85, 0x0d, 0xea, 0xd3, 0x73, 0xb8, 0xf7, 0xb2, 0xf4, 0xa3, 0xa8, 0x44, 0x7f,
	0xfb, 0xcb, 0x34, 0x5c, 0xce, 0xae, 0x59, 0x22, 0x0d, 0x7d, 0x33, 0xe0, 0xd2, 0xd0, 0x74, 0x45,
	0xeb, 0xd5, 0xdf, 0x00, 0x66, 0xa7, 0x12, 0x43, 0xdf, 0xd4, 0x5b, 0x9f, 0x7f, 0xfc, 0xf9, 0x5a,
	0xbb, 0x8e, 0xae, 0xa9, 0x97, 0x6e, 0xf6, 0xd2, 0x0c, 0x33, 0xa5, 0xc7, 0x06, 0xcc, 0x0e, 0x0f,
	0x12, 0xd4, 0x39, 0x87, 0x39, 0x6a, 0x6e, 0x54, 0x83, 0x68, 0xf9, 0x77, 0x95, 0xfc, 0xdb, 0xd6,
	0x38, 0xf9, 0x6b, 0xd3, 0xe9, 0x12, 0x7d, 0x37, 0xa0, 0x91, 0xbb, 0x7a, 0xe8, 0x49, 0x39, 0x01,
	0xff, 0xf6, 0xb3, 0xf9, 0xb4, 0x02, 0x41, 0xeb, 0xbf, 0xa3, 0xf4, 0xb7, 0xd0, 0xcd, 0x31, 0xfa,
	0xed, 0x0f, 0xf2, 0xba, 0x7d, 0x44, 0xbf, 0x0d, 0x98, 0x1d, 0x9e, 0x35, 0x65, 0xab, 0x50, 0x38,
	0x6a, 0xcb, 0x56, 0x61, 0xcc, 0xb8, 0x5b, 0x57, 0x59, 0x3c, 0x32, 0x1f, 0x8c, 0xcf, 0x22, 0x25,
	0xc8, 0x18, 0x3d, 0x22, 0x08, 0x56, 0x79, 0xe5, 0xca, 0xf3, 0xd3, 0x80, 0xd9, 0xe1, 0x96, 0x2a,
	0x9b, 0x61, 0x61, 0x6f, 0x97, 0xcd, 0xb0, 0xb8, 0xab, 0xad, 0x15, 0x95, 0x61, 0xfb, 0x5e, 0xe9,
	0x0c, 0xd7, 0x97, 0xdf, 0xda, 0x67, 0xfc, 0xe0, 0x4c, 0x09, 0x7b, 0x53, 0xea, 0x63, 0xf3, 0xe1,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xeb, 0xe9, 0x83, 0x5f, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotifierServiceClient is the client API for NotifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotifierServiceClient interface {
	ListNotifiers(ctx context.Context, in *ListNotifiersRequest, opts ...grpc.CallOption) (*ListNotifiersResponse, error)
	CreateNotifier(ctx context.Context, in *CreateNotifierRequest, opts ...grpc.CallOption) (*CreateNotifierResponse, error)
	GetNotifier(ctx context.Context, in *GetNotifierRequest, opts ...grpc.CallOption) (*GetNotifierResponse, error)
	UpdateNotifier(ctx context.Context, in *UpdateNotifierRequest, opts ...grpc.CallOption) (*UpdateNotifierResponse, error)
	DeleteNotifier(ctx context.Context, in *DeleteNotifierRequest, opts ...grpc.CallOption) (*DeleteNotifierResponse, error)
}

type notifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotifierServiceClient(cc *grpc.ClientConn) NotifierServiceClient {
	return &notifierServiceClient{cc}
}

func (c *notifierServiceClient) ListNotifiers(ctx context.Context, in *ListNotifiersRequest, opts ...grpc.CallOption) (*ListNotifiersResponse, error) {
	out := new(ListNotifiersResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.notifier.NotifierService/ListNotifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierServiceClient) CreateNotifier(ctx context.Context, in *CreateNotifierRequest, opts ...grpc.CallOption) (*CreateNotifierResponse, error) {
	out := new(CreateNotifierResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.notifier.NotifierService/CreateNotifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierServiceClient) GetNotifier(ctx context.Context, in *GetNotifierRequest, opts ...grpc.CallOption) (*GetNotifierResponse, error) {
	out := new(GetNotifierResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.notifier.NotifierService/GetNotifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierServiceClient) UpdateNotifier(ctx context.Context, in *UpdateNotifierRequest, opts ...grpc.CallOption) (*UpdateNotifierResponse, error) {
	out := new(UpdateNotifierResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.notifier.NotifierService/UpdateNotifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierServiceClient) DeleteNotifier(ctx context.Context, in *DeleteNotifierRequest, opts ...grpc.CallOption) (*DeleteNotifierResponse, error) {
	out := new(DeleteNotifierResponse)
	err := c.cc.Invoke(ctx, "/github.com.metaprov.modeld.services.notifier.NotifierService/DeleteNotifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServiceServer is the server API for NotifierService service.
type NotifierServiceServer interface {
	ListNotifiers(context.Context, *ListNotifiersRequest) (*ListNotifiersResponse, error)
	CreateNotifier(context.Context, *CreateNotifierRequest) (*CreateNotifierResponse, error)
	GetNotifier(context.Context, *GetNotifierRequest) (*GetNotifierResponse, error)
	UpdateNotifier(context.Context, *UpdateNotifierRequest) (*UpdateNotifierResponse, error)
	DeleteNotifier(context.Context, *DeleteNotifierRequest) (*DeleteNotifierResponse, error)
}

// UnimplementedNotifierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotifierServiceServer struct {
}

func (*UnimplementedNotifierServiceServer) ListNotifiers(ctx context.Context, req *ListNotifiersRequest) (*ListNotifiersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifiers not implemented")
}
func (*UnimplementedNotifierServiceServer) CreateNotifier(ctx context.Context, req *CreateNotifierRequest) (*CreateNotifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotifier not implemented")
}
func (*UnimplementedNotifierServiceServer) GetNotifier(ctx context.Context, req *GetNotifierRequest) (*GetNotifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifier not implemented")
}
func (*UnimplementedNotifierServiceServer) UpdateNotifier(ctx context.Context, req *UpdateNotifierRequest) (*UpdateNotifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotifier not implemented")
}
func (*UnimplementedNotifierServiceServer) DeleteNotifier(ctx context.Context, req *DeleteNotifierRequest) (*DeleteNotifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotifier not implemented")
}

func RegisterNotifierServiceServer(s *grpc.Server, srv NotifierServiceServer) {
	s.RegisterService(&_NotifierService_serviceDesc, srv)
}

func _NotifierService_ListNotifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServiceServer).ListNotifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.notifier.NotifierService/ListNotifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServiceServer).ListNotifiers(ctx, req.(*ListNotifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifierService_CreateNotifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServiceServer).CreateNotifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.notifier.NotifierService/CreateNotifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServiceServer).CreateNotifier(ctx, req.(*CreateNotifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifierService_GetNotifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServiceServer).GetNotifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.notifier.NotifierService/GetNotifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServiceServer).GetNotifier(ctx, req.(*GetNotifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifierService_UpdateNotifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServiceServer).UpdateNotifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.notifier.NotifierService/UpdateNotifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServiceServer).UpdateNotifier(ctx, req.(*UpdateNotifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifierService_DeleteNotifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServiceServer).DeleteNotifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.metaprov.modeld.services.notifier.NotifierService/DeleteNotifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServiceServer).DeleteNotifier(ctx, req.(*DeleteNotifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.metaprov.modeld.services.notifier.NotifierService",
	HandlerType: (*NotifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNotifiers",
			Handler:    _NotifierService_ListNotifiers_Handler,
		},
		{
			MethodName: "CreateNotifier",
			Handler:    _NotifierService_CreateNotifier_Handler,
		},
		{
			MethodName: "GetNotifier",
			Handler:    _NotifierService_GetNotifier_Handler,
		},
		{
			MethodName: "UpdateNotifier",
			Handler:    _NotifierService_UpdateNotifier_Handler,
		},
		{
			MethodName: "DeleteNotifier",
			Handler:    _NotifierService_DeleteNotifier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/notifier/notifier.proto",
}
