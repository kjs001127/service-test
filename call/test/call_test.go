//go:build !test
// +build !test

package main

import (
	"context"
	"fmt"
	"net"
	"testing"

	proto "github.com/channel-io/ch-proto/call"
	"google.golang.org/grpc"
)

func Test_call(t *testing.T) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	call := New()
	proto.RegisterCallServer(s, call)

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	client := proto.NewCallClient(conn)
	phoneCall := &proto.PhoneCallRequest{
		FromNumber: "1",
		ToNumber:   "2",
	}

	response, err := client.CreatePhoneCall(context.Background(), phoneCall)
	if err != nil {
		t.Error("failed to createPhoneCall")
	}

	if response.Session == "" {
		t.Error("session field have to exist")
	}

	conn.Close()
}
