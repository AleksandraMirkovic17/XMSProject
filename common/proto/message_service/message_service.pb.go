// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: message_service.proto

package dislinked

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMultipleMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMessages []*UserMessage `protobuf:"bytes,1,rep,name=userMessages,proto3" json:"userMessages,omitempty"`
}

func (x *GetMultipleMessages) Reset() {
	*x = GetMultipleMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMultipleMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMultipleMessages) ProtoMessage() {}

func (x *GetMultipleMessages) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMultipleMessages.ProtoReflect.Descriptor instead.
func (*GetMultipleMessages) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetMultipleMessages) GetUserMessages() []*UserMessage {
	if x != nil {
		return x.UserMessages
	}
	return nil
}

type NewUserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMessage *UserMessage `protobuf:"bytes,1,opt,name=userMessage,proto3" json:"userMessage,omitempty"`
}

func (x *NewUserMessage) Reset() {
	*x = NewUserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUserMessage) ProtoMessage() {}

func (x *NewUserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUserMessage.ProtoReflect.Descriptor instead.
func (*NewUserMessage) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewUserMessage) GetUserMessage() *UserMessage {
	if x != nil {
		return x.UserMessage
	}
	return nil
}

type GetMessagesByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetMessagesByUserRequest) Reset() {
	*x = GetMessagesByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesByUserRequest) ProtoMessage() {}

func (x *GetMessagesByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesByUserRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesByUserRequest) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessagesByUserRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromUser string                 `protobuf:"bytes,2,opt,name=fromUser,proto3" json:"fromUser,omitempty"`
	ToUser   string                 `protobuf:"bytes,3,opt,name=toUser,proto3" json:"toUser,omitempty"`
	Content  string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Date     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_message_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserMessage) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *UserMessage) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *UserMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UserMessage) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_message_service_proto protoreflect.FileDescriptor

var file_message_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x30, 0x0a,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x40, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32,
	0xbc, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x19, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x12, 0x4e,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e,
	0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f,
	0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x3a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x16,
	0x5a, 0x14, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x73,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_service_proto_rawDescOnce sync.Once
	file_message_service_proto_rawDescData = file_message_service_proto_rawDesc
)

func file_message_service_proto_rawDescGZIP() []byte {
	file_message_service_proto_rawDescOnce.Do(func() {
		file_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_service_proto_rawDescData)
	})
	return file_message_service_proto_rawDescData
}

var file_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_service_proto_goTypes = []interface{}{
	(*GetMultipleMessages)(nil),      // 0: GetMultipleMessages
	(*NewUserMessage)(nil),           // 1: NewUserMessage
	(*GetMessagesByUserRequest)(nil), // 2: GetMessagesByUserRequest
	(*UserMessage)(nil),              // 3: UserMessage
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_message_service_proto_depIdxs = []int32{
	3, // 0: GetMultipleMessages.userMessages:type_name -> UserMessage
	3, // 1: NewUserMessage.userMessage:type_name -> UserMessage
	4, // 2: UserMessage.date:type_name -> google.protobuf.Timestamp
	2, // 3: MessageService.GetByUser:input_type -> GetMessagesByUserRequest
	1, // 4: MessageService.SendMessage:input_type -> NewUserMessage
	0, // 5: MessageService.GetByUser:output_type -> GetMultipleMessages
	1, // 6: MessageService.SendMessage:output_type -> NewUserMessage
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_service_proto_init() }
func file_message_service_proto_init() {
	if File_message_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMultipleMessages); i {
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
		file_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUserMessage); i {
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
		file_message_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesByUserRequest); i {
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
		file_message_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessage); i {
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
			RawDescriptor: file_message_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_service_proto_goTypes,
		DependencyIndexes: file_message_service_proto_depIdxs,
		MessageInfos:      file_message_service_proto_msgTypes,
	}.Build()
	File_message_service_proto = out.File
	file_message_service_proto_rawDesc = nil
	file_message_service_proto_goTypes = nil
	file_message_service_proto_depIdxs = nil
}
