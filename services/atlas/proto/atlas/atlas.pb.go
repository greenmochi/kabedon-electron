// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atlas.proto

package atlas

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/greenmochi/ultimate/services/atlas/proto/atlas/message"
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

type PingRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ebe89c2f4b7977, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ebe89c2f4b7977, []int{1}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

func (m *PingReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "atlas.PingRequest")
	proto.RegisterType((*PingReply)(nil), "atlas.PingReply")
}

func init() { proto.RegisterFile("atlas.proto", fileDescriptor_e8ebe89c2f4b7977) }

var fileDescriptor_e8ebe89c2f4b7977 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0xc9, 0x49,
	0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x84, 0x73, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0xf5, 0x91, 0xe4, 0x94, 0xd4, 0xb9, 0xb8, 0x03, 0x32, 0xf3, 0xd2, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0xa1, 0xaa, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x55, 0x2e, 0x4e, 0x88, 0xc2, 0x82, 0x9c, 0x4a, 0xdc, 0xca,
	0x8c, 0x4c, 0xb9, 0x58, 0x1d, 0x41, 0xc6, 0x0b, 0xe9, 0x70, 0xb1, 0x80, 0xd4, 0x0b, 0x09, 0xe9,
	0x41, 0xac, 0x43, 0xb2, 0x45, 0x4a, 0x00, 0x45, 0xac, 0x20, 0xa7, 0xd2, 0xc9, 0x3a, 0xca, 0x32,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xbd, 0x28, 0x35, 0x35, 0x2f,
	0x37, 0x3f, 0x39, 0x23, 0x53, 0xbf, 0x34, 0xa7, 0x24, 0x33, 0x37, 0xb1, 0x24, 0x55, 0xbf, 0x38,
	0xb5, 0xa8, 0x2c, 0x33, 0x39, 0xb5, 0x18, 0xe2, 0x7a, 0x7d, 0xb0, 0xeb, 0x21, 0xec, 0x24, 0x36,
	0x30, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x85, 0xc3, 0xb3, 0x05, 0xf5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AtlasClient is the client API for Atlas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AtlasClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type atlasClient struct {
	cc *grpc.ClientConn
}

func NewAtlasClient(cc *grpc.ClientConn) AtlasClient {
	return &atlasClient{cc}
}

func (c *atlasClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/atlas.Atlas/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AtlasServer is the server API for Atlas service.
type AtlasServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

// UnimplementedAtlasServer can be embedded to have forward compatible implementations.
type UnimplementedAtlasServer struct {
}

func (*UnimplementedAtlasServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterAtlasServer(s *grpc.Server, srv AtlasServer) {
	s.RegisterService(&_Atlas_serviceDesc, srv)
}

func _Atlas_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtlasServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.Atlas/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtlasServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Atlas_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atlas.Atlas",
	HandlerType: (*AtlasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Atlas_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atlas.proto",
}
