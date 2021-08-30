// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.4.0
// source: Exercise.proto

package Exercise

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

type Empty1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty1) Reset() {
	*x = Empty1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Exercise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty1) ProtoMessage() {}

func (x *Empty1) ProtoReflect() protoreflect.Message {
	mi := &file_Exercise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty1.ProtoReflect.Descriptor instead.
func (*Empty1) Descriptor() ([]byte, []int) {
	return file_Exercise_proto_rawDescGZIP(), []int{0}
}

type I struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I int32 `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *I) Reset() {
	*x = I{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Exercise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I) ProtoMessage() {}

func (x *I) ProtoReflect() protoreflect.Message {
	mi := &file_Exercise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I.ProtoReflect.Descriptor instead.
func (*I) Descriptor() ([]byte, []int) {
	return file_Exercise_proto_rawDescGZIP(), []int{1}
}

func (x *I) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

type ExerciseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                  //数据自身ID
	Classid    int32  `protobuf:"varint,2,opt,name=classid,proto3" json:"classid,omitempty"`        //所属班级
	Ownerid    int32  `protobuf:"varint,3,opt,name=ownerid,proto3" json:"ownerid,omitempty"`        //发布人
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`         //内容
	Typ        uint32 `protobuf:"varint,5,opt,name=typ,proto3" json:"typ,omitempty"`                //类型起始日期与截止日期、无时间限制、单次限时
	Begin      int64  `protobuf:"varint,6,opt,name=begin,proto3" json:"begin,omitempty"`            //起始日期
	End        int64  `protobuf:"varint,7,opt,name=end,proto3" json:"end,omitempty"`                //截止日期
	Duration   int64  `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`      //持续时长
	Name       string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`               //考试名
	Discipline uint32 `protobuf:"varint,10,opt,name=Discipline,proto3" json:"Discipline,omitempty"` //科目
}

func (x *ExerciseData) Reset() {
	*x = ExerciseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Exercise_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExerciseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseData) ProtoMessage() {}

func (x *ExerciseData) ProtoReflect() protoreflect.Message {
	mi := &file_Exercise_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseData.ProtoReflect.Descriptor instead.
func (*ExerciseData) Descriptor() ([]byte, []int) {
	return file_Exercise_proto_rawDescGZIP(), []int{2}
}

func (x *ExerciseData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExerciseData) GetClassid() int32 {
	if x != nil {
		return x.Classid
	}
	return 0
}

func (x *ExerciseData) GetOwnerid() int32 {
	if x != nil {
		return x.Ownerid
	}
	return 0
}

func (x *ExerciseData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ExerciseData) GetTyp() uint32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

func (x *ExerciseData) GetBegin() int64 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *ExerciseData) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *ExerciseData) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *ExerciseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExerciseData) GetDiscipline() uint32 {
	if x != nil {
		return x.Discipline
	}
	return 0
}

type Submit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 //数据本身的ID-固定！！！
	Uploaderid int32  `protobuf:"varint,2,opt,name=uploaderid,proto3" json:"uploaderid,omitempty"` //上传者ID
	Exerciseid int32  `protobuf:"varint,3,opt,name=exerciseid,proto3" json:"exerciseid,omitempty"` //考试ID
	Contents   string `protobuf:"bytes,4,opt,name=contents,proto3" json:"contents,omitempty"`      //提交内容
	Value      int32  `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`           //得分
}

func (x *Submit) Reset() {
	*x = Submit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Exercise_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Submit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Submit) ProtoMessage() {}

func (x *Submit) ProtoReflect() protoreflect.Message {
	mi := &file_Exercise_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Submit.ProtoReflect.Descriptor instead.
func (*Submit) Descriptor() ([]byte, []int) {
	return file_Exercise_proto_rawDescGZIP(), []int{3}
}

func (x *Submit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Submit) GetUploaderid() int32 {
	if x != nil {
		return x.Uploaderid
	}
	return 0
}

func (x *Submit) GetExerciseid() int32 {
	if x != nil {
		return x.Exerciseid
	}
	return 0
}

func (x *Submit) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

func (x *Submit) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submitid int32 `protobuf:"varint,1,opt,name=submitid,proto3" json:"submitid,omitempty"` //提交记录
	Judge    int32 `protobuf:"varint,2,opt,name=judge,proto3" json:"judge,omitempty"`       //打分人
	Value    int32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`       //分值
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Exercise_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_Exercise_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_Exercise_proto_rawDescGZIP(), []int{4}
}

func (x *Score) GetSubmitid() int32 {
	if x != nil {
		return x.Submitid
	}
	return 0
}

func (x *Score) GetJudge() int32 {
	if x != nil {
		return x.Judge
	}
	return 0
}

