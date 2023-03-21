// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: doc/goctl/rpc/product.proto

package product

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

//
//Common
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{0}
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Total    int64 `protobuf:"varint,3,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Page) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Page) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// List
type ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category int32  `protobuf:"varint,3,opt,name=category,proto3" json:"category,omitempty"`
	Stock    int64  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *ListItem) Reset() {
	*x = ListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItem) ProtoMessage() {}

func (x *ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItem.ProtoReflect.Descriptor instead.
func (*ListItem) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{2}
}

func (x *ListItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListItem) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *ListItem) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search   string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Category int32  `protobuf:"varint,2,opt,name=category,proto3" json:"category,omitempty"`
	Page     int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{3}
}

func (x *ListReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListReq) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *ListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Page *Page       `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{4}
}

func (x *ListReply) GetList() []*ListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListReply) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

// Detail
type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{5}
}

func (x *IdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category    int32  `protobuf:"varint,3,opt,name=category,proto3" json:"category,omitempty"`
	Stock       int64  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *DetailReply) Reset() {
	*x = DetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailReply) ProtoMessage() {}

func (x *DetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailReply.ProtoReflect.Descriptor instead.
func (*DetailReply) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{6}
}

func (x *DetailReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DetailReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DetailReply) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *DetailReply) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *DetailReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DetailReply) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Edit
type EditReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category    int32  `protobuf:"varint,3,opt,name=category,proto3" json:"category,omitempty"`
	Stock       int64  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EditReq) Reset() {
	*x = EditReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditReq) ProtoMessage() {}

func (x *EditReq) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditReq.ProtoReflect.Descriptor instead.
func (*EditReq) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{7}
}

func (x *EditReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditReq) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *EditReq) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *EditReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type IdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReply) Reset() {
	*x = IdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_goctl_rpc_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReply) ProtoMessage() {}

func (x *IdReply) ProtoReflect() protoreflect.Message {
	mi := &file_doc_goctl_rpc_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReply.ProtoReflect.Descriptor instead.
func (*IdReply) Descriptor() ([]byte, []int) {
	return file_doc_goctl_rpc_product_proto_rawDescGZIP(), []int{8}
}

func (x *IdReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_doc_goctl_rpc_product_proto protoreflect.FileDescriptor

var file_doc_goctl_rpc_product_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x6f, 0x63, 0x2f, 0x67, 0x6f, 0x63, 0x74, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x4c, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x62, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x22, 0x6d, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x55, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x45, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19,
	0x0a, 0x07, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbd, 0x01, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x28, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doc_goctl_rpc_product_proto_rawDescOnce sync.Once
	file_doc_goctl_rpc_product_proto_rawDescData = file_doc_goctl_rpc_product_proto_rawDesc
)

func file_doc_goctl_rpc_product_proto_rawDescGZIP() []byte {
	file_doc_goctl_rpc_product_proto_rawDescOnce.Do(func() {
		file_doc_goctl_rpc_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_doc_goctl_rpc_product_proto_rawDescData)
	})
	return file_doc_goctl_rpc_product_proto_rawDescData
}

var file_doc_goctl_rpc_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_doc_goctl_rpc_product_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: product.Empty
	(*Page)(nil),        // 1: product.Page
	(*ListItem)(nil),    // 2: product.ListItem
	(*ListReq)(nil),     // 3: product.ListReq
	(*ListReply)(nil),   // 4: product.ListReply
	(*IdReq)(nil),       // 5: product.IdReq
	(*DetailReply)(nil), // 6: product.DetailReply
	(*EditReq)(nil),     // 7: product.EditReq
	(*IdReply)(nil),     // 8: product.IdReply
}
var file_doc_goctl_rpc_product_proto_depIdxs = []int32{
	2, // 0: product.ListReply.list:type_name -> product.ListItem
	1, // 1: product.ListReply.page:type_name -> product.Page
	3, // 2: product.product.List:input_type -> product.ListReq
	5, // 3: product.product.Detail:input_type -> product.IdReq
	7, // 4: product.product.Edit:input_type -> product.EditReq
	5, // 5: product.product.Remove:input_type -> product.IdReq
	4, // 6: product.product.List:output_type -> product.ListReply
	6, // 7: product.product.Detail:output_type -> product.DetailReply
	8, // 8: product.product.Edit:output_type -> product.IdReply
	0, // 9: product.product.Remove:output_type -> product.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_doc_goctl_rpc_product_proto_init() }
func file_doc_goctl_rpc_product_proto_init() {
	if File_doc_goctl_rpc_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doc_goctl_rpc_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItem); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReply); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailReply); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditReq); i {
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
		file_doc_goctl_rpc_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReply); i {
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
			RawDescriptor: file_doc_goctl_rpc_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_doc_goctl_rpc_product_proto_goTypes,
		DependencyIndexes: file_doc_goctl_rpc_product_proto_depIdxs,
		MessageInfos:      file_doc_goctl_rpc_product_proto_msgTypes,
	}.Build()
	File_doc_goctl_rpc_product_proto = out.File
	file_doc_goctl_rpc_product_proto_rawDesc = nil
	file_doc_goctl_rpc_product_proto_goTypes = nil
	file_doc_goctl_rpc_product_proto_depIdxs = nil
}
