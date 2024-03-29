// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stop_visits.proto

package external_models

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

type ExternalStopVisit struct {
	Objectid             string                   `protobuf:"bytes,1,opt,name=objectid,proto3" json:"objectid,omitempty"`
	StopAreaRef          string                   `protobuf:"bytes,2,opt,name=stopAreaRef,proto3" json:"stopAreaRef,omitempty"`
	PassageOrder         string                   `protobuf:"bytes,3,opt,name=passageOrder,proto3" json:"passageOrder,omitempty"`
	VehicleJourneyRef    string                   `protobuf:"bytes,4,opt,name=vehicleJourneyRef,proto3" json:"vehicleJourneyRef,omitempty"`
	Monitored            bool                     `protobuf:"varint,5,opt,name=monitored,proto3" json:"monitored,omitempty"`
	ArrivalTimes         *ExternalStopVisit_Times `protobuf:"bytes,6,opt,name=arrival_times,json=arrivalTimes,proto3" json:"arrival_times,omitempty"`
	DepartureTimes       *ExternalStopVisit_Times `protobuf:"bytes,7,opt,name=departure_times,json=departureTimes,proto3" json:"departure_times,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExternalStopVisit) Reset()         { *m = ExternalStopVisit{} }
func (m *ExternalStopVisit) String() string { return proto.CompactTextString(m) }
func (*ExternalStopVisit) ProtoMessage()    {}
func (*ExternalStopVisit) Descriptor() ([]byte, []int) {
	return fileDescriptor_17309d9b77d9e691, []int{0}
}

func (m *ExternalStopVisit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalStopVisit.Unmarshal(m, b)
}
func (m *ExternalStopVisit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalStopVisit.Marshal(b, m, deterministic)
}
func (m *ExternalStopVisit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalStopVisit.Merge(m, src)
}
func (m *ExternalStopVisit) XXX_Size() int {
	return xxx_messageInfo_ExternalStopVisit.Size(m)
}
func (m *ExternalStopVisit) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalStopVisit.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalStopVisit proto.InternalMessageInfo

func (m *ExternalStopVisit) GetObjectid() string {
	if m != nil {
		return m.Objectid
	}
	return ""
}

func (m *ExternalStopVisit) GetStopAreaRef() string {
	if m != nil {
		return m.StopAreaRef
	}
	return ""
}

func (m *ExternalStopVisit) GetPassageOrder() string {
	if m != nil {
		return m.PassageOrder
	}
	return ""
}

func (m *ExternalStopVisit) GetVehicleJourneyRef() string {
	if m != nil {
		return m.VehicleJourneyRef
	}
	return ""
}

func (m *ExternalStopVisit) GetMonitored() bool {
	if m != nil {
		return m.Monitored
	}
	return false
}

func (m *ExternalStopVisit) GetArrivalTimes() *ExternalStopVisit_Times {
	if m != nil {
		return m.ArrivalTimes
	}
	return nil
}

func (m *ExternalStopVisit) GetDepartureTimes() *ExternalStopVisit_Times {
	if m != nil {
		return m.DepartureTimes
	}
	return nil
}

type ExternalStopVisit_Times struct {
	Aimed                *timestamp.Timestamp `protobuf:"bytes,1,opt,name=aimed,proto3" json:"aimed,omitempty"`
	Expected             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expected,proto3" json:"expected,omitempty"`
	Actual               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=actual,proto3" json:"actual,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExternalStopVisit_Times) Reset()         { *m = ExternalStopVisit_Times{} }
func (m *ExternalStopVisit_Times) String() string { return proto.CompactTextString(m) }
func (*ExternalStopVisit_Times) ProtoMessage()    {}
func (*ExternalStopVisit_Times) Descriptor() ([]byte, []int) {
	return fileDescriptor_17309d9b77d9e691, []int{0, 0}
}

func (m *ExternalStopVisit_Times) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalStopVisit_Times.Unmarshal(m, b)
}
func (m *ExternalStopVisit_Times) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalStopVisit_Times.Marshal(b, m, deterministic)
}
func (m *ExternalStopVisit_Times) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalStopVisit_Times.Merge(m, src)
}
func (m *ExternalStopVisit_Times) XXX_Size() int {
	return xxx_messageInfo_ExternalStopVisit_Times.Size(m)
}
func (m *ExternalStopVisit_Times) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalStopVisit_Times.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalStopVisit_Times proto.InternalMessageInfo

func (m *ExternalStopVisit_Times) GetAimed() *timestamp.Timestamp {
	if m != nil {
		return m.Aimed
	}
	return nil
}

func (m *ExternalStopVisit_Times) GetExpected() *timestamp.Timestamp {
	if m != nil {
		return m.Expected
	}
	return nil
}

