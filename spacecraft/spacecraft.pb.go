// Code generated by protoc-gen-go.
// source: spacecraft.proto
// DO NOT EDIT!

/*
Package spacecraft is a generated protocol buffer package.

It is generated from these files:
	spacecraft.proto

It has these top-level messages:
	SvnUpParam
	VersionNum
	SvnCheckoutParams
	SvnUpToRevisionParams
	ResponseStr
	SpecifiedCommandParams
*/
package spacecraft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// svnUp params
type SvnUpParam struct {
	Dir string `protobuf:"bytes,1,opt,name=dir" json:"dir,omitempty"`
}

func (m *SvnUpParam) Reset()                    { *m = SvnUpParam{} }
func (m *SvnUpParam) String() string            { return proto.CompactTextString(m) }
func (*SvnUpParam) ProtoMessage()               {}
func (*SvnUpParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// svnUp response
type VersionNum struct {
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionNum) Reset()                    { *m = VersionNum{} }
func (m *VersionNum) String() string            { return proto.CompactTextString(m) }
func (*VersionNum) ProtoMessage()               {}
func (*VersionNum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// svnCheckout params
type SvnCheckoutParams struct {
	SvnUrl string `protobuf:"bytes,1,opt,name=svnUrl" json:"svnUrl,omitempty"`
	Dir    string `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
}

func (m *SvnCheckoutParams) Reset()                    { *m = SvnCheckoutParams{} }
func (m *SvnCheckoutParams) String() string            { return proto.CompactTextString(m) }
func (*SvnCheckoutParams) ProtoMessage()               {}
func (*SvnCheckoutParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// svnUpToRevision params
type SvnUpToRevisionParams struct {
	Dir     string `protobuf:"bytes,1,opt,name=dir" json:"dir,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *SvnUpToRevisionParams) Reset()                    { *m = SvnUpToRevisionParams{} }
func (m *SvnUpToRevisionParams) String() string            { return proto.CompactTextString(m) }
func (*SvnUpToRevisionParams) ProtoMessage()               {}
func (*SvnUpToRevisionParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ResponseStr struct {
	String_ string `protobuf:"bytes,1,opt,name=string" json:"string,omitempty"`
}

func (m *ResponseStr) Reset()                    { *m = ResponseStr{} }
func (m *ResponseStr) String() string            { return proto.CompactTextString(m) }
func (*ResponseStr) ProtoMessage()               {}
func (*ResponseStr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// specifiedCommand params
type SpecifiedCommandParams struct {
	Command string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Dir     string `protobuf:"bytes,2,opt,name=dir" json:"dir,omitempty"`
}

func (m *SpecifiedCommandParams) Reset()                    { *m = SpecifiedCommandParams{} }
func (m *SpecifiedCommandParams) String() string            { return proto.CompactTextString(m) }
func (*SpecifiedCommandParams) ProtoMessage()               {}
func (*SpecifiedCommandParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*SvnUpParam)(nil), "spacecraft.SvnUpParam")
	proto.RegisterType((*VersionNum)(nil), "spacecraft.VersionNum")
	proto.RegisterType((*SvnCheckoutParams)(nil), "spacecraft.SvnCheckoutParams")
	proto.RegisterType((*SvnUpToRevisionParams)(nil), "spacecraft.SvnUpToRevisionParams")
	proto.RegisterType((*ResponseStr)(nil), "spacecraft.ResponseStr")
	proto.RegisterType((*SpecifiedCommandParams)(nil), "spacecraft.SpecifiedCommandParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Spacecraft service

type SpacecraftClient interface {
	SvnUp(ctx context.Context, in *SvnUpParam, opts ...grpc.CallOption) (*VersionNum, error)
	SvnCheckout(ctx context.Context, in *SvnCheckoutParams, opts ...grpc.CallOption) (*VersionNum, error)
	SvnUpToRevision(ctx context.Context, in *SvnUpToRevisionParams, opts ...grpc.CallOption) (*VersionNum, error)
	SvnInfo(ctx context.Context, in *SvnUpParam, opts ...grpc.CallOption) (*ResponseStr, error)
	SpecifiedCommand(ctx context.Context, in *SpecifiedCommandParams, opts ...grpc.CallOption) (*ResponseStr, error)
	ComplexCommand(ctx context.Context, in *SpecifiedCommandParams, opts ...grpc.CallOption) (*ResponseStr, error)
}

type spacecraftClient struct {
	cc *grpc.ClientConn
}

func NewSpacecraftClient(cc *grpc.ClientConn) SpacecraftClient {
	return &spacecraftClient{cc}
}

func (c *spacecraftClient) SvnUp(ctx context.Context, in *SvnUpParam, opts ...grpc.CallOption) (*VersionNum, error) {
	out := new(VersionNum)
	err := grpc.Invoke(ctx, "/spacecraft.Spacecraft/svnUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacecraftClient) SvnCheckout(ctx context.Context, in *SvnCheckoutParams, opts ...grpc.CallOption) (*VersionNum, error) {
	out := new(VersionNum)
	err := grpc.Invoke(ctx, "/spacecraft.Spacecraft/svnCheckout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacecraftClient) SvnUpToRevision(ctx context.Context, in *SvnUpToRevisionParams, opts ...grpc.CallOption) (*VersionNum, error) {
	out := new(VersionNum)
	err := grpc.Invoke(ctx, "/spacecraft.Spacecraft/svnUpToRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacecraftClient) SvnInfo(ctx context.Context, in *SvnUpParam, opts ...grpc.CallOption) (*ResponseStr, error) {
	out := new(ResponseStr)
	err := grpc.Invoke(ctx, "/spacecraft.Spacecraft/svnInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacecraftClient) SpecifiedCommand(ctx context.Context, in *SpecifiedCommandParams, opts ...grpc.CallOption) (*ResponseStr, error) {
	out := new(ResponseStr)
	err := grpc.Invoke(ctx, "/spacecraft.Spacecraft/specifiedCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spacecraftClient) ComplexCommand(ctx context.Context, in *SpecifiedCommandParams, opts ...grpc.CallOption) (*ResponseStr, error) {
	out := new(ResponseStr)
	err := grpc.Invoke(ctx, "/spacecraft.Spacecraft/complexCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Spacecraft service

type SpacecraftServer interface {
	SvnUp(context.Context, *SvnUpParam) (*VersionNum, error)
	SvnCheckout(context.Context, *SvnCheckoutParams) (*VersionNum, error)
	SvnUpToRevision(context.Context, *SvnUpToRevisionParams) (*VersionNum, error)
	SvnInfo(context.Context, *SvnUpParam) (*ResponseStr, error)
	SpecifiedCommand(context.Context, *SpecifiedCommandParams) (*ResponseStr, error)
	ComplexCommand(context.Context, *SpecifiedCommandParams) (*ResponseStr, error)
}

func RegisterSpacecraftServer(s *grpc.Server, srv SpacecraftServer) {
	s.RegisterService(&_Spacecraft_serviceDesc, srv)
}

func _Spacecraft_SvnUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SvnUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SpacecraftServer).SvnUp(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Spacecraft_SvnCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SvnCheckoutParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SpacecraftServer).SvnCheckout(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Spacecraft_SvnUpToRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SvnUpToRevisionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SpacecraftServer).SvnUpToRevision(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Spacecraft_SvnInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SvnUpParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SpacecraftServer).SvnInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Spacecraft_SpecifiedCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SpecifiedCommandParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SpacecraftServer).SpecifiedCommand(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Spacecraft_ComplexCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SpecifiedCommandParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SpacecraftServer).ComplexCommand(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Spacecraft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacecraft.Spacecraft",
	HandlerType: (*SpacecraftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "svnUp",
			Handler:    _Spacecraft_SvnUp_Handler,
		},
		{
			MethodName: "svnCheckout",
			Handler:    _Spacecraft_SvnCheckout_Handler,
		},
		{
			MethodName: "svnUpToRevision",
			Handler:    _Spacecraft_SvnUpToRevision_Handler,
		},
		{
			MethodName: "svnInfo",
			Handler:    _Spacecraft_SvnInfo_Handler,
		},
		{
			MethodName: "specifiedCommand",
			Handler:    _Spacecraft_SpecifiedCommand_Handler,
		},
		{
			MethodName: "complexCommand",
			Handler:    _Spacecraft_ComplexCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x41, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0xdd, 0x96, 0x6d, 0xf1, 0x2d, 0xd1, 0xd9, 0x44, 0x5c, 0x4c, 0x34, 0xda, 0x44, 0xe3,
	0x69, 0x07, 0x3d, 0x99, 0xe8, 0x09, 0x63, 0xe2, 0x65, 0x2a, 0xa8, 0x77, 0x84, 0xa2, 0x8d, 0xa3,
	0x6d, 0x5a, 0x46, 0xfc, 0x7a, 0x7e, 0x33, 0x4b, 0x07, 0x83, 0x0e, 0xe5, 0xb4, 0x1b, 0xef, 0x85,
	0xfe, 0xde, 0xbf, 0xbf, 0x07, 0x30, 0x56, 0x22, 0x08, 0x49, 0x28, 0x83, 0x38, 0x9d, 0x0a, 0xc9,
	0x53, 0x8e, 0xa0, 0xea, 0xe0, 0x63, 0x00, 0x3f, 0x63, 0xaf, 0xe2, 0x29, 0x90, 0x41, 0x82, 0xc6,
	0xd0, 0x8b, 0xa8, 0x9c, 0x74, 0x4e, 0x3a, 0x17, 0xdb, 0x5e, 0xfe, 0x88, 0xcf, 0x01, 0xde, 0x88,
	0x54, 0x94, 0xb3, 0xd9, 0x22, 0x41, 0x13, 0x18, 0x66, 0xcb, 0xca, 0xbc, 0xd3, 0xf7, 0xca, 0x12,
	0xdf, 0xc2, 0x9e, 0xe6, 0xb8, 0x9f, 0x24, 0xfc, 0xe2, 0x8b, 0xd4, 0xd0, 0x14, 0x72, 0x60, 0xa0,
	0x34, 0x5c, 0xce, 0x0b, 0x62, 0x51, 0x95, 0x63, 0xba, 0xd5, 0x18, 0x17, 0xf6, 0x4d, 0x8c, 0x17,
	0xee, 0x91, 0x8c, 0xe6, 0xc4, 0x02, 0xd1, 0x48, 0x54, 0xcf, 0xd0, 0xb5, 0x33, 0x9c, 0xc1, 0xc8,
	0x23, 0x4a, 0x70, 0xa6, 0x88, 0x9f, 0x4a, 0x33, 0x3d, 0x95, 0x94, 0x7d, 0xac, 0xa6, 0x9b, 0x0a,
	0xdf, 0x81, 0xe3, 0x0b, 0x12, 0xd2, 0x98, 0x92, 0xc8, 0xe5, 0x49, 0x12, 0xb0, 0xa8, 0x18, 0xa6,
	0xd1, 0xe1, 0xb2, 0x51, 0x1c, 0x29, 0xcb, 0x66, 0xe2, 0xcb, 0x9f, 0x9e, 0x36, 0xb7, 0xf2, 0x88,
	0xae, 0xa1, 0x9f, 0x5f, 0x4e, 0x20, 0x67, 0x5a, 0xf3, 0x5d, 0xa9, 0x3d, 0xb4, 0xfa, 0x95, 0x52,
	0xbc, 0x85, 0xee, 0x61, 0xa4, 0x2a, 0x75, 0xe8, 0x68, 0x0d, 0x60, 0x3b, 0x6d, 0xe1, 0xcc, 0x60,
	0x57, 0xd9, 0x0e, 0xd1, 0x69, 0x23, 0xcc, 0xba, 0xe0, 0x16, 0xde, 0x0d, 0x0c, 0x35, 0xef, 0x81,
	0xc5, 0xfc, 0xdf, 0x4b, 0x1d, 0xd4, 0xfb, 0x35, 0xf7, 0xfa, 0xf4, 0x73, 0xfe, 0xe1, 0xd9, 0x96,
	0x11, 0xb6, 0x30, 0x7f, 0xee, 0xa0, 0x0d, 0xf9, 0x08, 0x3b, 0x7a, 0x1f, 0x62, 0x4e, 0xbe, 0x37,
	0x03, 0x7c, 0x1f, 0x98, 0xff, 0xe1, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x72, 0x09, 0x53, 0x0e,
	0x23, 0x03, 0x00, 0x00,
}
