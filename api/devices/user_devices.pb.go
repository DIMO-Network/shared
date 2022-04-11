// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/devices/user_devices.proto

package devices

import (
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

type GetUserDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserDeviceRequest) Reset() {
	*x = GetUserDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devices_user_devices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDeviceRequest) ProtoMessage() {}

func (x *GetUserDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devices_user_devices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDeviceRequest.ProtoReflect.Descriptor instead.
func (*GetUserDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_devices_user_devices_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserDeviceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserDevice) Reset() {
	*x = UserDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devices_user_devices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDevice) ProtoMessage() {}

func (x *UserDevice) ProtoReflect() protoreflect.Message {
	mi := &file_api_devices_user_devices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDevice.ProtoReflect.Descriptor instead.
func (*UserDevice) Descriptor() ([]byte, []int) {
	return file_api_devices_user_devices_proto_rawDescGZIP(), []int{1}
}

func (x *UserDevice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserDevice) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListUserDevicesForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListUserDevicesForUserRequest) Reset() {
	*x = ListUserDevicesForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devices_user_devices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserDevicesForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserDevicesForUserRequest) ProtoMessage() {}

func (x *ListUserDevicesForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devices_user_devices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserDevicesForUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserDevicesForUserRequest) Descriptor() ([]byte, []int) {
	return file_api_devices_user_devices_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserDevicesForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListUserDevicesForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserDevices []*UserDevice `protobuf:"bytes,1,rep,name=user_devices,json=userDevices,proto3" json:"user_devices,omitempty"`
}

func (x *ListUserDevicesForUserResponse) Reset() {
	*x = ListUserDevicesForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devices_user_devices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserDevicesForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserDevicesForUserResponse) ProtoMessage() {}

func (x *ListUserDevicesForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_devices_user_devices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserDevicesForUserResponse.ProtoReflect.Descriptor instead.
func (*ListUserDevicesForUserResponse) Descriptor() ([]byte, []int) {
	return file_api_devices_user_devices_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserDevicesForUserResponse) GetUserDevices() []*UserDevice {
	if x != nil {
		return x.UserDevices
	}
	return nil
}

var File_api_devices_user_devices_proto protoreflect.FileDescriptor

var file_api_devices_user_devices_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x35, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x58, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xc3, 0x01, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_devices_user_devices_proto_rawDescOnce sync.Once
	file_api_devices_user_devices_proto_rawDescData = file_api_devices_user_devices_proto_rawDesc
)

func file_api_devices_user_devices_proto_rawDescGZIP() []byte {
	file_api_devices_user_devices_proto_rawDescOnce.Do(func() {
		file_api_devices_user_devices_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_devices_user_devices_proto_rawDescData)
	})
	return file_api_devices_user_devices_proto_rawDescData
}

var file_api_devices_user_devices_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_devices_user_devices_proto_goTypes = []interface{}{
	(*GetUserDeviceRequest)(nil),           // 0: devices.GetUserDeviceRequest
	(*UserDevice)(nil),                     // 1: devices.UserDevice
	(*ListUserDevicesForUserRequest)(nil),  // 2: devices.ListUserDevicesForUserRequest
	(*ListUserDevicesForUserResponse)(nil), // 3: devices.ListUserDevicesForUserResponse
}
var file_api_devices_user_devices_proto_depIdxs = []int32{
	1, // 0: devices.ListUserDevicesForUserResponse.user_devices:type_name -> devices.UserDevice
	0, // 1: devices.UserDeviceService.GetUserDevice:input_type -> devices.GetUserDeviceRequest
	2, // 2: devices.UserDeviceService.ListUserDevicesForUser:input_type -> devices.ListUserDevicesForUserRequest
	1, // 3: devices.UserDeviceService.GetUserDevice:output_type -> devices.UserDevice
	3, // 4: devices.UserDeviceService.ListUserDevicesForUser:output_type -> devices.ListUserDevicesForUserResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_devices_user_devices_proto_init() }
func file_api_devices_user_devices_proto_init() {
	if File_api_devices_user_devices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_devices_user_devices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDeviceRequest); i {
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
		file_api_devices_user_devices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDevice); i {
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
		file_api_devices_user_devices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserDevicesForUserRequest); i {
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
		file_api_devices_user_devices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserDevicesForUserResponse); i {
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
			RawDescriptor: file_api_devices_user_devices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_devices_user_devices_proto_goTypes,
		DependencyIndexes: file_api_devices_user_devices_proto_depIdxs,
		MessageInfos:      file_api_devices_user_devices_proto_msgTypes,
	}.Build()
	File_api_devices_user_devices_proto = out.File
	file_api_devices_user_devices_proto_rawDesc = nil
	file_api_devices_user_devices_proto_goTypes = nil
	file_api_devices_user_devices_proto_depIdxs = nil
}