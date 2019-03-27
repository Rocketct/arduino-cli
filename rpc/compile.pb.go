// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compile.proto

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

type CompileReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	SketchPath           string    `protobuf:"bytes,3,opt,name=sketchPath,proto3" json:"sketchPath,omitempty"`
	ShowProperties       bool      `protobuf:"varint,4,opt,name=showProperties,proto3" json:"showProperties,omitempty"`
	Preprocess           bool      `protobuf:"varint,5,opt,name=preprocess,proto3" json:"preprocess,omitempty"`
	BuildCachePath       string    `protobuf:"bytes,6,opt,name=buildCachePath,proto3" json:"buildCachePath,omitempty"`
	BuildPath            string    `protobuf:"bytes,7,opt,name=buildPath,proto3" json:"buildPath,omitempty"`
	BuildProperties      []string  `protobuf:"bytes,8,rep,name=buildProperties,proto3" json:"buildProperties,omitempty"`
	Warnings             string    `protobuf:"bytes,9,opt,name=warnings,proto3" json:"warnings,omitempty"`
	Verbose              bool      `protobuf:"varint,10,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Quiet                bool      `protobuf:"varint,11,opt,name=quiet,proto3" json:"quiet,omitempty"`
	VidPid               string    `protobuf:"bytes,12,opt,name=vidPid,proto3" json:"vidPid,omitempty"`
	ExportFile           string    `protobuf:"bytes,13,opt,name=exportFile,proto3" json:"exportFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CompileReq) Reset()         { *m = CompileReq{} }
func (m *CompileReq) String() string { return proto.CompactTextString(m) }
func (*CompileReq) ProtoMessage()    {}
func (*CompileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_compile_14c4f101697f20d5, []int{0}
}
func (m *CompileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileReq.Unmarshal(m, b)
}
func (m *CompileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileReq.Marshal(b, m, deterministic)
}
func (dst *CompileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileReq.Merge(dst, src)
}
func (m *CompileReq) XXX_Size() int {
	return xxx_messageInfo_CompileReq.Size(m)
}
func (m *CompileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileReq.DiscardUnknown(m)
}

var xxx_messageInfo_CompileReq proto.InternalMessageInfo

func (m *CompileReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CompileReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *CompileReq) GetSketchPath() string {
	if m != nil {
		return m.SketchPath
	}
	return ""
}

func (m *CompileReq) GetShowProperties() bool {
	if m != nil {
		return m.ShowProperties
	}
	return false
}

func (m *CompileReq) GetPreprocess() bool {
	if m != nil {
		return m.Preprocess
	}
	return false
}

func (m *CompileReq) GetBuildCachePath() string {
	if m != nil {
		return m.BuildCachePath
	}
	return ""
}

func (m *CompileReq) GetBuildPath() string {
	if m != nil {
		return m.BuildPath
	}
	return ""
}

func (m *CompileReq) GetBuildProperties() []string {
	if m != nil {
		return m.BuildProperties
	}
	return nil
}

func (m *CompileReq) GetWarnings() string {
	if m != nil {
		return m.Warnings
	}
	return ""
}

func (m *CompileReq) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *CompileReq) GetQuiet() bool {
	if m != nil {
		return m.Quiet
	}
	return false
}

func (m *CompileReq) GetVidPid() string {
	if m != nil {
		return m.VidPid
	}
	return ""
}

func (m *CompileReq) GetExportFile() string {
	if m != nil {
		return m.ExportFile
	}
	return ""
}

type CompileResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompileResp) Reset()         { *m = CompileResp{} }
func (m *CompileResp) String() string { return proto.CompactTextString(m) }
func (*CompileResp) ProtoMessage()    {}
func (*CompileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_compile_14c4f101697f20d5, []int{1}
}
func (m *CompileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileResp.Unmarshal(m, b)
}
func (m *CompileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileResp.Marshal(b, m, deterministic)
}
func (dst *CompileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileResp.Merge(dst, src)
}
func (m *CompileResp) XXX_Size() int {
	return xxx_messageInfo_CompileResp.Size(m)
}
func (m *CompileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileResp.DiscardUnknown(m)
}

var xxx_messageInfo_CompileResp proto.InternalMessageInfo

func (m *CompileResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileReq)(nil), "arduino.CompileReq")
	proto.RegisterType((*CompileResp)(nil), "arduino.CompileResp")
}

func init() { proto.RegisterFile("compile.proto", fileDescriptor_compile_14c4f101697f20d5) }

