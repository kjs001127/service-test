// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sip/sip.proto

package sip

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

type CallSessionResponse_Result int32

const (
	CallSessionResponse_EXISTS    CallSessionResponse_Result = 0
	CallSessionResponse_NOT_FOUND CallSessionResponse_Result = 1
)

// Enum value maps for CallSessionResponse_Result.
var (
	CallSessionResponse_Result_name = map[int32]string{
		0: "EXISTS",
		1: "NOT_FOUND",
	}
	CallSessionResponse_Result_value = map[string]int32{
		"EXISTS":    0,
		"NOT_FOUND": 1,
	}
)

func (x CallSessionResponse_Result) Enum() *CallSessionResponse_Result {
	p := new(CallSessionResponse_Result)
	*p = x
	return p
}

func (x CallSessionResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CallSessionResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_sip_sip_proto_enumTypes[0].Descriptor()
}

func (CallSessionResponse_Result) Type() protoreflect.EnumType {
	return &file_sip_sip_proto_enumTypes[0]
}

func (x CallSessionResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CallSessionResponse_Result.Descriptor instead.
func (CallSessionResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{5, 0}
}

type SipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
}

func (x *SipRequest) Reset() {
	*x = SipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_sip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipRequest) ProtoMessage() {}

func (x *SipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sip_sip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipRequest.ProtoReflect.Descriptor instead.
func (*SipRequest) Descriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{0}
}

func (x *SipRequest) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carrier           string  `protobuf:"bytes,1,opt,name=carrier,proto3" json:"carrier,omitempty"`
	CallId            string  `protobuf:"bytes,2,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
	From              string  `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                string  `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Sdp               string  `protobuf:"bytes,5,opt,name=sdp,proto3" json:"sdp,omitempty"`
	PAssertedIdentity *string `protobuf:"bytes,6,opt,name=p_asserted_identity,json=pAssertedIdentity,proto3,oneof" json:"p_asserted_identity,omitempty"`
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_sip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sip_sip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest.ProtoReflect.Descriptor instead.
func (*InviteRequest) Descriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{1}
}

func (x *InviteRequest) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

func (x *InviteRequest) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *InviteRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *InviteRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *InviteRequest) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

func (x *InviteRequest) GetPAssertedIdentity() string {
	if x != nil && x.PAssertedIdentity != nil {
		return *x.PAssertedIdentity
	}
	return ""
}

type AcceptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
	Sdp    string `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *AcceptRequest) Reset() {
	*x = AcceptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_sip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptRequest) ProtoMessage() {}

func (x *AcceptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sip_sip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptRequest.ProtoReflect.Descriptor instead.
func (*AcceptRequest) Descriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptRequest) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *AcceptRequest) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type SipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
}

func (x *SipResponse) Reset() {
	*x = SipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_sip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipResponse) ProtoMessage() {}

func (x *SipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sip_sip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipResponse.ProtoReflect.Descriptor instead.
func (*SipResponse) Descriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{3}
}

func (x *SipResponse) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

type CallSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
}

func (x *CallSessionRequest) Reset() {
	*x = CallSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_sip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallSessionRequest) ProtoMessage() {}

func (x *CallSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sip_sip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallSessionRequest.ProtoReflect.Descriptor instead.
func (*CallSessionRequest) Descriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{4}
}

func (x *CallSessionRequest) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

type CallSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result CallSessionResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=sip.CallSessionResponse_Result" json:"result,omitempty"`
}

func (x *CallSessionResponse) Reset() {
	*x = CallSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_sip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallSessionResponse) ProtoMessage() {}

