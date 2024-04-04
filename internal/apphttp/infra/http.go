package infra

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/apphttp/svc"
)

type HttpRequester struct {
	cli *resty.Client
}

func NewHttpRequester(cli *resty.Client) *HttpRequester {
	return &HttpRequester{cli: cli}
}

func (h HttpRequester) Request(ctx context.Context, req svc.HttpRequest) ([]byte, error) {
	r := h.cli.R()
	r.SetContext(ctx)
	r.SetBody(req.Body)
	r.SetHeaders(req.Headers)

	var resp *resty.Response
	var err error
	switch req.Method {
	case http.MethodGet:
		resp, err = r.Get(req.Url)
	case http.MethodPost:
		resp, err = r.Post(req.Url)
	case http.MethodPut:
		resp, err = r.Put(req.Url)
	case http.MethodPatch:
		resp, err = r.Patch(req.Url)
	default:
		return nil, errors.New("unsupported method")
	}
	if err != nil {
		return nil, errors.Wrap(err, "error while requesting to app server")
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 400 {
		return resp.Body(), fmt.Errorf("http response fail, %d", resp.StatusCode())
	}

	return resp.Body(), nil
}