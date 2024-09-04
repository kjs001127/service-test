package proxyapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/util"
)

const (
	contentTypeHeader = "Content-Type"
	mimeTypeJson      = "application/json"
)

var configs = []routeConfig{CHANNEL_API_ROUTES, DOCUMENT_API_ROUTES}

type APIRouter struct {
	router router
	resty  *resty.Client
}

func (a *APIRouter) RegisterTo(registry native.FunctionRegistry) {
	// @TODO camel add version support for native fn, omit version for now
	for fnName, _ := range a.router {
		registry.Register(string(fnName), a.Handle)
	}
}

func NewProxyAPI(svcMap util.ServiceMap, resty *resty.Client) *APIRouter {
	return &APIRouter{router: newRouter(svcMap, configs...), resty: resty}
}

func (a *APIRouter) Handle(ctx context.Context, token native.Token, fnReq native.FunctionRequest) native.FunctionResponse {

	// use defaultVersion for now
	url, ok := a.router.route(function(fnReq.Method), defaultVersion)
	if !ok {
		return native.WrapCommonErr(errors.New("routeConfig not found"))
	}

	r := a.resty.R()

	if token.Exists && len(token.Value) > 0 {
		r.SetHeader("x-access-token", token.Value)
	}

	r.SetHeader(contentTypeHeader, mimeTypeJson)
	r.SetBody(fnReq.Params)
	r.SetContext(ctx)

	resp, err := r.Post(url)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	if resp.IsError() {
		return native.WrapCommonErr(fmt.Errorf("request failed, body: %s", resp.Body()))
	}

	return native.FunctionResponse{
		Result: resp.Body(),
	}
}
