package domain

import (
	"context"
	"io"
)

type Invoker[RES any] struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

func NewInvoker[RES any](appChRepo AppChannelRepository, appRepo AppRepository) *Invoker[RES] {
	return &Invoker[RES]{appChRepo: appChRepo, appRepo: appRepo}
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
	Channel struct {
		ID string `json:"id"`
	}
	Chat struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
}

func (i *Invoker[RES]) InvokeChannelFunction(
	ctx context.Context,
	channelID string,
	request FunctionRequest,
) (RES, error) {
	var res RES

	_, err := i.appChRepo.Fetch(ctx, Install{
		AppID:     request.AppID,
		ChannelID: channelID,
	})
	if err != nil {
		return res, nil
	}

	return i.InvokeFunction(ctx, request)
}

func (i *Invoker[RES]) InvokeFunction(
	ctx context.Context,
	request FunctionRequest,
) (RES, error) {
	var res RES

	installedApp, err := i.appRepo.FindApp(ctx, request.AppID)
	if err != nil {
		return res, err
	}

	appReq := AppRequest{
		FunctionName: request.FunctionName,
		Body:         request.Body,
	}

	if err := installedApp.Invoke(ctx, appReq, &res); err != nil {
		return res, err
	}

	return res, err
}

type FileStreamer struct {
	appRepo AppRepository
}

func NewFileStreamer(appRepo AppRepository) *FileStreamer {
	return &FileStreamer{appRepo: appRepo}
}

type StreamRequest struct {
	Writer io.Writer
	Path   string
	AppID  string
}

func (i *FileStreamer) StreamFile(ctx context.Context, req StreamRequest) error {

	installedApp, err := i.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return err
	}

	return installedApp.StreamFile(ctx, req.Path, req.Writer)
}
