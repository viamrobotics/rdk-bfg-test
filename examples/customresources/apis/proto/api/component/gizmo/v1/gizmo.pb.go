// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api/component/gizmo/v1/gizmo.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DoOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arg1 string `protobuf:"bytes,2,opt,name=arg1,proto3" json:"arg1,omitempty"`
}

func (x *DoOneRequest) Reset() {
	*x = DoOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneRequest) ProtoMessage() {}

func (x *DoOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneRequest.ProtoReflect.Descriptor instead.
func (*DoOneRequest) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{0}
}

func (x *DoOneRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoOneRequest) GetArg1() string {
	if x != nil {
		return x.Arg1
	}
	return ""
}

type DoOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret1 bool `protobuf:"varint,1,opt,name=ret1,proto3" json:"ret1,omitempty"`
}

func (x *DoOneResponse) Reset() {
	*x = DoOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneResponse) ProtoMessage() {}

func (x *DoOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneResponse.ProtoReflect.Descriptor instead.
func (*DoOneResponse) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{1}
}

func (x *DoOneResponse) GetRet1() bool {
	if x != nil {
		return x.Ret1
	}
	return false
}

type DoOneServerStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arg1 string `protobuf:"bytes,2,opt,name=arg1,proto3" json:"arg1,omitempty"`
}

func (x *DoOneServerStreamRequest) Reset() {
	*x = DoOneServerStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneServerStreamRequest) ProtoMessage() {}

func (x *DoOneServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneServerStreamRequest.ProtoReflect.Descriptor instead.
func (*DoOneServerStreamRequest) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{2}
}

func (x *DoOneServerStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoOneServerStreamRequest) GetArg1() string {
	if x != nil {
		return x.Arg1
	}
	return ""
}

type DoOneServerStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret1 bool `protobuf:"varint,1,opt,name=ret1,proto3" json:"ret1,omitempty"`
}

func (x *DoOneServerStreamResponse) Reset() {
	*x = DoOneServerStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneServerStreamResponse) ProtoMessage() {}

func (x *DoOneServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneServerStreamResponse.ProtoReflect.Descriptor instead.
func (*DoOneServerStreamResponse) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{3}
}

func (x *DoOneServerStreamResponse) GetRet1() bool {
	if x != nil {
		return x.Ret1
	}
	return false
}

type DoOneClientStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arg1 string `protobuf:"bytes,2,opt,name=arg1,proto3" json:"arg1,omitempty"`
}

func (x *DoOneClientStreamRequest) Reset() {
	*x = DoOneClientStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneClientStreamRequest) ProtoMessage() {}

func (x *DoOneClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneClientStreamRequest.ProtoReflect.Descriptor instead.
func (*DoOneClientStreamRequest) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{4}
}

func (x *DoOneClientStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoOneClientStreamRequest) GetArg1() string {
	if x != nil {
		return x.Arg1
	}
	return ""
}

type DoOneClientStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret1 bool `protobuf:"varint,1,opt,name=ret1,proto3" json:"ret1,omitempty"`
}

func (x *DoOneClientStreamResponse) Reset() {
	*x = DoOneClientStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneClientStreamResponse) ProtoMessage() {}

func (x *DoOneClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneClientStreamResponse.ProtoReflect.Descriptor instead.
func (*DoOneClientStreamResponse) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{5}
}

func (x *DoOneClientStreamResponse) GetRet1() bool {
	if x != nil {
		return x.Ret1
	}
	return false
}

type DoOneBiDiStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arg1 string `protobuf:"bytes,2,opt,name=arg1,proto3" json:"arg1,omitempty"`
}

func (x *DoOneBiDiStreamRequest) Reset() {
	*x = DoOneBiDiStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneBiDiStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneBiDiStreamRequest) ProtoMessage() {}

func (x *DoOneBiDiStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneBiDiStreamRequest.ProtoReflect.Descriptor instead.
func (*DoOneBiDiStreamRequest) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{6}
}

func (x *DoOneBiDiStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoOneBiDiStreamRequest) GetArg1() string {
	if x != nil {
		return x.Arg1
	}
	return ""
}

type DoOneBiDiStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret1 bool `protobuf:"varint,1,opt,name=ret1,proto3" json:"ret1,omitempty"`
}

func (x *DoOneBiDiStreamResponse) Reset() {
	*x = DoOneBiDiStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoOneBiDiStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoOneBiDiStreamResponse) ProtoMessage() {}

func (x *DoOneBiDiStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoOneBiDiStreamResponse.ProtoReflect.Descriptor instead.
func (*DoOneBiDiStreamResponse) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{7}
}

func (x *DoOneBiDiStreamResponse) GetRet1() bool {
	if x != nil {
		return x.Ret1
	}
	return false
}

type DoTwoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arg1 bool   `protobuf:"varint,2,opt,name=arg1,proto3" json:"arg1,omitempty"`
}

func (x *DoTwoRequest) Reset() {
	*x = DoTwoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTwoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTwoRequest) ProtoMessage() {}

