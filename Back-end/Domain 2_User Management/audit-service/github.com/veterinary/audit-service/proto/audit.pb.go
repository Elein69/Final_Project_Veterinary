// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v6.31.1
// source: proto/audit.proto

package proto

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

type AuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Action    string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AuditRequest) Reset() {
	*x = AuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditRequest) ProtoMessage() {}

func (x *AuditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditRequest.ProtoReflect.Descriptor instead.
func (*AuditRequest) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{0}
}

func (x *AuditRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuditRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AuditRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type AuditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AuditResponse) Reset() {
	*x = AuditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_audit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditResponse) ProtoMessage() {}

func (x *AuditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_audit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditResponse.ProtoReflect.Descriptor instead.
func (*AuditResponse) Descriptor() ([]byte, []int) {
	return file_proto_audit_proto_rawDescGZIP(), []int{1}
}

func (x *AuditResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_audit_proto protoreflect.FileDescriptor

var file_proto_audit_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0x5c, 0x0a, 0x0c, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x27, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x45, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_audit_proto_rawDescOnce sync.Once
	file_proto_audit_proto_rawDescData = file_proto_audit_proto_rawDesc
)

func file_proto_audit_proto_rawDescGZIP() []byte {
	file_proto_audit_proto_rawDescOnce.Do(func() {
		file_proto_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_audit_proto_rawDescData)
	})
	return file_proto_audit_proto_rawDescData
}

var file_proto_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_audit_proto_goTypes = []any{
	(*AuditRequest)(nil),  // 0: audit.AuditRequest
	(*AuditResponse)(nil), // 1: audit.AuditResponse
}
var file_proto_audit_proto_depIdxs = []int32{
	0, // 0: audit.AuditService.LogEvent:input_type -> audit.AuditRequest
	1, // 1: audit.AuditService.LogEvent:output_type -> audit.AuditResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_audit_proto_init() }
func file_proto_audit_proto_init() {
	if File_proto_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_audit_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AuditRequest); i {
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
		file_proto_audit_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AuditResponse); i {
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
			RawDescriptor: file_proto_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_audit_proto_goTypes,
		DependencyIndexes: file_proto_audit_proto_depIdxs,
		MessageInfos:      file_proto_audit_proto_msgTypes,
	}.Build()
	File_proto_audit_proto = out.File
	file_proto_audit_proto_rawDesc = nil
	file_proto_audit_proto_goTypes = nil
	file_proto_audit_proto_depIdxs = nil
}
