// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atlas.proto

package atlas

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	message "github.com/greenmochi/ultimate/services/atlas/proto/atlas/message"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x59, 0xb1, 0x4a, 0xb3, 0x97, 0x1a, 0x41, 0xcb, 0xa2, 0x20, 0x01, 0x51, 0x7a, 0x48,
	0x40, 0x4f, 0xd6, 0x93, 0x22, 0x78, 0x2d, 0xbd, 0xe9, 0x2d, 0x5d, 0x86, 0x74, 0x20, 0x9b, 0xc4,
	0xcd, 0xac, 0xb0, 0x57, 0x5f, 0xc1, 0x97, 0xf0, 0x7d, 0x7c, 0x05, 0x1f, 0x44, 0xf6, 0x1f, 0xac,
	0x48, 0x6f, 0xf9, 0x26, 0xdf, 0xfc, 0xbe, 0x8f, 0x61, 0xa9, 0x26, 0xab, 0xa3, 0x0c, 0xa5, 0x27,
	0xcf, 0x27, 0xad, 0xc8, 0xce, 0x8c, 0xf7, 0xc6, 0x82, 0xd2, 0x01, 0x95, 0x76, 0xce, 0x93, 0x26,
	0xf4, 0xae, 0x37, 0x65, 0xc7, 0x05, 0xc4, 0xa8, 0x0d, 0xa8, 0xd1, 0xa6, 0xb8, 0x62, 0xe9, 0x0a,
	0x9d, 0x59, 0xc3, 0x5b, 0x05, 0x91, 0xf8, 0x9c, 0x1d, 0xf6, 0xae, 0x79, 0x72, 0x91, 0x5c, 0x4f,
	0xd7, 0x83, 0x14, 0x97, 0x6c, 0xda, 0x19, 0x83, 0xad, 0x77, 0xdb, 0x6e, 0xbe, 0x12, 0x36, 0x79,
	0x68, 0xf8, 0xfc, 0x89, 0xed, 0x37, 0x0b, 0x9c, 0xcb, 0x2e, 0x6f, 0x14, 0x93, 0xcd, 0xfe, 0xcc,
	0x82, 0xad, 0xc5, 0xc9, 0xc7, 0xf7, 0xcf, 0xe7, 0xde, 0x4c, 0xa4, 0x5d, 0x3b, 0x15, 0xd0, 0x99,
	0x65, 0xb2, 0xe0, 0x2f, 0x2c, 0x7d, 0x06, 0x5a, 0x59, 0x5d, 0x5b, 0x6c, 0xfa, 0xc9, 0x3e, 0x48,
	0x0e, 0xa3, 0x01, 0x79, 0xf4, 0xef, 0x47, 0x9c, 0xb7, 0xcc, 0x53, 0xc1, 0x7b, 0xe6, 0x08, 0xb4,
	0x4c, 0x16, 0x8f, 0xf7, 0xaf, 0x77, 0x06, 0x69, 0x5b, 0x6d, 0x64, 0xee, 0x0b, 0x65, 0x4a, 0x00,
	0x57, 0xf8, 0x7c, 0x8b, 0xaa, 0xb2, 0x84, 0x85, 0x26, 0x50, 0x11, 0xca, 0x77, 0xcc, 0x21, 0x0e,
	0x9d, 0x9a, 0x8b, 0x75, 0xef, 0xcd, 0x41, 0x2b, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16,
	0x43, 0x3e, 0xaa, 0x87, 0x01, 0x00, 0x00,
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
	GetPlaylist(ctx context.Context, in *message.PlaylistRequest, opts ...grpc.CallOption) (*message.Playlist, error)
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

func (c *atlasClient) GetPlaylist(ctx context.Context, in *message.PlaylistRequest, opts ...grpc.CallOption) (*message.Playlist, error) {
	out := new(message.Playlist)
	err := c.cc.Invoke(ctx, "/atlas.Atlas/GetPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AtlasServer is the server API for Atlas service.
type AtlasServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetPlaylist(context.Context, *message.PlaylistRequest) (*message.Playlist, error)
}

// UnimplementedAtlasServer can be embedded to have forward compatible implementations.
type UnimplementedAtlasServer struct {
}

func (*UnimplementedAtlasServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAtlasServer) GetPlaylist(ctx context.Context, req *message.PlaylistRequest) (*message.Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylist not implemented")
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

func _Atlas_GetPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.PlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AtlasServer).GetPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atlas.Atlas/GetPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AtlasServer).GetPlaylist(ctx, req.(*message.PlaylistRequest))
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
		{
			MethodName: "GetPlaylist",
			Handler:    _Atlas_GetPlaylist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atlas.proto",
}
