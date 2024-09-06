// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
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

type AverageTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AverageTokens int64 `protobuf:"varint,1,opt,name=averageTokens,proto3" json:"averageTokens,omitempty"`
}

func (x *AverageTokensResponse) Reset() {
	*x = AverageTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageTokensResponse) ProtoMessage() {}

func (x *AverageTokensResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AverageTokensResponse.ProtoReflect.Descriptor instead.
func (*AverageTokensResponse) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{1}
}

func (x *AverageTokensResponse) GetAverageTokens() int64 {
	if x != nil {
		return x.AverageTokens
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
		mi := &file_api_rewards_rewards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQualifiedDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQualifiedDevicesRequest) ProtoMessage() {}

func (x *GetQualifiedDevicesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetQualifiedDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetQualifiedDevicesRequest) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{2}
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

type GetDeviceRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeviceRewardsRequest) Reset() {
	*x = GetDeviceRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceRewardsRequest) ProtoMessage() {}

func (x *GetDeviceRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceRewardsRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceRewardsRequest) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeviceRewardsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeviceRewardsWeek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndDate             string   `protobuf:"bytes,1,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Tokens              float64  `protobuf:"fixed64,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
	ConnectionStreak    int32    `protobuf:"varint,3,opt,name=connection_streak,json=connectionStreak,proto3" json:"connection_streak,omitempty"`
	DisconnectionStreak int32    `protobuf:"varint,4,opt,name=disconnection_streak,json=disconnectionStreak,proto3" json:"disconnection_streak,omitempty"`
	IntegrationIds      []string `protobuf:"bytes,5,rep,name=integration_ids,json=integrationIds,proto3" json:"integration_ids,omitempty"`
}

func (x *DeviceRewardsWeek) Reset() {
	*x = DeviceRewardsWeek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRewardsWeek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRewardsWeek) ProtoMessage() {}

func (x *DeviceRewardsWeek) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRewardsWeek.ProtoReflect.Descriptor instead.
func (*DeviceRewardsWeek) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceRewardsWeek) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *DeviceRewardsWeek) GetTokens() float64 {
	if x != nil {
		return x.Tokens
	}
	return 0
}

func (x *DeviceRewardsWeek) GetConnectionStreak() int32 {
	if x != nil {
		return x.ConnectionStreak
	}
	return 0
}

func (x *DeviceRewardsWeek) GetDisconnectionStreak() int32 {
	if x != nil {
		return x.DisconnectionStreak
	}
	return 0
}

func (x *DeviceRewardsWeek) GetIntegrationIds() []string {
	if x != nil {
		return x.IntegrationIds
	}
	return nil
}

type GetDeviceRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tokens float64              `protobuf:"fixed64,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
	Weeks  []*DeviceRewardsWeek `protobuf:"bytes,3,rep,name=weeks,proto3" json:"weeks,omitempty"`
}

func (x *GetDeviceRewardsResponse) Reset() {
	*x = GetDeviceRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceRewardsResponse) ProtoMessage() {}

func (x *GetDeviceRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceRewardsResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceRewardsResponse) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{6}
}

func (x *GetDeviceRewardsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetDeviceRewardsResponse) GetTokens() float64 {
	if x != nil {
		return x.Tokens
	}
	return 0
}

func (x *GetDeviceRewardsResponse) GetWeeks() []*DeviceRewardsWeek {
	if x != nil {
		return x.Weeks
	}
	return nil
}

type GetBlacklistStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddress []byte `protobuf:"bytes,1,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
}

func (x *GetBlacklistStatusRequest) Reset() {
	*x = GetBlacklistStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlacklistStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlacklistStatusRequest) ProtoMessage() {}

