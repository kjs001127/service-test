package domain

import (
	"context"
	"encoding/json"
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

type Invoker[REQ any, RES any] struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
	handler   InvokeHandler
}

func NewInvoker[REQ any, RES any](
	appChRepo AppChannelRepository,
	appRepo AppRepository,
	handler InvokeHandler,
) *Invoker[REQ, RES] {
	return &Invoker[REQ, RES]{
		appChRepo: appChRepo,
		appRepo:   appRepo,
		handler:   handler,
	}
}

type FunctionRequest[REQ any] struct {
	Endpoint
	Body[REQ]
}

type Endpoint struct {
	AppID        string
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
		Type string `json:"type"`
	} `json:"trigger"`
}

func (i *Invoker[REQ, RES]) InvokeChannelFunction(
	ctx context.Context,
	channelID string,
	request FunctionRequest[REQ],
) (RES, error) {
	var ret RES

	_, err := i.appChRepo.Fetch(ctx, Install{
		AppID:     request.AppID,
		ChannelID: channelID,
	})
	if err != nil {
		return ret, err
	}

	app, err := i.appRepo.FindApp(ctx, request.AppID)
	if err != nil {
		return ret, err
	}

	paramMarshaled, err := json.Marshal(request.Params)
	if err != nil {
		return ret, err
	}

	jsonRes, err := i.handler.Invoke(ctx, app, JsonFunctionRequest{
		Method:  request.FunctionName,
		Params:  paramMarshaled,
		Context: request.Context,
	})

	if err := json.Unmarshal(jsonRes, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

type FileStreamHandler interface {
	StreamFile(ctx context.Context, appID string, path string, writer io.Writer) error
}

type FileStreamer struct {
	appChRepo AppChannelRepository
	handler   FileStreamHandler
}

func NewFileStreamer(appChRepo AppChannelRepository, handler FileStreamHandler) *FileStreamer {
	return &FileStreamer{appChRepo: appChRepo, handler: handler}
}

type StreamRequest struct {
	Writer    io.Writer
	Path      string
	AppID     string
	ChannelID string
}

func (i *FileStreamer) StreamFile(ctx context.Context, req StreamRequest) error {

	_, err := i.appChRepo.Fetch(ctx, Install{
		AppID:     req.AppID,
		ChannelID: req.ChannelID,
	})
	if err != nil {
		return nil
	}

	return i.handler.StreamFile(ctx, req.AppID, req.Path, req.Writer)
}
