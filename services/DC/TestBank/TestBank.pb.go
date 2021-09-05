// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.4.0
// source: TestBank.proto

package TestBank

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

type Ans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ans string `protobuf:"bytes,1,opt,name=ans,proto3" json:"ans,omitempty"`
}

func (x *Ans) Reset() {
	*x = Ans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestBank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ans) ProtoMessage() {}

func (x *Ans) ProtoReflect() protoreflect.Message {
	mi := &file_TestBank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ans.ProtoReflect.Descriptor instead.
func (*Ans) Descriptor() ([]byte, []int) {
	return file_TestBank_proto_rawDescGZIP(), []int{0}
}

func (x *Ans) GetAns() string {
	if x != nil {
		return x.Ans
	}
	return ""
}

type Testid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Testid) Reset() {
	*x = Testid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestBank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testid) ProtoMessage() {}

func (x *Testid) ProtoReflect() protoreflect.Message {
	mi := &file_TestBank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Testid.ProtoReflect.Descriptor instead.
func (*Testid) Descriptor() ([]byte, []int) {
	return file_TestBank_proto_rawDescGZIP(), []int{1}
}

func (x *Testid) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Testconf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectiveItem uint32 `protobuf:"varint,1,opt,name=subjective_item,json=subjectiveItem,proto3" json:"subjective_item,omitempty"` //主观题的数量
	ObjectiveItem  uint32 `protobuf:"varint,2,opt,name=objective_item,json=objectiveItem,proto3" json:"objective_item,omitempty"`    //客观题的数量
	Discipline     uint32 `protobuf:"varint,3,opt,name=discipline,proto3" json:"discipline,omitempty"`                               //学科
}

func (x *Testconf) Reset() {
	*x = Testconf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestBank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testconf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testconf) ProtoMessage() {}

func (x *Testconf) ProtoReflect() protoreflect.Message {
	mi := &file_TestBank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Testconf.ProtoReflect.Descriptor instead.
func (*Testconf) Descriptor() ([]byte, []int) {
	return file_TestBank_proto_rawDescGZIP(), []int{2}
}

func (x *Testconf) GetSubjectiveItem() uint32 {
	if x != nil {
		return x.SubjectiveItem
	}
	return 0
}

func (x *Testconf) GetObjectiveItem() uint32 {
	if x != nil {
		return x.ObjectiveItem
	}
	return 0
}

func (x *Testconf) GetDiscipline() uint32 {
	if x != nil {
		return x.Discipline
	}
	return 0
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 //自身ID
	Typ        uint32 `protobuf:"varint,2,opt,name=typ,proto3" json:"typ,omitempty"`               //类型-主观题/客观题
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`        //内容
	Ans        string `protobuf:"bytes,4,opt,name=ans,proto3" json:"ans,omitempty"`                //答案(如果有)
	Name       string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`              //名字-题目描述
	Uploader   uint32 `protobuf:"varint,6,opt,name=uploader,proto3" json:"uploader,omitempty"`     //上传者
	Discipline uint32 `protobuf:"varint,7,opt,name=discipline,proto3" json:"discipline,omitempty"` //学科
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestBank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_TestBank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_TestBank_proto_rawDescGZIP(), []int{3}
}

func (x *Test) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Test) GetTyp() uint32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

func (x *Test) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Test) GetAns() string {
	if x != nil {
		return x.Ans
	}
	return ""
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Test) GetUploader() uint32 {
	if x != nil {
		return x.Uploader
	}
	return 0
}

func (x *Test) GetDiscipline() uint32 {
	if x != nil {
		return x.Discipline
	}
	return 0
}

var File_TestBank_proto protoreflect.FileDescriptor

var file_TestBank_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x22, 0x17, 0x0a, 0x03, 0x61, 0x6e,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x6e, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x69, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7a, 0x0a,
	0x08, 0x74, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73,
	0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64,
	0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x74, 0x79, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65,
	0x32, 0xa5, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x2a, 0x0a,
	0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x6e, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x6e, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61,
	0x6e, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a,
	0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x6f,
	0x6e, 0x66, 0x1a, 0x10, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x64, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x54, 0x65,
	0x73, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TestBank_proto_rawDescOnce sync.Once
	file_TestBank_proto_rawDescData = file_TestBank_proto_rawDesc
)

func file_TestBank_proto_rawDescGZIP() []byte {
	file_TestBank_proto_rawDescOnce.Do(func() {
		file_TestBank_proto_rawDescData = protoimpl.X.CompressGZIP(file_TestBank_proto_rawDescData)
	})
	return file_TestBank_proto_rawDescData
}

var file_TestBank_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_TestBank_proto_goTypes = []interface{}{
	(*Ans)(nil),      // 0: TestBank.ans
	(*Testid)(nil),   // 1: TestBank.testid
	(*Testconf)(nil), // 2: TestBank.testconf
	(*Test)(nil),     // 3: TestBank.test
}
var file_TestBank_proto_depIdxs = []int32{
	3, // 0: TestBank.TestBank.upload:input_type -> TestBank.test
	1, // 1: TestBank.TestBank.download:input_type -> TestBank.testid
	2, // 2: TestBank.TestBank.generate_test:input_type -> TestBank.testconf
	3, // 3: TestBank.TestBank.upload:output_type -> TestBank.test
	3, // 4: TestBank.TestBank.download:output_type -> TestBank.test
	1, // 5: TestBank.TestBank.generate_test:output_type -> TestBank.testid
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TestBank_proto_init() }
func file_TestBank_proto_init() {
	if File_TestBank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TestBank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ans); i {
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
		file_TestBank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Testid); i {
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
		file_TestBank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Testconf); i {
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
		file_TestBank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_TestBank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TestBank_proto_goTypes,
		DependencyIndexes: file_TestBank_proto_depIdxs,
		MessageInfos:      file_TestBank_proto_msgTypes,
	}.Build()
	File_TestBank_proto = out.File
	file_TestBank_proto_rawDesc = nil
	file_TestBank_proto_goTypes = nil
	file_TestBank_proto_depIdxs = nil
}
