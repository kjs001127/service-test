// +build tools

package tools

import (
    _ "github.com/googleapis/googleapis@v0.0.0-20220201063650-f78745822aad"
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
    _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1"
    _ "google.golang.org/protobuf/cmd/protoc-gen-go@v1.26"
)
