// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/opiapi.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetNetworkRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkRequest) Reset()         { *m = GetNetworkRequest{} }
func (m *GetNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkRequest) ProtoMessage()    {}
func (*GetNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aac41881a81f4b0, []int{0}
}

func (m *GetNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkRequest.Unmarshal(m, b)
}
func (m *GetNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkRequest.Marshal(b, m, deterministic)
}
func (m *GetNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkRequest.Merge(m, src)
}
func (m *GetNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_GetNetworkRequest.Size(m)
}
func (m *GetNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkRequest proto.InternalMessageInfo

func (m *GetNetworkRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetNetworkResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkResponse) Reset()         { *m = GetNetworkResponse{} }
func (m *GetNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*GetNetworkResponse) ProtoMessage()    {}
func (*GetNetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aac41881a81f4b0, []int{1}
}

func (m *GetNetworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkResponse.Unmarshal(m, b)
}
func (m *GetNetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkResponse.Marshal(b, m, deterministic)
}
func (m *GetNetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkResponse.Merge(m, src)
}
func (m *GetNetworkResponse) XXX_Size() int {
	return xxx_messageInfo_GetNetworkResponse.Size(m)
}
func (m *GetNetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkResponse proto.InternalMessageInfo

func (m *GetNetworkResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CreateNetworkRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNetworkRequest) Reset()         { *m = CreateNetworkRequest{} }
func (m *CreateNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNetworkRequest) ProtoMessage()    {}
func (*CreateNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aac41881a81f4b0, []int{2}
}

func (m *CreateNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNetworkRequest.Unmarshal(m, b)
}
func (m *CreateNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNetworkRequest.Marshal(b, m, deterministic)
}
func (m *CreateNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNetworkRequest.Merge(m, src)
}
func (m *CreateNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNetworkRequest.Size(m)
}
func (m *CreateNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNetworkRequest proto.InternalMessageInfo

func (m *CreateNetworkRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateNetworkRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6aac41881a81f4b0, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetNetworkRequest)(nil), "proto.GetNetworkRequest")
	proto.RegisterType((*GetNetworkResponse)(nil), "proto.GetNetworkResponse")
	proto.RegisterType((*CreateNetworkRequest)(nil), "proto.CreateNetworkRequest")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
}

func init() {
	proto.RegisterFile("proto/opiapi.proto", fileDescriptor_6aac41881a81f4b0)
}

var fileDescriptor_6aac41881a81f4b0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x2f, 0xc8, 0x4c, 0x2c, 0xc8, 0xd4, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x92,
	0x2a, 0x97, 0xa0, 0x7b, 0x6a, 0x89, 0x5f, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0x76, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x88, 0xa9, 0xa4, 0xc5, 0x25, 0x84, 0xac, 0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0xac, 0x92, 0x27, 0x08, 0xc2, 0x51, 0xb2,
	0xe3, 0x12, 0x71, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x25, 0x64, 0x2a, 0x42, 0x3f, 0x13, 0xb2, 0x7e,
	0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0xa3, 0x76, 0x46, 0x2e, 0x36, 0xff, 0x00, 0x4f,
	0xc7, 0x00, 0x4f, 0x21, 0x47, 0x2e, 0x2e, 0x84, 0xfd, 0x42, 0x12, 0x10, 0x3f, 0xe8, 0x61, 0xb8,
	0x5c, 0x4a, 0x12, 0x8b, 0x0c, 0xd4, 0xb1, 0x56, 0x5c, 0xbc, 0x28, 0xce, 0x12, 0x92, 0x86, 0xaa,
	0xc5, 0xe6, 0x58, 0x29, 0x1e, 0xa8, 0x24, 0xd8, 0x25, 0x4e, 0x9c, 0x51, 0xec, 0x7a, 0xfa, 0x60,
	0x81, 0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x8f, 0x65, 0x22, 0x54,
	0x01, 0x00, 0x00,
}
