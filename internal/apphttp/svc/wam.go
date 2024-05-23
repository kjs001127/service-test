package svc

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/apphttp/model"
)

type AppHttpProxy struct {
	repo              AppServerSettingRepository
	internalRequester http.RoundTripper
	externalRequester http.RoundTripper
}

func NewAppHttpProxy(
	repo AppServerSettingRepository,
	internal http.RoundTripper,
	external http.RoundTripper,
) *AppHttpProxy {
	return &AppHttpProxy{repo: repo, internalRequester: internal, externalRequester: external}
}

type WamProxyRequest struct {
	Req    *http.Request
	Writer http.ResponseWriter
	AppID  string
}

func (a *AppHttpProxy) Proxy(ctx context.Context, req WamProxyRequest) error {
	serverSetting, err := a.repo.Fetch(ctx, req.AppID)
	if err != nil {
		return errors.WithStack(err)
	}

	if serverSetting.WamURL == nil {
		return apierr.BadRequest(fmt.Errorf("wam url invalid for appID: %s", req.AppID))
	}

	wamUrl, err := url.Parse(*serverSetting.WamURL)
	if err != nil {
		return errors.Wrap(err, "error while parsing wamURL")
	}

	proxy := httputil.NewSingleHostReverseProxy(wamUrl)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = wamUrl.Host
	}

	switch serverSetting.AccessType {
	case model.AccessType_External:
		proxy.Transport = a.externalRequester
	case model.AccessType_Internal:
		proxy.Transport = a.internalRequester
	default:
		return errors.New("accessType invalid")
	}

	proxy.ServeHTTP(req.Writer, req.Req)
	return nil
}