func (x *Score) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_Exercise_proto protoreflect.FileDescriptor

var file_Exercise_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x31, 0x22, 0x11, 0x0a, 0x01, 0x69, 0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x69, 0x22, 0xf6, 0x01, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0x8a, 0x01, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4f, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb0,
	0x05, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x67,
	0x65, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x0b, 0x2e, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x63, 0x12, 0x0b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69,
	0x1a, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0d, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x5f, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x5f, 0x61, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x1a, 0x0b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x2e, 0x69, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x12, 0x0b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x10, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x0f,
	0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x10, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x31, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x0b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x0f, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x0b,
	0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x10, 0x2e, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x34, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x2e, 0x69, 0x1a, 0x0f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x0b, 0x2e, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2f,
	0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x0b,
	0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x10, 0x2e, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x31, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x5f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73,
	0x12, 0x0b, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x69, 0x1a, 0x10, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x31, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Exercise_proto_rawDescOnce sync.Once
	file_Exercise_proto_rawDescData = file_Exercise_proto_rawDesc
)

func file_Exercise_proto_rawDescGZIP() []byte {
	file_Exercise_proto_rawDescOnce.Do(func() {
		file_Exercise_proto_rawDescData = protoimpl.X.CompressGZIP(file_Exercise_proto_rawDescData)
	})
	return file_Exercise_proto_rawDescData
}

var file_Exercise_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Exercise_proto_goTypes = []interface{}{
	(*Empty1)(nil),       // 0: exercise.empty1
	(*I)(nil),            // 1: exercise.i
	(*ExerciseData)(nil), // 2: exercise.exerciseData
	(*Submit)(nil),       // 3: exercise.submit
	(*Score)(nil),        // 4: exercise.score
}
var file_Exercise_proto_depIdxs = []int32{
	1,  // 0: exercise.exercise.get_exercise:input_type -> exercise.i
	1,  // 1: exercise.exercise.get_exercisec:input_type -> exercise.i
	1,  // 2: exercise.exercise.get_exercises:input_type -> exercise.i
	2,  // 3: exercise.exercise.add_exercise:input_type -> exercise.exerciseData
	3,  // 4: exercise.exercise.submit_ans:input_type -> exercise.submit
	1,  // 5: exercise.exercise.get_key:input_type -> exercise.i
	4,  // 6: exercise.exercise.set_score:input_type -> exercise.score
	1,  // 7: exercise.exercise.get_score:input_type -> exercise.i
	1,  // 8: exercise.exercise.get_scores:input_type -> exercise.i
	1,  // 9: exercise.exercise.get_class_scores:input_type -> exercise.i
	1,  // 10: exercise.exercise.get_class_submit:input_type -> exercise.i
	1,  // 11: exercise.exercise.del_exercise:input_type -> exercise.i
	1,  // 12: exercise.exercise.del_exercises:input_type -> exercise.i
	2,  // 13: exercise.exercise.get_exercise:output_type -> exercise.exerciseData
	2,  // 14: exercise.exercise.get_exercisec:output_type -> exercise.exerciseData
	2,  // 15: exercise.exercise.get_exercises:output_type -> exercise.exerciseData
	2,  // 16: exercise.exercise.add_exercise:output_type -> exercise.exerciseData
	1,  // 17: exercise.exercise.submit_ans:output_type -> exercise.i
	3,  // 18: exercise.exercise.get_key:output_type -> exercise.submit
	0,  // 19: exercise.exercise.set_score:output_type -> exercise.empty1
	4,  // 20: exercise.exercise.get_score:output_type -> exercise.score
	3,  // 21: exercise.exercise.get_scores:output_type -> exercise.submit
	4,  // 22: exercise.exercise.get_class_scores:output_type -> exercise.score
	3,  // 23: exercise.exercise.get_class_submit:output_type -> exercise.submit
	0,  // 24: exercise.exercise.del_exercise:output_type -> exercise.empty1
	0,  // 25: exercise.exercise.del_exercises:output_type -> exercise.empty1
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_Exercise_proto_init() }
func file_Exercise_proto_init() {
	if File_Exercise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Exercise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty1); i {
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
		file_Exercise_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I); i {
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
		file_Exercise_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExerciseData); i {
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
		file_Exercise_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Submit); i {
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
		file_Exercise_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
			RawDescriptor: file_Exercise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Exercise_proto_goTypes,
		DependencyIndexes: file_Exercise_proto_depIdxs,
		MessageInfos:      file_Exercise_proto_msgTypes,
	}.Build()
	File_Exercise_proto = out.File
	file_Exercise_proto_rawDesc = nil
	file_Exercise_proto_goTypes = nil
	file_Exercise_proto_depIdxs = nil
}