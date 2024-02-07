package domain

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

const bufSize = 1024 * 2 // 2KB

var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, bufSize)
	},
}

func (a *RemoteApp) StreamFile(ctx context.Context, path string, writer io.Writer) error {
	if a.WamURL != nil {
		return apierr.BadRequest(errors.New("wam url invalid"))
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := *a.WamURL + path

	reader, err := a.requester.Request(ctx, HttpRequest{
		Method: http.MethodGet,
		Url:    url,
	})
	if err != nil {
		return err
	}
	defer reader.Close() // TODO add logging

	return doStream(reader, writer)
}

func doStream(from io.ReadCloser, to io.Writer) error {
	buf := bufPool.Get().([]byte)
	defer bufPool.Put(buf)

	var n int
	var err error
	for ; err == nil; n, err = from.Read(buf) {
		if n <= 0 {
			continue
		}

		if _, err := to.Write(buf[:n]); err != nil {
			return err
		}
	}

	if err != io.EOF {
		return err
	}

	return nil
}
