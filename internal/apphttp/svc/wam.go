package svc

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
)

type AppHttpProxy struct {
	repo      AppUrlRepository
	requester http.RoundTripper
}

func NewAppHttpProxy(repo AppUrlRepository, tripper http.RoundTripper) *AppHttpProxy {
	return &AppHttpProxy{repo: repo, requester: tripper}
}

type WamProxyRequest struct {
	Req    *http.Request
	Writer http.ResponseWriter
	AppID  string
}

func (a *AppHttpProxy) Proxy(ctx context.Context, req WamProxyRequest) error {
	urls, err := a.repo.Fetch(ctx, req.AppID)
	if err != nil {
		return errors.WithStack(err)
	}

	if urls.WamURL == nil {
		return apierr.BadRequest(fmt.Errorf("wam url invalid for appID: %s", req.AppID))
	}

	wamUrl, err := url.Parse(*urls.WamURL)
	if err != nil {
		return errors.Wrap(err, "error while parsing wamURL")
	}

	proxy := httputil.NewSingleHostReverseProxy(wamUrl)
	proxy.Transport = a.requester
	proxy.ServeHTTP(req.Writer, req.Req)
	return nil
}