// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/v1/camera.proto

package v1

import (
	v1 "go.viam.com/rdk/proto/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
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

type CameraServiceGetFrameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a camera
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Requested MIME type of response
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *CameraServiceGetFrameRequest) Reset() {
	*x = CameraServiceGetFrameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceGetFrameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceGetFrameRequest) ProtoMessage() {}

func (x *CameraServiceGetFrameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceGetFrameRequest.ProtoReflect.Descriptor instead.
func (*CameraServiceGetFrameRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{0}
}

func (x *CameraServiceGetFrameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CameraServiceGetFrameRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type CameraServiceGetFrameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actual MIME type of response
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Frame in bytes
	Frame []byte `protobuf:"bytes,2,opt,name=frame,proto3" json:"frame,omitempty"`
	// Width of frame in px
	WidthPx int64 `protobuf:"varint,3,opt,name=width_px,json=widthPx,proto3" json:"width_px,omitempty"`
	// Height of frame in px
	HeightPx int64 `protobuf:"varint,4,opt,name=height_px,json=heightPx,proto3" json:"height_px,omitempty"`
}

func (x *CameraServiceGetFrameResponse) Reset() {
	*x = CameraServiceGetFrameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceGetFrameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceGetFrameResponse) ProtoMessage() {}

func (x *CameraServiceGetFrameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceGetFrameResponse.ProtoReflect.Descriptor instead.
func (*CameraServiceGetFrameResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{1}
}

func (x *CameraServiceGetFrameResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *CameraServiceGetFrameResponse) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *CameraServiceGetFrameResponse) GetWidthPx() int64 {
	if x != nil {
		return x.WidthPx
	}
	return 0
}

func (x *CameraServiceGetFrameResponse) GetHeightPx() int64 {
	if x != nil {
		return x.HeightPx
	}
	return 0
}

type CameraServiceRenderFrameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a camera
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Requested MIME type of response
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *CameraServiceRenderFrameRequest) Reset() {
	*x = CameraServiceRenderFrameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceRenderFrameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceRenderFrameRequest) ProtoMessage() {}

