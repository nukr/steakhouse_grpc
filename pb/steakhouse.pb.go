// Code generated by protoc-gen-go.
// source: steakhouse.proto
// DO NOT EDIT!

/*
Package steakhouse is a generated protocol buffer package.

It is generated from these files:
	steakhouse.proto

It has these top-level messages:
	Cuisine
	Menu
	Intstr
	Order
	Query
	Status
*/
package steakhouse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type Cuisine struct {
	Name        string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string            `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Price       float32           `protobuf:"fixed32,3,opt,name=price" json:"price,omitempty"`
	Id          string            `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
	IsDelete    bool              `protobuf:"varint,5,opt,name=is_delete,json=isDelete" json:"is_delete,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Cuisine) Reset()                    { *m = Cuisine{} }
func (m *Cuisine) String() string            { return proto.CompactTextString(m) }
func (*Cuisine) ProtoMessage()               {}
func (*Cuisine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Cuisine) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Menu struct {
	Name        string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Cuisines    []*Cuisine `protobuf:"bytes,3,rep,name=cuisines" json:"cuisines,omitempty"`
	Id          string     `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
}

func (m *Menu) Reset()                    { *m = Menu{} }
func (m *Menu) String() string            { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()               {}
func (*Menu) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Menu) GetCuisines() []*Cuisine {
	if m != nil {
		return m.Cuisines
	}
	return nil
}

