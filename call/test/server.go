package main

import (
	"context"
	pb "github.com/channel-io/ch-proto/call"
)

type Server struct {
	pb.CallServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) CreatePhoneCall(context.Context, *pb.PhoneCallRequest) (*pb.CallResponse, error) {
	return &pb.CallResponse{
		Session: "1",
	}, nil
}
