// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: proto/payment-engine/payment-engine.proto

package payment

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

type PreAuthorizeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MerchantId    string                 `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Amount        int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PreAuthorizeRequest) Reset() {
	*x = PreAuthorizeRequest{}
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PreAuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthorizeRequest) ProtoMessage() {}

func (x *PreAuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthorizeRequest.ProtoReflect.Descriptor instead.
func (*PreAuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_engine_payment_engine_proto_rawDescGZIP(), []int{0}
}

func (x *PreAuthorizeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PreAuthorizeRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *PreAuthorizeRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PreAuthorizeResponse struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	PreAuthorizationId string                 `protobuf:"bytes,1,opt,name=pre_authorization_id,json=preAuthorizationId,proto3" json:"pre_authorization_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PreAuthorizeResponse) Reset() {
	*x = PreAuthorizeResponse{}
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PreAuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthorizeResponse) ProtoMessage() {}

func (x *PreAuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthorizeResponse.ProtoReflect.Descriptor instead.
func (*PreAuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_engine_payment_engine_proto_rawDescGZIP(), []int{1}
}

func (x *PreAuthorizeResponse) GetPreAuthorizationId() string {
	if x != nil {
		return x.PreAuthorizationId
	}
	return ""
}

type CaptureRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	PreAuthorizationId string                 `protobuf:"bytes,1,opt,name=pre_authorization_id,json=preAuthorizationId,proto3" json:"pre_authorization_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CaptureRequest) Reset() {
	*x = CaptureRequest{}
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CaptureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureRequest) ProtoMessage() {}

func (x *CaptureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureRequest.ProtoReflect.Descriptor instead.
func (*CaptureRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_engine_payment_engine_proto_rawDescGZIP(), []int{2}
}

func (x *CaptureRequest) GetPreAuthorizationId() string {
	if x != nil {
		return x.PreAuthorizationId
	}
	return ""
}

type CaptureResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	PaymentCaptureId string                 `protobuf:"bytes,1,opt,name=payment_capture_id,json=paymentCaptureId,proto3" json:"payment_capture_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CaptureResponse) Reset() {
	*x = CaptureResponse{}
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CaptureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureResponse) ProtoMessage() {}

func (x *CaptureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureResponse.ProtoReflect.Descriptor instead.
func (*CaptureResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_engine_payment_engine_proto_rawDescGZIP(), []int{3}
}

func (x *CaptureResponse) GetPaymentCaptureId() string {
	if x != nil {
		return x.PaymentCaptureId
	}
	return ""
}

type ReimburseRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	PreAuthorizationId string                 `protobuf:"bytes,1,opt,name=pre_authorization_id,json=preAuthorizationId,proto3" json:"pre_authorization_id,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ReimburseRequest) Reset() {
	*x = ReimburseRequest{}
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReimburseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReimburseRequest) ProtoMessage() {}

func (x *ReimburseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReimburseRequest.ProtoReflect.Descriptor instead.
func (*ReimburseRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_engine_payment_engine_proto_rawDescGZIP(), []int{4}
}

func (x *ReimburseRequest) GetPreAuthorizationId() string {
	if x != nil {
		return x.PreAuthorizationId
	}
	return ""
}

type ReimburseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReimburseResponse) Reset() {
	*x = ReimburseResponse{}
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReimburseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReimburseResponse) ProtoMessage() {}

func (x *ReimburseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_engine_payment_engine_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReimburseResponse.ProtoReflect.Descriptor instead.
func (*ReimburseResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_engine_payment_engine_proto_rawDescGZIP(), []int{5}
}

func (x *ReimburseResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_payment_engine_payment_engine_proto protoreflect.FileDescriptor

var file_proto_payment_engine_payment_engine_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x67, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x48, 0x0a,
	0x14, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x43,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x10,
	0x52, 0x65, 0x69, 0x6d, 0x62, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x70, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x52, 0x65, 0x69, 0x6d, 0x62, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xd8, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0c, 0x50,
	0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x52, 0x65, 0x69, 0x6d, 0x62, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65,
	0x69, 0x6d, 0x62, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x69, 0x6d, 0x62, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b, 0x6c, 0x72, 0x73, 0x6e, 0x2f,
	0x67, 0x52, 0x50, 0x43, 0x2d, 0x64, 0x65, 0x66, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_proto_payment_engine_payment_engine_proto_rawDescOnce sync.Once
	file_proto_payment_engine_payment_engine_proto_rawDescData []byte
)

func file_proto_payment_engine_payment_engine_proto_rawDescGZIP() []byte {
	file_proto_payment_engine_payment_engine_proto_rawDescOnce.Do(func() {
		file_proto_payment_engine_payment_engine_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_payment_engine_payment_engine_proto_rawDesc), len(file_proto_payment_engine_payment_engine_proto_rawDesc)))
	})
	return file_proto_payment_engine_payment_engine_proto_rawDescData
}

var file_proto_payment_engine_payment_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_payment_engine_payment_engine_proto_goTypes = []any{
	(*PreAuthorizeRequest)(nil),  // 0: payment.PreAuthorizeRequest
	(*PreAuthorizeResponse)(nil), // 1: payment.PreAuthorizeResponse
	(*CaptureRequest)(nil),       // 2: payment.CaptureRequest
	(*CaptureResponse)(nil),      // 3: payment.CaptureResponse
	(*ReimburseRequest)(nil),     // 4: payment.ReimburseRequest
	(*ReimburseResponse)(nil),    // 5: payment.ReimburseResponse
}
var file_proto_payment_engine_payment_engine_proto_depIdxs = []int32{
	0, // 0: payment.Payment.PreAuthorize:input_type -> payment.PreAuthorizeRequest
	2, // 1: payment.Payment.Capture:input_type -> payment.CaptureRequest
	4, // 2: payment.Payment.Reimburse:input_type -> payment.ReimburseRequest
	1, // 3: payment.Payment.PreAuthorize:output_type -> payment.PreAuthorizeResponse
	3, // 4: payment.Payment.Capture:output_type -> payment.CaptureResponse
	5, // 5: payment.Payment.Reimburse:output_type -> payment.ReimburseResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_payment_engine_payment_engine_proto_init() }
func file_proto_payment_engine_payment_engine_proto_init() {
	if File_proto_payment_engine_payment_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_payment_engine_payment_engine_proto_rawDesc), len(file_proto_payment_engine_payment_engine_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payment_engine_payment_engine_proto_goTypes,
		DependencyIndexes: file_proto_payment_engine_payment_engine_proto_depIdxs,
		MessageInfos:      file_proto_payment_engine_payment_engine_proto_msgTypes,
	}.Build()
	File_proto_payment_engine_payment_engine_proto = out.File
	file_proto_payment_engine_payment_engine_proto_goTypes = nil
	file_proto_payment_engine_payment_engine_proto_depIdxs = nil
}
