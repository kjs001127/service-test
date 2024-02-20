package domain

import (
	"context"
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

const bufSize = 1024 * 2 // 2KB

var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, bufSize)
	},
}

type FileStreamer struct {
	repo      AppUrlRepository
	requester HttpRequester
}

func NewFileStreamer(repo AppUrlRepository, requester HttpRequester) *FileStreamer {
	return &FileStreamer{repo: repo, requester: requester}
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
	proxy.ServeHTTP(req.Writer, req.Req)

	return nil
}
