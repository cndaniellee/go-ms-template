// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: doc/goctl/rpc/user.proto

package user

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

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_user_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_user_proto_rawDescGZIP(), []int{1}
}

func (x *AuthReply) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CurrentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CurrentReq) Reset() {
	*x = CurrentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentReq) ProtoMessage() {}

func (x *CurrentReq) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentReq.ProtoReflect.Descriptor instead.
func (*CurrentReq) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_user_proto_rawDescGZIP(), []int{2}
}

func (x *CurrentReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CurrentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CurrentReply) Reset() {
	*x = CurrentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentReply) ProtoMessage() {}

func (x *CurrentReply) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentReply.ProtoReflect.Descriptor instead.
func (*CurrentReply) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_user_proto_rawDescGZIP(), []int{3}
}

func (x *CurrentReply) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_doc_goctl_rpc_user_proto protoreflect.FileDescriptor

var file_doc_goctl_rpc_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x6f, 0x63, 0x2f, 0x67, 0x6f, 0x63, 0x74, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x41, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x23, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8c, 0x01, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doc_goctl_rpc_user_proto_rawDescOnce sync.Once
	file_doc_goctl_rpc_user_proto_rawDescData = file_doc_goctl_rpc_user_proto_rawDesc
)

func file_doc_goctl_rpc_user_proto_rawDescGZIP() []byte {
	file_doc_goctl_rpc_user_proto_rawDescOnce.Do(func() {
		file_doc_goctl_rpc_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_doc_goctl_rpc_user_proto_rawDescData)
	})
	return file_doc_goctl_rpc_user_proto_rawDescData
}

var file_doc_goctl_rpc_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_doc_goctl_rpc_user_proto_goTypes = []interface{}{
	(*AuthReq)(nil),      // 0: user.AuthReq
	(*AuthReply)(nil),    // 1: user.AuthReply
	(*CurrentReq)(nil),   // 2: user.CurrentReq
	(*CurrentReply)(nil), // 3: user.CurrentReply
}
var file_doc_goctl_rpc_user_proto_depIdxs = []int32{
	0, // 0: user.user.Login:input_type -> user.AuthReq
	0, // 1: user.user.Register:input_type -> user.AuthReq
	2, // 2: user.user.Current:input_type -> user.CurrentReq
	1, // 3: user.user.Login:output_type -> user.AuthReply
	1, // 4: user.user.Register:output_type -> user.AuthReply
	3, // 5: user.user.Current:output_type -> user.CurrentReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_doc_goctl_rpc_user_proto_init() }
func file_doc_goctl_rpc_user_proto_init() {
	if File_doc_goctl_rpc_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doc_goctl_rpc_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_doc_goctl_rpc_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_doc_goctl_rpc_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentReq); i {
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
		file_doc_goctl_rpc_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentReply); i {
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
			RawDescriptor: file_doc_goctl_rpc_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doc_goctl_rpc_user_proto_goTypes,
		DependencyIndexes: file_doc_goctl_rpc_user_proto_depIdxs,
		MessageInfos:      file_doc_goctl_rpc_user_proto_msgTypes,
	}.Build()
	File_doc_goctl_rpc_user_proto = out.File
	file_doc_goctl_rpc_user_proto_rawDesc = nil
	file_doc_goctl_rpc_user_proto_goTypes = nil
	file_doc_goctl_rpc_user_proto_depIdxs = nil
}