func (m *ExternalStopVisit_Times) GetActual() *timestamp.Timestamp {
	if m != nil {
		return m.Actual
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalStopVisit)(nil), "external_models.ExternalStopVisit")
	proto.RegisterType((*ExternalStopVisit_Times)(nil), "external_models.ExternalStopVisit.Times")
}

func init() { proto.RegisterFile("stop_visits.proto", fileDescriptor_17309d9b77d9e691) }

var fileDescriptor_17309d9b77d9e691 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xb6, 0xb5, 0x9d, 0x56, 0x4b, 0xf7, 0x14, 0x82, 0x60, 0xe8, 0x29, 0x07, 0x49,
	0xa5, 0x82, 0x77, 0x0f, 0x5e, 0x04, 0x11, 0xa3, 0x78, 0x2d, 0xdb, 0x64, 0x5a, 0x57, 0x36, 0xdd,
	0x65, 0x77, 0x52, 0xea, 0x67, 0xf0, 0x73, 0xf8, 0x3d, 0x25, 0xbb, 0xb1, 0xfe, 0x29, 0x14, 0x3c,
	0xce, 0xbc, 0xf7, 0x7e, 0xcb, 0xbe, 0x81, 0x91, 0x25, 0xa5, 0x67, 0x6b, 0x61, 0x05, 0xd9, 0x54,
	0x1b, 0x45, 0x8a, 0x0d, 0x71, 0x43, 0x68, 0x56, 0x5c, 0xce, 0x4a, 0x55, 0xa0, 0xb4, 0xd1, 0xd9,
	0x52, 0xa9, 0xa5, 0xc4, 0x89, 0x93, 0xe7, 0xd5, 0x62, 0x42, 0xa2, 0x44, 0x4b, 0xbc, 0xd4, 0x3e,
	0x31, 0x7e, 0x6f, 0xc1, 0xe8, 0xa6, 0x09, 0x3d, 0x92, 0xd2, 0xcf, 0x35, 0x8e, 0x45, 0xd0, 0x55,
	0xf3, 0x57, 0xcc, 0x49, 0x14, 0x61, 0x10, 0x07, 0x49, 0x2f, 0xdb, 0xce, 0x2c, 0x86, 0x7e, 0xfd,
	0xf0, 0xb5, 0x41, 0x9e, 0xe1, 0x22, 0x3c, 0x70, 0xf2, 0xcf, 0x15, 0x1b, 0xc3, 0x40, 0x73, 0x6b,
	0xf9, 0x12, 0xef, 0x4d, 0x81, 0x26, 0x3c, 0x74, 0x96, 0x5f, 0x3b, 0x76, 0x0e, 0xa3, 0x35, 0xbe,
	0x88, 0x5c, 0xe2, 0xad, 0xaa, 0xcc, 0x0a, 0xdf, 0x6a, 0x56, 0xcb, 0x19, 0x77, 0x05, 0x76, 0x0a,
	0xbd, 0x52, 0xad, 0x04, 0x29, 0x83, 0x45, 0xd8, 0x8e, 0x83, 0xa4, 0x9b, 0x7d, 0x2f, 0xd8, 0x1d,
	0x1c, 0x73, 0x63, 0xc4, 0x9a, 0xcb, 0x99, 0xfb, 0x5e, 0xd8, 0x89, 0x83, 0xa4, 0x3f, 0x4d, 0xd2,
	0x3f, 0x6d, 0xa4, 0x3b, 0x1f, 0x4d, 0x9f, 0x6a, 0x7f, 0x36, 0x68, 0xe2, 0x6e, 0x62, 0x0f, 0x30,
	0x2c, 0x50, 0x73, 0x43, 0x95, 0xc1, 0x06, 0x78, 0xf4, 0x4f, 0xe0, 0xc9, 0x16, 0xe0, 0xe6, 0xe8,
	0x23, 0x80, 0xb6, 0x87, 0x5f, 0x40, 0x9b, 0x8b, 0x12, 0x7d, 0xad, 0xfd, 0x69, 0x94, 0xfa, 0x03,
	0xa5, 0x5f, 0x07, 0xf2, 0x80, 0xfa, 0x40, 0x99, 0x37, 0xb2, 0x2b, 0xe8, 0xe2, 0x46, 0x63, 0x4e,
	0x58, 0xb8, 0xb2, 0xf7, 0x87, 0xb6, 0x5e, 0x36, 0x85, 0x0e, 0xcf, 0xa9, 0xe2, 0xd2, 0xf5, 0xbf,
	0x3f, 0xd5, 0x38, 0xe7, 0x1d, 0xa7, 0x5d, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x76, 0x73, 0x9a,
	0xe7, 0x5b, 0x02, 0x00, 0x00,
}
