package domain

import (
	"context"
	"encoding/json"
	"errors"
	"io"
)

type JsonFunctionRequest struct {
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Context ChannelContext  `json:"context"`
}

type JsonFunctionResponse json.RawMessage

type InvokeHandler interface {
	Invoke(ctx context.Context, app *App, request JsonFunctionRequest) (JsonFunctionResponse, error)
}

type FunctionRequest[REQ any] struct {
	Endpoint
	Body[REQ]
}

type Endpoint struct {
	AppID        string
	ChannelID    string
	FunctionName string
}

type Channel struct {
	ID string `json:"id"`
}
type Caller struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Body[REQ any] struct {
	Context ChannelContext `json:"context"`
	Params  REQ            `json:"params"`
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

type InvokeTyper[REQ any, RES any] struct {
	invoker *Invoker
}

func NewInvokeTyper[REQ any, RES any](
	invoker *Invoker,
) *InvokeTyper[REQ, RES] {
	return &InvokeTyper[REQ, RES]{invoker: invoker}
}

func (i *InvokeTyper[REQ, RES]) InvokeChannelFunction(
	ctx context.Context,
	request FunctionRequest[REQ],
) (RES, error) {
	var ret RES

	marshaled, err := json.Marshal(request.Params)
	if err != nil {
		return ret, err
	}

	res, err := i.invoker.doInvoke(ctx, request.AppID, request.ChannelID, JsonFunctionRequest{
		Method:  request.FunctionName,
		Params:  marshaled,
		Context: request.Context,
	})
	if err != nil {
		return ret, err
	}

	if err := json.Unmarshal(res, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

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

func (i *Invoker) doInvoke(ctx context.Context, appID string, channelID string, request JsonFunctionRequest) (JsonFunctionResponse, error) {
	_, err := i.appChRepo.Fetch(ctx, Install{
		AppID:     appID,
		ChannelID: channelID,
	})
	if err != nil {
		return nil, err
	}

	app, err := i.appRepo.FindApp(ctx, appID)
	if err != nil {
		return nil, err
	}

	paramMarshaled, err := json.Marshal(request.Params)
	if err != nil {
		return nil, err
	}

	h, ok := i.handlerMap[app.Type]
	if !ok {
		return nil, errors.New("no handler found")
	}

	jsonRes, err := h.Invoke(ctx, app, JsonFunctionRequest{
		Method:  request.Method,
		Params:  paramMarshaled,
		Context: request.Context,
	})

	if err != nil {
		return nil, err
	}

	return jsonRes, nil
}

type FileStreamHandler interface {
	StreamFile(ctx context.Context, appID string, path string, writer io.Writer) error
}

type FileStreamer struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository

	handlerMap map[AppType]FileStreamHandler
}

func NewFileStreamer(
	appChRepo AppChannelRepository,
	appRepo AppRepository,
	handlers []Typed[FileStreamHandler]) *FileStreamer {
	ret := &FileStreamer{
		appChRepo:  appChRepo,
		appRepo:    appRepo,
		handlerMap: make(map[AppType]FileStreamHandler),
	}
	for _, h := range handlers {
		ret.handlerMap[h.Type] = h.Handler
	}
	return ret
}

type StreamRequest struct {
	Writer    io.Writer
	Path      string
	AppID     string
	ChannelID string
}

func (i *FileStreamer) StreamFile(ctx context.Context, req StreamRequest) error {
	app, err := i.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return err
	}

	h, ok := i.handlerMap[app.Type]
	if !ok {
		return errors.New("no handler found")
	}

	return h.StreamFile(ctx, req.AppID, req.Path, req.Writer)
}
