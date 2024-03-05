// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: proto/private.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePrivateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type    uint32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreatePrivateDataRequest) Reset() {
	*x = CreatePrivateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrivateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivateDataRequest) ProtoMessage() {}

func (x *CreatePrivateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivateDataRequest.ProtoReflect.Descriptor instead.
func (*CreatePrivateDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePrivateDataRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePrivateDataRequest) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreatePrivateDataRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreatePrivateDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type    uint32                 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Updated *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted bool                   `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *CreatePrivateDataResponse) Reset() {
	*x = CreatePrivateDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePrivateDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivateDataResponse) ProtoMessage() {}

func (x *CreatePrivateDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivateDataResponse.ProtoReflect.Descriptor instead.
func (*CreatePrivateDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePrivateDataResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreatePrivateDataResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePrivateDataResponse) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreatePrivateDataResponse) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *CreatePrivateDataResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type GetPrivateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPrivateDataRequest) Reset() {
	*x = GetPrivateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateDataRequest) ProtoMessage() {}

func (x *GetPrivateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateDataRequest.ProtoReflect.Descriptor instead.
func (*GetPrivateDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{2}
}

func (x *GetPrivateDataRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPrivateDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type    uint32                 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Content []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Updated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted bool                   `protobuf:"varint,6,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *GetPrivateDataResponse) Reset() {
	*x = GetPrivateDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateDataResponse) ProtoMessage() {}

func (x *GetPrivateDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateDataResponse.ProtoReflect.Descriptor instead.
func (*GetPrivateDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{3}
}

func (x *GetPrivateDataResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPrivateDataResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetPrivateDataResponse) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GetPrivateDataResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *GetPrivateDataResponse) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *GetPrivateDataResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type UpdatePrivateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type      uint32                 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Content   []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdatePrivateDataRequest) Reset() {
	*x = UpdatePrivateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrivateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrivateDataRequest) ProtoMessage() {}

func (x *UpdatePrivateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrivateDataRequest.ProtoReflect.Descriptor instead.
func (*UpdatePrivateDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePrivateDataRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePrivateDataRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePrivateDataRequest) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *UpdatePrivateDataRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UpdatePrivateDataRequest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdatePrivateDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type    uint32                 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Content []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Updated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted bool                   `protobuf:"varint,6,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *UpdatePrivateDataResponse) Reset() {
	*x = UpdatePrivateDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrivateDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrivateDataResponse) ProtoMessage() {}

func (x *UpdatePrivateDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrivateDataResponse.ProtoReflect.Descriptor instead.
func (*UpdatePrivateDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePrivateDataResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePrivateDataResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePrivateDataResponse) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *UpdatePrivateDataResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UpdatePrivateDataResponse) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *UpdatePrivateDataResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type DeletePrivateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePrivateDataRequest) Reset() {
	*x = DeletePrivateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePrivateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePrivateDataRequest) ProtoMessage() {}

func (x *DeletePrivateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePrivateDataRequest.ProtoReflect.Descriptor instead.
func (*DeletePrivateDataRequest) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePrivateDataRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePrivateDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePrivateDataResponse) Reset() {
	*x = DeletePrivateDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePrivateDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePrivateDataResponse) ProtoMessage() {}

func (x *DeletePrivateDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePrivateDataResponse.ProtoReflect.Descriptor instead.
func (*DeletePrivateDataResponse) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{7}
}

type PrivateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type    uint32                 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Content []byte                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Updated *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted bool                   `protobuf:"varint,6,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *PrivateData) Reset() {
	*x = PrivateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateData) ProtoMessage() {}

func (x *PrivateData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateData.ProtoReflect.Descriptor instead.
func (*PrivateData) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{8}
}

