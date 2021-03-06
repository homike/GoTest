// Code generated by protoc-gen-go. DO NOT EDIT.
// source: etcd.proto

package pb

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

type PingReq struct {
	ReqInfo string `protobuf:"bytes,1,opt,name=reqInfo,proto3" json:"reqInfo,omitempty" redis:"reqInfo"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef3728b01153dcf4, []int{0}
}

func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (m *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(m, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

func (m *PingReq) GetReqInfo() string {
	if m != nil {
		return m.ReqInfo
	}
	return ""
}

type PingAck struct {
	ReqInfo string `protobuf:"bytes,1,opt,name=reqInfo,proto3" json:"reqInfo,omitempty" redis:"reqInfo"`
	AckInfo string `protobuf:"bytes,2,opt,name=ackInfo,proto3" json:"ackInfo,omitempty" redis:"ackInfo"`
}

func (m *PingAck) Reset()         { *m = PingAck{} }
func (m *PingAck) String() string { return proto.CompactTextString(m) }
func (*PingAck) ProtoMessage()    {}
func (*PingAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef3728b01153dcf4, []int{1}
}

func (m *PingAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingAck.Unmarshal(m, b)
}
func (m *PingAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingAck.Marshal(b, m, deterministic)
}
func (m *PingAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingAck.Merge(m, src)
}
func (m *PingAck) XXX_Size() int {
	return xxx_messageInfo_PingAck.Size(m)
}
func (m *PingAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PingAck.DiscardUnknown(m)
}

var xxx_messageInfo_PingAck proto.InternalMessageInfo

func (m *PingAck) GetReqInfo() string {
	if m != nil {
		return m.ReqInfo
	}
	return ""
}

func (m *PingAck) GetAckInfo() string {
	if m != nil {
		return m.AckInfo
	}
	return ""
}

func init() {
	proto.RegisterType((*PingReq)(nil), "pb.PingReq")
	proto.RegisterType((*PingAck)(nil), "pb.PingAck")
}

func init() { proto.RegisterFile("etcd.proto", fileDescriptor_ef3728b01153dcf4) }

var fileDescriptor_ef3728b01153dcf4 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2d, 0x49, 0x4e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe6, 0x62, 0x0f, 0xc8,
	0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x14, 0x92, 0xe0, 0x62, 0x2f, 0x4a, 0x2d, 0xf4, 0xcc, 0x4b, 0xcb,
	0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0x6c, 0x21, 0x8a, 0x1c, 0x93, 0xb3,
	0x71, 0x2b, 0x02, 0xc9, 0x24, 0x26, 0x67, 0x83, 0x65, 0x98, 0x20, 0x32, 0x50, 0xae, 0x91, 0x21,
	0x17, 0xb7, 0x6b, 0x49, 0x72, 0x4a, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x12, 0x17,
	0x0b, 0xc8, 0x34, 0x21, 0x6e, 0xbd, 0x82, 0x24, 0x3d, 0xa8, 0xe5, 0x52, 0x70, 0x8e, 0x63, 0x72,
	0xb6, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x85, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x54,
	0x7a, 0xde, 0xaf, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EtcdServiceClient is the client API for EtcdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EtcdServiceClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingAck, error)
}

type etcdServiceClient struct {
	cc *grpc.ClientConn
}

func NewEtcdServiceClient(cc *grpc.ClientConn) EtcdServiceClient {
	return &etcdServiceClient{cc}
}

func (c *etcdServiceClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingAck, error) {
	out := new(PingAck)
	err := c.cc.Invoke(ctx, "/pb.EtcdService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EtcdServiceServer is the server API for EtcdService service.
type EtcdServiceServer interface {
	Ping(context.Context, *PingReq) (*PingAck, error)
}

// UnimplementedEtcdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEtcdServiceServer struct {
}

func (*UnimplementedEtcdServiceServer) Ping(ctx context.Context, req *PingReq) (*PingAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterEtcdServiceServer(s *grpc.Server, srv EtcdServiceServer) {
	s.RegisterService(&_EtcdService_serviceDesc, srv)
}

func _EtcdService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EtcdServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EtcdService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EtcdServiceServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EtcdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EtcdService",
	HandlerType: (*EtcdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _EtcdService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "etcd.proto",
}

func Get_EtcdService_serviceDesc() *grpc.ServiceDesc {
	return &_EtcdService_serviceDesc
}
