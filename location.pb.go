// Code generated by protoc-gen-go. DO NOT EDIT.
// source: location.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Location struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               uint64               `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Lat                  float32              `protobuf:"fixed32,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Long                 float32              `protobuf:"fixed32,4,opt,name=long,proto3" json:"long,omitempty"`
	Hdrop                uint32               `protobuf:"varint,5,opt,name=hdrop,proto3" json:"hdrop,omitempty"`
	Altitude             float32              `protobuf:"fixed32,6,opt,name=altitude,proto3" json:"altitude,omitempty"`
	Speed                uint32               `protobuf:"varint,7,opt,name=speed,proto3" json:"speed,omitempty"`
	Message              string               `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0f35158dcf9f2c, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Location) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Location) GetLong() float32 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Location) GetHdrop() uint32 {
	if m != nil {
		return m.Hdrop
	}
	return 0
}

func (m *Location) GetAltitude() float32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSpeed() uint32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Location) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Location) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "main.Location")
}

func init() { proto.RegisterFile("location.proto", fileDescriptor_4f0f35158dcf9f2c) }

var fileDescriptor_4f0f35158dcf9f2c = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x49, 0xb7, 0xdb, 0x3f, 0x23, 0x2e, 0x32, 0x08, 0x0e, 0xbd, 0x18, 0x3c, 0xe5, 0xd4,
	0x05, 0xbd, 0xf8, 0x15, 0x04, 0x4f, 0xc1, 0xbb, 0x64, 0x4d, 0xac, 0x81, 0xb4, 0x29, 0x4d, 0xfa,
	0xd1, 0xbd, 0x4b, 0x53, 0xdb, 0xbd, 0xcd, 0xef, 0xcd, 0xef, 0x1d, 0x1e, 0x9c, 0x9c, 0xff, 0x52,
	0xd1, 0xfa, 0xa1, 0x1d, 0x27, 0x1f, 0x3d, 0xe6, 0xbd, 0xb2, 0x43, 0xf3, 0xd8, 0x79, 0xdf, 0x39,
	0x73, 0x4e, 0xd9, 0x65, 0xfe, 0x3e, 0x47, 0xdb, 0x9b, 0x10, 0x55, 0x3f, 0xae, 0xda, 0xd3, 0x2f,
	0x83, 0xea, 0xfd, 0xbf, 0x89, 0x27, 0xc8, 0xac, 0x26, 0xc6, 0x99, 0xc8, 0x65, 0x66, 0x35, 0x3e,
	0x40, 0x39, 0x07, 0x33, 0x7d, 0x5a, 0x4d, 0x59, 0x0a, 0x8b, 0x05, 0xdf, 0x34, 0xde, 0xc1, 0xc1,
	0xa9, 0x48, 0x07, 0xce, 0x44, 0x26, 0x97, 0x13, 0x11, 0x72, 0xe7, 0x87, 0x8e, 0xf2, 0x14, 0xa5,
	0x1b, 0xef, 0xe1, 0xf8, 0xa3, 0x27, 0x3f, 0xd2, 0x91, 0x33, 0x71, 0x2b, 0x57, 0xc0, 0x06, 0x2a,
	0xe5, 0xa2, 0x8d, 0xb3, 0x36, 0x54, 0x24, 0x7b, 0xe7, 0xa5, 0x11, 0x46, 0x63, 0x34, 0x95, 0x6b,
	0x23, 0x01, 0x12, 0x94, 0xbd, 0x09, 0x41, 0x75, 0x86, 0x2a, 0xce, 0x44, 0x2d, 0x37, 0xc4, 0x57,
	0xa8, 0xf7, 0x41, 0x54, 0x73, 0x26, 0x6e, 0x9e, 0x9b, 0x76, 0x9d, 0xdc, 0x6e, 0x93, 0xdb, 0x8f,
	0xcd, 0x90, 0x57, 0xf9, 0x52, 0xa4, 0xf7, 0xcb, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x96,
	0x2d, 0x02, 0x37, 0x01, 0x00, 0x00,
}