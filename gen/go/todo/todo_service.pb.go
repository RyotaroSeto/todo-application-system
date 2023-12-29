// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: todo_service/todo_service.proto

package todo

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TodoStatus int32

const (
	TodoStatus_TODO_STATUS_UNSPECIFIED TodoStatus = 0
	TodoStatus_TODO_STATUS_DOING       TodoStatus = 1
	TodoStatus_TODO_STATUS_DONE        TodoStatus = 2
)

// Enum value maps for TodoStatus.
var (
	TodoStatus_name = map[int32]string{
		0: "TODO_STATUS_UNSPECIFIED",
		1: "TODO_STATUS_DOING",
		2: "TODO_STATUS_DONE",
	}
	TodoStatus_value = map[string]int32{
		"TODO_STATUS_UNSPECIFIED": 0,
		"TODO_STATUS_DOING":       1,
		"TODO_STATUS_DONE":        2,
	}
)

func (x TodoStatus) Enum() *TodoStatus {
	p := new(TodoStatus)
	*p = x
	return p
}

func (x TodoStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_service_todo_service_proto_enumTypes[0].Descriptor()
}

func (TodoStatus) Type() protoreflect.EnumType {
	return &file_todo_service_todo_service_proto_enumTypes[0]
}

func (x TodoStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoStatus.Descriptor instead.
func (TodoStatus) EnumDescriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{0}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{0}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Todo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetList() []*Todo {
	if x != nil {
		return x.List
	}
	return nil
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status TodoStatus `protobuf:"varint,2,opt,name=status,proto3,enum=todo_service.TodoStatus" json:"status,omitempty"`
}

func (x *UpdateStatusRequest) Reset() {
	*x = UpdateStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusRequest) ProtoMessage() {}

func (x *UpdateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStatusRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStatusRequest) GetStatus() TodoStatus {
	if x != nil {
		return x.Status
	}
	return TodoStatus_TODO_STATUS_UNSPECIFIED
}

type UpdateStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateStatusResponse) Reset() {
	*x = UpdateStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusResponse) ProtoMessage() {}

func (x *UpdateStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateStatusResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{5}
}

type UpdateTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateTitleRequest) Reset() {
	*x = UpdateTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTitleRequest) ProtoMessage() {}

func (x *UpdateTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTitleRequest.ProtoReflect.Descriptor instead.
func (*UpdateTitleRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTitleRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTitleResponse) Reset() {
	*x = UpdateTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTitleResponse) ProtoMessage() {}

func (x *UpdateTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTitleResponse.ProtoReflect.Descriptor instead.
func (*UpdateTitleResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{7}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{9}
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Status     TodoStatus `protobuf:"varint,3,opt,name=status,proto3,enum=todo_service.TodoStatus" json:"status,omitempty"`
	StatusName string     `protobuf:"bytes,4,opt,name=status_name,json=statusName,proto3" json:"status_name,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_service_todo_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_service_todo_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_service_todo_service_proto_rawDescGZIP(), []int{10}
}

func (x *Todo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetStatus() TodoStatus {
	if x != nil {
		return x.Status
	}
	return TodoStatus_TODO_STATUS_UNSPECIFIED
}

func (x *Todo) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

var File_todo_service_todo_service_proto protoreflect.FileDescriptor

var file_todo_service_todo_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f,
	0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x2a,
	0x56, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x17, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f,
	0x44, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x32, 0xf8, 0x01, 0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f,
	0x41, 0x70, 0x69, 0x12, 0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x49,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x07, 0x22, 0x05, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x57, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0x7d, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x10, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0b,
	0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x0b, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x17, 0x54, 0x6f, 0x64, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_service_todo_service_proto_rawDescOnce sync.Once
	file_todo_service_todo_service_proto_rawDescData = file_todo_service_todo_service_proto_rawDesc
)

func file_todo_service_todo_service_proto_rawDescGZIP() []byte {
	file_todo_service_todo_service_proto_rawDescOnce.Do(func() {
		file_todo_service_todo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_service_todo_service_proto_rawDescData)
	})
	return file_todo_service_todo_service_proto_rawDescData
}

var file_todo_service_todo_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_service_todo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_todo_service_todo_service_proto_goTypes = []interface{}{
	(TodoStatus)(0),              // 0: todo_service.TodoStatus
	(*GetRequest)(nil),           // 1: todo_service.GetRequest
	(*GetResponse)(nil),          // 2: todo_service.GetResponse
	(*AddRequest)(nil),           // 3: todo_service.AddRequest
	(*AddResponse)(nil),          // 4: todo_service.AddResponse
	(*UpdateStatusRequest)(nil),  // 5: todo_service.UpdateStatusRequest
	(*UpdateStatusResponse)(nil), // 6: todo_service.UpdateStatusResponse
	(*UpdateTitleRequest)(nil),   // 7: todo_service.UpdateTitleRequest
	(*UpdateTitleResponse)(nil),  // 8: todo_service.UpdateTitleResponse
	(*DeleteRequest)(nil),        // 9: todo_service.DeleteRequest
	(*DeleteResponse)(nil),       // 10: todo_service.DeleteResponse
	(*Todo)(nil),                 // 11: todo_service.Todo
}
var file_todo_service_todo_service_proto_depIdxs = []int32{
	11, // 0: todo_service.GetResponse.list:type_name -> todo_service.Todo
	11, // 1: todo_service.AddRequest.todo:type_name -> todo_service.Todo
	0,  // 2: todo_service.UpdateStatusRequest.status:type_name -> todo_service.TodoStatus
	0,  // 3: todo_service.Todo.status:type_name -> todo_service.TodoStatus
	1,  // 4: todo_service.TodoApi.Get:input_type -> todo_service.GetRequest
	3,  // 5: todo_service.TodoApi.Add:input_type -> todo_service.AddRequest
	9,  // 6: todo_service.TodoApi.Delete:input_type -> todo_service.DeleteRequest
	2,  // 7: todo_service.TodoApi.Get:output_type -> todo_service.GetResponse
	4,  // 8: todo_service.TodoApi.Add:output_type -> todo_service.AddResponse
	10, // 9: todo_service.TodoApi.Delete:output_type -> todo_service.DeleteResponse
	7,  // [7:10] is the sub-list for method output_type
	4,  // [4:7] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_todo_service_todo_service_proto_init() }
func file_todo_service_todo_service_proto_init() {
	if File_todo_service_todo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_service_todo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_todo_service_todo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_todo_service_todo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_todo_service_todo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_todo_service_todo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatusRequest); i {
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
		file_todo_service_todo_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatusResponse); i {
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
		file_todo_service_todo_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTitleRequest); i {
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
		file_todo_service_todo_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTitleResponse); i {
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
		file_todo_service_todo_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_todo_service_todo_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_todo_service_todo_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_todo_service_todo_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_service_todo_service_proto_goTypes,
		DependencyIndexes: file_todo_service_todo_service_proto_depIdxs,
		EnumInfos:         file_todo_service_todo_service_proto_enumTypes,
		MessageInfos:      file_todo_service_todo_service_proto_msgTypes,
	}.Build()
	File_todo_service_todo_service_proto = out.File
	file_todo_service_todo_service_proto_rawDesc = nil
	file_todo_service_todo_service_proto_goTypes = nil
	file_todo_service_todo_service_proto_depIdxs = nil
}