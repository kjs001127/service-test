package domain

import (
	"context"
	"io"
)

type FunctionResponse any

type InvokeHandler interface {
	Invoke(ctx context.Context, request FunctionRequest, response FunctionResponse) error
}

type Invoker[RES any] struct {
	appChRepo AppChannelRepository
	handler   InvokeHandler
}

func NewInvoker[RES any](appChRepo AppChannelRepository, handler InvokeHandler) *Invoker[RES] {
	return &Invoker[RES]{appChRepo: appChRepo, handler: handler}
}

type FunctionRequest struct {
	Endpoint
	Body
}

type Endpoint struct {
	AppID        string
	FunctionName string
}

type Caller struct {
	Type string
	ID   string
}

type Body struct {
	Caller  Caller
	Context ChannelContext
	Params  any
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

func (i *Invoker[RES]) InvokeChannelFunction(
	ctx context.Context,
	channelID string,
	request FunctionRequest,
) (RES, error) {
	var ret RES

	_, err := i.appChRepo.Fetch(ctx, Install{
		AppID:     request.AppID,
		ChannelID: channelID,
	})
	if err != nil {
		return ret, nil
	}

	if err := i.handler.Invoke(ctx, request, &ret); err != nil {
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
