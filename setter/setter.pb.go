// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: setter.proto

package setter

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

var E_AllMessages = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         90010,
	Name:          "setter.all_messages",
	Tag:           "varint,90010,opt,name=all_messages",
	Filename:      "setter.proto",
}

var E_AllFields = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         90020,
	Name:          "setter.all_fields",
	Tag:           "varint,90020,opt,name=all_fields",
	Filename:      "setter.proto",
}

var E_Include = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         90030,
	Name:          "setter.include",
	Tag:           "varint,90030,opt,name=include",
	Filename:      "setter.proto",
}

var E_Exclude = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         90031,
	Name:          "setter.exclude",
	Tag:           "varint,90031,opt,name=exclude",
	Filename:      "setter.proto",
}

func init() {
	proto.RegisterExtension(E_AllMessages)
	proto.RegisterExtension(E_AllFields)
	proto.RegisterExtension(E_Include)
	proto.RegisterExtension(E_Exclude)
}

func init() { proto.RegisterFile("setter.proto", fileDescriptor_f225ba13d7bdf8ad) }

var fileDescriptor_f225ba13d7bdf8ad = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x29,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x14, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xa2, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45,
	0x99, 0x05, 0x25, 0xf9, 0x50, 0x95, 0x56, 0x8e, 0x5c, 0x3c, 0x89, 0x39, 0x39, 0xf1, 0xb9, 0xa9,
	0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0xc5, 0x42, 0x32, 0x7a, 0x10, 0x2d, 0x7a, 0x30, 0x2d, 0x7a, 0x6e,
	0x99, 0x39, 0xa9, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x12, 0xb3, 0xf6, 0xb3, 0x2a, 0x30,
	0x6a, 0x70, 0x04, 0x71, 0x27, 0xe6, 0xe4, 0xf8, 0x42, 0xb5, 0x58, 0x39, 0x70, 0x71, 0x81, 0x8c,
	0x48, 0xcb, 0x4c, 0xcd, 0x49, 0x29, 0x16, 0x92, 0xc7, 0x30, 0x00, 0xaa, 0x10, 0x66, 0xc6, 0x12,
	0xa8, 0x19, 0x9c, 0x89, 0x39, 0x39, 0x6e, 0x60, 0x3d, 0x56, 0x96, 0x5c, 0xec, 0x99, 0x79, 0xc9,
	0x39, 0xa5, 0x29, 0xa9, 0x42, 0xb2, 0x58, 0xec, 0x4f, 0xcd, 0x49, 0x81, 0x69, 0x5e, 0x07, 0xd5,
	0x0c, 0x53, 0x0f, 0xd2, 0x9a, 0x5a, 0x41, 0x94, 0xd6, 0xf5, 0x30, 0xad, 0x50, 0xf5, 0x49, 0x6c,
	0x60, 0x75, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x38, 0xf0, 0x32, 0x3b, 0x01, 0x00,
	0x00,
}
