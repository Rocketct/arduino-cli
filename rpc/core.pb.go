// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package rpc // import "github.com/arduino/arduino-cli/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PlatformInstallReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Version              string    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformInstallReq) Reset()         { *m = PlatformInstallReq{} }
func (m *PlatformInstallReq) String() string { return proto.CompactTextString(m) }
func (*PlatformInstallReq) ProtoMessage()    {}
func (*PlatformInstallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_5a4e192f333a13f3, []int{0}
}
func (m *PlatformInstallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformInstallReq.Unmarshal(m, b)
}
func (m *PlatformInstallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformInstallReq.Marshal(b, m, deterministic)
}
func (dst *PlatformInstallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformInstallReq.Merge(dst, src)
}
func (m *PlatformInstallReq) XXX_Size() int {
	return xxx_messageInfo_PlatformInstallReq.Size(m)
}
func (m *PlatformInstallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformInstallReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformInstallReq proto.InternalMessageInfo

func (m *PlatformInstallReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformInstallReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformInstallReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *PlatformInstallReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PlatformInstallResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformInstallResp) Reset()         { *m = PlatformInstallResp{} }
func (m *PlatformInstallResp) String() string { return proto.CompactTextString(m) }
func (*PlatformInstallResp) ProtoMessage()    {}
func (*PlatformInstallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_5a4e192f333a13f3, []int{1}
}
func (m *PlatformInstallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformInstallResp.Unmarshal(m, b)
}
func (m *PlatformInstallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformInstallResp.Marshal(b, m, deterministic)
}
func (dst *PlatformInstallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformInstallResp.Merge(dst, src)
}
func (m *PlatformInstallResp) XXX_Size() int {
	return xxx_messageInfo_PlatformInstallResp.Size(m)
}
func (m *PlatformInstallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformInstallResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformInstallResp proto.InternalMessageInfo

type PlatformDownloadReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Version              string    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformDownloadReq) Reset()         { *m = PlatformDownloadReq{} }
func (m *PlatformDownloadReq) String() string { return proto.CompactTextString(m) }
func (*PlatformDownloadReq) ProtoMessage()    {}
func (*PlatformDownloadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_5a4e192f333a13f3, []int{2}
}
func (m *PlatformDownloadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformDownloadReq.Unmarshal(m, b)
}
func (m *PlatformDownloadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformDownloadReq.Marshal(b, m, deterministic)
}
func (dst *PlatformDownloadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformDownloadReq.Merge(dst, src)
}
func (m *PlatformDownloadReq) XXX_Size() int {
	return xxx_messageInfo_PlatformDownloadReq.Size(m)
}
func (m *PlatformDownloadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformDownloadReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformDownloadReq proto.InternalMessageInfo

func (m *PlatformDownloadReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformDownloadReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformDownloadReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *PlatformDownloadReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PlatformDownloadResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformDownloadResp) Reset()         { *m = PlatformDownloadResp{} }
func (m *PlatformDownloadResp) String() string { return proto.CompactTextString(m) }
func (*PlatformDownloadResp) ProtoMessage()    {}
func (*PlatformDownloadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_5a4e192f333a13f3, []int{3}
}
func (m *PlatformDownloadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformDownloadResp.Unmarshal(m, b)
}
func (m *PlatformDownloadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformDownloadResp.Marshal(b, m, deterministic)
}
func (dst *PlatformDownloadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformDownloadResp.Merge(dst, src)
}
func (m *PlatformDownloadResp) XXX_Size() int {
	return xxx_messageInfo_PlatformDownloadResp.Size(m)
}
func (m *PlatformDownloadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformDownloadResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformDownloadResp proto.InternalMessageInfo

type PlatformUninstallReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	PlatformPackage      string    `protobuf:"bytes,2,opt,name=platform_package,json=platformPackage,proto3" json:"platform_package,omitempty"`
	Architecture         string    `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Version              string    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlatformUninstallReq) Reset()         { *m = PlatformUninstallReq{} }
func (m *PlatformUninstallReq) String() string { return proto.CompactTextString(m) }
func (*PlatformUninstallReq) ProtoMessage()    {}
func (*PlatformUninstallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_5a4e192f333a13f3, []int{4}
}
func (m *PlatformUninstallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformUninstallReq.Unmarshal(m, b)
}
func (m *PlatformUninstallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformUninstallReq.Marshal(b, m, deterministic)
}
func (dst *PlatformUninstallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformUninstallReq.Merge(dst, src)
}
func (m *PlatformUninstallReq) XXX_Size() int {
	return xxx_messageInfo_PlatformUninstallReq.Size(m)
}
func (m *PlatformUninstallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformUninstallReq.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformUninstallReq proto.InternalMessageInfo

func (m *PlatformUninstallReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *PlatformUninstallReq) GetPlatformPackage() string {
	if m != nil {
		return m.PlatformPackage
	}
	return ""
}

func (m *PlatformUninstallReq) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *PlatformUninstallReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PlatformUninstallResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformUninstallResp) Reset()         { *m = PlatformUninstallResp{} }
func (m *PlatformUninstallResp) String() string { return proto.CompactTextString(m) }
func (*PlatformUninstallResp) ProtoMessage()    {}
func (*PlatformUninstallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_5a4e192f333a13f3, []int{5}
}
func (m *PlatformUninstallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformUninstallResp.Unmarshal(m, b)
}
func (m *PlatformUninstallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformUninstallResp.Marshal(b, m, deterministic)
}
func (dst *PlatformUninstallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformUninstallResp.Merge(dst, src)
}
func (m *PlatformUninstallResp) XXX_Size() int {
	return xxx_messageInfo_PlatformUninstallResp.Size(m)
}
func (m *PlatformUninstallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformUninstallResp.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformUninstallResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PlatformInstallReq)(nil), "arduino.PlatformInstallReq")
	proto.RegisterType((*PlatformInstallResp)(nil), "arduino.PlatformInstallResp")
	proto.RegisterType((*PlatformDownloadReq)(nil), "arduino.PlatformDownloadReq")
	proto.RegisterType((*PlatformDownloadResp)(nil), "arduino.PlatformDownloadResp")
	proto.RegisterType((*PlatformUninstallReq)(nil), "arduino.PlatformUninstallReq")
	proto.RegisterType((*PlatformUninstallResp)(nil), "arduino.PlatformUninstallResp")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_core_5a4e192f333a13f3) }

var fileDescriptor_core_5a4e192f333a13f3 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2c, 0x4a, 0x29, 0xcd, 0xcc, 0xcb, 0x97,
	0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x08, 0x2b, 0xad, 0x61, 0xe4, 0x12, 0x0a, 0xc8,
	0x49, 0x2c, 0x49, 0xcb, 0x2f, 0xca, 0xf5, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0xc9, 0x09, 0x4a, 0x2d,
	0x14, 0xd2, 0xe5, 0xe2, 0xc8, 0x04, 0xf1, 0xf2, 0x92, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0x04, 0xf5, 0xa0, 0x06, 0xe8, 0x79, 0x42, 0x25, 0x82, 0xe0, 0x4a, 0x84, 0x34, 0xb9, 0x04,
	0x0a, 0xa0, 0x86, 0xc4, 0x17, 0x24, 0x26, 0x67, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x01, 0xb5, 0x71,
	0x06, 0xf1, 0xc3, 0xc4, 0x03, 0x20, 0xc2, 0x42, 0x4a, 0x5c, 0x3c, 0x89, 0x45, 0xc9, 0x19, 0x99,
	0x25, 0xa9, 0xc9, 0x25, 0xa5, 0x45, 0xa9, 0x12, 0xcc, 0x60, 0x65, 0x28, 0x62, 0x42, 0x12, 0x5c,
	0xec, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x2c, 0x60, 0x69, 0x18, 0x57, 0x49, 0x94,
	0x4b, 0x18, 0xc3, 0xb5, 0xc5, 0x05, 0x4a, 0x6b, 0x19, 0x11, 0xe2, 0x2e, 0xf9, 0xe5, 0x79, 0x39,
	0xf9, 0x89, 0x29, 0x83, 0xd9, 0x1b, 0x62, 0x5c, 0x22, 0x98, 0xce, 0x05, 0xfa, 0x63, 0x1d, 0x23,
	0x42, 0x22, 0x34, 0x2f, 0x73, 0xf0, 0xc7, 0x87, 0x38, 0x97, 0x28, 0x16, 0xf7, 0x16, 0x17, 0x38,
	0xa9, 0x44, 0x29, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x01, 0x93, 0x9b, 0x3e, 0xd4, 0xa9,
	0x30, 0x5a, 0x37, 0x39, 0x27, 0x53, 0xbf, 0xa8, 0x20, 0x39, 0x89, 0x0d, 0x9c, 0x08, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x15, 0xb3, 0x58, 0xa9, 0x02, 0x00, 0x00,
}