// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Auction

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

// AuctionHouseClient is the client API for AuctionHouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionHouseClient interface {
	OpenConnection(ctx context.Context, in *Connect, opts ...grpc.CallOption) (AuctionHouse_OpenConnectionClient, error)
	CloseConnection(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error)
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*BidReply, error)
	Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ResultMessage, error)
	Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error)
	Replicate(ctx context.Context, in *ReplicateMessage, opts ...grpc.CallOption) (*BidReply, error)
	GetID(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PortIndex, error)
	RegisterPulse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error)
	RingElection(ctx context.Context, in *PortsAndClocks, opts ...grpc.CallOption) (*Void, error)
	SelectNewLeader(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ElectionPorts, error)
	BroadcastNewLeader(ctx context.Context, in *ElectionPorts, opts ...grpc.CallOption) (*Void, error)
	CutOfReplicate(ctx context.Context, in *CutOfMessage, opts ...grpc.CallOption) (*Void, error)
	PromptTimeAndDuration(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TimeMessage, error)
	//Ricart And Agrawala
	AccessCritical(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ReplyMessage, error)
	ReceiveRequest(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*Void, error)
	ReceiveReply(ctx context.Context, in *ReplyMessage, opts ...grpc.CallOption) (*Void, error)
}

type auctionHouseClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionHouseClient(cc grpc.ClientConnInterface) AuctionHouseClient {
	return &auctionHouseClient{cc}
}

func (c *auctionHouseClient) OpenConnection(ctx context.Context, in *Connect, opts ...grpc.CallOption) (AuctionHouse_OpenConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuctionHouse_ServiceDesc.Streams[0], "/Auction.AuctionHouse/OpenConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &auctionHouseOpenConnectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuctionHouse_OpenConnectionClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type auctionHouseOpenConnectionClient struct {
	grpc.ClientStream
}

func (x *auctionHouseOpenConnectionClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *auctionHouseClient) CloseConnection(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/CloseConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*BidReply, error) {
	out := new(BidReply)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) Broadcast(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) Replicate(ctx context.Context, in *ReplicateMessage, opts ...grpc.CallOption) (*BidReply, error) {
	out := new(BidReply)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/Replicate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) GetID(ctx context.Context, in *Void, opts ...grpc.CallOption) (*PortIndex, error) {
	out := new(PortIndex)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/getID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) RegisterPulse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/RegisterPulse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) RingElection(ctx context.Context, in *PortsAndClocks, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/RingElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) SelectNewLeader(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ElectionPorts, error) {
	out := new(ElectionPorts)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/SelectNewLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) BroadcastNewLeader(ctx context.Context, in *ElectionPorts, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/BroadcastNewLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) CutOfReplicate(ctx context.Context, in *CutOfMessage, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/CutOfReplicate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) PromptTimeAndDuration(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TimeMessage, error) {
	out := new(TimeMessage)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/PromptTimeAndDuration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) AccessCritical(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/AccessCritical", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) ReceiveRequest(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/ReceiveRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionHouseClient) ReceiveReply(ctx context.Context, in *ReplyMessage, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/Auction.AuctionHouse/ReceiveReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionHouseServer is the server API for AuctionHouse service.
// All implementations must embed UnimplementedAuctionHouseServer
// for forward compatibility
type AuctionHouseServer interface {
	OpenConnection(*Connect, AuctionHouse_OpenConnectionServer) error
	CloseConnection(context.Context, *Message) (*Void, error)
	Bid(context.Context, *BidMessage) (*BidReply, error)
	Result(context.Context, *Void) (*ResultMessage, error)
	Broadcast(context.Context, *Message) (*Void, error)
	Replicate(context.Context, *ReplicateMessage) (*BidReply, error)
	GetID(context.Context, *Void) (*PortIndex, error)
	RegisterPulse(context.Context, *Message) (*Void, error)
	RingElection(context.Context, *PortsAndClocks) (*Void, error)
	SelectNewLeader(context.Context, *Void) (*ElectionPorts, error)
	BroadcastNewLeader(context.Context, *ElectionPorts) (*Void, error)
	CutOfReplicate(context.Context, *CutOfMessage) (*Void, error)
	PromptTimeAndDuration(context.Context, *Void) (*TimeMessage, error)
	//Ricart And Agrawala
	AccessCritical(context.Context, *RequestMessage) (*ReplyMessage, error)
	ReceiveRequest(context.Context, *RequestMessage) (*Void, error)
	ReceiveReply(context.Context, *ReplyMessage) (*Void, error)
	mustEmbedUnimplementedAuctionHouseServer()
}

// UnimplementedAuctionHouseServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionHouseServer struct {
}

