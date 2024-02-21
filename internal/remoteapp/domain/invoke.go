package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJson   = "application/json"
)

type Invoker struct {
	requester HttpRequester
	repo      AppUrlRepository
}

func NewInvoker(requester HttpRequester, repo AppUrlRepository) *Invoker {
	return &Invoker{requester: requester, repo: repo}
}

func (a *Invoker) Invoke(ctx context.Context, target *app.App, request app.JsonFunctionRequest) app.JsonFunctionResponse {
	urls, err := a.repo.Fetch(ctx, target.ID)
	if err != nil {
		return app.WrapErr(err)
	}

	if urls.FunctionURL == nil {
		return app.WrapErr(errors.New("function url empty"))
	}

	marshaled, err := json.Marshal(request)
	if err != nil {
		return app.WrapErr(err)
	}

	ret, err := a.requestWithHttp(ctx, *urls.FunctionURL, marshaled)
	if err != nil {
		return app.WrapErr(err)
	}

	fmt.Printf("requesting remote app function %s : %s \n %s: %s", "request", marshaled, "response", ret)

	var jsonResp app.JsonFunctionResponse
	if err = json.Unmarshal(ret, &jsonResp); err != nil {
		return app.WrapErr(err)
	}

	return jsonResp
}

func (a *Invoker) requestWithHttp(ctx context.Context, url string, body []byte) ([]byte, error) {
	reader, err := a.requester.Request(ctx, HttpRequest{
		Body:   body,
		Method: http.MethodPut,
		Headers: map[string]string{
			contentTypeHeader: contentTypeJson,
		},
		Url: url,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error while requesting to app server")
	}
	defer reader.Close()

	ret, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "error while reading body")
	}

	return ret, nil
}
