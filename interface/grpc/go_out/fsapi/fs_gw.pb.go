// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fs_gw.proto

package fsapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PigstyInfo struct {
	//猪圈ID
	PigstyId string `protobuf:"bytes,1,opt,name=PigstyId,proto3" json:"PigstyId,omitempty"`
	//猪数量
	PigNum uint32 `protobuf:"varint,2,opt,name=PigNum,proto3" json:"PigNum,omitempty"`
	//猪平均重量,单位为“克”
	AverageWeight uint32 `protobuf:"varint,3,opt,name=AverageWeight,proto3" json:"AverageWeight,omitempty"`
	//喂料推荐值，单位为“克”
	AdviseWeight uint32 `protobuf:"varint,4,opt,name=AdviseWeight,proto3" json:"AdviseWeight,omitempty"`
	//上次喂料时间戳 unix_timestamp
	LastFed uint64 `protobuf:"varint,5,opt,name=LastFed,proto3" json:"LastFed,omitempty"`
	//猪的id列表
	PigId []string `protobuf:"bytes,6,rep,name=PigId,proto3" json:"PigId,omitempty"`
	//猪圈名称
	StyName string `protobuf:"bytes,7,opt,name=StyName,proto3" json:"StyName,omitempty"`
	//猪的品种
	PigSpecies           string   `protobuf:"bytes,8,opt,name=PigSpecies,proto3" json:"PigSpecies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PigstyInfo) Reset()         { *m = PigstyInfo{} }
func (m *PigstyInfo) String() string { return proto.CompactTextString(m) }
func (*PigstyInfo) ProtoMessage()    {}
func (*PigstyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{0}
}

func (m *PigstyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PigstyInfo.Unmarshal(m, b)
}
func (m *PigstyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PigstyInfo.Marshal(b, m, deterministic)
}
func (m *PigstyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PigstyInfo.Merge(m, src)
}
func (m *PigstyInfo) XXX_Size() int {
	return xxx_messageInfo_PigstyInfo.Size(m)
}
func (m *PigstyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PigstyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PigstyInfo proto.InternalMessageInfo

func (m *PigstyInfo) GetPigstyId() string {
	if m != nil {
		return m.PigstyId
	}
	return ""
}

func (m *PigstyInfo) GetPigNum() uint32 {
	if m != nil {
		return m.PigNum
	}
	return 0
}

func (m *PigstyInfo) GetAverageWeight() uint32 {
	if m != nil {
		return m.AverageWeight
	}
	return 0
}

func (m *PigstyInfo) GetAdviseWeight() uint32 {
	if m != nil {
		return m.AdviseWeight
	}
	return 0
}

func (m *PigstyInfo) GetLastFed() uint64 {
	if m != nil {
		return m.LastFed
	}
	return 0
}

func (m *PigstyInfo) GetPigId() []string {
	if m != nil {
		return m.PigId
	}
	return nil
}

func (m *PigstyInfo) GetStyName() string {
	if m != nil {
		return m.StyName
	}
	return ""
}

func (m *PigstyInfo) GetPigSpecies() string {
	if m != nil {
		return m.PigSpecies
	}
	return ""
}

type PigHouseInfo struct {
	//猪舍ID
	HouseId string `protobuf:"bytes,1,opt,name=HouseId,proto3" json:"HouseId,omitempty"`
	//日龄
	AGE string `protobuf:"bytes,2,opt,name=AGE,proto3" json:"AGE,omitempty"`
	//猪圈信息
	PigstyInfo           []*PigstyInfo `protobuf:"bytes,3,rep,name=pigstyInfo,proto3" json:"pigstyInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PigHouseInfo) Reset()         { *m = PigHouseInfo{} }
func (m *PigHouseInfo) String() string { return proto.CompactTextString(m) }
func (*PigHouseInfo) ProtoMessage()    {}
func (*PigHouseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{1}
}

func (m *PigHouseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PigHouseInfo.Unmarshal(m, b)
}
func (m *PigHouseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PigHouseInfo.Marshal(b, m, deterministic)
}
func (m *PigHouseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PigHouseInfo.Merge(m, src)
}
func (m *PigHouseInfo) XXX_Size() int {
	return xxx_messageInfo_PigHouseInfo.Size(m)
}
func (m *PigHouseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PigHouseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PigHouseInfo proto.InternalMessageInfo

func (m *PigHouseInfo) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func (m *PigHouseInfo) GetAGE() string {
	if m != nil {
		return m.AGE
	}
	return ""
}

func (m *PigHouseInfo) GetPigstyInfo() []*PigstyInfo {
	if m != nil {
		return m.PigstyInfo
	}
	return nil
}

type DevInfoReq struct {
	//Version ,Devid ,Timestamp
	ReqHeader            *ReqHeader `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DevInfoReq) Reset()         { *m = DevInfoReq{} }
func (m *DevInfoReq) String() string { return proto.CompactTextString(m) }
func (*DevInfoReq) ProtoMessage()    {}
func (*DevInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{2}
}

func (m *DevInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevInfoReq.Unmarshal(m, b)
}
func (m *DevInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevInfoReq.Marshal(b, m, deterministic)
}
func (m *DevInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevInfoReq.Merge(m, src)
}
func (m *DevInfoReq) XXX_Size() int {
	return xxx_messageInfo_DevInfoReq.Size(m)
}
func (m *DevInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DevInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DevInfoReq proto.InternalMessageInfo

func (m *DevInfoReq) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

//
type PigstyInfoRes struct {
	//Version ,Devid ,Timestamp
	ReqHeader *ReqHeader `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	//猪场ID
	PigFarmId string `protobuf:"bytes,2,opt,name=PigFarmId,proto3" json:"PigFarmId,omitempty"`
	//猪场主人
	Farmer string `protobuf:"bytes,3,opt,name=farmer,proto3" json:"farmer,omitempty"`
	//猪舍信息
	PigHouseInfo []*PigHouseInfo `protobuf:"bytes,4,rep,name=pigHouseInfo,proto3" json:"pigHouseInfo,omitempty"`
	//猪场版本信息
	FarmVer              string   `protobuf:"bytes,5,opt,name=FarmVer,proto3" json:"FarmVer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PigstyInfoRes) Reset()         { *m = PigstyInfoRes{} }
func (m *PigstyInfoRes) String() string { return proto.CompactTextString(m) }
func (*PigstyInfoRes) ProtoMessage()    {}
func (*PigstyInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{3}
}

func (m *PigstyInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PigstyInfoRes.Unmarshal(m, b)
}
func (m *PigstyInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PigstyInfoRes.Marshal(b, m, deterministic)
}
func (m *PigstyInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PigstyInfoRes.Merge(m, src)
}
func (m *PigstyInfoRes) XXX_Size() int {
	return xxx_messageInfo_PigstyInfoRes.Size(m)
}
func (m *PigstyInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PigstyInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_PigstyInfoRes proto.InternalMessageInfo

func (m *PigstyInfoRes) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

func (m *PigstyInfoRes) GetPigFarmId() string {
	if m != nil {
		return m.PigFarmId
	}
	return ""
}

func (m *PigstyInfoRes) GetFarmer() string {
	if m != nil {
		return m.Farmer
	}
	return ""
}

func (m *PigstyInfoRes) GetPigHouseInfo() []*PigHouseInfo {
	if m != nil {
		return m.PigHouseInfo
	}
	return nil
}

func (m *PigstyInfoRes) GetFarmVer() string {
	if m != nil {
		return m.FarmVer
	}
	return ""
}

//ChoosePigstyReq
//称在一秒内6次上传当前重量,重量选取算法：
//如果一秒内的重量变化（小于0.5kg），那么上传当前重量。
//如果下一秒的数据无变化，那么不上传数据。
type ScaleDevRawData struct {
	//当前猪圈
	PigstyId int32 `protobuf:"varint,1,opt,name=PigstyId,proto3" json:"PigstyId,omitempty"`
	//当前重量，重量采用“克”为单位
	CurrentWeight        float64  `protobuf:"fixed64,2,opt,name=CurrentWeight,proto3" json:"CurrentWeight,omitempty"`
	FeedType             FeedType `protobuf:"varint,3,opt,name=FeedType,proto3,enum=fsapi.FeedType" json:"FeedType,omitempty"`
	Timestamp            uint64   `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleDevRawData) Reset()         { *m = ScaleDevRawData{} }
func (m *ScaleDevRawData) String() string { return proto.CompactTextString(m) }
func (*ScaleDevRawData) ProtoMessage()    {}
func (*ScaleDevRawData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{4}
}

func (m *ScaleDevRawData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleDevRawData.Unmarshal(m, b)
}
func (m *ScaleDevRawData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleDevRawData.Marshal(b, m, deterministic)
}
func (m *ScaleDevRawData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleDevRawData.Merge(m, src)
}
func (m *ScaleDevRawData) XXX_Size() int {
	return xxx_messageInfo_ScaleDevRawData.Size(m)
}
func (m *ScaleDevRawData) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleDevRawData.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleDevRawData proto.InternalMessageInfo

func (m *ScaleDevRawData) GetPigstyId() int32 {
	if m != nil {
		return m.PigstyId
	}
	return 0
}

func (m *ScaleDevRawData) GetCurrentWeight() float64 {
	if m != nil {
		return m.CurrentWeight
	}
	return 0
}

func (m *ScaleDevRawData) GetFeedType() FeedType {
	if m != nil {
		return m.FeedType
	}
	return FeedType_LOAD
}

func (m *ScaleDevRawData) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type UploadDevDateReq struct {
	//Version ,Devid ,Timestamp
	ReqHeader  *ReqHeader         `protobuf:"bytes,1,opt,name=ReqHeader,proto3" json:"ReqHeader,omitempty"`
	DevRawData []*ScaleDevRawData `protobuf:"bytes,2,rep,name=DevRawData,proto3" json:"DevRawData,omitempty"`
	//包序号
	SerialNo             int32    `protobuf:"varint,3,opt,name=SerialNo,proto3" json:"SerialNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadDevDateReq) Reset()         { *m = UploadDevDateReq{} }
func (m *UploadDevDateReq) String() string { return proto.CompactTextString(m) }
func (*UploadDevDateReq) ProtoMessage()    {}
func (*UploadDevDateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1a5b6d85b3d9f79, []int{5}
}

func (m *UploadDevDateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadDevDateReq.Unmarshal(m, b)
}
func (m *UploadDevDateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadDevDateReq.Marshal(b, m, deterministic)
}
func (m *UploadDevDateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadDevDateReq.Merge(m, src)
}
func (m *UploadDevDateReq) XXX_Size() int {
	return xxx_messageInfo_UploadDevDateReq.Size(m)
}
func (m *UploadDevDateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadDevDateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadDevDateReq proto.InternalMessageInfo

func (m *UploadDevDateReq) GetReqHeader() *ReqHeader {
	if m != nil {
		return m.ReqHeader
	}
	return nil
}

func (m *UploadDevDateReq) GetDevRawData() []*ScaleDevRawData {
	if m != nil {
		return m.DevRawData
	}
	return nil
}

func (m *UploadDevDateReq) GetSerialNo() int32 {
	if m != nil {
		return m.SerialNo
	}
	return 0
}

func init() {
	proto.RegisterType((*PigstyInfo)(nil), "fsapi.PigstyInfo")
	proto.RegisterType((*PigHouseInfo)(nil), "fsapi.PigHouseInfo")
	proto.RegisterType((*DevInfoReq)(nil), "fsapi.DevInfoReq")
	proto.RegisterType((*PigstyInfoRes)(nil), "fsapi.PigstyInfoRes")
	proto.RegisterType((*ScaleDevRawData)(nil), "fsapi.ScaleDevRawData")
	proto.RegisterType((*UploadDevDateReq)(nil), "fsapi.UploadDevDateReq")
}

func init() { proto.RegisterFile("fs_gw.proto", fileDescriptor_d1a5b6d85b3d9f79) }

var fileDescriptor_d1a5b6d85b3d9f79 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x4e, 0xdb, 0x40,
	0x10, 0xc6, 0x31, 0x49, 0x20, 0x1e, 0x48, 0x49, 0xb7, 0x88, 0xae, 0xa2, 0xaa, 0x8a, 0x2c, 0x0e,
	0x91, 0xaa, 0x3a, 0x6a, 0xaa, 0x96, 0x0b, 0x17, 0x68, 0x1a, 0x40, 0x42, 0xc8, 0xda, 0xd0, 0x56,
	0xea, 0x05, 0x2d, 0xf6, 0x78, 0x59, 0x29, 0x8e, 0xcd, 0xae, 0x93, 0x08, 0xa9, 0xcf, 0xd1, 0x63,
	0x9f, 0xa7, 0x0f, 0xd4, 0x07, 0xa8, 0xbc, 0xf1, 0x9f, 0x04, 0x4e, 0xed, 0x2d, 0xdf, 0xb7, 0x33,
	0xeb, 0xfd, 0x7d, 0x33, 0x0a, 0xec, 0x84, 0xfa, 0x46, 0x2c, 0xdc, 0x44, 0xc5, 0x69, 0x4c, 0x1a,
	0xa1, 0xe6, 0x89, 0xec, 0xb4, 0xc5, 0x24, 0xbe, 0xe5, 0x93, 0x9b, 0x00, 0xc3, 0xe5, 0x81, 0xf3,
	0xc7, 0x02, 0xf0, 0xa4, 0xd0, 0xe9, 0xc3, 0xc5, 0x34, 0x8c, 0x49, 0x07, 0x9a, 0xb9, 0x0a, 0xa8,
	0xd5, 0xb5, 0x7a, 0x36, 0x2b, 0x35, 0x39, 0x80, 0x2d, 0x4f, 0x8a, 0xab, 0x59, 0x44, 0x37, 0xbb,
	0x56, 0xaf, 0xc5, 0x72, 0x45, 0x0e, 0xa1, 0x75, 0x32, 0x47, 0xc5, 0x05, 0x7e, 0x43, 0x29, 0xee,
	0x52, 0x5a, 0x33, 0xc7, 0xeb, 0x26, 0x71, 0x60, 0xf7, 0x24, 0x98, 0x4b, 0x5d, 0x14, 0xd5, 0x4d,
	0xd1, 0x9a, 0x47, 0x28, 0x6c, 0x5f, 0x72, 0x9d, 0x8e, 0x30, 0xa0, 0x8d, 0xae, 0xd5, 0xab, 0xb3,
	0x42, 0x92, 0x7d, 0x68, 0x78, 0x52, 0x5c, 0x04, 0x74, 0xab, 0x5b, 0xeb, 0xd9, 0x6c, 0x29, 0xb2,
	0xfa, 0x71, 0xfa, 0x70, 0xc5, 0x23, 0xa4, 0xdb, 0xe6, 0xb1, 0x85, 0x24, 0xaf, 0x0d, 0xd5, 0x38,
	0x41, 0x5f, 0xa2, 0xa6, 0x4d, 0x73, 0xb8, 0xe2, 0x38, 0x11, 0xec, 0x7a, 0x52, 0x9c, 0xc7, 0x33,
	0x8d, 0x86, 0x9b, 0xc2, 0xf6, 0x52, 0x14, 0xd8, 0x85, 0x24, 0x6d, 0xa8, 0x9d, 0x9c, 0x7d, 0x36,
	0xc8, 0x36, 0xcb, 0x7e, 0x92, 0x77, 0x00, 0x49, 0x99, 0x18, 0xad, 0x75, 0x6b, 0xbd, 0x9d, 0xc1,
	0x73, 0xd7, 0x04, 0xec, 0x56, 0x51, 0xb2, 0x95, 0x22, 0xe7, 0x18, 0x60, 0x88, 0x73, 0x63, 0xe3,
	0x3d, 0x71, 0xc1, 0x66, 0x78, 0x7f, 0x8e, 0x3c, 0x40, 0x65, 0x3e, 0xb7, 0x33, 0x68, 0xe7, 0xfd,
	0xa5, 0xcf, 0xaa, 0x12, 0xe7, 0xb7, 0x05, 0xad, 0x95, 0x8b, 0x51, 0xff, 0xeb, 0x0d, 0xe4, 0x15,
	0xd8, 0x9e, 0x14, 0x23, 0xae, 0xa2, 0x8b, 0x20, 0x47, 0xa9, 0x8c, 0x6c, 0xb0, 0x21, 0x57, 0x11,
	0x2a, 0x33, 0x39, 0x9b, 0xe5, 0x8a, 0x1c, 0xc1, 0x6e, 0xb2, 0x12, 0x12, 0xad, 0x1b, 0xd4, 0x17,
	0x15, 0x6a, 0x79, 0xc4, 0xd6, 0x0a, 0xb3, 0x34, 0xb3, 0xab, 0xbf, 0xa2, 0x32, 0x73, 0xb4, 0x59,
	0x21, 0x9d, 0x5f, 0x16, 0xec, 0x8d, 0x7d, 0x3e, 0xc1, 0x21, 0xce, 0x19, 0x5f, 0x0c, 0x79, 0xca,
	0x9f, 0xec, 0x5c, 0x63, 0x65, 0xe7, 0x0e, 0xa1, 0xf5, 0x69, 0xa6, 0x14, 0x4e, 0xd3, 0x7c, 0x6d,
	0xb2, 0xc7, 0x5b, 0x6c, 0xdd, 0x24, 0x6f, 0xa0, 0x39, 0x42, 0x0c, 0xae, 0x1f, 0x12, 0x34, 0x08,
	0xcf, 0x06, 0x7b, 0xf9, 0x23, 0x0b, 0x9b, 0x95, 0x05, 0x59, 0x16, 0xd7, 0x32, 0x42, 0x9d, 0xf2,
	0x28, 0x31, 0x5b, 0x58, 0x67, 0x95, 0xe1, 0xfc, 0xb4, 0xa0, 0xfd, 0x25, 0x99, 0xc4, 0x3c, 0x18,
	0xe2, 0x7c, 0xc8, 0x53, 0xfc, 0x8f, 0x81, 0x91, 0x8f, 0x66, 0xdc, 0x39, 0x1f, 0xdd, 0x34, 0xb1,
	0x1d, 0xe4, 0x0d, 0x8f, 0xe8, 0x19, 0xac, 0x27, 0x31, 0x46, 0x25, 0xf9, 0xe4, 0x2a, 0x36, 0x1c,
	0x0d, 0x56, 0xea, 0xc1, 0x0f, 0x68, 0x8c, 0xb4, 0xc7, 0x03, 0xf2, 0x01, 0x9a, 0x1e, 0x0f, 0x2e,
	0x63, 0x21, 0xa7, 0xa4, 0x58, 0xbb, 0x6a, 0xb9, 0x3a, 0xfb, 0x4f, 0x37, 0x11, 0xb5, 0xb3, 0x41,
	0x8e, 0xa1, 0xb5, 0xe4, 0x62, 0x7c, 0x61, 0x86, 0xf4, 0x32, 0x2f, 0x7c, 0x4c, 0xdb, 0xa9, 0xd0,
	0x74, 0xbe, 0x80, 0x1b, 0xa7, 0x02, 0xa8, 0x3f, 0x75, 0x65, 0x9c, 0x48, 0xe1, 0x86, 0x88, 0x81,
	0xab, 0x33, 0x0c, 0x57, 0xa8, 0xc4, 0x3f, 0x6d, 0x67, 0xd1, 0x66, 0x58, 0x3e, 0x9e, 0xa9, 0xc4,
	0x3f, 0x5b, 0x78, 0xd6, 0xf7, 0x23, 0x21, 0xd3, 0xbb, 0xd9, 0xad, 0xeb, 0xc7, 0x51, 0xdf, 0x74,
	0xf5, 0xb3, 0xae, 0xb7, 0xa6, 0xab, 0x2f, 0xa7, 0x29, 0xaa, 0x90, 0xfb, 0xd8, 0xcf, 0xfa, 0xfb,
	0x22, 0xbe, 0x89, 0x67, 0x69, 0xdf, 0x7c, 0xf4, 0x76, 0xcb, 0xfc, 0x2d, 0xbd, 0xff, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x25, 0xd9, 0xb5, 0x94, 0xbe, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FsPadClient is the client API for FsPad service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FsPadClient interface {
	//PAD 登录
	PadLogin(ctx context.Context, in *DevInfoReq, opts ...grpc.CallOption) (*PigstyInfoRes, error)
	//上传当前重量
	UploadRawInfo(ctx context.Context, in *UploadDevDateReq, opts ...grpc.CallOption) (*ResHeader, error)
}

type fsPadClient struct {
	cc *grpc.ClientConn
}

func NewFsPadClient(cc *grpc.ClientConn) FsPadClient {
	return &fsPadClient{cc}
}

func (c *fsPadClient) PadLogin(ctx context.Context, in *DevInfoReq, opts ...grpc.CallOption) (*PigstyInfoRes, error) {
	out := new(PigstyInfoRes)
	err := c.cc.Invoke(ctx, "/fsapi.FsPad/PadLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsPadClient) UploadRawInfo(ctx context.Context, in *UploadDevDateReq, opts ...grpc.CallOption) (*ResHeader, error) {
	out := new(ResHeader)
	err := c.cc.Invoke(ctx, "/fsapi.FsPad/UploadRawInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FsPadServer is the server API for FsPad service.
type FsPadServer interface {
	//PAD 登录
	PadLogin(context.Context, *DevInfoReq) (*PigstyInfoRes, error)
	//上传当前重量
	UploadRawInfo(context.Context, *UploadDevDateReq) (*ResHeader, error)
}

// UnimplementedFsPadServer can be embedded to have forward compatible implementations.
type UnimplementedFsPadServer struct {
}

func (*UnimplementedFsPadServer) PadLogin(ctx context.Context, req *DevInfoReq) (*PigstyInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PadLogin not implemented")
}
func (*UnimplementedFsPadServer) UploadRawInfo(ctx context.Context, req *UploadDevDateReq) (*ResHeader, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadRawInfo not implemented")
}

func RegisterFsPadServer(s *grpc.Server, srv FsPadServer) {
	s.RegisterService(&_FsPad_serviceDesc, srv)
}

func _FsPad_PadLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsPadServer).PadLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fsapi.FsPad/PadLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsPadServer).PadLogin(ctx, req.(*DevInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsPad_UploadRawInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDevDateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsPadServer).UploadRawInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fsapi.FsPad/UploadRawInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsPadServer).UploadRawInfo(ctx, req.(*UploadDevDateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FsPad_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fsapi.FsPad",
	HandlerType: (*FsPadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PadLogin",
			Handler:    _FsPad_PadLogin_Handler,
		},
		{
			MethodName: "UploadRawInfo",
			Handler:    _FsPad_UploadRawInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fs_gw.proto",
}
