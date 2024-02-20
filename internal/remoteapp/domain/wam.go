package domain

import (
	"context"
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

type FileStreamer struct {
	repo      AppUrlRepository
	requester http.RoundTripper
}

func NewFileStreamer(repo AppUrlRepository, tripper http.RoundTripper) *FileStreamer {
	return &FileStreamer{repo: repo, requester: tripper}
}

type AppProxyRequest struct {
	Req    *http.Request
	Writer http.ResponseWriter
	AppID  string
}

func (a *FileStreamer) StreamFile(ctx context.Context, req AppProxyRequest) error {
	urls, err := a.repo.Fetch(ctx, req.AppID)
	if err != nil {
		return err
	}

	if urls.WamURL == nil {
		return apierr.BadRequest(errors.New("wam url invalid"))
	}

	wamUrl, err := url.Parse(*urls.WamURL)
	if err != nil {
		return err
	}

	proxy := httputil.NewSingleHostReverseProxy(wamUrl)
	proxy.Transport = a.requester
	proxy.ServeHTTP(req.Writer, req.Req)

	return nil
}
