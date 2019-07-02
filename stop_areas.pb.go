// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stop_areas.proto

package external_models

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

type ExternalStopArea struct {
	Objectid             string   `protobuf:"bytes,1,opt,name=objectid,proto3" json:"objectid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PositionX            int32    `protobuf:"varint,3,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	PositionY            int32    `protobuf:"varint,4,opt,name=position_y,json=positionY,proto3" json:"position_y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExternalStopArea) Reset()         { *m = ExternalStopArea{} }
func (m *ExternalStopArea) String() string { return proto.CompactTextString(m) }
func (*ExternalStopArea) ProtoMessage()    {}
func (*ExternalStopArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_82d1dca9db066224, []int{0}
}

func (m *ExternalStopArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalStopArea.Unmarshal(m, b)
}
func (m *ExternalStopArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalStopArea.Marshal(b, m, deterministic)
}
func (m *ExternalStopArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalStopArea.Merge(m, src)
}
func (m *ExternalStopArea) XXX_Size() int {
	return xxx_messageInfo_ExternalStopArea.Size(m)
}
func (m *ExternalStopArea) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalStopArea.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalStopArea proto.InternalMessageInfo

func (m *ExternalStopArea) GetObjectid() string {
	if m != nil {
		return m.Objectid
	}
	return ""
}

func (m *ExternalStopArea) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExternalStopArea) GetPositionX() int32 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *ExternalStopArea) GetPositionY() int32 {
	if m != nil {
		return m.PositionY
	}
	return 0
}

type ExternalStopAreas struct {
	StopAreas            []*ExternalStopArea `protobuf:"bytes,1,rep,name=stop_areas,json=stopAreas,proto3" json:"stop_areas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExternalStopAreas) Reset()         { *m = ExternalStopAreas{} }
func (m *ExternalStopAreas) String() string { return proto.CompactTextString(m) }
func (*ExternalStopAreas) ProtoMessage()    {}
func (*ExternalStopAreas) Descriptor() ([]byte, []int) {
	return fileDescriptor_82d1dca9db066224, []int{1}
}

func (m *ExternalStopAreas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalStopAreas.Unmarshal(m, b)
}
func (m *ExternalStopAreas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalStopAreas.Marshal(b, m, deterministic)
}
func (m *ExternalStopAreas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalStopAreas.Merge(m, src)
}
func (m *ExternalStopAreas) XXX_Size() int {
	return xxx_messageInfo_ExternalStopAreas.Size(m)
}
func (m *ExternalStopAreas) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalStopAreas.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalStopAreas proto.InternalMessageInfo

func (m *ExternalStopAreas) GetStopAreas() []*ExternalStopArea {
	if m != nil {
		return m.StopAreas
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalStopArea)(nil), "external_models.ExternalStopArea")
	proto.RegisterType((*ExternalStopAreas)(nil), "external_models.ExternalStopAreas")
}

func init() { proto.RegisterFile("stop_areas.proto", fileDescriptor_82d1dca9db066224) }

var fileDescriptor_82d1dca9db066224 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0xc9, 0x2f,
	0x88, 0x4f, 0x2c, 0x4a, 0x4d, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0xad,
	0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x89, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x56, 0x6a, 0x60,
	0xe4, 0x12, 0x70, 0x85, 0x8a, 0x05, 0x97, 0xe4, 0x17, 0x38, 0x16, 0xa5, 0x26, 0x0a, 0x49, 0x71,
	0x71, 0xe4, 0x27, 0x65, 0xa5, 0x26, 0x97, 0x64, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x71, 0x30, 0x5b,
	0x48, 0x96, 0x8b, 0xab, 0x20, 0xbf, 0x38, 0xb3, 0x24, 0x33, 0x3f, 0x2f, 0xbe, 0x42, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x35, 0x88, 0x13, 0x26, 0x12, 0x81, 0x22, 0x5d, 0x29, 0xc1, 0x82, 0x2a, 0x1d,
	0xa9, 0x14, 0xca, 0x25, 0x88, 0xee, 0x82, 0x62, 0x21, 0x07, 0x2e, 0x2e, 0x84, 0xe3, 0x25, 0x18,
	0x15, 0x98, 0x35, 0xb8, 0x8d, 0x14, 0xf5, 0xd0, 0x5c, 0xaf, 0x87, 0xae, 0x2f, 0x88, 0xb3, 0x18,
	0x66, 0x42, 0x12, 0x1b, 0xd8, 0xc7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x46, 0x94,
	0xa0, 0x05, 0x01, 0x00, 0x00,
}