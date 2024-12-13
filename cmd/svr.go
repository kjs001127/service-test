package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/durationpb"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	testv1 "github.com/channel-io/service-test/proto/test/v1"
	"github.com/channel-io/service-test/proto/test/v1/testv1connect"
)

type TestServer struct{}

func (t TestServer) Test(ctx context.Context, c *connect.Request[testv1.TestRequest]) (*connect.Response[testv1.TestResponse], error) {
	log.Println("Request received")

	err := connect.NewError(
		connect.CodeUnavailable,
		errors.New("back off and retry"),
	)

	retryInfo := &errdetails.RetryInfo{
		RetryDelay: durationpb.New(10 * time.Second),
	}

	if detail, detailErr := connect.NewErrorDetail(retryInfo); detailErr == nil {
		err.AddDetail(detail)
	}

	return nil, err
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
