// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: chat.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GroupMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	SyncSeq int64 `protobuf:"varint,3,opt,name=sync_seq,json=syncSeq,proto3" json:"sync_seq,omitempty"`
	Limit   int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GroupMessageRequest) Reset() {
	*x = GroupMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMessageRequest) ProtoMessage() {}

func (x *GroupMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMessageRequest.ProtoReflect.Descriptor instead.
func (*GroupMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMessageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GroupMessageRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupMessageRequest) GetSyncSeq() int64 {
	if x != nil {
		return x.SyncSeq
	}
	return 0
}

func (x *GroupMessageRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GroupMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId  int64                `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Data     string               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	SenderId int64                `protobuf:"varint,4,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Seq      int64                `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	ReplyTo  int64                `protobuf:"varint,6,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	Type     int64                `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	Time     *timestamp.Timestamp `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GroupMessageResponse) Reset() {
	*x = GroupMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMessageResponse) ProtoMessage() {}

func (x *GroupMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMessageResponse.ProtoReflect.Descriptor instead.
func (*GroupMessageResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GroupMessageResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GroupMessageResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupMessageResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GroupMessageResponse) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *GroupMessageResponse) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *GroupMessageResponse) GetReplyTo() int64 {
	if x != nil {
		return x.ReplyTo
	}
	return 0
}

func (x *GroupMessageResponse) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GroupMessageResponse) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type UserJoinGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *UserJoinGroupRequest) Reset() {
	*x = UserJoinGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserJoinGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserJoinGroupRequest) ProtoMessage() {}