func (UnimplementedAuctionHouseServer) OpenConnection(*Connect, AuctionHouse_OpenConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenConnection not implemented")
}
func (UnimplementedAuctionHouseServer) CloseConnection(context.Context, *Message) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseConnection not implemented")
}
func (UnimplementedAuctionHouseServer) Bid(context.Context, *BidMessage) (*BidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionHouseServer) Result(context.Context, *Void) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionHouseServer) Broadcast(context.Context, *Message) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedAuctionHouseServer) Replicate(context.Context, *ReplicateMessage) (*BidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replicate not implemented")
}
func (UnimplementedAuctionHouseServer) GetID(context.Context, *Void) (*PortIndex, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetID not implemented")
}
func (UnimplementedAuctionHouseServer) RegisterPulse(context.Context, *Message) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPulse not implemented")
}
func (UnimplementedAuctionHouseServer) RingElection(context.Context, *PortsAndClocks) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RingElection not implemented")
}
func (UnimplementedAuctionHouseServer) SelectNewLeader(context.Context, *Void) (*ElectionPorts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectNewLeader not implemented")
}
func (UnimplementedAuctionHouseServer) BroadcastNewLeader(context.Context, *ElectionPorts) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastNewLeader not implemented")
}
func (UnimplementedAuctionHouseServer) CutOfReplicate(context.Context, *CutOfMessage) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CutOfReplicate not implemented")
}
func (UnimplementedAuctionHouseServer) PromptTimeAndDuration(context.Context, *Void) (*TimeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromptTimeAndDuration not implemented")
}
func (UnimplementedAuctionHouseServer) AccessCritical(context.Context, *RequestMessage) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessCritical not implemented")
}
func (UnimplementedAuctionHouseServer) ReceiveRequest(context.Context, *RequestMessage) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveRequest not implemented")
}
func (UnimplementedAuctionHouseServer) ReceiveReply(context.Context, *ReplyMessage) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveReply not implemented")
}
func (UnimplementedAuctionHouseServer) mustEmbedUnimplementedAuctionHouseServer() {}

// UnsafeAuctionHouseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionHouseServer will
// result in compilation errors.
type UnsafeAuctionHouseServer interface {
	mustEmbedUnimplementedAuctionHouseServer()
}

func RegisterAuctionHouseServer(s grpc.ServiceRegistrar, srv AuctionHouseServer) {
	s.RegisterService(&AuctionHouse_ServiceDesc, srv)
}

func _AuctionHouse_OpenConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Connect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuctionHouseServer).OpenConnection(m, &auctionHouseOpenConnectionServer{stream})
}

type AuctionHouse_OpenConnectionServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type auctionHouseOpenConnectionServer struct {
	grpc.ServerStream
}

func (x *auctionHouseOpenConnectionServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _AuctionHouse_CloseConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).CloseConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/CloseConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).CloseConnection(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Result(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Broadcast(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_Replicate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).Replicate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/Replicate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).Replicate(ctx, req.(*ReplicateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/getID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).GetID(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_RegisterPulse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).RegisterPulse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/RegisterPulse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).RegisterPulse(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_RingElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortsAndClocks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).RingElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/RingElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).RingElection(ctx, req.(*PortsAndClocks))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_SelectNewLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).SelectNewLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/SelectNewLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).SelectNewLeader(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_BroadcastNewLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionPorts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).BroadcastNewLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/BroadcastNewLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).BroadcastNewLeader(ctx, req.(*ElectionPorts))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_CutOfReplicate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CutOfMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).CutOfReplicate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/CutOfReplicate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).CutOfReplicate(ctx, req.(*CutOfMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_PromptTimeAndDuration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).PromptTimeAndDuration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/PromptTimeAndDuration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).PromptTimeAndDuration(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_AccessCritical_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).AccessCritical(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/AccessCritical",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).AccessCritical(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_ReceiveRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).ReceiveRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/ReceiveRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).ReceiveRequest(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionHouse_ReceiveReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionHouseServer).ReceiveReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auction.AuctionHouse/ReceiveReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionHouseServer).ReceiveReply(ctx, req.(*ReplyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionHouse_ServiceDesc is the grpc.ServiceDesc for AuctionHouse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionHouse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auction.AuctionHouse",
	HandlerType: (*AuctionHouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseConnection",
			Handler:    _AuctionHouse_CloseConnection_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _AuctionHouse_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _AuctionHouse_Result_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _AuctionHouse_Broadcast_Handler,
		},
		{
			MethodName: "Replicate",
			Handler:    _AuctionHouse_Replicate_Handler,
		},
		{
			MethodName: "getID",
			Handler:    _AuctionHouse_GetID_Handler,
		},
		{
			MethodName: "RegisterPulse",
			Handler:    _AuctionHouse_RegisterPulse_Handler,
		},
		{
			MethodName: "RingElection",
			Handler:    _AuctionHouse_RingElection_Handler,
		},
		{
			MethodName: "SelectNewLeader",
			Handler:    _AuctionHouse_SelectNewLeader_Handler,
		},
		{
			MethodName: "BroadcastNewLeader",
			Handler:    _AuctionHouse_BroadcastNewLeader_Handler,
		},
		{
			MethodName: "CutOfReplicate",
			Handler:    _AuctionHouse_CutOfReplicate_Handler,
		},
		{
			MethodName: "PromptTimeAndDuration",
			Handler:    _AuctionHouse_PromptTimeAndDuration_Handler,
		},
		{
			MethodName: "AccessCritical",
			Handler:    _AuctionHouse_AccessCritical_Handler,
		},
		{
			MethodName: "ReceiveRequest",
			Handler:    _AuctionHouse_ReceiveRequest_Handler,
		},
		{
			MethodName: "ReceiveReply",
			Handler:    _AuctionHouse_ReceiveReply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenConnection",
			Handler:       _AuctionHouse_OpenConnection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Auction/auction.proto",
}
