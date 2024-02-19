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

type FileStreamHandler struct {
	repo      AppUrlRepository
	requester HttpRequester
}

func NewFileStreamHandler(repo AppUrlRepository, requester HttpRequester) *FileStreamHandler {
	return &FileStreamHandler{repo: repo, requester: requester}
}

func (a *FileStreamHandler) StreamFile(ctx context.Context, appID string, path string, writer io.Writer) error {
	urls, err := a.repo.Fetch(ctx, appID)
	if err != nil {
		return err
	}

	if urls.WamURL == nil {
		return apierr.BadRequest(errors.New("wam url invalid"))
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	url := *urls.WamURL + path

	reader, err := a.requester.Request(ctx, HttpRequest{
		Method: http.MethodGet,
		Headers: map[string]string{
			"Accept": "text/html,image/webp,image/apng",
		},
		Url: url,
	})
	if err != nil {
		return err
	}

	return doStream(reader, writer)
}

func doStream(from io.ReadCloser, to io.Writer) error {
	defer from.Close() // TODO add logging

	buf := bufPool.Get().([]byte)
	defer bufPool.Put(buf)

	for {
		n, err := from.Read(buf)
		if n > 0 {
			if _, err := to.Write(buf[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}