func (x *CameraServiceRenderFrameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceRenderFrameRequest.ProtoReflect.Descriptor instead.
func (*CameraServiceRenderFrameRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{2}
}

func (x *CameraServiceRenderFrameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CameraServiceRenderFrameRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type CameraServiceGetPointCloudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a camera
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Requested MIME type of response
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *CameraServiceGetPointCloudRequest) Reset() {
	*x = CameraServiceGetPointCloudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceGetPointCloudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceGetPointCloudRequest) ProtoMessage() {}

func (x *CameraServiceGetPointCloudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceGetPointCloudRequest.ProtoReflect.Descriptor instead.
func (*CameraServiceGetPointCloudRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{3}
}

func (x *CameraServiceGetPointCloudRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CameraServiceGetPointCloudRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type CameraServiceGetPointCloudResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actual MIME type of response
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Frame in bytes
	Frame []byte `protobuf:"bytes,2,opt,name=frame,proto3" json:"frame,omitempty"`
}

func (x *CameraServiceGetPointCloudResponse) Reset() {
	*x = CameraServiceGetPointCloudResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceGetPointCloudResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceGetPointCloudResponse) ProtoMessage() {}

func (x *CameraServiceGetPointCloudResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceGetPointCloudResponse.ProtoReflect.Descriptor instead.
func (*CameraServiceGetPointCloudResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{4}
}

func (x *CameraServiceGetPointCloudResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *CameraServiceGetPointCloudResponse) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

type CameraServiceGetObjectPointCloudsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of a camera
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Requested MIME type of response
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Minimum points in plane
	MinPointsInPlane int64 `protobuf:"varint,3,opt,name=min_points_in_plane,json=minPointsInPlane,proto3" json:"min_points_in_plane,omitempty"`
	// Minimum points in segment
	MinPointsInSegment int64 `protobuf:"varint,4,opt,name=min_points_in_segment,json=minPointsInSegment,proto3" json:"min_points_in_segment,omitempty"`
	// Clustering radius in mm
	ClusteringRadiusMm float64 `protobuf:"fixed64,5,opt,name=clustering_radius_mm,json=clusteringRadiusMm,proto3" json:"clustering_radius_mm,omitempty"`
}

func (x *CameraServiceGetObjectPointCloudsRequest) Reset() {
	*x = CameraServiceGetObjectPointCloudsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceGetObjectPointCloudsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceGetObjectPointCloudsRequest) ProtoMessage() {}

func (x *CameraServiceGetObjectPointCloudsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceGetObjectPointCloudsRequest.ProtoReflect.Descriptor instead.
func (*CameraServiceGetObjectPointCloudsRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{5}
}

func (x *CameraServiceGetObjectPointCloudsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CameraServiceGetObjectPointCloudsRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *CameraServiceGetObjectPointCloudsRequest) GetMinPointsInPlane() int64 {
	if x != nil {
		return x.MinPointsInPlane
	}
	return 0
}

func (x *CameraServiceGetObjectPointCloudsRequest) GetMinPointsInSegment() int64 {
	if x != nil {
		return x.MinPointsInSegment
	}
	return 0
}

func (x *CameraServiceGetObjectPointCloudsRequest) GetClusteringRadiusMm() float64 {
	if x != nil {
		return x.ClusteringRadiusMm
	}
	return 0
}

type CameraServiceGetObjectPointCloudsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actual MIME type of response
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// List of objects in the scene
	Objects []*PointCloudObject `protobuf:"bytes,2,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *CameraServiceGetObjectPointCloudsResponse) Reset() {
	*x = CameraServiceGetObjectPointCloudsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraServiceGetObjectPointCloudsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraServiceGetObjectPointCloudsResponse) ProtoMessage() {}

func (x *CameraServiceGetObjectPointCloudsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraServiceGetObjectPointCloudsResponse.ProtoReflect.Descriptor instead.
func (*CameraServiceGetObjectPointCloudsResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{6}
}

func (x *CameraServiceGetObjectPointCloudsResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *CameraServiceGetObjectPointCloudsResponse) GetObjects() []*PointCloudObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

type PointCloudObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Frame of object in bytes
	Frame []byte `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	// Center of object in mm
	CenterCoordinatesMm *v1.Vector3 `protobuf:"bytes,2,opt,name=center_coordinates_mm,json=centerCoordinatesMm,proto3" json:"center_coordinates_mm,omitempty"`
	// Bounding box of object in mm
	BoundingBoxMm *v1.BoxGeometry `protobuf:"bytes,3,opt,name=bounding_box_mm,json=boundingBoxMm,proto3" json:"bounding_box_mm,omitempty"`
}

func (x *PointCloudObject) Reset() {
	*x = PointCloudObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_camera_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointCloudObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointCloudObject) ProtoMessage() {}

func (x *PointCloudObject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_camera_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointCloudObject.ProtoReflect.Descriptor instead.
func (*PointCloudObject) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_camera_proto_rawDescGZIP(), []int{7}
}

func (x *PointCloudObject) GetFrame() []byte {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *PointCloudObject) GetCenterCoordinatesMm() *v1.Vector3 {
	if x != nil {
		return x.CenterCoordinatesMm
	}
	return nil
}

func (x *PointCloudObject) GetBoundingBoxMm() *v1.BoxGeometry {
	if x != nil {
		return x.BoundingBoxMm
	}
	return nil
}

var File_proto_api_component_v1_camera_proto protoreflect.FileDescriptor

var file_proto_api_component_v1_camera_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x1c, 0x43, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x1d, 0x43, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x70, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x77, 0x69, 0x64, 0x74, 0x68, 0x50, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x70, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x50, 0x78, 0x22, 0x52, 0x0a, 0x1f, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a, 0x21, 0x43, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x57, 0x0a, 0x22, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x22, 0xef, 0x01, 0x0a, 0x28, 0x43, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x49,
	0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x49, 0x6e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x5f, 0x6d,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x4d, 0x6d, 0x22, 0x8c, 0x01, 0x0a, 0x29,
	0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x10, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x15, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6d, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x33, 0x52, 0x13, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x4d, 0x6d, 0x12, 0x48, 0x0a, 0x0f, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x6d, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x78, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x4d,
	0x6d, 0x32, 0xe6, 0x05, 0x0a, 0x0d, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x92, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x34, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0xbb, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x12, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12,
	0xd8, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x12, 0x40, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x42, 0x4d, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_api_component_v1_camera_proto_rawDescOnce sync.Once
	file_proto_api_component_v1_camera_proto_rawDescData = file_proto_api_component_v1_camera_proto_rawDesc
)

func file_proto_api_component_v1_camera_proto_rawDescGZIP() []byte {
	file_proto_api_component_v1_camera_proto_rawDescOnce.Do(func() {
		file_proto_api_component_v1_camera_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_v1_camera_proto_rawDescData)
	})
	return file_proto_api_component_v1_camera_proto_rawDescData
}