func (x *UserJoinGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserJoinGroupRequest.ProtoReflect.Descriptor instead.
func (*UserJoinGroupRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *UserJoinGroupRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserJoinGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type UserJoinGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserNum      int64                `protobuf:"varint,1,opt,name=user_num,json=userNum,proto3" json:"user_num,omitempty"`
	GroupId      int64                `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Name         string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction string               `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	OwnerId      int64                `protobuf:"varint,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	CreateTime   *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *UserJoinGroupResponse) Reset() {
	*x = UserJoinGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserJoinGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserJoinGroupResponse) ProtoMessage() {}

func (x *UserJoinGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserJoinGroupResponse.ProtoReflect.Descriptor instead.
func (*UserJoinGroupResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *UserJoinGroupResponse) GetUserNum() int64 {
	if x != nil {
		return x.UserNum
	}
	return 0
}

func (x *UserJoinGroupResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserJoinGroupResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserJoinGroupResponse) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *UserJoinGroupResponse) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *UserJoinGroupResponse) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type PackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are assignable to Payload:
	//	*PackageRequest_UserJoinGroupRequest
	//	*PackageRequest_GroupMessageRequest
	Payload isPackageRequest_Payload `protobuf_oneof:"payload"`
}

func (x *PackageRequest) Reset() {
	*x = PackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageRequest) ProtoMessage() {}

func (x *PackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageRequest.ProtoReflect.Descriptor instead.
func (*PackageRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *PackageRequest) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (m *PackageRequest) GetPayload() isPackageRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *PackageRequest) GetUserJoinGroupRequest() *UserJoinGroupRequest {
	if x, ok := x.GetPayload().(*PackageRequest_UserJoinGroupRequest); ok {
		return x.UserJoinGroupRequest
	}
	return nil
}

func (x *PackageRequest) GetGroupMessageRequest() *GroupMessageRequest {
	if x, ok := x.GetPayload().(*PackageRequest_GroupMessageRequest); ok {
		return x.GroupMessageRequest
	}
	return nil
}

type isPackageRequest_Payload interface {
	isPackageRequest_Payload()
}

type PackageRequest_UserJoinGroupRequest struct {
	UserJoinGroupRequest *UserJoinGroupRequest `protobuf:"bytes,2,opt,name=user_join_group_request,json=userJoinGroupRequest,proto3,oneof"`
}

type PackageRequest_GroupMessageRequest struct {
	GroupMessageRequest *GroupMessageRequest `protobuf:"bytes,3,opt,name=group_message_request,json=groupMessageRequest,proto3,oneof"`
}

func (*PackageRequest_UserJoinGroupRequest) isPackageRequest_Payload() {}

func (*PackageRequest_GroupMessageRequest) isPackageRequest_Payload() {}

type PackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int64 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are assignable to Payload:
	//	*PackageResponse_UserJoinGroupResponse
	//	*PackageResponse_GroupMessageResponse
	Payload isPackageResponse_Payload `protobuf_oneof:"payload"`
}

func (x *PackageResponse) Reset() {
	*x = PackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageResponse) ProtoMessage() {}

func (x *PackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageResponse.ProtoReflect.Descriptor instead.
func (*PackageResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *PackageResponse) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (m *PackageResponse) GetPayload() isPackageResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *PackageResponse) GetUserJoinGroupResponse() *UserJoinGroupResponse {
	if x, ok := x.GetPayload().(*PackageResponse_UserJoinGroupResponse); ok {
		return x.UserJoinGroupResponse
	}
	return nil
}

func (x *PackageResponse) GetGroupMessageResponse() *GroupMessageResponse {
	if x, ok := x.GetPayload().(*PackageResponse_GroupMessageResponse); ok {
		return x.GroupMessageResponse
	}
	return nil
}

type isPackageResponse_Payload interface {
	isPackageResponse_Payload()
}

type PackageResponse_UserJoinGroupResponse struct {
	UserJoinGroupResponse *UserJoinGroupResponse `protobuf:"bytes,2,opt,name=user_join_group_response,json=userJoinGroupResponse,proto3,oneof"`
}

type PackageResponse_GroupMessageResponse struct {
	GroupMessageResponse *GroupMessageResponse `protobuf:"bytes,3,opt,name=group_message_response,json=groupMessageResponse,proto3,oneof"`
}

func (*PackageResponse_UserJoinGroupResponse) isPackageResponse_Payload() {}

func (*PackageResponse_GroupMessageResponse) isPackageResponse_Payload() {}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x13, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0xdd,
	0x01, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xdd,
	0x01, 0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x57, 0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6a, 0x6f,
	0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x14, 0x75, 0x73, 0x65, 0x72, 0x4a, 0x6f,
	0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53,
	0x0a, 0x15, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x13,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe4,
	0x01, 0x0a, 0x0f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5a, 0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6a,
	0x6f, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x15, 0x75, 0x73, 0x65,
	0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x56, 0x0a, 0x16, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x14, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x69, 0x76, 0x6f, 0x74, 0x2d, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2f, 0x70, 0x69, 0x76, 0x6f, 0x74, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chat_proto_goTypes = []interface{}{
	(*GroupMessageRequest)(nil),   // 0: tutorial.GroupMessageRequest
	(*GroupMessageResponse)(nil),  // 1: tutorial.GroupMessageResponse
	(*UserJoinGroupRequest)(nil),  // 2: tutorial.UserJoinGroupRequest
	(*UserJoinGroupResponse)(nil), // 3: tutorial.UserJoinGroupResponse
	(*PackageRequest)(nil),        // 4: tutorial.PackageRequest
	(*PackageResponse)(nil),       // 5: tutorial.PackageResponse
	(*timestamp.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_chat_proto_depIdxs = []int32{
	6, // 0: tutorial.GroupMessageResponse.time:type_name -> google.protobuf.Timestamp
	6, // 1: tutorial.UserJoinGroupResponse.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: tutorial.PackageRequest.user_join_group_request:type_name -> tutorial.UserJoinGroupRequest
	0, // 3: tutorial.PackageRequest.group_message_request:type_name -> tutorial.GroupMessageRequest
	3, // 4: tutorial.PackageResponse.user_join_group_response:type_name -> tutorial.UserJoinGroupResponse
	1, // 5: tutorial.PackageResponse.group_message_response:type_name -> tutorial.GroupMessageResponse
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMessageRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMessageResponse); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserJoinGroupRequest); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserJoinGroupResponse); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageRequest); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageResponse); i {
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
	file_chat_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PackageRequest_UserJoinGroupRequest)(nil),
		(*PackageRequest_GroupMessageRequest)(nil),
	}
	file_chat_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*PackageResponse_UserJoinGroupResponse)(nil),
		(*PackageResponse_GroupMessageResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
