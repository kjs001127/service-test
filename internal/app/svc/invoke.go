package svc

import (
	"context"
	"encoding/json"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type Invoker struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
	handler   InvokeHandler

	listeners []FunctionRequestListener
}

func NewInvoker(
	appChRepo AppChannelRepository,
	appRepo AppRepository,
	handler InvokeHandler,
	listeners []FunctionRequestListener,
) *Invoker {
	return &Invoker{appChRepo: appChRepo, appRepo: appRepo, handler: handler, listeners: listeners}
}

func (i *Invoker) Invoke(ctx context.Context, appID string, req JsonFunctionRequest) JsonFunctionResponse {
	_, err := i.appChRepo.Fetch(ctx, model.AppChannelID{
		AppID:     appID,
		ChannelID: req.Context.Channel.ID,
	})
	if err != nil {
		return WrapCommonErr(err)
	}

	app, err := i.appRepo.FindApp(ctx, appID)
	if err != nil {
		return WrapCommonErr(err)
	}

	res := i.handler.Invoke(ctx, app, req)

	event := FunctionInvokeEvent{
		Request:  req,
		Response: res,
		AppID:    appID,
	}
	i.callListeners(ctx, event)

	return res
}

func (i *Invoker) callListeners(ctx context.Context, event FunctionInvokeEvent) {
	for _, listener := range i.listeners {
		listener.OnInvoke(ctx, event)
	}
}

type InvokeHandler interface {
	Invoke(ctx context.Context, app *model.App, request JsonFunctionRequest) JsonFunctionResponse
}

type JsonFunctionRequest struct {
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Context ChannelContext  `json:"context"`
}

type ChannelContext struct {
	Caller  Caller  `json:"caller"`
	Channel Channel `json:"channel"`
}

type Channel struct {
	ID string `json:"id"`
}

type Caller struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type JsonFunctionResponse struct {
	Error  *Error          `json:"error"`
	Result json.RawMessage `json:"result"`
}

func (r *JsonFunctionResponse) IsError() bool {
	if r.Error == nil {
		return false
	}
	if len(r.Error.Type) <= 0 {
		return false
	}

	return true
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func WrapCommonErr(err error) JsonFunctionResponse {
	return JsonFunctionResponse{Error: &Error{Type: "common", Message: err.Error()}}
}

type FunctionInvokeEvent struct {
	AppID    string
	Request  JsonFunctionRequest
	Response JsonFunctionResponse
}

type FunctionRequestListener interface {
	OnInvoke(ctx context.Context, event FunctionInvokeEvent)
}