var fileDescriptor_compile_14c4f101697f20d5 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xdf, 0x6a, 0xea, 0x40,
	0x10, 0x87, 0xf1, 0xa8, 0x31, 0x8e, 0x7a, 0xe4, 0x2c, 0x87, 0xb2, 0x48, 0x29, 0x22, 0xa5, 0xf5,
	0xc6, 0x08, 0x2d, 0xf4, 0x01, 0x2a, 0x14, 0x7a, 0x27, 0xb9, 0xec, 0x5d, 0xfe, 0x4c, 0xcd, 0xd2,
	0x98, 0x5d, 0x77, 0x37, 0xda, 0x77, 0xea, 0x4b, 0x76, 0x9d, 0xc4, 0x28, 0x5e, 0x25, 0xf3, 0xcd,
	0x97, 0xd9, 0x49, 0x7e, 0x81, 0x51, 0x22, 0xb7, 0x4a, 0xe4, 0x18, 0x28, 0x2d, 0xad, 0x64, 0xbd,
	0x48, 0xa7, 0xa5, 0x28, 0xe4, 0x64, 0xe8, 0xf8, 0x56, 0x16, 0x15, 0x9e, 0xfd, 0xb4, 0x01, 0x56,
	0x95, 0x18, 0xe2, 0x8e, 0x2d, 0xc0, 0x17, 0x85, 0xb1, 0x51, 0x91, 0x20, 0x6f, 0x4d, 0x5b, 0xf3,
	0xc1, 0xd3, 0xbf, 0xa0, 0x7e, 0x30, 0x78, 0xaf, 0x1b, 0x61, 0xa3, 0x30, 0x06, 0x9d, 0xcf, 0x5d,
	0x5c, 0xf0, 0x3f, 0x4e, 0xed, 0x87, 0x74, 0xcf, 0xee, 0x00, 0xcc, 0x17, 0xda, 0x24, 0x5b, 0x47,
	0x36, 0xe3, 0x6d, 0xea, 0x5c, 0x10, 0xf6, 0x00, 0x7f, 0x4d, 0x26, 0x0f, 0x6b, 0x2d, 0x15, 0x6a,
	0x2b, 0xd0, 0xf0, 0x8e, 0x73, 0xfc, 0xf0, 0x8a, 0x1e, 0xe7, 0x28, 0x8d, 0x6e, 0xcb, 0x04, 0x8d,
	0xe1, 0x5d, 0x72, 0x2e, 0xc8, 0x71, 0x4e, 0x5c, 0x8a, 0x3c, 0x5d, 0x45, 0x49, 0x86, 0x74, 0x96,
	0x47, 0x67, 0x5d, 0x51, 0x76, 0x0b, 0x7d, 0x22, 0xa4, 0xf4, 0x48, 0x39, 0x03, 0x36, 0x87, 0x71,
	0x55, 0x9c, 0xd7, 0xf1, 0xa7, 0x6d, 0xe7, 0x5c, 0x63, 0x36, 0x01, 0xff, 0x10, 0xe9, 0x42, 0x14,
	0x1b, 0xc3, 0xfb, 0x34, 0xa6, 0xa9, 0x19, 0x87, 0xde, 0x1e, 0x75, 0x2c, 0x0d, 0x72, 0xa0, 0x45,
	0x4f, 0x25, 0xfb, 0x0f, 0xdd, 0x5d, 0x29, 0xd0, 0xf2, 0x01, 0xf1, 0xaa, 0x60, 0x37, 0xe0, 0xed,
	0x45, 0xba, 0x16, 0x29, 0x1f, 0xd2, 0xa4, 0xba, 0x3a, 0xbe, 0x33, 0x7e, 0x2b, 0xa9, 0xed, 0x9b,
	0xcb, 0x83, 0x8f, 0xaa, 0x6f, 0x77, 0x26, 0xb3, 0x17, 0x18, 0x34, 0x61, 0x19, 0xc5, 0x1e, 0xc1,
	0xd3, 0x68, 0xca, 0xdc, 0xd6, 0x59, 0x8d, 0x9b, 0xac, 0x42, 0xc2, 0x61, 0xdd, 0x7e, 0xbd, 0xff,
	0x98, 0x6d, 0x84, 0xcd, 0xca, 0x38, 0x70, 0xe1, 0x2f, 0x6b, 0xe9, 0x74, 0x5d, 0x24, 0xb9, 0x58,
	0x6a, 0x95, 0xc4, 0x1e, 0xfd, 0x12, 0xcf, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0xa4, 0xd9,
	0x97, 0x3a, 0x02, 0x00, 0x00,
}