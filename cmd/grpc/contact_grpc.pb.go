// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/proto/contact.proto

package grpc

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
	ContactService_CreateContact_FullMethodName = "/go_test_crm_service.ContactService/CreateContact"
	ContactService_GetContact_FullMethodName    = "/go_test_crm_service.ContactService/GetContact"
	ContactService_UpdateContact_FullMethodName = "/go_test_crm_service.ContactService/UpdateContact"
	ContactService_DeleteContact_FullMethodName = "/go_test_crm_service.ContactService/DeleteContact"
	ContactService_ListContacts_FullMethodName  = "/go_test_crm_service.ContactService/ListContacts"
)

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactServiceClient interface {
	CreateContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error)
	GetContact(ctx context.Context, in *ContactId, opts ...grpc.CallOption) (*ContactResponse, error)
	UpdateContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error)
	DeleteContact(ctx context.Context, in *ContactId, opts ...grpc.CallOption) (*Empty, error)
	ListContacts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ContactList, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) CreateContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := c.cc.Invoke(ctx, ContactService_CreateContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetContact(ctx context.Context, in *ContactId, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := c.cc.Invoke(ctx, ContactService_GetContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) UpdateContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*ContactResponse, error) {
	out := new(ContactResponse)
	err := c.cc.Invoke(ctx, ContactService_UpdateContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) DeleteContact(ctx context.Context, in *ContactId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ContactService_DeleteContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) ListContacts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ContactList, error) {
	out := new(ContactList)
	err := c.cc.Invoke(ctx, ContactService_ListContacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
// All implementations must embed UnimplementedContactServiceServer
// for forward compatibility
type ContactServiceServer interface {
	CreateContact(context.Context, *Contact) (*ContactResponse, error)
	GetContact(context.Context, *ContactId) (*ContactResponse, error)
	UpdateContact(context.Context, *Contact) (*ContactResponse, error)
	DeleteContact(context.Context, *ContactId) (*Empty, error)
	ListContacts(context.Context, *Empty) (*ContactList, error)
	mustEmbedUnimplementedContactServiceServer()
}

// UnimplementedContactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (UnimplementedContactServiceServer) CreateContact(context.Context, *Contact) (*ContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContact not implemented")
}
func (UnimplementedContactServiceServer) GetContact(context.Context, *ContactId) (*ContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedContactServiceServer) UpdateContact(context.Context, *Contact) (*ContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContact not implemented")
}
func (UnimplementedContactServiceServer) DeleteContact(context.Context, *ContactId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContact not implemented")
}
func (UnimplementedContactServiceServer) ListContacts(context.Context, *Empty) (*ContactList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContacts not implemented")
}
func (UnimplementedContactServiceServer) mustEmbedUnimplementedContactServiceServer() {}

// UnsafeContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServiceServer will
// result in compilation errors.
type UnsafeContactServiceServer interface {
	mustEmbedUnimplementedContactServiceServer()
}

func RegisterContactServiceServer(s grpc.ServiceRegistrar, srv ContactServiceServer) {
	s.RegisterService(&ContactService_ServiceDesc, srv)
}

func _ContactService_CreateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).CreateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_CreateContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).CreateContact(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_GetContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetContact(ctx, req.(*ContactId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_UpdateContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).UpdateContact(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_DeleteContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).DeleteContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_DeleteContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).DeleteContact(ctx, req.(*ContactId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_ListContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).ListContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_ListContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).ListContacts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactService_ServiceDesc is the grpc.ServiceDesc for ContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_test_crm_service.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContact",
			Handler:    _ContactService_CreateContact_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _ContactService_GetContact_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _ContactService_UpdateContact_Handler,
		},
		{
			MethodName: "DeleteContact",
			Handler:    _ContactService_DeleteContact_Handler,
		},
		{
			MethodName: "ListContacts",
			Handler:    _ContactService_ListContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/contact.proto",
}