var file_proto_api_component_v1_camera_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_api_component_v1_camera_proto_goTypes = []interface{}{
	(*CameraServiceGetFrameRequest)(nil),              // 0: proto.api.component.v1.CameraServiceGetFrameRequest
	(*CameraServiceGetFrameResponse)(nil),             // 1: proto.api.component.v1.CameraServiceGetFrameResponse
	(*CameraServiceRenderFrameRequest)(nil),           // 2: proto.api.component.v1.CameraServiceRenderFrameRequest
	(*CameraServiceGetPointCloudRequest)(nil),         // 3: proto.api.component.v1.CameraServiceGetPointCloudRequest
	(*CameraServiceGetPointCloudResponse)(nil),        // 4: proto.api.component.v1.CameraServiceGetPointCloudResponse
	(*CameraServiceGetObjectPointCloudsRequest)(nil),  // 5: proto.api.component.v1.CameraServiceGetObjectPointCloudsRequest
	(*CameraServiceGetObjectPointCloudsResponse)(nil), // 6: proto.api.component.v1.CameraServiceGetObjectPointCloudsResponse
	(*PointCloudObject)(nil),                          // 7: proto.api.component.v1.PointCloudObject
	(*v1.Vector3)(nil),                                // 8: proto.api.common.v1.Vector3
	(*v1.BoxGeometry)(nil),                            // 9: proto.api.common.v1.BoxGeometry
	(*httpbody.HttpBody)(nil),                         // 10: google.api.HttpBody
}
var file_proto_api_component_v1_camera_proto_depIdxs = []int32{
	7,  // 0: proto.api.component.v1.CameraServiceGetObjectPointCloudsResponse.objects:type_name -> proto.api.component.v1.PointCloudObject
	8,  // 1: proto.api.component.v1.PointCloudObject.center_coordinates_mm:type_name -> proto.api.common.v1.Vector3
	9,  // 2: proto.api.component.v1.PointCloudObject.bounding_box_mm:type_name -> proto.api.common.v1.BoxGeometry
	0,  // 3: proto.api.component.v1.CameraService.GetFrame:input_type -> proto.api.component.v1.CameraServiceGetFrameRequest
	2,  // 4: proto.api.component.v1.CameraService.RenderFrame:input_type -> proto.api.component.v1.CameraServiceRenderFrameRequest
	3,  // 5: proto.api.component.v1.CameraService.GetPointCloud:input_type -> proto.api.component.v1.CameraServiceGetPointCloudRequest
	5,  // 6: proto.api.component.v1.CameraService.GetObjectPointClouds:input_type -> proto.api.component.v1.CameraServiceGetObjectPointCloudsRequest
	1,  // 7: proto.api.component.v1.CameraService.GetFrame:output_type -> proto.api.component.v1.CameraServiceGetFrameResponse
	10, // 8: proto.api.component.v1.CameraService.RenderFrame:output_type -> google.api.HttpBody
	4,  // 9: proto.api.component.v1.CameraService.GetPointCloud:output_type -> proto.api.component.v1.CameraServiceGetPointCloudResponse
	6,  // 10: proto.api.component.v1.CameraService.GetObjectPointClouds:output_type -> proto.api.component.v1.CameraServiceGetObjectPointCloudsResponse
	7,  // [7:11] is the sub-list for method output_type
	3,  // [3:7] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_api_component_v1_camera_proto_init() }
func file_proto_api_component_v1_camera_proto_init() {
	if File_proto_api_component_v1_camera_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_v1_camera_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceGetFrameRequest); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceGetFrameResponse); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceRenderFrameRequest); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceGetPointCloudRequest); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceGetPointCloudResponse); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceGetObjectPointCloudsRequest); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraServiceGetObjectPointCloudsResponse); i {
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
		file_proto_api_component_v1_camera_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointCloudObject); i {
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
			RawDescriptor: file_proto_api_component_v1_camera_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_v1_camera_proto_goTypes,
		DependencyIndexes: file_proto_api_component_v1_camera_proto_depIdxs,
		MessageInfos:      file_proto_api_component_v1_camera_proto_msgTypes,
	}.Build()
	File_proto_api_component_v1_camera_proto = out.File
	file_proto_api_component_v1_camera_proto_rawDesc = nil
	file_proto_api_component_v1_camera_proto_goTypes = nil
	file_proto_api_component_v1_camera_proto_depIdxs = nil
}
