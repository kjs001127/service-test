package infra

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

type HttpRequester struct {
}

func NewHttpRequester() *HttpRequester {
	return &HttpRequester{}
}

func (h HttpRequester) Request(ctx context.Context, req domain.HttpRequest) (io.ReadCloser, error) {
	resp, err := http.Post(req.Url, req.ContentType, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
