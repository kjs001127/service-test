package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	testv1 "github.com/channel-io/service-test/proto/test/v1"
	"github.com/channel-io/service-test/proto/test/v1/testv1connect"
)

type TestServer struct{}

func (t TestServer) Test(ctx context.Context, c *connect.Request[testv1.TestRequest]) (*connect.Response[testv1.TestResponse], error) {
	log.Println("Request received")
	return connect.NewResponse(&testv1.TestResponse{SomeStr: c.Msg.SomeStr}), nil
}

func main() {
	test := TestServer{}
	mux := http.NewServeMux()
	mux.Handle(testv1connect.NewTestServiceHandler(test))
	err := http.ListenAndServe(
		"localhost:10010",
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		log.Fatalln("error:", err)
	}
}
