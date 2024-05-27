// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: tron.proto

package chain

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

type ContractParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data            []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	OwnerAddress    []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ContractAddress []byte `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (x *ContractParameter) Reset() {
	*x = ContractParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tron_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractParameter) ProtoMessage() {}

func (x *ContractParameter) ProtoReflect() protoreflect.Message {
	mi := &file_tron_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractParameter.ProtoReflect.Descriptor instead.
func (*ContractParameter) Descriptor() ([]byte, []int) {
	return file_tron_proto_rawDescGZIP(), []int{0}
}

func (x *ContractParameter) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ContractParameter) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ContractParameter) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

var File_tron_proto protoreflect.FileDescriptor

var file_tron_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x72,
	0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70,
	0x75, 0x73, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x75, 0x73, 0x68, 0x65,
	0x72, 0x2f, 0x74, 0x72, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tron_proto_rawDescOnce sync.Once
	file_tron_proto_rawDescData = file_tron_proto_rawDesc
)

func file_tron_proto_rawDescGZIP() []byte {
	file_tron_proto_rawDescOnce.Do(func() {
		file_tron_proto_rawDescData = protoimpl.X.CompressGZIP(file_tron_proto_rawDescData)
	})
	return file_tron_proto_rawDescData
}

var file_tron_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tron_proto_goTypes = []interface{}{
	(*ContractParameter)(nil), // 0: tron.ContractParameter
}
var file_tron_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tron_proto_init() }
func file_tron_proto_init() {
	if File_tron_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tron_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractParameter); i {
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
			RawDescriptor: file_tron_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tron_proto_goTypes,
		DependencyIndexes: file_tron_proto_depIdxs,
		MessageInfos:      file_tron_proto_msgTypes,
	}.Build()
	File_tron_proto = out.File
	file_tron_proto_rawDesc = nil
	file_tron_proto_goTypes = nil
	file_tron_proto_depIdxs = nil
}
