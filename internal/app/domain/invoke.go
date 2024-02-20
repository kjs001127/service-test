package domain

import (
	"context"
	"encoding/json"
)

type Invoker struct {
	appChRepo  AppChannelRepository
	appRepo    AppRepository
	handlerMap map[AppType]InvokeHandler
}

func NewInvoker(appChRepo AppChannelRepository, appRepo AppRepository, handlers []Typed[InvokeHandler]) *Invoker {
	handlerMap := make(map[AppType]InvokeHandler)
	for _, h := range handlers {
		handlerMap[h.Type] = h.Handler
	}
	return &Invoker{appChRepo: appChRepo, appRepo: appRepo, handlerMap: handlerMap}
}

func (i *Invoker) Invoke(ctx context.Context, appID string, channelID string, request JsonFunctionRequest) JsonFunctionResponse {
	_, err := i.appChRepo.Fetch(ctx, Install{
		AppID:     appID,
		ChannelID: channelID,
	})
	if err != nil {
		return WrapErr(err)
	}

	app, err := i.appRepo.FindApp(ctx, appID)
	if err != nil {
		return WrapErr(err)
	}

	paramMarshaled, err := json.Marshal(request.Params)
	if err != nil {
		return WrapErr(err)
	}

	h, ok := i.handlerMap[app.Type]
	if !ok {
		return WrapErr(err)
	}

	return h.Invoke(ctx, app, JsonFunctionRequest{
		Method:  request.Method,
		Params:  paramMarshaled,
		Context: request.Context,
	})
}

type InvokeHandler interface {
	Invoke(ctx context.Context, app *App, request JsonFunctionRequest) JsonFunctionResponse
}

type JsonFunctionRequest struct {
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Context ChannelContext  `json:"context"`
}

type ChannelContext struct {
	Caller  Caller  `json:"caller"`
	Channel Channel `json:"channel"`
	Chat    struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"chat"`
	Trigger struct {
		Type       string            `json:"type"`
		Attributes map[string]string `json:"attributes"`
	} `json:"trigger"`
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

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func WrapErr(err error) JsonFunctionResponse {
	return JsonFunctionResponse{Error: &Error{Type: "common", Message: err.Error()}}
}

type TypedInvoker[REQ any, RES any] struct {
	invoker *Invoker
}

func NewTypedInvoker[REQ any, RES any](
	invoker *Invoker,
) *TypedInvoker[REQ, RES] {
	return &TypedInvoker[REQ, RES]{invoker: invoker}
}

func (i *TypedInvoker[REQ, RES]) Invoke(
	ctx context.Context,
	request TypedRequest[REQ],
) TypedResponse[RES] {
	var ret RES

	marshaled, err := json.Marshal(request.Params)
	if err != nil {
		return TypedResponse[RES]{Error: &Error{Type: "appstore", Message: err.Error()}}
	}

	res := i.invoker.Invoke(ctx, request.AppID, request.ChannelID, JsonFunctionRequest{
		Method:  request.FunctionName,
		Params:  marshaled,
		Context: request.Context,
	})
	if res.Error != nil {
		return TypedResponse[RES]{Error: res.Error}
	}

	if err := json.Unmarshal(res.Result, &ret); err != nil {
		return TypedResponse[RES]{Error: &Error{Type: "appstore", Message: err.Error()}}
	}

	return TypedResponse[RES]{Result: ret}
}

type TypedRequest[REQ any] struct {
	Endpoint
	Body[REQ]
}

type Endpoint struct {
	AppID        string
	ChannelID    string
	FunctionName string
}

type Body[REQ any] struct {
	Context ChannelContext `json:"context"`
	Params  REQ            `json:"params"`
}

type TypedResponse[REQ any] struct {
	Result REQ
	Error  *Error
}
