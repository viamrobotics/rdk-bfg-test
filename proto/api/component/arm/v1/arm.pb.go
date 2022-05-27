// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/arm/v1/arm.proto

package v1

import (
	v1 "go.viam.com/rdk/proto/api/common/v1"
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

type GetEndPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an arm
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetEndPositionRequest) Reset() {
	*x = GetEndPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndPositionRequest) ProtoMessage() {}

func (x *GetEndPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndPositionRequest.ProtoReflect.Descriptor instead.
func (*GetEndPositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{0}
}

func (x *GetEndPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetEndPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returns 6d pose of the end effector relative to the base, represented by X,Y,Z coordinates which express
	// millimeters and theta, ox, oy, oz coordinates which express an orientation vector
	Pose *v1.Pose `protobuf:"bytes,1,opt,name=pose,proto3" json:"pose,omitempty"`
}

func (x *GetEndPositionResponse) Reset() {
	*x = GetEndPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEndPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndPositionResponse) ProtoMessage() {}

func (x *GetEndPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndPositionResponse.ProtoReflect.Descriptor instead.
func (*GetEndPositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{1}
}

func (x *GetEndPositionResponse) GetPose() *v1.Pose {
	if x != nil {
		return x.Pose
	}
	return nil
}

type JointPositions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of joint positions represented in degrees
	// The numbers are ordered spatially from the base toward the end effector
	// This is used in GetJointPositionsResponse and MoveToJointPositionsRequest
	Degrees []float64 `protobuf:"fixed64,1,rep,packed,name=degrees,proto3" json:"degrees,omitempty"`
}

func (x *JointPositions) Reset() {
	*x = JointPositions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JointPositions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JointPositions) ProtoMessage() {}

func (x *JointPositions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JointPositions.ProtoReflect.Descriptor instead.
func (*JointPositions) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{2}
}

func (x *JointPositions) GetDegrees() []float64 {
	if x != nil {
		return x.Degrees
	}
	return nil
}

type GetJointPositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an arm
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetJointPositionsRequest) Reset() {
	*x = GetJointPositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJointPositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJointPositionsRequest) ProtoMessage() {}