func (x *GetBlacklistStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlacklistStatusRequest.ProtoReflect.Descriptor instead.
func (*GetBlacklistStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlacklistStatusRequest) GetEthereumAddress() []byte {
	if x != nil {
		return x.EthereumAddress
	}
	return nil
}

type GetBlacklistStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsBlacklisted bool                   `protobuf:"varint,1,opt,name=is_blacklisted,json=isBlacklisted,proto3" json:"is_blacklisted,omitempty"`
	Note          string                 `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetBlacklistStatusResponse) Reset() {
	*x = GetBlacklistStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlacklistStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlacklistStatusResponse) ProtoMessage() {}

func (x *GetBlacklistStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlacklistStatusResponse.ProtoReflect.Descriptor instead.
func (*GetBlacklistStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{8}
}

func (x *GetBlacklistStatusResponse) GetIsBlacklisted() bool {
	if x != nil {
		return x.IsBlacklisted
	}
	return false
}

func (x *GetBlacklistStatusResponse) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *GetBlacklistStatusResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type SetBlacklistStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddress []byte `protobuf:"bytes,1,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	IsBlacklisted   bool   `protobuf:"varint,2,opt,name=is_blacklisted,json=isBlacklisted,proto3" json:"is_blacklisted,omitempty"`
	// A note describing why someone was blacklisted, typically containing the name of the
	// person doing it and a link to an issue. Ignored if is_blacklisted is false.
	Note string `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *SetBlacklistStatusRequest) Reset() {
	*x = SetBlacklistStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBlacklistStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBlacklistStatusRequest) ProtoMessage() {}

func (x *SetBlacklistStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBlacklistStatusRequest.ProtoReflect.Descriptor instead.
func (*SetBlacklistStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{9}
}

func (x *SetBlacklistStatusRequest) GetEthereumAddress() []byte {
	if x != nil {
		return x.EthereumAddress
	}
	return nil
}

func (x *SetBlacklistStatusRequest) GetIsBlacklisted() bool {
	if x != nil {
		return x.IsBlacklisted
	}
	return false
}

func (x *SetBlacklistStatusRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

type SetBlacklistStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetBlacklistStatusResponse) Reset() {
	*x = SetBlacklistStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rewards_rewards_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBlacklistStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBlacklistStatusResponse) ProtoMessage() {}

func (x *SetBlacklistStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_rewards_rewards_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBlacklistStatusResponse.ProtoReflect.Descriptor instead.
func (*SetBlacklistStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_rewards_rewards_proto_rawDescGZIP(), []int{10}
}

var File_api_rewards_rewards_proto protoreflect.FileDescriptor

var file_api_rewards_rewards_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3d,
	0x0a, 0x15, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x7c, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x54, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcf, 0x01, 0x0a,
	0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x57, 0x65,
	0x65, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6b, 0x12, 0x31, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x74,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x05, 0x77,
	0x65, 0x65, 0x6b, 0x73, 0x22, 0x46, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69,
	0x73, 0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x81, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73,
	0x5f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xa0, 0x04, 0x0a, 0x0e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x30, 0x01, 0x12,
	0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22,
	0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x42,
	0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x49, 0x4d, 0x4f, 0x2d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_rewards_rewards_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_rewards_rewards_proto_goTypes = []any{
	(*GetTotalPointsResponse)(nil),     // 0: rewards.GetTotalPointsResponse
	(*AverageTokensResponse)(nil),      // 1: rewards.AverageTokensResponse
	(*GetQualifiedDevicesRequest)(nil), // 2: rewards.GetQualifiedDevicesRequest
	(*GetQualifiedDevicesDevice)(nil),  // 3: rewards.GetQualifiedDevicesDevice
	(*GetDeviceRewardsRequest)(nil),    // 4: rewards.GetDeviceRewardsRequest
	(*DeviceRewardsWeek)(nil),          // 5: rewards.DeviceRewardsWeek
	(*GetDeviceRewardsResponse)(nil),   // 6: rewards.GetDeviceRewardsResponse
	(*GetBlacklistStatusRequest)(nil),  // 7: rewards.GetBlacklistStatusRequest
	(*GetBlacklistStatusResponse)(nil), // 8: rewards.GetBlacklistStatusResponse
	(*SetBlacklistStatusRequest)(nil),  // 9: rewards.SetBlacklistStatusRequest
	(*SetBlacklistStatusResponse)(nil), // 10: rewards.SetBlacklistStatusResponse
	(*timestamppb.Timestamp)(nil),      // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 12: google.protobuf.Empty
}
var file_api_rewards_rewards_proto_depIdxs = []int32{
	11, // 0: rewards.GetQualifiedDevicesRequest.start:type_name -> google.protobuf.Timestamp
	11, // 1: rewards.GetQualifiedDevicesRequest.end:type_name -> google.protobuf.Timestamp
	5,  // 2: rewards.GetDeviceRewardsResponse.weeks:type_name -> rewards.DeviceRewardsWeek
	11, // 3: rewards.GetBlacklistStatusResponse.created_at:type_name -> google.protobuf.Timestamp
	12, // 4: rewards.RewardsService.GetTotalPoints:input_type -> google.protobuf.Empty
	12, // 5: rewards.RewardsService.GetAverageTokens:input_type -> google.protobuf.Empty
	2,  // 6: rewards.RewardsService.GetQualifiedDevices:input_type -> rewards.GetQualifiedDevicesRequest
	4,  // 7: rewards.RewardsService.GetDeviceRewards:input_type -> rewards.GetDeviceRewardsRequest
	7,  // 8: rewards.RewardsService.GetBlacklistStatus:input_type -> rewards.GetBlacklistStatusRequest
	9,  // 9: rewards.RewardsService.SetBlacklistStatus:input_type -> rewards.SetBlacklistStatusRequest
	0,  // 10: rewards.RewardsService.GetTotalPoints:output_type -> rewards.GetTotalPointsResponse
	1,  // 11: rewards.RewardsService.GetAverageTokens:output_type -> rewards.AverageTokensResponse
	3,  // 12: rewards.RewardsService.GetQualifiedDevices:output_type -> rewards.GetQualifiedDevicesDevice
	6,  // 13: rewards.RewardsService.GetDeviceRewards:output_type -> rewards.GetDeviceRewardsResponse
	8,  // 14: rewards.RewardsService.GetBlacklistStatus:output_type -> rewards.GetBlacklistStatusResponse
	10, // 15: rewards.RewardsService.SetBlacklistStatus:output_type -> rewards.SetBlacklistStatusResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_rewards_rewards_proto_init() }
func file_api_rewards_rewards_proto_init() {
	if File_api_rewards_rewards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rewards_rewards_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_api_rewards_rewards_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AverageTokensResponse); i {
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
		file_api_rewards_rewards_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_api_rewards_rewards_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_api_rewards_rewards_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetDeviceRewardsRequest); i {
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
		file_api_rewards_rewards_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceRewardsWeek); i {
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
		file_api_rewards_rewards_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetDeviceRewardsResponse); i {
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
		file_api_rewards_rewards_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetBlacklistStatusRequest); i {
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
		file_api_rewards_rewards_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetBlacklistStatusResponse); i {
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
		file_api_rewards_rewards_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SetBlacklistStatusRequest); i {
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
		file_api_rewards_rewards_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SetBlacklistStatusResponse); i {
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
			NumMessages:   11,
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
