// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.0
// source: loyalty_service.proto

package pb

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

// Запрос c user_id, order_id, count
type ConcludeUserBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Count         int64                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConcludeUserBalanceRequest) Reset() {
	*x = ConcludeUserBalanceRequest{}
	mi := &file_loyalty_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConcludeUserBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConcludeUserBalanceRequest) ProtoMessage() {}

func (x *ConcludeUserBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loyalty_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConcludeUserBalanceRequest.ProtoReflect.Descriptor instead.
func (*ConcludeUserBalanceRequest) Descriptor() ([]byte, []int) {
	return file_loyalty_service_proto_rawDescGZIP(), []int{0}
}

func (x *ConcludeUserBalanceRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ConcludeUserBalanceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ConcludeUserBalanceRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ConcludeUserBalanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConcludeUserBalanceResponse) Reset() {
	*x = ConcludeUserBalanceResponse{}
	mi := &file_loyalty_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConcludeUserBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConcludeUserBalanceResponse) ProtoMessage() {}

func (x *ConcludeUserBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loyalty_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConcludeUserBalanceResponse.ProtoReflect.Descriptor instead.
func (*ConcludeUserBalanceResponse) Descriptor() ([]byte, []int) {
	return file_loyalty_service_proto_rawDescGZIP(), []int{1}
}

func (x *ConcludeUserBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ConcludeUserBalanceResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type CreateLoyaltyBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLoyaltyBalanceRequest) Reset() {
	*x = CreateLoyaltyBalanceRequest{}
	mi := &file_loyalty_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLoyaltyBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoyaltyBalanceRequest) ProtoMessage() {}

func (x *CreateLoyaltyBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loyalty_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoyaltyBalanceRequest.ProtoReflect.Descriptor instead.
func (*CreateLoyaltyBalanceRequest) Descriptor() ([]byte, []int) {
	return file_loyalty_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLoyaltyBalanceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateLoyaltyBalanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateLoyaltyBalanceResponse) Reset() {
	*x = CreateLoyaltyBalanceResponse{}
	mi := &file_loyalty_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateLoyaltyBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoyaltyBalanceResponse) ProtoMessage() {}

func (x *CreateLoyaltyBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loyalty_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoyaltyBalanceResponse.ProtoReflect.Descriptor instead.
func (*CreateLoyaltyBalanceResponse) Descriptor() ([]byte, []int) {
	return file_loyalty_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLoyaltyBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateLoyaltyBalanceResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_loyalty_service_proto protoreflect.FileDescriptor

var file_loyalty_service_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79,
	0x22, 0x66, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d,
	0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd7, 0x01,
	0x0a, 0x0e, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x63, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74,
	0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x79, 0x61, 0x6c,
	0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x6c,
	0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x78, 0x65, 0x6c, 0x37, 0x39, 0x31, 0x2f, 0x6c, 0x6f,
	0x79, 0x61, 0x6c, 0x74, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_loyalty_service_proto_rawDescOnce sync.Once
	file_loyalty_service_proto_rawDescData []byte
)

func file_loyalty_service_proto_rawDescGZIP() []byte {
	file_loyalty_service_proto_rawDescOnce.Do(func() {
		file_loyalty_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_loyalty_service_proto_rawDesc), len(file_loyalty_service_proto_rawDesc)))
	})
	return file_loyalty_service_proto_rawDescData
}

var file_loyalty_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_loyalty_service_proto_goTypes = []any{
	(*ConcludeUserBalanceRequest)(nil),   // 0: loyalty.ConcludeUserBalanceRequest
	(*ConcludeUserBalanceResponse)(nil),  // 1: loyalty.ConcludeUserBalanceResponse
	(*CreateLoyaltyBalanceRequest)(nil),  // 2: loyalty.CreateLoyaltyBalanceRequest
	(*CreateLoyaltyBalanceResponse)(nil), // 3: loyalty.CreateLoyaltyBalanceResponse
}
var file_loyalty_service_proto_depIdxs = []int32{
	2, // 0: loyalty.LoyaltyService.CreateLoyaltyBalance:input_type -> loyalty.CreateLoyaltyBalanceRequest
	0, // 1: loyalty.LoyaltyService.ConcludeUserBalance:input_type -> loyalty.ConcludeUserBalanceRequest
	3, // 2: loyalty.LoyaltyService.CreateLoyaltyBalance:output_type -> loyalty.CreateLoyaltyBalanceResponse
	1, // 3: loyalty.LoyaltyService.ConcludeUserBalance:output_type -> loyalty.ConcludeUserBalanceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loyalty_service_proto_init() }
func file_loyalty_service_proto_init() {
	if File_loyalty_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_loyalty_service_proto_rawDesc), len(file_loyalty_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loyalty_service_proto_goTypes,
		DependencyIndexes: file_loyalty_service_proto_depIdxs,
		MessageInfos:      file_loyalty_service_proto_msgTypes,
	}.Build()
	File_loyalty_service_proto = out.File
	file_loyalty_service_proto_goTypes = nil
	file_loyalty_service_proto_depIdxs = nil
}
