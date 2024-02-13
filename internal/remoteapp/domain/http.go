package domain

import (
	"context"
	"io"
)

type HttpRequest struct {
	Method  string
	Url     string
	Body    []byte
	Headers map[string]string
}

type HttpRequester interface {
	Request(ctx context.Context, request HttpRequest) (io.ReadCloser, error)
}
