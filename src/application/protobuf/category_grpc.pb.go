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

// CategoryClient is the client API for Category service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryClient interface {
	Create(ctx context.Context, in *CategoryCreate, opts ...grpc.CallOption) (*protobuf.Response, error)
	Edit(ctx context.Context, in *CategoryEdit, opts ...grpc.CallOption) (*protobuf.Response, error)
	Delete(ctx context.Context, in *CategoryDelete, opts ...grpc.CallOption) (*protobuf.Response, error)
	GetTree(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*protobuf.Response, error)
}

type categoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryClient(cc grpc.ClientConnInterface) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) Create(ctx context.Context, in *CategoryCreate, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Category/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) Edit(ctx context.Context, in *CategoryEdit, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Category/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) Delete(ctx context.Context, in *CategoryDelete, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Category/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetTree(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*protobuf.Response, error) {
	out := new(protobuf.Response)
	err := c.cc.Invoke(ctx, "/protobuf.Category/GetTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
// All implementations must embed UnimplementedCategoryServer
// for forward compatibility
type CategoryServer interface {
	Create(context.Context, *CategoryCreate) (*protobuf.Response, error)
	Edit(context.Context, *CategoryEdit) (*protobuf.Response, error)
	Delete(context.Context, *CategoryDelete) (*protobuf.Response, error)
	GetTree(context.Context, *protobuf.Empty) (*protobuf.Response, error)
	mustEmbedUnimplementedCategoryServer()
}

// UnimplementedCategoryServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServer struct {
}

func (UnimplementedCategoryServer) Create(context.Context, *CategoryCreate) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryServer) Edit(context.Context, *CategoryEdit) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedCategoryServer) Delete(context.Context, *CategoryDelete) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoryServer) GetTree(context.Context, *protobuf.Empty) (*protobuf.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTree not implemented")
}
func (UnimplementedCategoryServer) mustEmbedUnimplementedCategoryServer() {}

// UnsafeCategoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServer will
// result in compilation errors.
type UnsafeCategoryServer interface {
	mustEmbedUnimplementedCategoryServer()
}

func RegisterCategoryServer(s grpc.ServiceRegistrar, srv CategoryServer) {
	s.RegisterService(&_Category_serviceDesc, srv)
}

func _Category_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Category/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).Create(ctx, req.(*CategoryCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryEdit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Category/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).Edit(ctx, req.(*CategoryEdit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Category/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).Delete(ctx, req.(*CategoryDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Category/GetTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetTree(ctx, req.(*protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Category_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Category_Create_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _Category_Edit_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Category_Delete_Handler,
		},
		{
			MethodName: "GetTree",
			Handler:    _Category_GetTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}
