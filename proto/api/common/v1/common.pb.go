// common.proto should hold messages that are used across multiple resource subtypes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/common/v1/common.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// millimeters of the end effector from the base
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	// ox, oy, oz, theta represents an orientation vector
	// Structured similarly to an angle axis, an orientation vector works differently. Rather than representing an orientation
	// with an arbitrary axis and a rotation around it from an origin, an orientation vector represents orientation
	// such that the ox/oy/oz components represent the point on the cartesian unit sphere at which your end effector is pointing
	// from the origin, and that unit vector forms an axis around which theta rotates. This means that incrementing/decrementing
	// theta will perform an in-line rotation of the end effector.
	// Theta is defined as rotation between two planes: the plane defined by the origin, the point (0,0,1), and the rx,ry,rz
	// point, and the plane defined by the origin, the rx,ry,rz point, and the new local Z axis. So if theta is kept at
	// zero as the north/south pole is circled, the Roll will correct itself to remain in-line.
	// Theta in pb.Pose should be degrees. It will be converted to radians in the internal OrientationVec.
	OX    float64 `protobuf:"fixed64,4,opt,name=o_x,json=oX,proto3" json:"o_x,omitempty"`
	OY    float64 `protobuf:"fixed64,5,opt,name=o_y,json=oY,proto3" json:"o_y,omitempty"`
	OZ    float64 `protobuf:"fixed64,6,opt,name=o_z,json=oZ,proto3" json:"o_z,omitempty"`
	Theta float64 `protobuf:"fixed64,7,opt,name=theta,proto3" json:"theta,omitempty"`
}

func (x *Pose) Reset() {
	*x = Pose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pose) ProtoMessage() {}

func (x *Pose) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pose.ProtoReflect.Descriptor instead.
func (*Pose) Descriptor() ([]byte, []int) {
	return file_proto_api_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Pose) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Pose) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Pose) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Pose) GetOX() float64 {
	if x != nil {
		return x.OX
	}
	return 0
}

func (x *Pose) GetOY() float64 {
	if x != nil {
		return x.OY
	}
	return 0
}

func (x *Pose) GetOZ() float64 {
	if x != nil {
		return x.OZ
	}
	return 0
}

func (x *Pose) GetTheta() float64 {
	if x != nil {
		return x.Theta
	}
	return 0
}

type Vector3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vector3) Reset() {
	*x = Vector3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector3) ProtoMessage() {}

func (x *Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector3.ProtoReflect.Descriptor instead.
func (*Vector3) Descriptor() ([]byte, []int) {
	return file_proto_api_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *Vector3) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector3) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vector3) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type BoxGeometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  float64 `protobuf:"fixed64,1,opt,name=width,proto3" json:"width,omitempty"`
	Length float64 `protobuf:"fixed64,2,opt,name=length,proto3" json:"length,omitempty"`
	Depth  float64 `protobuf:"fixed64,3,opt,name=depth,proto3" json:"depth,omitempty"`
}

func (x *BoxGeometry) Reset() {
	*x = BoxGeometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_common_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxGeometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxGeometry) ProtoMessage() {}

func (x *BoxGeometry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_common_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxGeometry.ProtoReflect.Descriptor instead.
func (*BoxGeometry) Descriptor() ([]byte, []int) {
	return file_proto_api_common_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *BoxGeometry) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BoxGeometry) GetLength() float64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *BoxGeometry) GetDepth() float64 {
	if x != nil {
		return x.Depth
	}
	return 0
}

var File_proto_api_common_v1_common_proto protoreflect.FileDescriptor

var file_proto_api_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x79, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x12, 0x0f, 0x0a, 0x03, 0x6f, 0x5f, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x6f, 0x58, 0x12, 0x0f, 0x0a, 0x03, 0x6f, 0x5f,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x6f, 0x59, 0x12, 0x0f, 0x0a, 0x03, 0x6f,
	0x5f, 0x7a, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x6f, 0x5a, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x68, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x68, 0x65,
	0x74, 0x61, 0x22, 0x33, 0x0a, 0x07, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x22, 0x51, 0x0a, 0x0b, 0x42, 0x6f, 0x78, 0x47, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x42, 0x47, 0x0a, 0x20, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x23,
	0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_common_v1_common_proto_rawDescOnce sync.Once
	file_proto_api_common_v1_common_proto_rawDescData = file_proto_api_common_v1_common_proto_rawDesc
)

func file_proto_api_common_v1_common_proto_rawDescGZIP() []byte {
	file_proto_api_common_v1_common_proto_rawDescOnce.Do(func() {
		file_proto_api_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_common_v1_common_proto_rawDescData)
	})
	return file_proto_api_common_v1_common_proto_rawDescData
}

var file_proto_api_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_api_common_v1_common_proto_goTypes = []interface{}{
	(*Pose)(nil),        // 0: proto.api.common.v1.Pose
	(*Vector3)(nil),     // 1: proto.api.common.v1.Vector3
	(*BoxGeometry)(nil), // 2: proto.api.common.v1.BoxGeometry
}
var file_proto_api_common_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_api_common_v1_common_proto_init() }
func file_proto_api_common_v1_common_proto_init() {
	if File_proto_api_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pose); i {
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
		file_proto_api_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector3); i {
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
		file_proto_api_common_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxGeometry); i {
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
			RawDescriptor: file_proto_api_common_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_api_common_v1_common_proto_goTypes,
		DependencyIndexes: file_proto_api_common_v1_common_proto_depIdxs,
		MessageInfos:      file_proto_api_common_v1_common_proto_msgTypes,
	}.Build()
	File_proto_api_common_v1_common_proto = out.File
	file_proto_api_common_v1_common_proto_rawDesc = nil
	file_proto_api_common_v1_common_proto_goTypes = nil
	file_proto_api_common_v1_common_proto_depIdxs = nil
}
