// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/GachaData.proto

package pb

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

type GachaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units []*GachaDataUnit `protobuf:"bytes,1,rep,name=Units,proto3" json:"Units,omitempty"`
}

func (x *GachaData) Reset() {
	*x = GachaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GachaData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaData) ProtoMessage() {}

func (x *GachaData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GachaData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaData.ProtoReflect.Descriptor instead.
func (*GachaData) Descriptor() ([]byte, []int) {
	return file_proto_GachaData_proto_rawDescGZIP(), []int{0}
}

func (x *GachaData) GetUnits() []*GachaDataUnit {
	if x != nil {
		return x.Units
	}
	return nil
}

type GachaDataUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type         int64               `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	StartTime    int64               `protobuf:"varint,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime      int64               `protobuf:"varint,8,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Name         *LanguageStringData `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	GunUpItem    string              `protobuf:"bytes,26,opt,name=gunUpItem,proto3" json:"gunUpItem,omitempty"`
	WeaponUpItem string              `protobuf:"bytes,28,opt,name=weaponUpItem,proto3" json:"weaponUpItem,omitempty"`
}

func (x *GachaDataUnit) Reset() {
	*x = GachaDataUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_GachaData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaDataUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaDataUnit) ProtoMessage() {}

func (x *GachaDataUnit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_GachaData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaDataUnit.ProtoReflect.Descriptor instead.
func (*GachaDataUnit) Descriptor() ([]byte, []int) {
	return file_proto_GachaData_proto_rawDescGZIP(), []int{1}
}

func (x *GachaDataUnit) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GachaDataUnit) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GachaDataUnit) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GachaDataUnit) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GachaDataUnit) GetName() *LanguageStringData {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GachaDataUnit) GetGunUpItem() string {
	if x != nil {
		return x.GunUpItem
	}
	return ""
}

func (x *GachaDataUnit) GetWeaponUpItem() string {
	if x != nil {
		return x.WeaponUpItem
	}
	return ""
}

var File_proto_GachaData_proto protoreflect.FileDescriptor

var file_proto_GachaData_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x09, 0x47, 0x61, 0x63, 0x68, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x47, 0x61, 0x63, 0x68, 0x61, 0x44, 0x61, 0x74, 0x61, 0x55,
	0x6e, 0x69, 0x74, 0x52, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x47,
	0x61, 0x63, 0x68, 0x61, 0x44, 0x61, 0x74, 0x61, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x75, 0x6e, 0x55, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x75, 0x6e, 0x55, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x22, 0x0a, 0x0c, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x70, 0x49,
	0x74, 0x65, 0x6d, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_GachaData_proto_rawDescOnce sync.Once
	file_proto_GachaData_proto_rawDescData = file_proto_GachaData_proto_rawDesc
)

func file_proto_GachaData_proto_rawDescGZIP() []byte {
	file_proto_GachaData_proto_rawDescOnce.Do(func() {
		file_proto_GachaData_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_GachaData_proto_rawDescData)
	})
	return file_proto_GachaData_proto_rawDescData
}

var file_proto_GachaData_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_GachaData_proto_goTypes = []interface{}{
	(*GachaData)(nil),          // 0: GachaData
	(*GachaDataUnit)(nil),      // 1: GachaDataUnit
	(*LanguageStringData)(nil), // 2: LanguageStringData
}
var file_proto_GachaData_proto_depIdxs = []int32{
	1, // 0: GachaData.Units:type_name -> GachaDataUnit
	2, // 1: GachaDataUnit.name:type_name -> LanguageStringData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_GachaData_proto_init() }
func file_proto_GachaData_proto_init() {
	if File_proto_GachaData_proto != nil {
		return
	}
	file_proto_LanguageStringData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_GachaData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaData); i {
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
		file_proto_GachaData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaDataUnit); i {
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
			RawDescriptor: file_proto_GachaData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_GachaData_proto_goTypes,
		DependencyIndexes: file_proto_GachaData_proto_depIdxs,
		MessageInfos:      file_proto_GachaData_proto_msgTypes,
	}.Build()
	File_proto_GachaData_proto = out.File
	file_proto_GachaData_proto_rawDesc = nil
	file_proto_GachaData_proto_goTypes = nil
	file_proto_GachaData_proto_depIdxs = nil
}
