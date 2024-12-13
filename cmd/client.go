package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"

	testv1 "github.com/channel-io/service-test/proto/test/v1"
	"github.com/channel-io/service-test/proto/test/v1/testv1connect"
)

func main() {
	client := testv1connect.NewTestServiceClient(
		http.DefaultClient,
		"http://localhost:10010",
	)
	res, err := client.Test(
		context.Background(),
		connect.NewRequest(&testv1.TestRequest{SomeStr: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.SomeStr)

}
