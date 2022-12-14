// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: wechaty/puppet/room-member.proto

package puppet

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

type RoomMemberPayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MemberId string `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *RoomMemberPayloadRequest) Reset() {
	*x = RoomMemberPayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_room_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMemberPayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMemberPayloadRequest) ProtoMessage() {}

func (x *RoomMemberPayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_room_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMemberPayloadRequest.ProtoReflect.Descriptor instead.
func (*RoomMemberPayloadRequest) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_room_member_proto_rawDescGZIP(), []int{0}
}

func (x *RoomMemberPayloadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomMemberPayloadRequest) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type RoomMemberPayloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomAlias string `protobuf:"bytes,2,opt,name=room_alias,json=roomAlias,proto3" json:"room_alias,omitempty"`
	InviterId string `protobuf:"bytes,3,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
	Avatar    string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RoomMemberPayloadResponse) Reset() {
	*x = RoomMemberPayloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_room_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMemberPayloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMemberPayloadResponse) ProtoMessage() {}

func (x *RoomMemberPayloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_room_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMemberPayloadResponse.ProtoReflect.Descriptor instead.
func (*RoomMemberPayloadResponse) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_room_member_proto_rawDescGZIP(), []int{1}
}

func (x *RoomMemberPayloadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomMemberPayloadResponse) GetRoomAlias() string {
	if x != nil {
		return x.RoomAlias
	}
	return ""
}

func (x *RoomMemberPayloadResponse) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *RoomMemberPayloadResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RoomMemberPayloadResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RoomMemberListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoomMemberListRequest) Reset() {
	*x = RoomMemberListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_room_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMemberListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMemberListRequest) ProtoMessage() {}

func (x *RoomMemberListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_room_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMemberListRequest.ProtoReflect.Descriptor instead.
func (*RoomMemberListRequest) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_room_member_proto_rawDescGZIP(), []int{2}
}

func (x *RoomMemberListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RoomMemberListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberIds []string `protobuf:"bytes,1,rep,name=member_ids,json=memberIds,proto3" json:"member_ids,omitempty"`
}

func (x *RoomMemberListResponse) Reset() {
	*x = RoomMemberListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_room_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomMemberListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMemberListResponse) ProtoMessage() {}

func (x *RoomMemberListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_room_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMemberListResponse.ProtoReflect.Descriptor instead.
func (*RoomMemberListResponse) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_room_member_proto_rawDescGZIP(), []int{3}
}

func (x *RoomMemberListResponse) GetMemberIds() []string {
	if x != nil {
		return x.MemberIds
	}
	return nil
}

var File_wechaty_puppet_room_member_proto protoreflect.FileDescriptor

var file_wechaty_puppet_room_member_proto_rawDesc = []byte{
	0x0a, 0x20, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74,
	0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70,
	0x65, 0x74, 0x22, 0x47, 0x0a, 0x18, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x19,
	0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x16,
	0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x42, 0x67, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65,
	0x74, 0xaa, 0x02, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wechaty_puppet_room_member_proto_rawDescOnce sync.Once
	file_wechaty_puppet_room_member_proto_rawDescData = file_wechaty_puppet_room_member_proto_rawDesc
)

func file_wechaty_puppet_room_member_proto_rawDescGZIP() []byte {
	file_wechaty_puppet_room_member_proto_rawDescOnce.Do(func() {
		file_wechaty_puppet_room_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_wechaty_puppet_room_member_proto_rawDescData)
	})
	return file_wechaty_puppet_room_member_proto_rawDescData
}

var file_wechaty_puppet_room_member_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wechaty_puppet_room_member_proto_goTypes = []interface{}{
	(*RoomMemberPayloadRequest)(nil),  // 0: wechaty.puppet.RoomMemberPayloadRequest
	(*RoomMemberPayloadResponse)(nil), // 1: wechaty.puppet.RoomMemberPayloadResponse
	(*RoomMemberListRequest)(nil),     // 2: wechaty.puppet.RoomMemberListRequest
	(*RoomMemberListResponse)(nil),    // 3: wechaty.puppet.RoomMemberListResponse
}
var file_wechaty_puppet_room_member_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wechaty_puppet_room_member_proto_init() }
func file_wechaty_puppet_room_member_proto_init() {
	if File_wechaty_puppet_room_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wechaty_puppet_room_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMemberPayloadRequest); i {
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
		file_wechaty_puppet_room_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMemberPayloadResponse); i {
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
		file_wechaty_puppet_room_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMemberListRequest); i {
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
		file_wechaty_puppet_room_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomMemberListResponse); i {
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
			RawDescriptor: file_wechaty_puppet_room_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wechaty_puppet_room_member_proto_goTypes,
		DependencyIndexes: file_wechaty_puppet_room_member_proto_depIdxs,
		MessageInfos:      file_wechaty_puppet_room_member_proto_msgTypes,
	}.Build()
	File_wechaty_puppet_room_member_proto = out.File
	file_wechaty_puppet_room_member_proto_rawDesc = nil
	file_wechaty_puppet_room_member_proto_goTypes = nil
	file_wechaty_puppet_room_member_proto_depIdxs = nil
}
