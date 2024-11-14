package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/httpfn/model"
	signutil "github.com/channel-io/ch-app-store/internal/httpfn/util"
	"github.com/channel-io/ch-app-store/lib/log"

	"github.com/pkg/errors"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJson   = "application/json"
	xSignatureHeader  = "X-Signature"
)

type RequesterMap map[model.AccessType]HttpRequester

type Invoker struct {
	requesterMap RequesterMap
	repo         AppServerSettingRepository
	logger       log.ContextAwareLogger
}

func NewInvoker(requesterMap RequesterMap, repo AppServerSettingRepository, logger log.ContextAwareLogger) *Invoker {
	return &Invoker{requesterMap: requesterMap, repo: repo, logger: logger}
}

func (a *Invoker) Invoke(ctx context.Context, target *appmodel.App, request app.JsonFunctionRequest) app.JsonFunctionResponse {
	serverSetting, err := a.repo.Fetch(ctx, target.ID)
	if err != nil {
		return app.WrapCommonErr(err)
	}

	if serverSetting.FunctionURL == nil {
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

	ret, err := a.requestWithHttp(ctx, serverSetting, marshaled)
	if err != nil {
		a.logger.Warnw(ctx, "function http request failed", "appID", target.ID, "error", err)
		return app.WrapCommonErr(err)
	}

	a.logger.Debugw(ctx, "function response", "appID", target.ID, "response", json.RawMessage(ret))

	if len(ret) <= 0 {
		return app.JsonFunctionResponse{}
	}

	var jsonResp app.JsonFunctionResponse
	if err = json.Unmarshal(ret, &jsonResp); err != nil {
		a.logger.Warnw(ctx, "function response cannot be unmarshalled", "appID", target.ID, "error", err)
		return app.WrapCommonErr(fmt.Errorf("unmarshaling function response to JsonResp, cause: %w", err))
	}

	if jsonResp.IsError() {
		a.logger.Warnw(ctx, "function returned err", "appID", target.ID, "error", jsonResp.Error)
	}

	return jsonResp
}

func (a *Invoker) requestWithHttp(ctx context.Context, serverSetting model.ServerSetting, body []byte) ([]byte, error) {
	headers := map[string]string{
		contentTypeHeader: contentTypeJson,
	}
	if serverSetting.SigningKey != nil {
		signature, err := signutil.Sign(*serverSetting.SigningKey, body)
		if err != nil {
			return nil, err
		}
		headers[xSignatureHeader] = signature
	}
	req := HttpRequest{
		Body:    body,
		Method:  http.MethodPut,
		Headers: headers,
		Url:     *serverSetting.FunctionURL,
	}

	requester, exists := a.requesterMap[serverSetting.AccessType]
	if !exists {
		return nil, errors.New("invalid accessType")
	}
	return requester.Request(ctx, req)
}
