package svc

import (
	"context"

	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/model"
)

type Invoker struct {
	repository        CommandRepository
	requester         app.TypedInvoker[CommandBody, Action]
	activationChecker ActivationSvc
	listeners         []CommandRequestListener
}

func NewInvoker(
	repository CommandRepository,
	requester app.TypedInvoker[CommandBody, Action],
	activationSvc ActivationSvc,
	listeners []CommandRequestListener,
) *Invoker {
	return &Invoker{repository: repository, requester: requester, listeners: listeners, activationChecker: activationSvc}
}

func (r *Invoker) Invoke(ctx context.Context, request CommandRequest) (Action, error) {

	if err := r.checkActivated(ctx, request); err != nil {
		return Action{}, err
	}

	cmd, err := r.repository.Fetch(ctx, request.CommandKey)
	if err != nil {
		return Action{}, errors.WithStack(err)
	}

	cmdReq := app.TypedRequest[CommandBody]{
		FunctionName: cmd.ActionFunctionName,
		Params:       request.CommandBody,
		Context: app.ChannelContext{
			Channel: app.Channel{
				ID: request.ChannelID,
			},
			Caller: app.Caller{
				ID:   request.Caller.ID,
				Type: request.Caller.Type,
			},
		},
	}

	ret := r.requester.Invoke(ctx, cmd.AppID, cmdReq)

	// call command invoke event listeners
	event := CommandInvokeEvent{
		Request: request,
		Result:  ret.Result,
		ID:      cmd.ID,
		Err:     nil,
	}
	if ret.IsError() {
		event.Err = ret.Error
	}
	r.callListeners(ctx, event)

	if ret.IsError() {
		return Action{}, ret.Error
	}

	return ret.Result, nil
}

func (r *Invoker) checkActivated(ctx context.Context, request CommandRequest) error {
	activated, err := r.activationChecker.Check(ctx, request.CommandKey, request.ChannelID)
	if err != nil {
		return err
	}

	if !activated {
		return errors.New("command inactive on this channel")
	}

	return nil
}

func (r *Invoker) callListeners(ctx context.Context, event CommandInvokeEvent) {
	for _, listener := range r.listeners {
		listener.OnInvoke(ctx, event)
	}
}

type CommandRequestListener interface {
	OnInvoke(ctx context.Context, event CommandInvokeEvent)
}

type CommandInvokeEvent struct {
	ID      string
	Err     error
	Result  Action
	Request CommandRequest
}

type CommandRequest struct {
	model.CommandKey
	CommandBody
	ChannelID string
	Caller    Caller
}

type ParamInput map[model.ParamName]any

type CommandBody struct {
	Chat     Chat       `json:"chat"`
	Trigger  Trigger    `json:"trigger"`
	Input    ParamInput `json:"input"`
	Language string     `json:"language"`
}

type Chat struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Trigger struct {
	Type       string            `json:"type"`
	Attributes map[string]string `json:"attributes"`
}

type Caller struct {
	Type app.CallerType `json:"type"`
	ID   string         `json:"id"`
}

type ActionType string
type Action struct {
	Type       ActionType     `json:"type"`
	Attributes map[string]any `json:"attributes"`
}
