package domain

import (
	"context"
	"io"
)

type HttpRequest struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Body    []byte            `json:"body"`
	Headers map[string]string `json:"headers"`
}

type HttpRequester interface {
	Request(ctx context.Context, request HttpRequest) (io.ReadCloser, error)
}