func (x *PrivateData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PrivateData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PrivateData) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PrivateData) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *PrivateData) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *PrivateData) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type GetPrivateDataByTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId uint32 `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
}

func (x *GetPrivateDataByTypeRequest) Reset() {
	*x = GetPrivateDataByTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateDataByTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateDataByTypeRequest) ProtoMessage() {}

func (x *GetPrivateDataByTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateDataByTypeRequest.ProtoReflect.Descriptor instead.
func (*GetPrivateDataByTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{9}
}

func (x *GetPrivateDataByTypeRequest) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type GetPrivateDataByTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*PrivateData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPrivateDataByTypeResponse) Reset() {
	*x = GetPrivateDataByTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_private_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateDataByTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateDataByTypeResponse) ProtoMessage() {}

func (x *GetPrivateDataByTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_private_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateDataByTypeResponse.ProtoReflect.Descriptor instead.
func (*GetPrivateDataByTypeResponse) Descriptor() ([]byte, []int) {
	return file_proto_private_proto_rawDescGZIP(), []int{10}
}

func (x *GetPrivateDataByTypeResponse) GetData() []*PrivateData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_private_proto protoreflect.FileDescriptor

var file_proto_private_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xa5, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbc,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xa9, 0x01,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x22, 0x46, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc1, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x70, 0x72, 0x69,
	0x79, 0x61, 0x6e, 0x6f, 0x76, 0x6b, 0x6b, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_private_proto_rawDescOnce sync.Once
	file_proto_private_proto_rawDescData = file_proto_private_proto_rawDesc
)

func file_proto_private_proto_rawDescGZIP() []byte {
	file_proto_private_proto_rawDescOnce.Do(func() {
		file_proto_private_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_private_proto_rawDescData)
	})
	return file_proto_private_proto_rawDescData
}

var file_proto_private_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_private_proto_goTypes = []interface{}{
	(*CreatePrivateDataRequest)(nil),     // 0: proto.CreatePrivateDataRequest
	(*CreatePrivateDataResponse)(nil),    // 1: proto.CreatePrivateDataResponse
	(*GetPrivateDataRequest)(nil),        // 2: proto.GetPrivateDataRequest
	(*GetPrivateDataResponse)(nil),       // 3: proto.GetPrivateDataResponse
	(*UpdatePrivateDataRequest)(nil),     // 4: proto.UpdatePrivateDataRequest
	(*UpdatePrivateDataResponse)(nil),    // 5: proto.UpdatePrivateDataResponse
	(*DeletePrivateDataRequest)(nil),     // 6: proto.DeletePrivateDataRequest
	(*DeletePrivateDataResponse)(nil),    // 7: proto.DeletePrivateDataResponse
	(*PrivateData)(nil),                  // 8: proto.PrivateData
	(*GetPrivateDataByTypeRequest)(nil),  // 9: proto.GetPrivateDataByTypeRequest
	(*GetPrivateDataByTypeResponse)(nil), // 10: proto.GetPrivateDataByTypeResponse
	(*timestamppb.Timestamp)(nil),        // 11: google.protobuf.Timestamp
}
var file_proto_private_proto_depIdxs = []int32{
	11, // 0: proto.CreatePrivateDataResponse.updated:type_name -> google.protobuf.Timestamp
	11, // 1: proto.GetPrivateDataResponse.updated:type_name -> google.protobuf.Timestamp
	11, // 2: proto.UpdatePrivateDataRequest.updated_at:type_name -> google.protobuf.Timestamp
	11, // 3: proto.UpdatePrivateDataResponse.updated:type_name -> google.protobuf.Timestamp
	11, // 4: proto.PrivateData.updated:type_name -> google.protobuf.Timestamp
	8,  // 5: proto.GetPrivateDataByTypeResponse.data:type_name -> proto.PrivateData
	0,  // 6: proto.Private.CreatePrivateData:input_type -> proto.CreatePrivateDataRequest
	2,  // 7: proto.Private.GetPrivateData:input_type -> proto.GetPrivateDataRequest
	4,  // 8: proto.Private.UpdatePrivateData:input_type -> proto.UpdatePrivateDataRequest
	6,  // 9: proto.Private.DeletePrivateData:input_type -> proto.DeletePrivateDataRequest
	9,  // 10: proto.Private.GetPrivateDataByType:input_type -> proto.GetPrivateDataByTypeRequest
	1,  // 11: proto.Private.CreatePrivateData:output_type -> proto.CreatePrivateDataResponse
	3,  // 12: proto.Private.GetPrivateData:output_type -> proto.GetPrivateDataResponse
	5,  // 13: proto.Private.UpdatePrivateData:output_type -> proto.UpdatePrivateDataResponse
	7,  // 14: proto.Private.DeletePrivateData:output_type -> proto.DeletePrivateDataResponse
	10, // 15: proto.Private.GetPrivateDataByType:output_type -> proto.GetPrivateDataByTypeResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_private_proto_init() }
func file_proto_private_proto_init() {
	if File_proto_private_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_private_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePrivateDataRequest); i {
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
		file_proto_private_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePrivateDataResponse); i {
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
		file_proto_private_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateDataRequest); i {
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
		file_proto_private_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateDataResponse); i {
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
		file_proto_private_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrivateDataRequest); i {
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
		file_proto_private_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrivateDataResponse); i {
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
		file_proto_private_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePrivateDataRequest); i {
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
		file_proto_private_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePrivateDataResponse); i {
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
		file_proto_private_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateData); i {
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
		file_proto_private_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateDataByTypeRequest); i {
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
		file_proto_private_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateDataByTypeResponse); i {
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
			RawDescriptor: file_proto_private_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_private_proto_goTypes,
		DependencyIndexes: file_proto_private_proto_depIdxs,
		MessageInfos:      file_proto_private_proto_msgTypes,
	}.Build()
	File_proto_private_proto = out.File
	file_proto_private_proto_rawDesc = nil
	file_proto_private_proto_goTypes = nil
	file_proto_private_proto_depIdxs = nil
}
