package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJson   = "application/json"
)

type Invoker struct {
	requester HttpRequester
	repo      AppUrlRepository
	logger    log.ContextAwareLogger
}

func NewInvoker(requester HttpRequester, repo AppUrlRepository, logger log.ContextAwareLogger) *Invoker {
	return &Invoker{requester: requester, repo: repo, logger: logger}
}

func (a *Invoker) Invoke(ctx context.Context, target *appmodel.App, request app.JsonFunctionRequest) app.JsonFunctionResponse {
	urls, err := a.repo.Fetch(ctx, target.ID)
	if err != nil {
		return app.WrapCommonErr(err)
	}

	if urls.FunctionURL == nil {
		a.logger.Debugw(ctx, "function url is nil", "appID", target.ID)
		return app.WrapCommonErr(errors.New("function url empty"))
	}

	marshaled, err := json.Marshal(request)
	if err != nil {
		a.logger.Debugw(ctx, "function request cannot be marshalled",
			"appID", target.ID,
			"request", request,
		)
		return app.WrapCommonErr(err)
	}

	a.logger.Debugw(ctx, "function request", "appID", target.ID, "request", json.RawMessage(marshaled))

	ret, err := a.requestWithHttp(ctx, *urls.FunctionURL, marshaled)
	if err != nil {
		a.logger.Warnw(ctx, "function http request failed", "appID", target.ID, "error", err)
		return app.WrapCommonErr(err)
	}

	a.logger.Debugw(ctx, "function response", "appID", target.ID, "response", json.RawMessage(ret))

	var jsonResp app.JsonFunctionResponse
	if err = json.Unmarshal(ret, &jsonResp); err != nil {
		a.logger.Warnw(ctx, "function response cannot be unmarshalled", "appID", target.ID, "error", err)
		return app.WrapCommonErr(fmt.Errorf("unmarshaling function response to JsonResp, cause: %w", err))
	}

	if jsonResp.Error != nil && len(jsonResp.Error.Type) > 0 {
		a.logger.Warnw(ctx, "function returned err", "appID", target.ID, "error", jsonResp.Error)
	}

	return jsonResp
}

func (a *Invoker) requestWithHttp(ctx context.Context, url string, body []byte) ([]byte, error) {
	return a.requester.Request(ctx, HttpRequest{
		Body:   body,
		Method: http.MethodPut,
		Headers: map[string]string{
			contentTypeHeader: contentTypeJson,
		},
		Url: url,
	})
}
