// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	protobuf "github.com/Etpmls/Etpmls-Micro/v2/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// VariableClient is the client API for Variable service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VariableClient interface {
	Create(ctx context.Context, in *VariableCreate, opts ...grpc.CallOption) (*protobuf.Response, error)
	GetAll(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*protobuf.Response, error)
}

type variableClient struct {
	cc grpc.ClientConnInterface
}

func NewVariableClient(cc grpc.ClientConnInterface) VariableClient {
	return &variableClient{cc}
}

func (c *variableClient) Create(ctx context.Context, in *VariableCreate, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Variable/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variableClient) GetAll(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Variable/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VariableServer is the server API for Variable service.
// All implementations must embed UnimplementedVariableServer
// for forward compatibility
type VariableServer interface {
	Create(context.Context, *VariableCreate) (*protobuf.Response, error)
	GetAll(context.Context, *protobuf.Empty) (*protobuf.Response, error)
	mustEmbedUnimplementedVariableServer()
}

// UnimplementedVariableServer must be embedded to have forward compatible implementations.
type UnimplementedVariableServer struct {
}

func (UnimplementedVariableServer) Create(context.Context, *VariableCreate) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVariableServer) GetAll(context.Context, *protobuf.Empty) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedVariableServer) mustEmbedUnimplementedVariableServer() {}

// UnsafeVariableServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VariableServer will
// result in compilation errors.
type UnsafeVariableServer interface {
	mustEmbedUnimplementedVariableServer()
}

func RegisterVariableServer(s grpc.ServiceRegistrar, srv VariableServer) {
	s.RegisterService(&_Variable_serviceDesc, srv)
}

func _Variable_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VariableCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Variable/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServer).Create(ctx, req.(*VariableCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Variable_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VariableServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Variable/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VariableServer).GetAll(ctx, req.(*protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Variable_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Variable",
	HandlerType: (*VariableServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Variable_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Variable_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "variable.proto",
}
