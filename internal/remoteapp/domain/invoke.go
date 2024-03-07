package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
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

func (a *Invoker) Invoke(ctx context.Context, target *app.App, request app.JsonFunctionRequest) app.JsonFunctionResponse {
	urls, err := a.repo.Fetch(ctx, target.ID)
	if err != nil {
		return app.WrapCommonErr(err)
	}

	if urls.FunctionURL == nil {
		return app.WrapCommonErr(errors.New("function url empty"))
	}

	marshaled, err := json.Marshal(request)
	if err != nil {
		return app.WrapCommonErr(err)
	}

	id := uid.New()
	a.logger.Debugw(ctx, "function request", "id", id, "appID", target.ID, "request", request)

	ret, err := a.requestWithHttp(ctx, *urls.FunctionURL, marshaled)
	if err != nil {
		return app.WrapCommonErr(err)
	}

	a.logger.Debugw(ctx, "function response", "id", id, "appID", target.ID, "response", string(ret))

	var jsonResp app.JsonFunctionResponse
	if err = json.Unmarshal(ret, &jsonResp); err != nil {
		return app.WrapCommonErr(fmt.Errorf("unmarshaling function response to JsonResp, cause: %w", err))
	}

	if jsonResp.Error != nil && len(jsonResp.Error.Type) > 0 {
		a.logger.Warnw(ctx, "function returned err", "id", id, "appID", target.ID, "error", jsonResp.Error)
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
