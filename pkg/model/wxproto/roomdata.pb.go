// v3 & v4 通用，可能会有部分字段差异

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: roomdata.proto

package wxproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RoomData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*RoomDataUser        `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	RoomCap       *int32                 `protobuf:"varint,5,opt,name=roomCap,proto3,oneof" json:"roomCap,omitempty"` // 只在第一份数据中出现，值为500
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomData) Reset() {
	*x = RoomData{}
	mi := &file_roomdata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomData) ProtoMessage() {}

func (x *RoomData) ProtoReflect() protoreflect.Message {
	mi := &file_roomdata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomData.ProtoReflect.Descriptor instead.
func (*RoomData) Descriptor() ([]byte, []int) {
	return file_roomdata_proto_rawDescGZIP(), []int{0}
}

func (x *RoomData) GetUsers() []*RoomDataUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *RoomData) GetRoomCap() int32 {
	if x != nil && x.RoomCap != nil {
		return *x.RoomCap
	}
	return 0
}

type RoomDataUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`             // 用户ID或名称
	DisplayName   *string                `protobuf:"bytes,2,opt,name=displayName,proto3,oneof" json:"displayName,omitempty"` // 显示名称，可能是UTF-8编码的中文，部分记录可能为空
	Status        int32                  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`                // 状态码，值范围0-9
	Inviter       *string                `protobuf:"bytes,4,opt,name=inviter,proto3,oneof" json:"inviter,omitempty"`         // 邀请人
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomDataUser) Reset() {
	*x = RoomDataUser{}
	mi := &file_roomdata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomDataUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomDataUser) ProtoMessage() {}

func (x *RoomDataUser) ProtoReflect() protoreflect.Message {
	mi := &file_roomdata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomDataUser.ProtoReflect.Descriptor instead.
func (*RoomDataUser) Descriptor() ([]byte, []int) {
	return file_roomdata_proto_rawDescGZIP(), []int{1}
}

func (x *RoomDataUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RoomDataUser) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *RoomDataUser) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RoomDataUser) GetInviter() string {
	if x != nil && x.Inviter != nil {
		return *x.Inviter
	}
	return ""
}

var File_roomdata_proto protoreflect.FileDescriptor

var file_roomdata_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x67,
	0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x61, 0x70, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x61, 0x70, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x3b, 0x77, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_roomdata_proto_rawDescOnce sync.Once
	file_roomdata_proto_rawDescData []byte
)

func file_roomdata_proto_rawDescGZIP() []byte {
	file_roomdata_proto_rawDescOnce.Do(func() {
		file_roomdata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_roomdata_proto_rawDesc), len(file_roomdata_proto_rawDesc)))
	})
	return file_roomdata_proto_rawDescData
}

var file_roomdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_roomdata_proto_goTypes = []any{
	(*RoomData)(nil),     // 0: app.protobuf.RoomData
	(*RoomDataUser)(nil), // 1: app.protobuf.RoomDataUser
}
var file_roomdata_proto_depIdxs = []int32{
	1, // 0: app.protobuf.RoomData.users:type_name -> app.protobuf.RoomDataUser
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_roomdata_proto_init() }
func file_roomdata_proto_init() {
	if File_roomdata_proto != nil {
		return
	}
	file_roomdata_proto_msgTypes[0].OneofWrappers = []any{}
	file_roomdata_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_roomdata_proto_rawDesc), len(file_roomdata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_roomdata_proto_goTypes,
		DependencyIndexes: file_roomdata_proto_depIdxs,
		MessageInfos:      file_roomdata_proto_msgTypes,
	}.Build()
	File_roomdata_proto = out.File
	file_roomdata_proto_goTypes = nil
	file_roomdata_proto_depIdxs = nil
}
