// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myanimelist.proto

package myanimelist

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	message "github.com/greenmochi/ultimate/services/gateway/proto/myanimelist/message"
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
	return fileDescriptor_7365bf8e2ee52b0c, []int{0}
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
	return fileDescriptor_7365bf8e2ee52b0c, []int{1}
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
	proto.RegisterType((*PingRequest)(nil), "myanimelist.PingRequest")
	proto.RegisterType((*PingReply)(nil), "myanimelist.PingReply")
}

func init() { proto.RegisterFile("myanimelist.proto", fileDescriptor_7365bf8e2ee52b0c) }

var fileDescriptor_7365bf8e2ee52b0c = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x89, 0x8a, 0x72, 0x27, 0x72, 0xb9, 0x1d, 0xe4, 0x92, 0xa6, 0xad, 0x94, 0xf1, 0x5f,
	0xe9, 0x22, 0x03, 0xba, 0xeb, 0xae, 0xb5, 0x50, 0x0a, 0x15, 0xb4, 0x62, 0x17, 0xe2, 0x66, 0x1a,
	0x0e, 0xd3, 0xc1, 0x64, 0x26, 0x66, 0x26, 0x4a, 0xb6, 0xbe, 0x82, 0x8f, 0xe6, 0x2b, 0xb8, 0xf3,
	0x25, 0x64, 0x26, 0xb5, 0x9d, 0xb4, 0xb7, 0xcb, 0xf3, 0x3b, 0x5f, 0x7e, 0xdf, 0x09, 0x0c, 0xea,
	0xe4, 0x35, 0x93, 0x22, 0x87, 0x4c, 0x68, 0x93, 0x14, 0xa5, 0x32, 0x0a, 0x87, 0x1e, 0x8a, 0xfb,
	0x5c, 0x29, 0x9e, 0x01, 0x65, 0x85, 0xa0, 0x4c, 0x4a, 0x65, 0x98, 0x11, 0x4a, 0xea, 0x26, 0x1a,
	0x77, 0x73, 0xd0, 0x9a, 0x71, 0xa0, 0x67, 0x16, 0xf2, 0x0a, 0x85, 0xef, 0x85, 0xe4, 0x6b, 0xf8,
	0x56, 0x81, 0x36, 0x38, 0x42, 0x8f, 0xf6, 0xd9, 0x28, 0x18, 0x06, 0xa3, 0xab, 0xf5, 0xff, 0x91,
	0xbc, 0x40, 0x57, 0x4d, 0xb0, 0xc8, 0xea, 0xcb, 0xb1, 0xd7, 0x7f, 0xef, 0xa3, 0xf0, 0x5d, 0x3d,
	0xb5, 0x2d, 0x2b, 0xa1, 0x0d, 0xde, 0xa0, 0x07, 0xf6, 0x33, 0x1c, 0x25, 0x7e, 0xb7, 0x57, 0x19,
	0xdf, 0xde, 0xb1, 0x29, 0xb2, 0x9a, 0xf4, 0x7f, 0xfe, 0xfe, 0xf3, 0xeb, 0xde, 0x2d, 0xe9, 0xf8,
	0x57, 0xd3, 0x42, 0x48, 0x3e, 0x09, 0xc6, 0x98, 0xa3, 0x9b, 0x05, 0x98, 0x4f, 0x1a, 0xca, 0x63,
	0x57, 0x27, 0xd9, 0x5f, 0x91, 0x58, 0x2e, 0x59, 0x0e, 0x56, 0xee, 0xa1, 0x43, 0x94, 0x8c, 0x9c,
	0x9c, 0x90, 0x41, 0x4b, 0x7e, 0x6a, 0xb4, 0x45, 0x80, 0xc2, 0x8f, 0xc0, 0xca, 0x74, 0xe7, 0x28,
	0x7e, 0x72, 0x10, 0x36, 0xf4, 0x43, 0x05, 0x65, 0x1d, 0xf7, 0x0e, 0xd4, 0xa5, 0x9a, 0xd5, 0x1a,
	0x74, 0x95, 0x19, 0x4d, 0x9e, 0xb9, 0xae, 0x01, 0x89, 0x5a, 0x5d, 0x9e, 0xd4, 0xd6, 0x7c, 0x41,
	0xd7, 0x0b, 0x30, 0x6e, 0x9c, 0xd5, 0x2b, 0x21, 0xbf, 0x62, 0xdc, 0x76, 0x5a, 0x16, 0x5f, 0xb7,
	0x19, 0x79, 0xe9, 0xd4, 0x43, 0xd2, 0x3b, 0xfd, 0x0d, 0x4f, 0x64, 0xed, 0x1b, 0xf4, 0xf8, 0x08,
	0x97, 0x73, 0x7c, 0xd3, 0xf6, 0x2c, 0xe7, 0x67, 0xe6, 0xe7, 0xce, 0xfc, 0x94, 0x74, 0x2f, 0x98,
	0x97, 0xf3, 0x49, 0x30, 0x9e, 0xbd, 0xfd, 0x3c, 0xe5, 0xc2, 0xec, 0xaa, 0x6d, 0x92, 0xaa, 0x9c,
	0xf2, 0x12, 0x40, 0xe6, 0x2a, 0xdd, 0x09, 0x5a, 0x65, 0x46, 0xe4, 0xcc, 0x00, 0xd5, 0x50, 0x7e,
	0x17, 0x29, 0x68, 0xca, 0x99, 0x81, 0x1f, 0xac, 0xa6, 0xee, 0xd9, 0xf9, 0xd2, 0xed, 0x43, 0x87,
	0xde, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x3b, 0x55, 0x68, 0xe4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MyAnimeListClient is the client API for MyAnimeList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyAnimeListClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetUserAnimeList(ctx context.Context, in *message.Username, opts ...grpc.CallOption) (*message.UserAnimeList, error)
	SearchAnime(ctx context.Context, in *message.SearchQuery, opts ...grpc.CallOption) (*message.AnimeSearchResults, error)
	GetAnimeByLink(ctx context.Context, in *message.AnimeLink, opts ...grpc.CallOption) (*message.Anime, error)
	GetAnimeByID(ctx context.Context, in *message.AnimeID, opts ...grpc.CallOption) (*message.Anime, error)
}