func (x *DoTwoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTwoRequest.ProtoReflect.Descriptor instead.
func (*DoTwoRequest) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{8}
}

func (x *DoTwoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoTwoRequest) GetArg1() bool {
	if x != nil {
		return x.Arg1
	}
	return false
}

type DoTwoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret1 string `protobuf:"bytes,1,opt,name=ret1,proto3" json:"ret1,omitempty"`
}

func (x *DoTwoResponse) Reset() {
	*x = DoTwoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTwoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTwoResponse) ProtoMessage() {}

func (x *DoTwoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTwoResponse.ProtoReflect.Descriptor instead.
func (*DoTwoResponse) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{9}
}

func (x *DoTwoResponse) GetRet1() string {
	if x != nil {
		return x.Ret1
	}
	return ""
}

type DoCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command *structpb.Struct `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *DoCommandRequest) Reset() {
	*x = DoCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoCommandRequest) ProtoMessage() {}

func (x *DoCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoCommandRequest.ProtoReflect.Descriptor instead.
func (*DoCommandRequest) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{10}
}

func (x *DoCommandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoCommandRequest) GetCommand() *structpb.Struct {
	if x != nil {
		return x.Command
	}
	return nil
}

type DoCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *structpb.Struct `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DoCommandResponse) Reset() {
	*x = DoCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoCommandResponse) ProtoMessage() {}

func (x *DoCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_component_gizmo_v1_gizmo_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoCommandResponse.ProtoReflect.Descriptor instead.
func (*DoCommandResponse) Descriptor() ([]byte, []int) {
	return file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP(), []int{11}
}

func (x *DoCommandResponse) GetResult() *structpb.Struct {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_component_gizmo_v1_gizmo_proto protoreflect.FileDescriptor

var file_api_component_gizmo_v1_gizmo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f,
	0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0c, 0x44, 0x6f, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x67, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67,
	0x31, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x72, 0x65, 0x74, 0x31, 0x22, 0x42, 0x0a, 0x18, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x31, 0x22, 0x2f, 0x0a, 0x19, 0x44, 0x6f,
	0x4f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x74, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x74, 0x31, 0x22, 0x42, 0x0a, 0x18, 0x44,
	0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x31, 0x22,
	0x2f, 0x0a, 0x19, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x74, 0x31,
	0x22, 0x40, 0x0a, 0x16, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x42, 0x69, 0x44, 0x69, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x31, 0x22, 0x2d, 0x0a, 0x17, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x42, 0x69, 0x44, 0x69, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x65, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x74,
	0x31, 0x22, 0x36, 0x0a, 0x0c, 0x44, 0x6f, 0x54, 0x77, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x31, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x61, 0x72, 0x67, 0x31, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x6f, 0x54,
	0x77, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65,
	0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x74, 0x31, 0x22, 0x59,
	0x0a, 0x10, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x44, 0x0a, 0x11, 0x44, 0x6f, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0xbb, 0x06, 0x0a, 0x0c, 0x47, 0x69, 0x7a, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x8a, 0x01, 0x0a, 0x05, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x63, 0x6d,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2c, 0x22, 0x2a, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x5f, 0x6f, 0x6e, 0x65, 0x12, 0x7c, 0x0a,
	0x11, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x31, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x4f,
	0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x7c, 0x0a, 0x11, 0x44,
	0x6f, 0x4f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x31, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x4f, 0x6e, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f,
	0x4f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x78, 0x0a, 0x0f, 0x44, 0x6f, 0x4f,
	0x6e, 0x65, 0x42, 0x69, 0x44, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x2e, 0x61,
	0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69,
	0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x42, 0x69, 0x44, 0x69,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x4f, 0x6e, 0x65, 0x42, 0x69, 0x44,
	0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x8a, 0x01, 0x0a, 0x05, 0x44, 0x6f, 0x54, 0x77, 0x6f, 0x12, 0x25, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x54, 0x77, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6f, 0x54, 0x77, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2c, 0x22, 0x2a, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x69, 0x7a,
	0x6d, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x5f, 0x74, 0x77, 0x6f,
	0x12, 0x9a, 0x01, 0x0a, 0x09, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x29,
	0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x63, 0x6d, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x69, 0x7a, 0x6d, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2e, 0x2f,
	0x61, 0x63, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x6f, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x67, 0x69, 0x7a, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_component_gizmo_v1_gizmo_proto_rawDescOnce sync.Once
	file_api_component_gizmo_v1_gizmo_proto_rawDescData = file_api_component_gizmo_v1_gizmo_proto_rawDesc
)

func file_api_component_gizmo_v1_gizmo_proto_rawDescGZIP() []byte {
	file_api_component_gizmo_v1_gizmo_proto_rawDescOnce.Do(func() {
		file_api_component_gizmo_v1_gizmo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_component_gizmo_v1_gizmo_proto_rawDescData)
	})
	return file_api_component_gizmo_v1_gizmo_proto_rawDescData
}