func (x *GetJointPositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJointPositionsRequest.ProtoReflect.Descriptor instead.
func (*GetJointPositionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{3}
}

func (x *GetJointPositionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetJointPositionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//a list JointPositions
	PositionDegs *JointPositions `protobuf:"bytes,1,opt,name=position_degs,json=positionDegs,proto3" json:"position_degs,omitempty"`
}

func (x *GetJointPositionsResponse) Reset() {
	*x = GetJointPositionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJointPositionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJointPositionsResponse) ProtoMessage() {}

func (x *GetJointPositionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJointPositionsResponse.ProtoReflect.Descriptor instead.
func (*GetJointPositionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{4}
}

func (x *GetJointPositionsResponse) GetPositionDegs() *JointPositions {
	if x != nil {
		return x.PositionDegs
	}
	return nil
}

type MoveToPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an arm
	Name       string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	To         *v1.Pose       `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	WorldState *v1.WorldState `protobuf:"bytes,3,opt,name=world_state,json=worldState,proto3,oneof" json:"world_state,omitempty"`
}

func (x *MoveToPositionRequest) Reset() {
	*x = MoveToPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveToPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToPositionRequest) ProtoMessage() {}

func (x *MoveToPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToPositionRequest.ProtoReflect.Descriptor instead.
func (*MoveToPositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{5}
}

func (x *MoveToPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MoveToPositionRequest) GetTo() *v1.Pose {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MoveToPositionRequest) GetWorldState() *v1.WorldState {
	if x != nil {
		return x.WorldState
	}
	return nil
}

type MoveToPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveToPositionResponse) Reset() {
	*x = MoveToPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveToPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToPositionResponse) ProtoMessage() {}

func (x *MoveToPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToPositionResponse.ProtoReflect.Descriptor instead.
func (*MoveToPositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{6}
}

type MoveToJointPositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an arm
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of joint positions represented in degrees
	// There should be 1 entry in the list per joint, ordered spatially from the base toward the end effector
	PositionDegs *JointPositions `protobuf:"bytes,2,opt,name=position_degs,json=positionDegs,proto3" json:"position_degs,omitempty"`
}

func (x *MoveToJointPositionsRequest) Reset() {
	*x = MoveToJointPositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveToJointPositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToJointPositionsRequest) ProtoMessage() {}

func (x *MoveToJointPositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToJointPositionsRequest.ProtoReflect.Descriptor instead.
func (*MoveToJointPositionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{7}
}

func (x *MoveToJointPositionsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MoveToJointPositionsRequest) GetPositionDegs() *JointPositions {
	if x != nil {
		return x.PositionDegs
	}
	return nil
}

type MoveToJointPositionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveToJointPositionsResponse) Reset() {
	*x = MoveToJointPositionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveToJointPositionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveToJointPositionsResponse) ProtoMessage() {}

func (x *MoveToJointPositionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveToJointPositionsResponse.ProtoReflect.Descriptor instead.
func (*MoveToJointPositionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{8}
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an arm
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{9}
}

func (x *StopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{10}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndPosition    *v1.Pose        `protobuf:"bytes,1,opt,name=end_position,json=endPosition,proto3" json:"end_position,omitempty"`
	JointPositions *JointPositions `protobuf:"bytes,2,opt,name=joint_positions,json=jointPositions,proto3" json:"joint_positions,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_arm_v1_arm_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_api_component_arm_v1_arm_proto_rawDescGZIP(), []int{11}
}

func (x *Status) GetEndPosition() *v1.Pose {
	if x != nil {
		return x.EndPosition
	}
	return nil
}

func (x *Status) GetJointPositions() *JointPositions {
	if x != nil {
		return x.JointPositions
	}
	return nil
}

var File_proto_api_component_arm_v1_arm_proto protoreflect.FileDescriptor

var file_proto_api_component_arm_v1_arm_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x47, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x6f, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x65, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x4a, 0x6f, 0x69, 0x6e,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x64, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x67, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x45, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x82, 0x01, 0x0a,
	0x1b, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x4f, 0x0a, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x67,
	0x73, 0x22, 0x1e, 0x0a, 0x1c, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3c, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x65,
	0x52, 0x0b, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a,
	0x0f, 0x6a, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0e, 0x6a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0xf9, 0x06, 0x0a, 0x0a, 0x41, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xab, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x6d, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0xab, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2c, 0x1a, 0x2a, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x6d, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xbb, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x31, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x61, 0x72, 0x6d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6a, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xc4, 0x01, 0x0a, 0x14,
	0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x54,
	0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x1a,
	0x31, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x6d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x27, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x26, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61,
	0x72, 0x6d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x55,
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x2a, 0x67, 0x6f, 0x2e, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x61,
	0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_component_arm_v1_arm_proto_rawDescOnce sync.Once
	file_proto_api_component_arm_v1_arm_proto_rawDescData = file_proto_api_component_arm_v1_arm_proto_rawDesc
)

func file_proto_api_component_arm_v1_arm_proto_rawDescGZIP() []byte {
	file_proto_api_component_arm_v1_arm_proto_rawDescOnce.Do(func() {
		file_proto_api_component_arm_v1_arm_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_arm_v1_arm_proto_rawDescData)
	})
	return file_proto_api_component_arm_v1_arm_proto_rawDescData
}

