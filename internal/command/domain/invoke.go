package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type InvokeSvc struct {
	repository CommandRepository

	requester app.Invoker[Action]
	validator ParamValidator
}

func NewInvokeSvc(
	repository CommandRepository,
	requester app.Invoker[Action],
	validator ParamValidator,
) *InvokeSvc {
	return &InvokeSvc{repository: repository, requester: requester, validator: validator}
}

type CommandRequest struct {
	Key
	ChannelID string
	app.Body
}

func (r *InvokeSvc) Invoke(ctx context.Context, request CommandRequest) (Action, error) {
	cmd, err := r.repository.Fetch(ctx, request.Key)
	if err != nil {
		return Action{}, err
	}

	ctxReq := app.FunctionRequest{
		Endpoint: app.Endpoint{
			AppID:        cmd.AppID,
			FunctionName: cmd.ActionFunctionName,
		},
		Body: request.Body,
	}

	return r.requester.InvokeChannelFunction(ctx, request.ChannelID, ctxReq)
}
