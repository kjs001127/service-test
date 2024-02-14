package domain

import (
	"context"
	"encoding/json"
	"io"
)

type JsonFunctionRequest struct {
	FunctionName string
	Body         json.RawMessage
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
	handler InvokeHandler,
) *Invoker[REQ, RES] {
	return &Invoker[REQ, RES]{appChRepo: appChRepo, handler: handler}
}

type FunctionRequest[REQ any] struct {
	Endpoint
	Body[REQ]
}

type Endpoint struct {
	AppID        string
	FunctionName string
}

type Caller struct {
	Type string
	ID   string
}

type Body[REQ any] struct {
	Caller  Caller         `json:"caller"`
	Context ChannelContext `json:"context"`
	Params  REQ            `json:"params"`
}

type ChannelContext struct {
	Caller struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
	Channel struct {
		ID       string `json:"id"`
		Language string `json:"language"`
	}
	Chat struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
	Trigger struct {
		Type string `json:"type"`
	}
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

	jsonReq, err := json.Marshal(request.Body)
	if err != nil {
		return ret, err
	}

	app, err := i.appRepo.FindApp(ctx, request.AppID)
	if err != nil {
		return ret, err
	}

	jsonRes, err := i.handler.Invoke(ctx, app, JsonFunctionRequest{
		Body:         jsonReq,
		FunctionName: request.FunctionName,
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
