// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: artist/artistProto/artist.proto

package artistProto

import (
	context "context"
	gatewayProto "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArtistUseCaseClient is the client API for ArtistUseCase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistUseCaseClient interface {
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArtistsResponse, error)
	GetLastId(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*gatewayProto.IntResponse, error)
	Create(ctx context.Context, in *Artist, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *Artist, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *gatewayProto.IdArg, opts ...grpc.CallOption) (*empty.Empty, error)
	GetById(ctx context.Context, in *gatewayProto.IdArg, opts ...grpc.CallOption) (*ArtistDataTransfer, error)
	GetPopular(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArtistsResponse, error)
	GetSize(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*gatewayProto.IntResponse, error)
	SearchByName(ctx context.Context, in *gatewayProto.StringArg, opts ...grpc.CallOption) (*ArtistsResponse, error)
	GetFavorites(ctx context.Context, in *gatewayProto.IdArg, opts ...grpc.CallOption) (*ArtistsResponse, error)
	AddToFavorites(ctx context.Context, in *gatewayProto.UserIdArtistIdArg, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveFromFavorites(ctx context.Context, in *gatewayProto.UserIdArtistIdArg, opts ...grpc.CallOption) (*empty.Empty, error)
}

type artistUseCaseClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistUseCaseClient(cc grpc.ClientConnInterface) ArtistUseCaseClient {
	return &artistUseCaseClient{cc}
}

func (c *artistUseCaseClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArtistsResponse, error) {
	out := new(ArtistsResponse)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) GetLastId(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*gatewayProto.IntResponse, error) {
	out := new(gatewayProto.IntResponse)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/GetLastId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) Create(ctx context.Context, in *Artist, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) Update(ctx context.Context, in *Artist, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) Delete(ctx context.Context, in *gatewayProto.IdArg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) GetById(ctx context.Context, in *gatewayProto.IdArg, opts ...grpc.CallOption) (*ArtistDataTransfer, error) {
	out := new(ArtistDataTransfer)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) GetPopular(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ArtistsResponse, error) {
	out := new(ArtistsResponse)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/GetPopular", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) GetSize(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*gatewayProto.IntResponse, error) {
	out := new(gatewayProto.IntResponse)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/GetSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) SearchByName(ctx context.Context, in *gatewayProto.StringArg, opts ...grpc.CallOption) (*ArtistsResponse, error) {
	out := new(ArtistsResponse)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/SearchByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) GetFavorites(ctx context.Context, in *gatewayProto.IdArg, opts ...grpc.CallOption) (*ArtistsResponse, error) {
	out := new(ArtistsResponse)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/GetFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) AddToFavorites(ctx context.Context, in *gatewayProto.UserIdArtistIdArg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/AddToFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistUseCaseClient) RemoveFromFavorites(ctx context.Context, in *gatewayProto.UserIdArtistIdArg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/track.ArtistUseCase/RemoveFromFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistUseCaseServer is the server API for ArtistUseCase service.
// All implementations must embed UnimplementedArtistUseCaseServer
// for forward compatibility
type ArtistUseCaseServer interface {
	GetAll(context.Context, *empty.Empty) (*ArtistsResponse, error)
	GetLastId(context.Context, *empty.Empty) (*gatewayProto.IntResponse, error)
	Create(context.Context, *Artist) (*empty.Empty, error)
	Update(context.Context, *Artist) (*empty.Empty, error)
	Delete(context.Context, *gatewayProto.IdArg) (*empty.Empty, error)
	GetById(context.Context, *gatewayProto.IdArg) (*ArtistDataTransfer, error)
	GetPopular(context.Context, *empty.Empty) (*ArtistsResponse, error)
	GetSize(context.Context, *empty.Empty) (*gatewayProto.IntResponse, error)
	SearchByName(context.Context, *gatewayProto.StringArg) (*ArtistsResponse, error)
	GetFavorites(context.Context, *gatewayProto.IdArg) (*ArtistsResponse, error)
	AddToFavorites(context.Context, *gatewayProto.UserIdArtistIdArg) (*empty.Empty, error)
	RemoveFromFavorites(context.Context, *gatewayProto.UserIdArtistIdArg) (*empty.Empty, error)
	mustEmbedUnimplementedArtistUseCaseServer()
}

// UnimplementedArtistUseCaseServer must be embedded to have forward compatible implementations.
type UnimplementedArtistUseCaseServer struct {
}

func (UnimplementedArtistUseCaseServer) GetAll(context.Context, *empty.Empty) (*ArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedArtistUseCaseServer) GetLastId(context.Context, *empty.Empty) (*gatewayProto.IntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastId not implemented")
}
func (UnimplementedArtistUseCaseServer) Create(context.Context, *Artist) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArtistUseCaseServer) Update(context.Context, *Artist) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArtistUseCaseServer) Delete(context.Context, *gatewayProto.IdArg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtistUseCaseServer) GetById(context.Context, *gatewayProto.IdArg) (*ArtistDataTransfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedArtistUseCaseServer) GetPopular(context.Context, *empty.Empty) (*ArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopular not implemented")
}
func (UnimplementedArtistUseCaseServer) GetSize(context.Context, *empty.Empty) (*gatewayProto.IntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSize not implemented")
}
func (UnimplementedArtistUseCaseServer) SearchByName(context.Context, *gatewayProto.StringArg) (*ArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchByName not implemented")
}
func (UnimplementedArtistUseCaseServer) GetFavorites(context.Context, *gatewayProto.IdArg) (*ArtistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorites not implemented")
}
func (UnimplementedArtistUseCaseServer) AddToFavorites(context.Context, *gatewayProto.UserIdArtistIdArg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToFavorites not implemented")
}
func (UnimplementedArtistUseCaseServer) RemoveFromFavorites(context.Context, *gatewayProto.UserIdArtistIdArg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromFavorites not implemented")
}
func (UnimplementedArtistUseCaseServer) mustEmbedUnimplementedArtistUseCaseServer() {}

// UnsafeArtistUseCaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtistUseCaseServer will
// result in compilation errors.
type UnsafeArtistUseCaseServer interface {
	mustEmbedUnimplementedArtistUseCaseServer()
}

func RegisterArtistUseCaseServer(s grpc.ServiceRegistrar, srv ArtistUseCaseServer) {
	s.RegisterService(&ArtistUseCase_ServiceDesc, srv)
}

func _ArtistUseCase_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_GetLastId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).GetLastId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/GetLastId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).GetLastId(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Artist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).Create(ctx, req.(*Artist))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Artist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).Update(ctx, req.(*Artist))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gatewayProto.IdArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).Delete(ctx, req.(*gatewayProto.IdArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gatewayProto.IdArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).GetById(ctx, req.(*gatewayProto.IdArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_GetPopular_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).GetPopular(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/GetPopular",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).GetPopular(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_GetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).GetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/GetSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).GetSize(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_SearchByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gatewayProto.StringArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).SearchByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/SearchByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).SearchByName(ctx, req.(*gatewayProto.StringArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_GetFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gatewayProto.IdArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).GetFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/GetFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).GetFavorites(ctx, req.(*gatewayProto.IdArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_AddToFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gatewayProto.UserIdArtistIdArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).AddToFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/AddToFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).AddToFavorites(ctx, req.(*gatewayProto.UserIdArtistIdArg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistUseCase_RemoveFromFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(gatewayProto.UserIdArtistIdArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistUseCaseServer).RemoveFromFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/track.ArtistUseCase/RemoveFromFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistUseCaseServer).RemoveFromFavorites(ctx, req.(*gatewayProto.UserIdArtistIdArg))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtistUseCase_ServiceDesc is the grpc.ServiceDesc for ArtistUseCase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtistUseCase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "track.ArtistUseCase",
	HandlerType: (*ArtistUseCaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _ArtistUseCase_GetAll_Handler,
		},
		{
			MethodName: "GetLastId",
			Handler:    _ArtistUseCase_GetLastId_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ArtistUseCase_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ArtistUseCase_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ArtistUseCase_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ArtistUseCase_GetById_Handler,
		},
		{
			MethodName: "GetPopular",
			Handler:    _ArtistUseCase_GetPopular_Handler,
		},
		{
			MethodName: "GetSize",
			Handler:    _ArtistUseCase_GetSize_Handler,
		},
		{
			MethodName: "SearchByName",
			Handler:    _ArtistUseCase_SearchByName_Handler,
		},
		{
			MethodName: "GetFavorites",
			Handler:    _ArtistUseCase_GetFavorites_Handler,
		},
		{
			MethodName: "AddToFavorites",
			Handler:    _ArtistUseCase_AddToFavorites_Handler,
		},
		{
			MethodName: "RemoveFromFavorites",
			Handler:    _ArtistUseCase_RemoveFromFavorites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "artist/artistProto/artist.proto",
}
