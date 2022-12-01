// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: api/rewards/rewards.proto

package rewards

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetTotalPointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPoints int64 `protobuf:"varint,1,opt,name=totalPoints,proto3" json:"totalPoints,omitempty"`
}

func (x *GetTotalPointsResponse) Reset() {
	*x = GetTotalPointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalPointsResponse) ProtoMessage() {}

func (x *GetTotalPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalPointsResponse.ProtoReflect.Descriptor instead.
func (*GetTotalPointsResponse) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{0}
}

func (x *GetTotalPointsResponse) GetTotalPoints() int64 {
	if x != nil {
		return x.TotalPoints
	}
	return 0
}

type GetQualifiedDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetQualifiedDevicesRequest) Reset() {
	*x = GetQualifiedDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQualifiedDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQualifiedDevicesRequest) ProtoMessage() {}

func (x *GetQualifiedDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQualifiedDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetQualifiedDevicesRequest) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{1}
}

func (x *GetQualifiedDevicesRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetQualifiedDevicesRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type GetQualifiedDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*GetQualifiedDevicesDevice `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *GetQualifiedDevicesResponse) Reset() {
	*x = GetQualifiedDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQualifiedDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQualifiedDevicesResponse) ProtoMessage() {}

func (x *GetQualifiedDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQualifiedDevicesResponse.ProtoReflect.Descriptor instead.
func (*GetQualifiedDevicesResponse) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{2}
}

func (x *GetQualifiedDevicesResponse) GetDevices() []*GetQualifiedDevicesDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

type GetQualifiedDevicesDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IntegrationIds []string `protobuf:"bytes,2,rep,name=integration_ids,json=integrationIds,proto3" json:"integration_ids,omitempty"`
}

func (x *GetQualifiedDevicesDevice) Reset() {
	*x = GetQualifiedDevicesDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQualifiedDevicesDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQualifiedDevicesDevice) ProtoMessage() {}

func (x *GetQualifiedDevicesDevice) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQualifiedDevicesDevice.ProtoReflect.Descriptor instead.
func (*GetQualifiedDevicesDevice) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{3}
}

func (x *GetQualifiedDevicesDevice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetQualifiedDevicesDevice) GetIntegrationIds() []string {
	if x != nil {
		return x.IntegrationIds
	}
	return nil
}

var File_api_rewards_rewards_proto protoreflect.FileDescriptor

var file_api_rewards_rewards_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x7c, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x59, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x32, 0xb7, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x21,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rewards_rewards_proto_rawDescOnce sync.Once
	file_api_rewards_rewards_proto_rawDescData = file_api_rewards_rewards_proto_rawDesc
)

func file_api_rewards_rewards_proto_rawDescGZIP() []byte {
	file_api_rewards_rewards_proto_rawDescOnce.Do(func() {
		file_api_rewards_rewards_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rewards_rewards_proto_rawDescData)
	})
	return file_api_rewards_rewards_proto_rawDescData
}

var file_api_rewards_rewards_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_rewards_rewards_proto_goTypes = []interface{}{
	(*GetTotalPointsResponse)(nil),      // 0: users.GetTotalPointsResponse
	(*GetQualifiedDevicesRequest)(nil),  // 1: users.GetQualifiedDevicesRequest
	(*GetQualifiedDevicesResponse)(nil), // 2: users.GetQualifiedDevicesResponse
	(*GetQualifiedDevicesDevice)(nil),   // 3: users.GetQualifiedDevicesDevice
	(*timestamppb.Timestamp)(nil),       // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),               // 5: google.protobuf.Empty
}
var file_api_rewards_rewards_proto_depIdxs = []int32{
	4, // 0: users.GetQualifiedDevicesRequest.start:type_name -> google.protobuf.Timestamp
	4, // 1: users.GetQualifiedDevicesRequest.end:type_name -> google.protobuf.Timestamp
	3, // 2: users.GetQualifiedDevicesResponse.devices:type_name -> users.GetQualifiedDevicesDevice
	5, // 3: users.RewardsService.GetTotalPoints:input_type -> google.protobuf.Empty
	1, // 4: users.RewardsService.GetQualifiedDevices:input_type -> users.GetQualifiedDevicesRequest
	0, // 5: users.RewardsService.GetTotalPoints:output_type -> users.GetTotalPointsResponse
	2, // 6: users.RewardsService.GetQualifiedDevices:output_type -> users.GetQualifiedDevicesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_rewards_rewards_proto_init() }
func file_api_rewards_rewards_proto_init() {
	if File_api_rewards_rewards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rewards_rewards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalPointsResponse); i {
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
		file_api_rewards_rewards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQualifiedDevicesRequest); i {
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
		file_api_rewards_rewards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQualifiedDevicesResponse); i {
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
		file_api_rewards_rewards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQualifiedDevicesDevice); i {
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
			RawDescriptor: file_api_rewards_rewards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rewards_rewards_proto_goTypes,
		DependencyIndexes: file_api_rewards_rewards_proto_depIdxs,
		MessageInfos:      file_api_rewards_rewards_proto_msgTypes,
	}.Build()
	File_api_rewards_rewards_proto = out.File
	file_api_rewards_rewards_proto_rawDesc = nil
	file_api_rewards_rewards_proto_goTypes = nil
	file_api_rewards_rewards_proto_depIdxs = nil
}