func (x *CallSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sip_sip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallSessionResponse.ProtoReflect.Descriptor instead.
func (*CallSessionResponse) Descriptor() ([]byte, []int) {
	return file_sip_sip_proto_rawDescGZIP(), []int{5}
}

func (x *CallSessionResponse) GetResult() CallSessionResponse_Result {
	if x != nil {
		return x.Result
	}
	return CallSessionResponse_EXISTS
}

var File_sip_sip_proto protoreflect.FileDescriptor

var file_sip_sip_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x69, 0x70, 0x2f, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x73, 0x69, 0x70, 0x22, 0x25, 0x0a, 0x0a, 0x53, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0xc5, 0x01, 0x0a, 0x0d,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x64, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x12, 0x33, 0x0a, 0x13, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x70, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f,
	0x70, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x3a, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x64, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x22,
	0x26, 0x0a, 0x0b, 0x53, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x6c, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x73, 0x69, 0x70, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x32, 0xe0, 0x01, 0x0a, 0x03,
	0x53, 0x69, 0x70, 0x12, 0x32, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x12, 0x2e, 0x73, 0x69, 0x70, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x69, 0x70, 0x2e, 0x53, 0x69, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x73, 0x69, 0x70, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x69, 0x70, 0x2e,
	0x53, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x42,
	0x79, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x73, 0x69, 0x70, 0x2e, 0x53, 0x69, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x69, 0x70, 0x2e, 0x53, 0x69,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x73, 0x69,
	0x70, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x69, 0x70, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24,
	0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sip_sip_proto_rawDescOnce sync.Once
	file_sip_sip_proto_rawDescData = file_sip_sip_proto_rawDesc
)

func file_sip_sip_proto_rawDescGZIP() []byte {
	file_sip_sip_proto_rawDescOnce.Do(func() {
		file_sip_sip_proto_rawDescData = protoimpl.X.CompressGZIP(file_sip_sip_proto_rawDescData)
	})
	return file_sip_sip_proto_rawDescData
}

var file_sip_sip_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sip_sip_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sip_sip_proto_goTypes = []interface{}{
	(CallSessionResponse_Result)(0), // 0: sip.CallSessionResponse.Result
	(*SipRequest)(nil),              // 1: sip.SipRequest
	(*InviteRequest)(nil),           // 2: sip.InviteRequest
	(*AcceptRequest)(nil),           // 3: sip.AcceptRequest
	(*SipResponse)(nil),             // 4: sip.SipResponse
	(*CallSessionRequest)(nil),      // 5: sip.CallSessionRequest
	(*CallSessionResponse)(nil),     // 6: sip.CallSessionResponse
}
var file_sip_sip_proto_depIdxs = []int32{
	0, // 0: sip.CallSessionResponse.result:type_name -> sip.CallSessionResponse.Result
	2, // 1: sip.Sip.InviteCall:input_type -> sip.InviteRequest
	3, // 2: sip.Sip.AcceptCall:input_type -> sip.AcceptRequest
	1, // 3: sip.Sip.ByeCall:input_type -> sip.SipRequest
	5, // 4: sip.Sip.GetCallSession:input_type -> sip.CallSessionRequest
	4, // 5: sip.Sip.InviteCall:output_type -> sip.SipResponse
	4, // 6: sip.Sip.AcceptCall:output_type -> sip.SipResponse
	4, // 7: sip.Sip.ByeCall:output_type -> sip.SipResponse
	6, // 8: sip.Sip.GetCallSession:output_type -> sip.CallSessionResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sip_sip_proto_init() }
func file_sip_sip_proto_init() {
	if File_sip_sip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sip_sip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipRequest); i {
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
		file_sip_sip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest); i {
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
		file_sip_sip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptRequest); i {
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
		file_sip_sip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipResponse); i {
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
		file_sip_sip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallSessionRequest); i {
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
		file_sip_sip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallSessionResponse); i {
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
	file_sip_sip_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sip_sip_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sip_sip_proto_goTypes,
		DependencyIndexes: file_sip_sip_proto_depIdxs,
		EnumInfos:         file_sip_sip_proto_enumTypes,
		MessageInfos:      file_sip_sip_proto_msgTypes,
	}.Build()
	File_sip_sip_proto = out.File
	file_sip_sip_proto_rawDesc = nil
	file_sip_sip_proto_goTypes = nil
	file_sip_sip_proto_depIdxs = nil
}