type myAnimeListClient struct {
	cc *grpc.ClientConn
}

func NewMyAnimeListClient(cc *grpc.ClientConn) MyAnimeListClient {
	return &myAnimeListClient{cc}
}

func (c *myAnimeListClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/myanimelist.MyAnimeList/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAnimeListClient) GetUserAnimeList(ctx context.Context, in *message.Username, opts ...grpc.CallOption) (*message.UserAnimeList, error) {
	out := new(message.UserAnimeList)
	err := c.cc.Invoke(ctx, "/myanimelist.MyAnimeList/GetUserAnimeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAnimeListClient) SearchAnime(ctx context.Context, in *message.SearchQuery, opts ...grpc.CallOption) (*message.AnimeSearchResults, error) {
	out := new(message.AnimeSearchResults)
	err := c.cc.Invoke(ctx, "/myanimelist.MyAnimeList/SearchAnime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAnimeListClient) GetAnimeByLink(ctx context.Context, in *message.AnimeLink, opts ...grpc.CallOption) (*message.Anime, error) {
	out := new(message.Anime)
	err := c.cc.Invoke(ctx, "/myanimelist.MyAnimeList/GetAnimeByLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myAnimeListClient) GetAnimeByID(ctx context.Context, in *message.AnimeID, opts ...grpc.CallOption) (*message.Anime, error) {
	out := new(message.Anime)
	err := c.cc.Invoke(ctx, "/myanimelist.MyAnimeList/GetAnimeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyAnimeListServer is the server API for MyAnimeList service.
type MyAnimeListServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetUserAnimeList(context.Context, *message.Username) (*message.UserAnimeList, error)
	SearchAnime(context.Context, *message.SearchQuery) (*message.AnimeSearchResults, error)
	GetAnimeByLink(context.Context, *message.AnimeLink) (*message.Anime, error)
	GetAnimeByID(context.Context, *message.AnimeID) (*message.Anime, error)
}

// UnimplementedMyAnimeListServer can be embedded to have forward compatible implementations.
type UnimplementedMyAnimeListServer struct {
}

func (*UnimplementedMyAnimeListServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMyAnimeListServer) GetUserAnimeList(ctx context.Context, req *message.Username) (*message.UserAnimeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAnimeList not implemented")
}
func (*UnimplementedMyAnimeListServer) SearchAnime(ctx context.Context, req *message.SearchQuery) (*message.AnimeSearchResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAnime not implemented")
}
func (*UnimplementedMyAnimeListServer) GetAnimeByLink(ctx context.Context, req *message.AnimeLink) (*message.Anime, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimeByLink not implemented")
}
func (*UnimplementedMyAnimeListServer) GetAnimeByID(ctx context.Context, req *message.AnimeID) (*message.Anime, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimeByID not implemented")
}

func RegisterMyAnimeListServer(s *grpc.Server, srv MyAnimeListServer) {
	s.RegisterService(&_MyAnimeList_serviceDesc, srv)
}

func _MyAnimeList_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyAnimeListServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myanimelist.MyAnimeList/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyAnimeListServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyAnimeList_GetUserAnimeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyAnimeListServer).GetUserAnimeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myanimelist.MyAnimeList/GetUserAnimeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyAnimeListServer).GetUserAnimeList(ctx, req.(*message.Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyAnimeList_SearchAnime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.SearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyAnimeListServer).SearchAnime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myanimelist.MyAnimeList/SearchAnime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyAnimeListServer).SearchAnime(ctx, req.(*message.SearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyAnimeList_GetAnimeByLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AnimeLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyAnimeListServer).GetAnimeByLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myanimelist.MyAnimeList/GetAnimeByLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyAnimeListServer).GetAnimeByLink(ctx, req.(*message.AnimeLink))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyAnimeList_GetAnimeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.AnimeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyAnimeListServer).GetAnimeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myanimelist.MyAnimeList/GetAnimeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyAnimeListServer).GetAnimeByID(ctx, req.(*message.AnimeID))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyAnimeList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myanimelist.MyAnimeList",
	HandlerType: (*MyAnimeListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _MyAnimeList_Ping_Handler,
		},
		{
			MethodName: "GetUserAnimeList",
			Handler:    _MyAnimeList_GetUserAnimeList_Handler,
		},
		{
			MethodName: "SearchAnime",
			Handler:    _MyAnimeList_SearchAnime_Handler,
		},
		{
			MethodName: "GetAnimeByLink",
			Handler:    _MyAnimeList_GetAnimeByLink_Handler,
		},
		{
			MethodName: "GetAnimeByID",
			Handler:    _MyAnimeList_GetAnimeByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myanimelist.proto",
}
