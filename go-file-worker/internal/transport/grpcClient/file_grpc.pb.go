// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc2
// source: file.proto

package grpcClient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FileWorker_SaveFile_FullMethodName       = "/fileworker.FileWorker/SaveFile"
	FileWorker_SaveFiles_FullMethodName      = "/fileworker.FileWorker/SaveFiles"
	FileWorker_DeleteFile_FullMethodName     = "/fileworker.FileWorker/DeleteFile"
	FileWorker_GetFolderFiles_FullMethodName = "/fileworker.FileWorker/GetFolderFiles"
)

// FileWorkerClient is the client API for FileWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileWorkerClient interface {
	SaveFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*DefaultResponse, error)
	SaveFiles(ctx context.Context, in *Files, opts ...grpc.CallOption) (*DefaultResponse, error)
	DeleteFile(ctx context.Context, in *FileNames, opts ...grpc.CallOption) (*DefaultResponse, error)
	GetFolderFiles(ctx context.Context, in *FolderName, opts ...grpc.CallOption) (*Files, error)
}

type fileWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileWorkerClient(cc grpc.ClientConnInterface) FileWorkerClient {
	return &fileWorkerClient{cc}
}

func (c *fileWorkerClient) SaveFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, FileWorker_SaveFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWorkerClient) SaveFiles(ctx context.Context, in *Files, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, FileWorker_SaveFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWorkerClient) DeleteFile(ctx context.Context, in *FileNames, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, FileWorker_DeleteFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileWorkerClient) GetFolderFiles(ctx context.Context, in *FolderName, opts ...grpc.CallOption) (*Files, error) {
	out := new(Files)
	err := c.cc.Invoke(ctx, FileWorker_GetFolderFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileWorkerServer is the server API for FileWorker service.
// All implementations must embed UnimplementedFileWorkerServer
// for forward compatibility
type FileWorkerServer interface {
	SaveFile(context.Context, *File) (*DefaultResponse, error)
	SaveFiles(context.Context, *Files) (*DefaultResponse, error)
	DeleteFile(context.Context, *FileNames) (*DefaultResponse, error)
	GetFolderFiles(context.Context, *FolderName) (*Files, error)
	mustEmbedUnimplementedFileWorkerServer()
}

// UnimplementedFileWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedFileWorkerServer struct {
}

func (UnimplementedFileWorkerServer) SaveFile(context.Context, *File) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedFileWorkerServer) SaveFiles(context.Context, *Files) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFiles not implemented")
}
func (UnimplementedFileWorkerServer) DeleteFile(context.Context, *FileNames) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileWorkerServer) GetFolderFiles(context.Context, *FolderName) (*Files, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFolderFiles not implemented")
}
func (UnimplementedFileWorkerServer) mustEmbedUnimplementedFileWorkerServer() {}

// UnsafeFileWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileWorkerServer will
// result in compilation errors.
type UnsafeFileWorkerServer interface {
	mustEmbedUnimplementedFileWorkerServer()
}

func RegisterFileWorkerServer(s grpc.ServiceRegistrar, srv FileWorkerServer) {
	s.RegisterService(&FileWorker_ServiceDesc, srv)
}

func _FileWorker_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWorkerServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWorker_SaveFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWorkerServer).SaveFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWorker_SaveFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Files)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWorkerServer).SaveFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWorker_SaveFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWorkerServer).SaveFiles(ctx, req.(*Files))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWorker_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileNames)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWorkerServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWorker_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWorkerServer).DeleteFile(ctx, req.(*FileNames))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileWorker_GetFolderFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileWorkerServer).GetFolderFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileWorker_GetFolderFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileWorkerServer).GetFolderFiles(ctx, req.(*FolderName))
	}
	return interceptor(ctx, in, info, handler)
}

// FileWorker_ServiceDesc is the grpc.ServiceDesc for FileWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileworker.FileWorker",
	HandlerType: (*FileWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveFile",
			Handler:    _FileWorker_SaveFile_Handler,
		},
		{
			MethodName: "SaveFiles",
			Handler:    _FileWorker_SaveFiles_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileWorker_DeleteFile_Handler,
		},
		{
			MethodName: "GetFolderFiles",
			Handler:    _FileWorker_GetFolderFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
