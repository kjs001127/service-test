// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: sfu/sfu.proto

package sfu

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

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip         string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	SipId      string `protobuf:"bytes,2,opt,name=sipId,proto3" json:"sipId,omitempty"`
	SipPw      string `protobuf:"bytes,3,opt,name=sipPw,proto3" json:"sipPw,omitempty"`
	PublicIp   string `protobuf:"bytes,4,opt,name=publicIp,proto3" json:"publicIp,omitempty"`
	InternalIp string `protobuf:"bytes,5,opt,name=internalIp,proto3" json:"internalIp,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sfu_sfu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sfu_sfu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_sfu_sfu_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SubscribeRequest) GetSipId() string {
	if x != nil {
		return x.SipId
	}
	return ""
}

func (x *SubscribeRequest) GetSipPw() string {
	if x != nil {
		return x.SipPw
	}
	return ""
}

func (x *SubscribeRequest) GetPublicIp() string {
	if x != nil {
		return x.PublicIp
	}
	return ""
}

func (x *SubscribeRequest) GetInternalIp() string {
	if x != nil {
		return x.InternalIp
	}
	return ""
}

type MeetKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetId     string `protobuf:"bytes,1,opt,name=meetId,proto3" json:"meetId,omitempty"`
	PersonId   string `protobuf:"bytes,2,opt,name=personId,proto3" json:"personId,omitempty"`
	PersonType string `protobuf:"bytes,3,opt,name=personType,proto3" json:"personType,omitempty"`
}

func (x *MeetKey) Reset() {
	*x = MeetKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sfu_sfu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeetKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeetKey) ProtoMessage() {}

func (x *MeetKey) ProtoReflect() protoreflect.Message {
	mi := &file_sfu_sfu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeetKey.ProtoReflect.Descriptor instead.
func (*MeetKey) Descriptor() ([]byte, []int) {
	return file_sfu_sfu_proto_rawDescGZIP(), []int{1}
}

func (x *MeetKey) GetMeetId() string {
	if x != nil {
		return x.MeetId
	}
	return ""
}

func (x *MeetKey) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *MeetKey) GetPersonType() string {
	if x != nil {
		return x.PersonType
	}
	return ""
}

type WebRtcRelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method  *string  `protobuf:"bytes,1,opt,name=method,proto3,oneof" json:"method,omitempty"`
	Params  []byte   `protobuf:"bytes,2,opt,name=params,proto3,oneof" json:"params,omitempty"`
	MeetKey *MeetKey `protobuf:"bytes,3,opt,name=meetKey,proto3,oneof" json:"meetKey,omitempty"`
}

func (x *WebRtcRelay) Reset() {
	*x = WebRtcRelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sfu_sfu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebRtcRelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebRtcRelay) ProtoMessage() {}

func (x *WebRtcRelay) ProtoReflect() protoreflect.Message {
	mi := &file_sfu_sfu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebRtcRelay.ProtoReflect.Descriptor instead.
func (*WebRtcRelay) Descriptor() ([]byte, []int) {
	return file_sfu_sfu_proto_rawDescGZIP(), []int{2}
}

func (x *WebRtcRelay) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *WebRtcRelay) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *WebRtcRelay) GetMeetKey() *MeetKey {
	if x != nil {
		return x.MeetKey
	}
	return nil
}

var File_sfu_sfu_proto protoreflect.FileDescriptor

var file_sfu_sfu_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x66, 0x75, 0x2f, 0x73, 0x66, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x73, 0x66, 0x75, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x70,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x69, 0x70, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x69, 0x70, 0x50, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x69, 0x70, 0x50, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x70, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x65, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x96, 0x01, 0x0a, 0x0b, 0x57, 0x65, 0x62, 0x52, 0x74, 0x63, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x1b, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x66,
	0x75, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x48, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x32, 0x44, 0x0a, 0x03, 0x53, 0x66, 0x75,
	0x12, 0x3d, 0x0a, 0x0e, 0x53, 0x66, 0x75, 0x57, 0x65, 0x62, 0x52, 0x74, 0x63, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x12, 0x15, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x66, 0x75, 0x2e,
	0x57, 0x65, 0x62, 0x52, 0x74, 0x63, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x66, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sfu_sfu_proto_rawDescOnce sync.Once
	file_sfu_sfu_proto_rawDescData = file_sfu_sfu_proto_rawDesc
)

func file_sfu_sfu_proto_rawDescGZIP() []byte {
	file_sfu_sfu_proto_rawDescOnce.Do(func() {
		file_sfu_sfu_proto_rawDescData = protoimpl.X.CompressGZIP(file_sfu_sfu_proto_rawDescData)
	})
	return file_sfu_sfu_proto_rawDescData
}

var file_sfu_sfu_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sfu_sfu_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil), // 0: sfu.SubscribeRequest
	(*MeetKey)(nil),          // 1: sfu.MeetKey
	(*WebRtcRelay)(nil),      // 2: sfu.WebRtcRelay
}
var file_sfu_sfu_proto_depIdxs = []int32{
	1, // 0: sfu.WebRtcRelay.meetKey:type_name -> sfu.MeetKey
	0, // 1: sfu.Sfu.SfuWebRtcRelay:input_type -> sfu.SubscribeRequest
	2, // 2: sfu.Sfu.SfuWebRtcRelay:output_type -> sfu.WebRtcRelay
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sfu_sfu_proto_init() }
func file_sfu_sfu_proto_init() {
	if File_sfu_sfu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sfu_sfu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_sfu_sfu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeetKey); i {
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
		file_sfu_sfu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebRtcRelay); i {
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
	file_sfu_sfu_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sfu_sfu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sfu_sfu_proto_goTypes,
		DependencyIndexes: file_sfu_sfu_proto_depIdxs,
		MessageInfos:      file_sfu_sfu_proto_msgTypes,
	}.Build()
	File_sfu_sfu_proto = out.File
	file_sfu_sfu_proto_rawDesc = nil
	file_sfu_sfu_proto_goTypes = nil
	file_sfu_sfu_proto_depIdxs = nil
}
