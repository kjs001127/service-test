package infra

import (
	"context"
	"errors"
	"io"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

type HttpRequester struct {
	cli *resty.Client
}

func NewHttpRequester(cli *resty.Client) *HttpRequester {
	return &HttpRequester{cli: cli}
}

func (h HttpRequester) Request(ctx context.Context, req domain.HttpRequest) (io.ReadCloser, error) {
	r := h.cli.R()
	r.SetContext(ctx)
	r.SetBody(req.Body)
	r.SetHeaders(req.Headers)
	r.SetDoNotParseResponse(true)

	var resp *resty.Response
	var err error
	switch req.Method {
	case domain.MethodGet:
		resp, err = r.Get(req.Url)
	case domain.MethodPost:
		resp, err = r.Post(req.Url)
	default:
		return nil, errors.New("unsupported method")
	}
	if err != nil {
		return nil, err
	}

	return resp.RawBody(), nil
}
