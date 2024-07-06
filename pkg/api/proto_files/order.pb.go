// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto_files/order.proto

package api

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

type OrderStatus int32

const (
	OrderStatus_NULL_ORD_STATUS OrderStatus = 0
	OrderStatus_FULFILLED       OrderStatus = 1
	OrderStatus_PENDING         OrderStatus = 2
	OrderStatus_CANCELLED       OrderStatus = 3
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "NULL_ORD_STATUS",
		1: "FULFILLED",
		2: "PENDING",
		3: "CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"NULL_ORD_STATUS": 0,
		"FULFILLED":       1,
		"PENDING":         2,
		"CANCELLED":       3,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_files_order_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_proto_files_order_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId int32       `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	Count     int32       `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Status    OrderStatus `protobuf:"varint,4,opt,name=status,proto3,enum=com.store.OrderStatus" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Order) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_NULL_ORD_STATUS
}

type NewOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32       `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Status    OrderStatus `protobuf:"varint,3,opt,name=status,proto3,enum=com.store.OrderStatus" json:"status,omitempty"`
}

func (x *NewOrder) Reset() {
	*x = NewOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOrder) ProtoMessage() {}

func (x *NewOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOrder.ProtoReflect.Descriptor instead.
func (*NewOrder) Descriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{1}
}

func (x *NewOrder) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *NewOrder) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *NewOrder) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_NULL_ORD_STATUS
}

type OrderId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderId) Reset() {
	*x = OrderId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderId) ProtoMessage() {}

func (x *OrderId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderId.ProtoReflect.Descriptor instead.
func (*OrderId) Descriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32       `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Status    OrderStatus `protobuf:"varint,2,opt,name=status,proto3,enum=com.store.OrderStatus" json:"status,omitempty"`
}

func (x *OrderSearchRequest) Reset() {
	*x = OrderSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSearchRequest) ProtoMessage() {}

func (x *OrderSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSearchRequest.ProtoReflect.Descriptor instead.
func (*OrderSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{3}
}

func (x *OrderSearchRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderSearchRequest) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_NULL_ORD_STATUS
}

type OrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderListResponse) Reset() {
	*x = OrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResponse) ProtoMessage() {}

func (x *OrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResponse.ProtoReflect.Descriptor instead.
func (*OrderListResponse) Descriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderListResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_files_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_files_order_proto protoreflect.FileDescriptor

var file_proto_files_order_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x22, 0x7b, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x6e, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x19, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x12,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x3d, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x29, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x4d, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x55, 0x4c,
	0x4c, 0x5f, 0x4f, 0x52, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0xba, 0x02, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4a, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42,
	0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x73,
	0x70, 0x65, 0x63, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x62,
	0x66, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_order_proto_rawDescOnce sync.Once
	file_proto_files_order_proto_rawDescData = file_proto_files_order_proto_rawDesc
)

func file_proto_files_order_proto_rawDescGZIP() []byte {
	file_proto_files_order_proto_rawDescOnce.Do(func() {
		file_proto_files_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_order_proto_rawDescData)
	})
	return file_proto_files_order_proto_rawDescData
}

var file_proto_files_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_files_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_files_order_proto_goTypes = []interface{}{
	(OrderStatus)(0),           // 0: com.store.OrderStatus
	(*Order)(nil),              // 1: com.store.Order
	(*NewOrder)(nil),           // 2: com.store.NewOrder
	(*OrderId)(nil),            // 3: com.store.OrderId
	(*OrderSearchRequest)(nil), // 4: com.store.OrderSearchRequest
	(*OrderListResponse)(nil),  // 5: com.store.OrderListResponse
	(*OrderResponse)(nil),      // 6: com.store.OrderResponse
}
var file_proto_files_order_proto_depIdxs = []int32{
	0, // 0: com.store.Order.status:type_name -> com.store.OrderStatus
	0, // 1: com.store.NewOrder.status:type_name -> com.store.OrderStatus
	0, // 2: com.store.OrderSearchRequest.status:type_name -> com.store.OrderStatus
	1, // 3: com.store.OrderListResponse.orders:type_name -> com.store.Order
	4, // 4: com.store.OrderService.SearchOrders:input_type -> com.store.OrderSearchRequest
	3, // 5: com.store.OrderService.GetOrder:input_type -> com.store.OrderId
	2, // 6: com.store.OrderService.AddOrder:input_type -> com.store.NewOrder
	1, // 7: com.store.OrderService.UpdateOrder:input_type -> com.store.Order
	3, // 8: com.store.OrderService.DeleteOrder:input_type -> com.store.OrderId
	5, // 9: com.store.OrderService.SearchOrders:output_type -> com.store.OrderListResponse
	1, // 10: com.store.OrderService.GetOrder:output_type -> com.store.Order
	3, // 11: com.store.OrderService.AddOrder:output_type -> com.store.OrderId
	6, // 12: com.store.OrderService.UpdateOrder:output_type -> com.store.OrderResponse
	6, // 13: com.store.OrderService.DeleteOrder:output_type -> com.store.OrderResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_files_order_proto_init() }
func file_proto_files_order_proto_init() {
	if File_proto_files_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_files_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_files_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOrder); i {
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
		file_proto_files_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderId); i {
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
		file_proto_files_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSearchRequest); i {
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
		file_proto_files_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResponse); i {
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
		file_proto_files_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
			RawDescriptor: file_proto_files_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_files_order_proto_goTypes,
		DependencyIndexes: file_proto_files_order_proto_depIdxs,
		EnumInfos:         file_proto_files_order_proto_enumTypes,
		MessageInfos:      file_proto_files_order_proto_msgTypes,
	}.Build()
	File_proto_files_order_proto = out.File
	file_proto_files_order_proto_rawDesc = nil
	file_proto_files_order_proto_goTypes = nil
	file_proto_files_order_proto_depIdxs = nil
}
