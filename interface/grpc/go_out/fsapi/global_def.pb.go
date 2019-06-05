// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global_def.proto

package fsapi

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

type ErrCode int32

const (
	ErrCode_ERR_SUCCESS              ErrCode = 0
	ErrCode_ERR_FAILED               ErrCode = 1
	ErrCode_ERR_DEVID_NOT_EXIST      ErrCode = 2
	ErrCode_ERR_CONN_TO_DA_NOT_READY ErrCode = 3
)

var ErrCode_name = map[int32]string{
	0: "ERR_SUCCESS",
	1: "ERR_FAILED",
	2: "ERR_DEVID_NOT_EXIST",
	3: "ERR_CONN_TO_DA_NOT_READY",
}

var ErrCode_value = map[string]int32{
	"ERR_SUCCESS":              0,
	"ERR_FAILED":               1,
	"ERR_DEVID_NOT_EXIST":      2,
	"ERR_CONN_TO_DA_NOT_READY": 3,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{0}
}

type ReqHeader struct {
	// 现在版本默认为 1
	Version int32 `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	// 设备ID
	DevId string `protobuf:"bytes,2,opt,name=DevId,proto3" json:"DevId,omitempty"`
	//时间戳 unix_timestamp
	Ts                   uint32   `protobuf:"varint,3,opt,name=Ts,proto3" json:"Ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHeader) Reset()         { *m = ReqHeader{} }
func (m *ReqHeader) String() string { return proto.CompactTextString(m) }
func (*ReqHeader) ProtoMessage()    {}
func (*ReqHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{0}
}

func (m *ReqHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHeader.Unmarshal(m, b)
}
func (m *ReqHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHeader.Marshal(b, m, deterministic)
}
func (m *ReqHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHeader.Merge(m, src)
}
func (m *ReqHeader) XXX_Size() int {
	return xxx_messageInfo_ReqHeader.Size(m)
}
func (m *ReqHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHeader proto.InternalMessageInfo

func (m *ReqHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ReqHeader) GetDevId() string {
	if m != nil {
		return m.DevId
	}
	return ""
}

func (m *ReqHeader) GetTs() uint32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

type ResHeader struct {
	// 现在版本默认为 1
	Version              int32    `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	ErrCode              ErrCode  `protobuf:"varint,2,opt,name=ErrCode,proto3,enum=fsapi.ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,3,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	ReqId                string   `protobuf:"bytes,4,opt,name=ReqId,proto3" json:"ReqId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResHeader) Reset()         { *m = ResHeader{} }
func (m *ResHeader) String() string { return proto.CompactTextString(m) }
func (*ResHeader) ProtoMessage()    {}
func (*ResHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a8c9fe6e60032d, []int{1}
}

func (m *ResHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResHeader.Unmarshal(m, b)
}
func (m *ResHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResHeader.Marshal(b, m, deterministic)
}
func (m *ResHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResHeader.Merge(m, src)
}
func (m *ResHeader) XXX_Size() int {
	return xxx_messageInfo_ResHeader.Size(m)
}
func (m *ResHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResHeader proto.InternalMessageInfo

func (m *ResHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ResHeader) GetErrCode() ErrCode {
	if m != nil {
		return m.ErrCode
	}
	return ErrCode_ERR_SUCCESS
}

func (m *ResHeader) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ResHeader) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func init() {
	proto.RegisterEnum("fsapi.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*ReqHeader)(nil), "fsapi.ReqHeader")
	proto.RegisterType((*ResHeader)(nil), "fsapi.ResHeader")
}

func init() { proto.RegisterFile("global_def.proto", fileDescriptor_24a8c9fe6e60032d) }

var fileDescriptor_24a8c9fe6e60032d = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4f, 0xeb, 0x30,
	0x18, 0xc4, 0x5f, 0xd2, 0xd7, 0x56, 0xf5, 0xd3, 0x0b, 0x91, 0x41, 0xe0, 0x81, 0xa1, 0xea, 0x14,
	0x21, 0xe1, 0x48, 0x30, 0x30, 0x97, 0xc4, 0x88, 0x08, 0x68, 0x25, 0x27, 0x54, 0xc0, 0x62, 0xa5,
	0xc9, 0x97, 0x10, 0xa9, 0xd4, 0xa9, 0x9d, 0xb2, 0xf1, 0xbf, 0x23, 0xbb, 0x2d, 0x2b, 0xe3, 0xef,
	0xee, 0x74, 0x3e, 0xeb, 0x43, 0x7e, 0xbd, 0x92, 0xcb, 0x7c, 0x25, 0x4a, 0xa8, 0x68, 0xab, 0x64,
	0x27, 0x71, 0xbf, 0xd2, 0x79, 0xdb, 0x4c, 0x1e, 0xd0, 0x88, 0xc3, 0xe6, 0x1e, 0xf2, 0x12, 0x14,
	0x26, 0x68, 0xb8, 0x00, 0xa5, 0x1b, 0xb9, 0x26, 0xce, 0xd8, 0x09, 0xfa, 0xfc, 0x80, 0xf8, 0x04,
	0xf5, 0x63, 0xf8, 0x4c, 0x4a, 0xe2, 0x8e, 0x9d, 0x60, 0xc4, 0x77, 0x80, 0x3d, 0xe4, 0x66, 0x9a,
	0xf4, 0xc6, 0x4e, 0xf0, 0x9f, 0xbb, 0x99, 0x9e, 0x7c, 0x99, 0x32, 0xfd, 0x6b, 0x59, 0x80, 0x86,
	0x4c, 0xa9, 0x48, 0x96, 0x60, 0xeb, 0xbc, 0x2b, 0x8f, 0xda, 0x31, 0x74, 0xaf, 0xf2, 0x83, 0x8d,
	0x4f, 0xd1, 0x80, 0x29, 0xf5, 0xa4, 0x6b, 0xfb, 0xc8, 0x88, 0xef, 0xc9, 0xcc, 0xe1, 0xb0, 0x49,
	0x4a, 0xf2, 0x77, 0x37, 0xc7, 0xc2, 0x45, 0xfe, 0xd3, 0x8b, 0x8f, 0xd0, 0x3f, 0xc6, 0xb9, 0x48,
	0x9f, 0xa3, 0x88, 0xa5, 0xa9, 0xff, 0x07, 0x7b, 0x08, 0x19, 0xe1, 0x6e, 0x9a, 0x3c, 0xb2, 0xd8,
	0x77, 0xf0, 0x19, 0x3a, 0x36, 0x1c, 0xb3, 0x45, 0x12, 0x8b, 0xd9, 0x3c, 0x13, 0xec, 0x25, 0x49,
	0x33, 0xdf, 0xc5, 0xe7, 0x88, 0x18, 0x23, 0x9a, 0xcf, 0x66, 0x22, 0x9b, 0x8b, 0x78, 0x6a, 0x5d,
	0xce, 0xa6, 0xf1, 0xab, 0xdf, 0xbb, 0x4d, 0x11, 0x29, 0xd6, 0xb4, 0x91, 0x6d, 0x53, 0xd3, 0x0a,
	0xa0, 0xa4, 0xba, 0xc8, 0x57, 0x40, 0x6b, 0xd5, 0x16, 0x6f, 0x37, 0x75, 0xd3, 0xbd, 0x6f, 0x97,
	0xb4, 0x90, 0x1f, 0xa1, 0x4d, 0x84, 0x26, 0x71, 0x69, 0x13, 0x61, 0xb3, 0xee, 0x40, 0x55, 0x79,
	0x01, 0xa1, 0xc9, 0x86, 0xb5, 0x14, 0x72, 0xdb, 0x85, 0xf6, 0xdb, 0xcb, 0x81, 0xbd, 0xc8, 0xf5,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xf9, 0xc8, 0x45, 0xa5, 0x01, 0x00, 0x00,
}