var file_proto_api_component_arm_v1_arm_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_api_component_arm_v1_arm_proto_goTypes = []interface{}{
	(*GetEndPositionRequest)(nil),        // 0: proto.api.component.arm.v1.GetEndPositionRequest
	(*GetEndPositionResponse)(nil),       // 1: proto.api.component.arm.v1.GetEndPositionResponse
	(*JointPositions)(nil),               // 2: proto.api.component.arm.v1.JointPositions
	(*GetJointPositionsRequest)(nil),     // 3: proto.api.component.arm.v1.GetJointPositionsRequest
	(*GetJointPositionsResponse)(nil),    // 4: proto.api.component.arm.v1.GetJointPositionsResponse
	(*MoveToPositionRequest)(nil),        // 5: proto.api.component.arm.v1.MoveToPositionRequest
	(*MoveToPositionResponse)(nil),       // 6: proto.api.component.arm.v1.MoveToPositionResponse
	(*MoveToJointPositionsRequest)(nil),  // 7: proto.api.component.arm.v1.MoveToJointPositionsRequest
	(*MoveToJointPositionsResponse)(nil), // 8: proto.api.component.arm.v1.MoveToJointPositionsResponse
	(*StopRequest)(nil),                  // 9: proto.api.component.arm.v1.StopRequest
	(*StopResponse)(nil),                 // 10: proto.api.component.arm.v1.StopResponse
	(*Status)(nil),                       // 11: proto.api.component.arm.v1.Status
	(*v1.Pose)(nil),                      // 12: proto.api.common.v1.Pose
	(*v1.WorldState)(nil),                // 13: proto.api.common.v1.WorldState
}
var file_proto_api_component_arm_v1_arm_proto_depIdxs = []int32{
	12, // 0: proto.api.component.arm.v1.GetEndPositionResponse.pose:type_name -> proto.api.common.v1.Pose
	2,  // 1: proto.api.component.arm.v1.GetJointPositionsResponse.position_degs:type_name -> proto.api.component.arm.v1.JointPositions
	12, // 2: proto.api.component.arm.v1.MoveToPositionRequest.to:type_name -> proto.api.common.v1.Pose
	13, // 3: proto.api.component.arm.v1.MoveToPositionRequest.world_state:type_name -> proto.api.common.v1.WorldState
	2,  // 4: proto.api.component.arm.v1.MoveToJointPositionsRequest.position_degs:type_name -> proto.api.component.arm.v1.JointPositions
	12, // 5: proto.api.component.arm.v1.Status.end_position:type_name -> proto.api.common.v1.Pose
	2,  // 6: proto.api.component.arm.v1.Status.joint_positions:type_name -> proto.api.component.arm.v1.JointPositions
	0,  // 7: proto.api.component.arm.v1.ArmService.GetEndPosition:input_type -> proto.api.component.arm.v1.GetEndPositionRequest
	5,  // 8: proto.api.component.arm.v1.ArmService.MoveToPosition:input_type -> proto.api.component.arm.v1.MoveToPositionRequest
	3,  // 9: proto.api.component.arm.v1.ArmService.GetJointPositions:input_type -> proto.api.component.arm.v1.GetJointPositionsRequest
	7,  // 10: proto.api.component.arm.v1.ArmService.MoveToJointPositions:input_type -> proto.api.component.arm.v1.MoveToJointPositionsRequest
	9,  // 11: proto.api.component.arm.v1.ArmService.Stop:input_type -> proto.api.component.arm.v1.StopRequest
	1,  // 12: proto.api.component.arm.v1.ArmService.GetEndPosition:output_type -> proto.api.component.arm.v1.GetEndPositionResponse
	6,  // 13: proto.api.component.arm.v1.ArmService.MoveToPosition:output_type -> proto.api.component.arm.v1.MoveToPositionResponse
	4,  // 14: proto.api.component.arm.v1.ArmService.GetJointPositions:output_type -> proto.api.component.arm.v1.GetJointPositionsResponse
	8,  // 15: proto.api.component.arm.v1.ArmService.MoveToJointPositions:output_type -> proto.api.component.arm.v1.MoveToJointPositionsResponse
	10, // 16: proto.api.component.arm.v1.ArmService.Stop:output_type -> proto.api.component.arm.v1.StopResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_api_component_arm_v1_arm_proto_init() }
func file_proto_api_component_arm_v1_arm_proto_init() {
	if File_proto_api_component_arm_v1_arm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_arm_v1_arm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndPositionRequest); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEndPositionResponse); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JointPositions); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJointPositionsRequest); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJointPositionsResponse); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveToPositionRequest); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveToPositionResponse); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveToJointPositionsRequest); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveToJointPositionsResponse); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
		file_proto_api_component_arm_v1_arm_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
	file_proto_api_component_arm_v1_arm_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_component_arm_v1_arm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_arm_v1_arm_proto_goTypes,
		DependencyIndexes: file_proto_api_component_arm_v1_arm_proto_depIdxs,
		MessageInfos:      file_proto_api_component_arm_v1_arm_proto_msgTypes,
	}.Build()
	File_proto_api_component_arm_v1_arm_proto = out.File
	file_proto_api_component_arm_v1_arm_proto_rawDesc = nil
	file_proto_api_component_arm_v1_arm_proto_goTypes = nil
	file_proto_api_component_arm_v1_arm_proto_depIdxs = nil
}
