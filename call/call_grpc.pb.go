// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package call

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CallClient is the client API for Call service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallClient interface {
	WebRtcTrickle(ctx context.Context, in *WebRtcTrickleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	WebRtcNegotiation(ctx context.Context, in *WebRtcNegotiationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	WebRtcJoin(ctx context.Context, in *WebRtcJoinRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HangUp(ctx context.Context, in *MeetKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Reconnect(ctx context.Context, in *MeetKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// CreateMeet
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	CreateMeet(ctx context.Context, in *PhoneRequest, opts ...grpc.CallOption) (*MeetKey, error)
}

type callClient struct {
	cc grpc.ClientConnInterface
}

func NewCallClient(cc grpc.ClientConnInterface) CallClient {
	return &callClient{cc}
}

func (c *callClient) WebRtcTrickle(ctx context.Context, in *WebRtcTrickleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/call.Call/WebRtcTrickle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callClient) WebRtcNegotiation(ctx context.Context, in *WebRtcNegotiationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/call.Call/WebRtcNegotiation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callClient) WebRtcJoin(ctx context.Context, in *WebRtcJoinRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/call.Call/WebRtcJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callClient) HangUp(ctx context.Context, in *MeetKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/call.Call/HangUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callClient) Reconnect(ctx context.Context, in *MeetKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/call.Call/Reconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callClient) CreateMeet(ctx context.Context, in *PhoneRequest, opts ...grpc.CallOption) (*MeetKey, error) {
	out := new(MeetKey)
	err := c.cc.Invoke(ctx, "/call.Call/CreateMeet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallServer is the server API for Call service.
// All implementations must embed UnimplementedCallServer
// for forward compatibility
type CallServer interface {
	WebRtcTrickle(context.Context, *WebRtcTrickleRequest) (*emptypb.Empty, error)
	WebRtcNegotiation(context.Context, *WebRtcNegotiationRequest) (*emptypb.Empty, error)
	WebRtcJoin(context.Context, *WebRtcJoinRequest) (*emptypb.Empty, error)
	HangUp(context.Context, *MeetKey) (*emptypb.Empty, error)
	Reconnect(context.Context, *MeetKey) (*emptypb.Empty, error)
	// CreateMeet
	//
	// {{.MethodDescriptorProto.Name}} is a call with the method(s) {{$first := true}}{{range .Bindings}}{{if $first}}{{$first = false}}{{else}}, {{end}}{{.HTTPMethod}}{{end}} within the "{{.Service.Name}}" service.
	// It takes in "{{.RequestType.Name}}" and returns a "{{.ResponseType.Name}}".
	//
	CreateMeet(context.Context, *PhoneRequest) (*MeetKey, error)
	mustEmbedUnimplementedCallServer()
}

// UnimplementedCallServer must be embedded to have forward compatible implementations.
type UnimplementedCallServer struct {
}

func (UnimplementedCallServer) WebRtcTrickle(context.Context, *WebRtcTrickleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebRtcTrickle not implemented")
}
func (UnimplementedCallServer) WebRtcNegotiation(context.Context, *WebRtcNegotiationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebRtcNegotiation not implemented")
}
func (UnimplementedCallServer) WebRtcJoin(context.Context, *WebRtcJoinRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebRtcJoin not implemented")
}
func (UnimplementedCallServer) HangUp(context.Context, *MeetKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HangUp not implemented")
}
func (UnimplementedCallServer) Reconnect(context.Context, *MeetKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reconnect not implemented")
}
func (UnimplementedCallServer) CreateMeet(context.Context, *PhoneRequest) (*MeetKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeet not implemented")
}
func (UnimplementedCallServer) mustEmbedUnimplementedCallServer() {}

// UnsafeCallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallServer will
// result in compilation errors.
type UnsafeCallServer interface {
	mustEmbedUnimplementedCallServer()
}

func RegisterCallServer(s grpc.ServiceRegistrar, srv CallServer) {
	s.RegisterService(&Call_ServiceDesc, srv)
}

func _Call_WebRtcTrickle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRtcTrickleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).WebRtcTrickle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/WebRtcTrickle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).WebRtcTrickle(ctx, req.(*WebRtcTrickleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Call_WebRtcNegotiation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRtcNegotiationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).WebRtcNegotiation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/WebRtcNegotiation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).WebRtcNegotiation(ctx, req.(*WebRtcNegotiationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Call_WebRtcJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRtcJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).WebRtcJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/WebRtcJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).WebRtcJoin(ctx, req.(*WebRtcJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Call_HangUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).HangUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/HangUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).HangUp(ctx, req.(*MeetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Call_Reconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).Reconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/Reconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).Reconnect(ctx, req.(*MeetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Call_CreateMeet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).CreateMeet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/CreateMeet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).CreateMeet(ctx, req.(*PhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Call_ServiceDesc is the grpc.ServiceDesc for Call service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Call_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "call.Call",
	HandlerType: (*CallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WebRtcTrickle",
			Handler:    _Call_WebRtcTrickle_Handler,
		},
		{
			MethodName: "WebRtcNegotiation",
			Handler:    _Call_WebRtcNegotiation_Handler,
		},
		{
			MethodName: "WebRtcJoin",
			Handler:    _Call_WebRtcJoin_Handler,
		},
		{
			MethodName: "HangUp",
			Handler:    _Call_HangUp_Handler,
		},
		{
			MethodName: "Reconnect",
			Handler:    _Call_Reconnect_Handler,
		},
		{
			MethodName: "CreateMeet",
			Handler:    _Call_CreateMeet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "call/call.proto",
}
