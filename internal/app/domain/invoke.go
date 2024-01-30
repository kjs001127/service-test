package domain

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/volatiletech/null/v8"
)

type Invoker[REQ any, RES any] struct {
	appChRepo AppChannelRepository
	appRepo   AppRepository
}

type ChannelFunctionRequest[REQ any] struct {
	ChannelID string
	FunctionRequest[REQ]
}

type FunctionRequest[REQ any] struct {
	AppID        string
	FunctionName string
	Body         REQ
}

func (i *Invoker[REQ, RES]) InvokeInChannel(
	ctx context.Context,
	request ChannelFunctionRequest[REQ],
) (RES, error) {
	var empty RES

	_, err := i.appChRepo.Fetch(ctx, Install{AppID: request.AppID, ChannelID: request.ChannelID})
	if err != nil {
		return empty, nil
	}

	return i.Invoke(ctx, request.FunctionRequest)
}

func (i *Invoker[REQ, RES]) Invoke(
	ctx context.Context,
	req FunctionRequest[REQ],
) (RES, error) {
	var ret RES

	installedApp, err := i.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return ret, err
	}

	marshaled, err := json.Marshal(req.Body)
	if err != nil {
		return ret, err
	}

	res, err := installedApp.Invoke(ctx, req.FunctionName, null.JSONFrom(marshaled))
	if err != nil {
		return ret, err
	}

	if !res.Valid {
		return ret, nil
	}

	if err := json.Unmarshal(res.JSON, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

type FileStreamer struct {
	appRepo AppRepository
}

type StreamRequest struct {
	Writer io.Writer
	Path   string
	AppID  string
}

func (i *FileStreamer) StreamFile(ctx context.Context, req StreamRequest) error {
	if !strings.HasPrefix(req.Path, "/") {
		req.Path = "/" + req.Path
	}

	installedApp, err := i.appRepo.FindApp(ctx, req.AppID)
	if err != nil {
		return err
	}

	return installedApp.StreamFile(ctx, req.Path, req.Writer)
}