type Intstr struct {
	Num  int32  `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Intstr) Reset()                    { *m = Intstr{} }
func (m *Intstr) String() string            { return proto.CompactTextString(m) }
func (*Intstr) ProtoMessage()               {}
func (*Intstr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Order struct {
	CreatedAt int64   `protobuf:"varint,1,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt int64   `protobuf:"varint,2,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	Discount  *Intstr `protobuf:"bytes,3,opt,name=discount" json:"discount,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Order) GetDiscount() *Intstr {
	if m != nil {
		return m.Discount
	}
	return nil
}

type Query struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Status struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Cuisine)(nil), "steakhouse.Cuisine")
	proto.RegisterType((*Menu)(nil), "steakhouse.menu")
	proto.RegisterType((*Intstr)(nil), "steakhouse.intstr")
	proto.RegisterType((*Order)(nil), "steakhouse.order")
	proto.RegisterType((*Query)(nil), "steakhouse.Query")
	proto.RegisterType((*Status)(nil), "steakhouse.status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Steakhouse service

type SteakhouseClient interface {
	GetCuisine(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Cuisine, error)
	GetCuisines(ctx context.Context, in *Query, opts ...grpc.CallOption) (Steakhouse_GetCuisinesClient, error)
	CreateCuisine(ctx context.Context, opts ...grpc.CallOption) (Steakhouse_CreateCuisineClient, error)
}

type steakhouseClient struct {
	cc *grpc.ClientConn
}

func NewSteakhouseClient(cc *grpc.ClientConn) SteakhouseClient {
	return &steakhouseClient{cc}
}

func (c *steakhouseClient) GetCuisine(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Cuisine, error) {
	out := new(Cuisine)
	err := grpc.Invoke(ctx, "/steakhouse.Steakhouse/GetCuisine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *steakhouseClient) GetCuisines(ctx context.Context, in *Query, opts ...grpc.CallOption) (Steakhouse_GetCuisinesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Steakhouse_serviceDesc.Streams[0], c.cc, "/steakhouse.Steakhouse/GetCuisines", opts...)
	if err != nil {
		return nil, err
	}
	x := &steakhouseGetCuisinesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Steakhouse_GetCuisinesClient interface {
	Recv() (*Cuisine, error)
	grpc.ClientStream
}

type steakhouseGetCuisinesClient struct {
	grpc.ClientStream
}

func (x *steakhouseGetCuisinesClient) Recv() (*Cuisine, error) {
	m := new(Cuisine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *steakhouseClient) CreateCuisine(ctx context.Context, opts ...grpc.CallOption) (Steakhouse_CreateCuisineClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Steakhouse_serviceDesc.Streams[1], c.cc, "/steakhouse.Steakhouse/CreateCuisine", opts...)
	if err != nil {
		return nil, err
	}
	x := &steakhouseCreateCuisineClient{stream}
	return x, nil
}

type Steakhouse_CreateCuisineClient interface {
	Send(*Cuisine) error
	Recv() (*Status, error)
	grpc.ClientStream
}

type steakhouseCreateCuisineClient struct {
	grpc.ClientStream
}

func (x *steakhouseCreateCuisineClient) Send(m *Cuisine) error {
	return x.ClientStream.SendMsg(m)
}

func (x *steakhouseCreateCuisineClient) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Steakhouse service

type SteakhouseServer interface {
	GetCuisine(context.Context, *Query) (*Cuisine, error)
	GetCuisines(*Query, Steakhouse_GetCuisinesServer) error
	CreateCuisine(Steakhouse_CreateCuisineServer) error
}

func RegisterSteakhouseServer(s *grpc.Server, srv SteakhouseServer) {
	s.RegisterService(&_Steakhouse_serviceDesc, srv)
}

func _Steakhouse_GetCuisine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SteakhouseServer).GetCuisine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/steakhouse.Steakhouse/GetCuisine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SteakhouseServer).GetCuisine(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Steakhouse_GetCuisines_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SteakhouseServer).GetCuisines(m, &steakhouseGetCuisinesServer{stream})
}

type Steakhouse_GetCuisinesServer interface {
	Send(*Cuisine) error
	grpc.ServerStream
}

type steakhouseGetCuisinesServer struct {
	grpc.ServerStream
}

func (x *steakhouseGetCuisinesServer) Send(m *Cuisine) error {
	return x.ServerStream.SendMsg(m)
}

func _Steakhouse_CreateCuisine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SteakhouseServer).CreateCuisine(&steakhouseCreateCuisineServer{stream})
}

type Steakhouse_CreateCuisineServer interface {
	Send(*Status) error
	Recv() (*Cuisine, error)
	grpc.ServerStream
}

type steakhouseCreateCuisineServer struct {
	grpc.ServerStream
}

func (x *steakhouseCreateCuisineServer) Send(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *steakhouseCreateCuisineServer) Recv() (*Cuisine, error) {
	m := new(Cuisine)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Steakhouse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "steakhouse.Steakhouse",
	HandlerType: (*SteakhouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCuisine",
			Handler:    _Steakhouse_GetCuisine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCuisines",
			Handler:       _Steakhouse_GetCuisines_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateCuisine",
			Handler:       _Steakhouse_CreateCuisine_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("steakhouse.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe5, 0x74, 0x2d, 0xed, 0x1b, 0x36, 0x15, 0x33, 0x41, 0x54, 0x40, 0x0a, 0x39, 0x45,
	0x1c, 0x92, 0x51, 0x2e, 0x68, 0x88, 0xc3, 0x34, 0x10, 0xa7, 0x21, 0x11, 0xb8, 0x4f, 0x26, 0x7e,
	0x55, 0xac, 0x35, 0x76, 0x14, 0xdb, 0x93, 0x2a, 0xb4, 0x03, 0x7c, 0x05, 0x3e, 0x1a, 0x5f, 0x81,
	0xcf, 0x81, 0x50, 0x1c, 0x37, 0xcd, 0x50, 0x0f, 0x88, 0xdb, 0xeb, 0xf7, 0xef, 0xef, 0x79, 0xda,
	0xc0, 0x5c, 0x1b, 0x64, 0x57, 0x5f, 0x94, 0xd5, 0x98, 0xd5, 0x8d, 0x32, 0x8a, 0xc2, 0x2e, 0xb3,
	0x78, 0xbc, 0x52, 0x6a, 0xb5, 0xc6, 0x9c, 0xd5, 0x22, 0x67, 0x52, 0x2a, 0xc3, 0x8c, 0x50, 0x52,
	0x77, 0x9d, 0xc9, 0x6f, 0x02, 0x77, 0xce, 0xad, 0xd0, 0x42, 0x22, 0xa5, 0x70, 0x20, 0x59, 0x85,
	0x11, 0x89, 0x49, 0x3a, 0x2b, 0x5c, 0x4c, 0x63, 0x08, 0x39, 0xea, 0xb2, 0x11, 0x75, 0x3b, 0x15,
	0x05, 0xae, 0x34, 0x4c, 0xd1, 0x63, 0x18, 0xd7, 0x8d, 0x28, 0x31, 0x1a, 0xc5, 0x24, 0x0d, 0x8a,
	0xee, 0x41, 0x8f, 0x20, 0x10, 0x3c, 0x3a, 0x70, 0xed, 0x81, 0xe0, 0xf4, 0x11, 0xcc, 0x84, 0xbe,
	0xe4, 0xb8, 0x46, 0x83, 0xd1, 0x38, 0x26, 0xe9, 0xb4, 0x98, 0x0a, 0xfd, 0xc6, 0xbd, 0xe9, 0x6b,
	0x98, 0x56, 0x68, 0x18, 0x67, 0x86, 0x45, 0x93, 0x78, 0x94, 0x86, 0xcb, 0xa7, 0xd9, 0x40, 0x93,
	0xe7, 0xcb, 0x2e, 0x7c, 0xcf, 0x5b, 0x69, 0x9a, 0x4d, 0xd1, 0x8f, 0x2c, 0x5e, 0xc1, 0xe1, 0xad,
	0x12, 0x9d, 0xc3, 0xe8, 0x0a, 0x37, 0x5e, 0x47, 0x1b, 0xb6, 0x90, 0xd7, 0x6c, 0x6d, 0xd1, 0x0b,
	0xe8, 0x1e, 0xa7, 0xc1, 0x4b, 0x92, 0xdc, 0xc0, 0x41, 0x85, 0xd2, 0xfe, 0xa7, 0xf8, 0x1c, 0xa6,
	0x65, 0x47, 0xa7, 0xa3, 0x91, 0x23, 0xbf, 0xbf, 0x87, 0xbc, 0xe8, 0x9b, 0xfe, 0xf6, 0x25, 0xc9,
	0x60, 0x22, 0xa4, 0xd1, 0xa6, 0x69, 0xa1, 0xa5, 0xad, 0xdc, 0xfd, 0x71, 0xd1, 0x86, 0x3d, 0x52,
	0xb0, 0x43, 0x4a, 0x2c, 0x8c, 0x55, 0xc3, 0xb1, 0xa1, 0x4f, 0x00, 0xca, 0x06, 0x99, 0x41, 0x7e,
	0xc9, 0x8c, 0x9b, 0x1a, 0x15, 0x33, 0x9f, 0x39, 0x33, 0x6d, 0xd9, 0xd6, 0x7c, 0x5b, 0x0e, 0xba,
	0xb2, 0xcf, 0x9c, 0x19, 0x9a, 0xc1, 0x94, 0x0b, 0x5d, 0x2a, 0x2b, 0x8d, 0xfb, 0xdd, 0xc2, 0x25,
	0x1d, 0x72, 0x77, 0x48, 0x45, 0xdf, 0x93, 0x3c, 0x84, 0xf1, 0x07, 0x8b, 0xcd, 0xc6, 0xf3, 0x93,
	0x21, 0xbf, 0x36, 0xcc, 0x58, 0xdd, 0xd2, 0x96, 0x8a, 0xa3, 0x17, 0xe0, 0xe2, 0x56, 0x53, 0xa5,
	0x57, 0x5e, 0x40, 0x1b, 0x2e, 0xbf, 0x05, 0x00, 0x1f, 0xfb, 0x43, 0xf4, 0x3d, 0xc0, 0x3b, 0x34,
	0xdb, 0x3f, 0xe0, 0xbd, 0x21, 0x83, 0xbb, 0xb7, 0xd8, 0x67, 0x67, 0x12, 0x7d, 0xff, 0xf9, 0xeb,
	0x47, 0x40, 0xe9, 0x3c, 0xbf, 0x7e, 0x9e, 0x7b, 0x6b, 0xf3, 0xaf, 0x82, 0xdf, 0xd0, 0x0b, 0x08,
	0x77, 0xfb, 0xf4, 0x3f, 0x2f, 0x3c, 0x76, 0x0b, 0x8f, 0xe8, 0xdd, 0xc1, 0x42, 0x7d, 0x42, 0xe8,
	0x27, 0x38, 0x3c, 0x77, 0x96, 0x6e, 0x09, 0xf7, 0x4d, 0x2f, 0x6e, 0x59, 0xd7, 0xb9, 0x91, 0x3c,
	0x70, 0x1b, 0xe7, 0x49, 0x38, 0xd8, 0x78, 0x4a, 0x9e, 0xa5, 0xe4, 0x84, 0x7c, 0x9e, 0xb8, 0x4f,
	0xef, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x40, 0x04, 0x3a, 0xb8, 0x03, 0x00, 0x00,
}
