package domain

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const contextParamName = ParamName("context")

type InvokeSvc interface {
	Invoke(ctx context.Context, request CommandRequest) (Action, error)
}

type InvokeSvcImpl struct {
	repository    CommandRepository
	argsValidator ArgsValidator

	requester app.ContextFnInvoker[Arguments, Action]
}

type CommandRequest struct {
	Key
	Params Arguments

	Token   app.AuthToken
	Context app.ChannelContext
}

func (r *InvokeSvcImpl) Invoke(ctx context.Context, request CommandRequest) (Action, error) {
	cmd, err := r.repository.Fetch(ctx, request.Key)
	if err != nil {
		return Action{}, err
	}

	if err := r.argsValidator.ValidateArgs(cmd.ParamDefinitions, request.Params); err != nil {
		return Action{}, err
	}

	ctxReq := app.Request[Arguments]{
		Token: request.Token,
		FunctionRequest: app.FunctionRequest[Arguments]{
			AppID:        request.AppID,
			Body:         argumentsOf(request),
			FunctionName: cmd.ActionFunctionName,
		},
	}

	return r.requester.Invoke(ctx, ctxReq)
}

func argumentsOf(request CommandRequest) Arguments {
	params := make(Arguments)
	params[contextParamName] = request.Context
	for key, val := range request.Params {
		params[key] = val
	}
	return params
}
