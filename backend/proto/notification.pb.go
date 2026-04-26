package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_UNKNOWN Status = 0
	Status_SUCCESS Status = 1
	Status_FAILED  Status = 2
)

var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILED",
	}
	Status_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILED":  2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_notification_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_notification_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_notification_proto_rawDescGZIP(), []int{0}
}

type NotificationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	mi := &file_proto_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_proto_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *NotificationRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NotificationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        Status                 `protobuf:"varint,1,opt,name=status,proto3,enum=notification.Status" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationResponse) Reset() {
	*x = NotificationResponse{}
	mi := &file_proto_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationResponse) ProtoMessage() {}

func (x *NotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return file_proto_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

var File_proto_notification_proto protoreflect.FileDescriptor

const file_proto_notification_proto_rawDesc = "" +
	"\n" +
	"\x18proto/notification.proto\x12\fnotification\"O\n" +
	"\x13NotificationRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x0e\n" +
	"\x02to\x18\x02 \x01(\tR\x02to\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"D\n" +
	"\x14NotificationResponse\x12,\n" +
	"\x06status\x18\x01 \x01(\x0e2\x14.notification.StatusR\x06status*.\n" +
	"\x06Status\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\v\n" +
	"\aSUCCESS\x10\x01\x12\n" +
	"\n" +
	"\x06FAILED\x10\x022p\n" +
	"\x13NotificationService\x12Y\n" +
	"\x10SendNotification\x12!.notification.NotificationRequest\x1a\".notification.NotificationResponseB\tZ\a./protob\x06proto3"

var (
	file_proto_notification_proto_rawDescOnce sync.Once
	file_proto_notification_proto_rawDescData []byte
)

func file_proto_notification_proto_rawDescGZIP() []byte {
	file_proto_notification_proto_rawDescOnce.Do(func() {
		file_proto_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_notification_proto_rawDesc), len(file_proto_notification_proto_rawDesc)))
	})
	return file_proto_notification_proto_rawDescData
}

var file_proto_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_notification_proto_goTypes = []any{
	(Status)(0),
	(*NotificationRequest)(nil),
	(*NotificationResponse)(nil),
}
var file_proto_notification_proto_depIdxs = []int32{
	0,
	1,
	2,
	2,
	1,
	1,
	1,
	0,
}

func init() { file_proto_notification_proto_init() }
func file_proto_notification_proto_init() {
	if File_proto_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_notification_proto_rawDesc), len(file_proto_notification_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notification_proto_goTypes,
		DependencyIndexes: file_proto_notification_proto_depIdxs,
		EnumInfos:         file_proto_notification_proto_enumTypes,
		MessageInfos:      file_proto_notification_proto_msgTypes,
	}.Build()
	File_proto_notification_proto = out.File
	file_proto_notification_proto_goTypes = nil
	file_proto_notification_proto_depIdxs = nil
}
