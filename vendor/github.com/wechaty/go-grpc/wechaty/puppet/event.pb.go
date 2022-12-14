// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: wechaty/puppet/event.proto

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

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED EventType = 0
	EventType_EVENT_TYPE_HEARTBEAT   EventType = 1
	EventType_EVENT_TYPE_MESSAGE     EventType = 2
	EventType_EVENT_TYPE_DONG        EventType = 3
	EventType_EVENT_TYPE_POST        EventType = 4
	EventType_EVENT_TYPE_ERROR       EventType = 16
	EventType_EVENT_TYPE_FRIENDSHIP  EventType = 17
	EventType_EVENT_TYPE_ROOM_INVITE EventType = 18
	EventType_EVENT_TYPE_ROOM_JOIN   EventType = 19
	EventType_EVENT_TYPE_ROOM_LEAVE  EventType = 20
	EventType_EVENT_TYPE_ROOM_TOPIC  EventType = 21
	EventType_EVENT_TYPE_SCAN        EventType = 22
	EventType_EVENT_TYPE_READY       EventType = 23
	EventType_EVENT_TYPE_RESET       EventType = 24
	EventType_EVENT_TYPE_LOGIN       EventType = 25
	EventType_EVENT_TYPE_LOGOUT      EventType = 26
	EventType_EVENT_TYPE_DIRTY       EventType = 27
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0:  "EVENT_TYPE_UNSPECIFIED",
		1:  "EVENT_TYPE_HEARTBEAT",
		2:  "EVENT_TYPE_MESSAGE",
		3:  "EVENT_TYPE_DONG",
		4:  "EVENT_TYPE_POST",
		16: "EVENT_TYPE_ERROR",
		17: "EVENT_TYPE_FRIENDSHIP",
		18: "EVENT_TYPE_ROOM_INVITE",
		19: "EVENT_TYPE_ROOM_JOIN",
		20: "EVENT_TYPE_ROOM_LEAVE",
		21: "EVENT_TYPE_ROOM_TOPIC",
		22: "EVENT_TYPE_SCAN",
		23: "EVENT_TYPE_READY",
		24: "EVENT_TYPE_RESET",
		25: "EVENT_TYPE_LOGIN",
		26: "EVENT_TYPE_LOGOUT",
		27: "EVENT_TYPE_DIRTY",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED": 0,
		"EVENT_TYPE_HEARTBEAT":   1,
		"EVENT_TYPE_MESSAGE":     2,
		"EVENT_TYPE_DONG":        3,
		"EVENT_TYPE_POST":        4,
		"EVENT_TYPE_ERROR":       16,
		"EVENT_TYPE_FRIENDSHIP":  17,
		"EVENT_TYPE_ROOM_INVITE": 18,
		"EVENT_TYPE_ROOM_JOIN":   19,
		"EVENT_TYPE_ROOM_LEAVE":  20,
		"EVENT_TYPE_ROOM_TOPIC":  21,
		"EVENT_TYPE_SCAN":        22,
		"EVENT_TYPE_READY":       23,
		"EVENT_TYPE_RESET":       24,
		"EVENT_TYPE_LOGIN":       25,
		"EVENT_TYPE_LOGOUT":      26,
		"EVENT_TYPE_DIRTY":       27,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_wechaty_puppet_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_wechaty_puppet_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_wechaty_puppet_event_proto_rawDescGZIP(), []int{0}
}

type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_event_proto_rawDescGZIP(), []int{0}
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=wechaty.puppet.EventType" json:"type,omitempty"`
	// TODO: Huan(202002) consider to use a PB Map?
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"` // JSON.stringify({ ... })
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wechaty_puppet_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wechaty_puppet_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_wechaty_puppet_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventResponse) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (x *EventResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_wechaty_puppet_event_proto protoreflect.FileDescriptor

var file_wechaty_puppet_event_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x65,
	0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x22, 0x0e, 0x0a, 0x0c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x0d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x77, 0x65,
	0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0xaa, 0x03, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48,
	0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x4f, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x10, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x48, 0x49, 0x50, 0x10, 0x11, 0x12, 0x1a, 0x0a,
	0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d,
	0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x10, 0x12, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4a, 0x4f, 0x49,
	0x4e, 0x10, 0x13, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x14, 0x12, 0x19,
	0x0a, 0x15, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x4f,
	0x4d, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x10, 0x16, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x18, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x19,
	0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c,
	0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x1a, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x54, 0x59, 0x10, 0x1b, 0x22, 0x04, 0x08,
	0x05, 0x10, 0x0f, 0x42, 0x67, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75,
	0x70, 0x70, 0x65, 0x74, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79, 0x2f, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0xaa,
	0x02, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x79,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x75, 0x70, 0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wechaty_puppet_event_proto_rawDescOnce sync.Once
	file_wechaty_puppet_event_proto_rawDescData = file_wechaty_puppet_event_proto_rawDesc
)

func file_wechaty_puppet_event_proto_rawDescGZIP() []byte {
	file_wechaty_puppet_event_proto_rawDescOnce.Do(func() {
		file_wechaty_puppet_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_wechaty_puppet_event_proto_rawDescData)
	})
	return file_wechaty_puppet_event_proto_rawDescData
}

var file_wechaty_puppet_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wechaty_puppet_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wechaty_puppet_event_proto_goTypes = []interface{}{
	(EventType)(0),        // 0: wechaty.puppet.EventType
	(*EventRequest)(nil),  // 1: wechaty.puppet.EventRequest
	(*EventResponse)(nil), // 2: wechaty.puppet.EventResponse
}
var file_wechaty_puppet_event_proto_depIdxs = []int32{
	0, // 0: wechaty.puppet.EventResponse.type:type_name -> wechaty.puppet.EventType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wechaty_puppet_event_proto_init() }
func file_wechaty_puppet_event_proto_init() {
	if File_wechaty_puppet_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wechaty_puppet_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
		file_wechaty_puppet_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
			RawDescriptor: file_wechaty_puppet_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wechaty_puppet_event_proto_goTypes,
		DependencyIndexes: file_wechaty_puppet_event_proto_depIdxs,
		EnumInfos:         file_wechaty_puppet_event_proto_enumTypes,
		MessageInfos:      file_wechaty_puppet_event_proto_msgTypes,
	}.Build()
	File_wechaty_puppet_event_proto = out.File
	file_wechaty_puppet_event_proto_rawDesc = nil
	file_wechaty_puppet_event_proto_goTypes = nil
	file_wechaty_puppet_event_proto_depIdxs = nil
}
