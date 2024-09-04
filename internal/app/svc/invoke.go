package svc

import (
	"context"
	"encoding/json"

	"github.com/channel-io/ch-app-store/internal/app/model"
)

type Invoker interface {
	Invoke(ctx context.Context, appID string, req JsonFunctionRequest) JsonFunctionResponse
}

type InvokerImpl struct {
	querySvc *InstalledAppQuerySvc
	appRepo  AppRepository
	handler  InvokeHandler

	listeners []FunctionRequestListener
}

func NewInvoker(
	querySvc *InstalledAppQuerySvc,
	appRepo AppRepository,
	handler InvokeHandler,
	listeners []FunctionRequestListener,
) *InvokerImpl {
	return &InvokerImpl{querySvc: querySvc, appRepo: appRepo, handler: handler, listeners: listeners}
}

func (i *InvokerImpl) Invoke(ctx context.Context, appID string, req JsonFunctionRequest) JsonFunctionResponse {
	app, err := i.querySvc.QueryInstalledApp(ctx, model.InstallationID{
		AppID:     appID,
		ChannelID: req.Context.Channel.ID,
	})
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

type InvokeHandler interface {
	Invoke(ctx context.Context, app *model.App, request JsonFunctionRequest) JsonFunctionResponse
}

type JsonFunctionRequest struct {
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	Context ChannelContext  `json:"context"`
}

type ChannelContext struct {
	Caller  Caller  `json:"caller"`
	Channel Channel `json:"channel"`
}

type Channel struct {
	ID string `json:"id"`
}

type CallerType string

const (
	CallerTypeSystem  = CallerType("system")
	CallerTypeUser    = CallerType("user")
	CallerTypeManager = CallerType("manager")
)

const CallerIDOmitted = "-"

type Caller struct {
	Type CallerType `json:"type"`
	ID   string     `json:"id,omitempty"`
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

func (i *InvokerImpl) callListeners(ctx context.Context, event FunctionInvokeEvent) {
	for _, listener := range i.listeners {
		listener.OnInvoke(ctx, event)
	}
}