var file_api_component_gizmo_v1_gizmo_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_component_gizmo_v1_gizmo_proto_goTypes = []interface{}{
	(*DoOneRequest)(nil),              // 0: acme.component.gizmo.v1.DoOneRequest
	(*DoOneResponse)(nil),             // 1: acme.component.gizmo.v1.DoOneResponse
	(*DoOneServerStreamRequest)(nil),  // 2: acme.component.gizmo.v1.DoOneServerStreamRequest
	(*DoOneServerStreamResponse)(nil), // 3: acme.component.gizmo.v1.DoOneServerStreamResponse
	(*DoOneClientStreamRequest)(nil),  // 4: acme.component.gizmo.v1.DoOneClientStreamRequest
	(*DoOneClientStreamResponse)(nil), // 5: acme.component.gizmo.v1.DoOneClientStreamResponse
	(*DoOneBiDiStreamRequest)(nil),    // 6: acme.component.gizmo.v1.DoOneBiDiStreamRequest
	(*DoOneBiDiStreamResponse)(nil),   // 7: acme.component.gizmo.v1.DoOneBiDiStreamResponse
	(*DoTwoRequest)(nil),              // 8: acme.component.gizmo.v1.DoTwoRequest
	(*DoTwoResponse)(nil),             // 9: acme.component.gizmo.v1.DoTwoResponse
	(*DoCommandRequest)(nil),          // 10: acme.component.gizmo.v1.DoCommandRequest
	(*DoCommandResponse)(nil),         // 11: acme.component.gizmo.v1.DoCommandResponse
	(*structpb.Struct)(nil),           // 12: google.protobuf.Struct
}
var file_api_component_gizmo_v1_gizmo_proto_depIdxs = []int32{
	12, // 0: acme.component.gizmo.v1.DoCommandRequest.command:type_name -> google.protobuf.Struct
	12, // 1: acme.component.gizmo.v1.DoCommandResponse.result:type_name -> google.protobuf.Struct
	0,  // 2: acme.component.gizmo.v1.GizmoService.DoOne:input_type -> acme.component.gizmo.v1.DoOneRequest
	4,  // 3: acme.component.gizmo.v1.GizmoService.DoOneClientStream:input_type -> acme.component.gizmo.v1.DoOneClientStreamRequest
	2,  // 4: acme.component.gizmo.v1.GizmoService.DoOneServerStream:input_type -> acme.component.gizmo.v1.DoOneServerStreamRequest
	6,  // 5: acme.component.gizmo.v1.GizmoService.DoOneBiDiStream:input_type -> acme.component.gizmo.v1.DoOneBiDiStreamRequest
	8,  // 6: acme.component.gizmo.v1.GizmoService.DoTwo:input_type -> acme.component.gizmo.v1.DoTwoRequest
	10, // 7: acme.component.gizmo.v1.GizmoService.DoCommand:input_type -> acme.component.gizmo.v1.DoCommandRequest
	1,  // 8: acme.component.gizmo.v1.GizmoService.DoOne:output_type -> acme.component.gizmo.v1.DoOneResponse
	5,  // 9: acme.component.gizmo.v1.GizmoService.DoOneClientStream:output_type -> acme.component.gizmo.v1.DoOneClientStreamResponse
	3,  // 10: acme.component.gizmo.v1.GizmoService.DoOneServerStream:output_type -> acme.component.gizmo.v1.DoOneServerStreamResponse
	7,  // 11: acme.component.gizmo.v1.GizmoService.DoOneBiDiStream:output_type -> acme.component.gizmo.v1.DoOneBiDiStreamResponse
	9,  // 12: acme.component.gizmo.v1.GizmoService.DoTwo:output_type -> acme.component.gizmo.v1.DoTwoResponse
	11, // 13: acme.component.gizmo.v1.GizmoService.DoCommand:output_type -> acme.component.gizmo.v1.DoCommandResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_component_gizmo_v1_gizmo_proto_init() }
func file_api_component_gizmo_v1_gizmo_proto_init() {
	if File_api_component_gizmo_v1_gizmo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneRequest); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneResponse); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneServerStreamRequest); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneServerStreamResponse); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneClientStreamRequest); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneClientStreamResponse); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneBiDiStreamRequest); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoOneBiDiStreamResponse); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTwoRequest); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTwoResponse); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoCommandRequest); i {
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
		file_api_component_gizmo_v1_gizmo_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoCommandResponse); i {
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
			RawDescriptor: file_api_component_gizmo_v1_gizmo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_component_gizmo_v1_gizmo_proto_goTypes,
		DependencyIndexes: file_api_component_gizmo_v1_gizmo_proto_depIdxs,
		MessageInfos:      file_api_component_gizmo_v1_gizmo_proto_msgTypes,
	}.Build()
	File_api_component_gizmo_v1_gizmo_proto = out.File
	file_api_component_gizmo_v1_gizmo_proto_rawDesc = nil
	file_api_component_gizmo_v1_gizmo_proto_goTypes = nil
	file_api_component_gizmo_v1_gizmo_proto_depIdxs = nil
}
