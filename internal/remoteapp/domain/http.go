package domain

import (
	"context"
	"io"
)

type Method string

const (
	MethodPost = Method("POST")
	MethodGet  = Method("GET")
)

type HttpRequest struct {
	Method  Method
	Url     string
	Body    []byte
	Headers map[string]string
}

type HttpRequester interface {
	Request(ctx context.Context, request HttpRequest) (io.ReadCloser, error)
}
