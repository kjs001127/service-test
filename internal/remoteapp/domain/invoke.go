package domain

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

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

func (a *Invoker) Invoke(ctx context.Context, app *app.App, request app.JsonFunctionRequest) (app.JsonFunctionResponse, error) {
	urls, err := a.repo.Fetch(ctx, app.ID)
	if err != nil {
		return nil, err
	}

	if urls.FunctionURL == nil {
		return nil, apierr.BadRequest(errors.New("function url invalid"))
	}

	marshaled, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reader, err := a.requester.Request(ctx, HttpRequest{
		Body:   marshaled,
		Method: http.MethodPut,
		Headers: map[string]string{
			contentTypeHeader: contentTypeJson,
		},
		Url: *urls.FunctionURL,
	})
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	ret, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return a.resultOf(ret)
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

type Response struct {
	Error  Error           `json:"error"`
	Result json.RawMessage `json:"result"`
}

func (a *Invoker) resultOf(ret []byte) (app.JsonFunctionResponse, error) {
	var jsonResp Response
	if err := json.Unmarshal(ret, &jsonResp); err != nil {
		return nil, err
	}

	if len(jsonResp.Error.Type) > 0 {
		return nil, &jsonResp.Error
	}

	return app.JsonFunctionResponse(jsonResp.Result), nil
}
