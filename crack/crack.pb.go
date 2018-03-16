// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crack.proto

/*
Package crack is a generated protocol buffer package.

It is generated from these files:
	crack.proto

It has these top-level messages:
	CrackRequest
	CrackResponse
*/
package crack

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CrackRequest struct {
	UserID  int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	TokenID int64 `protobuf:"varint,2,opt,name=TokenID" json:"TokenID,omitempty"`
}

func (m *CrackRequest) Reset()                    { *m = CrackRequest{} }
func (m *CrackRequest) String() string            { return proto.CompactTextString(m) }
func (*CrackRequest) ProtoMessage()               {}
func (*CrackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CrackRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CrackRequest) GetTokenID() int64 {
	if m != nil {
		return m.TokenID
	}
	return 0
}

type CrackResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *CrackResponse) Reset()                    { *m = CrackResponse{} }
func (m *CrackResponse) String() string            { return proto.CompactTextString(m) }
func (*CrackResponse) ProtoMessage()               {}
func (*CrackResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CrackResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CrackResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CrackRequest)(nil), "crack.CrackRequest")
	proto.RegisterType((*CrackResponse)(nil), "crack.CrackResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Crack service

type CrackClient interface {
	CrackToken(ctx context.Context, in *CrackRequest, opts ...grpc.CallOption) (*CrackResponse, error)
}

type crackClient struct {
	cc *grpc.ClientConn
}

func NewCrackClient(cc *grpc.ClientConn) CrackClient {
	return &crackClient{cc}
}

func (c *crackClient) CrackToken(ctx context.Context, in *CrackRequest, opts ...grpc.CallOption) (*CrackResponse, error) {
	out := new(CrackResponse)
	err := grpc.Invoke(ctx, "/crack.Crack/CrackToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Crack service

type CrackServer interface {
	CrackToken(context.Context, *CrackRequest) (*CrackResponse, error)
}

func RegisterCrackServer(s *grpc.Server, srv CrackServer) {
	s.RegisterService(&_Crack_serviceDesc, srv)
}

func _Crack_CrackToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrackServer).CrackToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crack.Crack/CrackToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrackServer).CrackToken(ctx, req.(*CrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Crack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crack.Crack",
	HandlerType: (*CrackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrackToken",
			Handler:    _Crack_CrackToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crack.proto",
}

func init() { proto.RegisterFile("crack.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2e, 0x4a, 0x4c,
	0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x1c, 0xb8, 0x78, 0x9c,
	0x41, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x31, 0x2e, 0xb6, 0xd0, 0xe2, 0xd4,
	0x22, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x28, 0x4f, 0x48, 0x82, 0x8b, 0x3d,
	0x24, 0x3f, 0x3b, 0x35, 0xcf, 0xd3, 0x45, 0x82, 0x09, 0x2c, 0x01, 0xe3, 0x2a, 0x39, 0x73, 0xf1,
	0x42, 0x4d, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05, 0x29, 0x0d, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d,
	0x2e, 0x06, 0x9b, 0xc1, 0x11, 0x04, 0xe3, 0x82, 0x64, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53,
	0xc1, 0x86, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0x4e, 0x5c, 0xac, 0x60, 0x43, 0x84, 0x2c, 0xb9, 0xb8,
	0xc0, 0x0c, 0xb0, 0xe9, 0x42, 0xc2, 0x7a, 0x10, 0x27, 0x23, 0x3b, 0x51, 0x4a, 0x04, 0x55, 0x10,
	0x62, 0xab, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x63, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75,
	0xc8, 0x83, 0x8d, 0xe7, 0x00, 0x00, 0x00,
}
