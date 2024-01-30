package domain

import (
	"context"

	"github.com/friendsofgo/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type AutocompleteClientRequest struct {
	Context app.ChannelContext
	Token   app.AuthToken

	Command Key
	Params  AutoCompleteArgs
}

type AutoCompleteArgs []AutoCompleteArg
type AutoCompleteArg struct {
	Focused bool
	Name    string
	Value   any
}

func (args AutoCompleteArgs) validate() error {
	for _, arg := range args {
		if len(arg.Name) <= 0 || arg.Value == nil {
			return errors.New("name and value must not be empty")
		}
	}
	return nil
}

type Choices []Choice
type Choice struct {
	Name  string
	Value any
}

func (choices Choices) validate() error {
	for _, c := range choices {
		if c.Value == nil || len(c.Name) == 0 {
			return errors.New("name and value of choice must not be empty")
		}
	}
	return nil
}

type AutoCompleteSvc interface {
	Invoke(ctx context.Context, request AutocompleteClientRequest) (Choices, error)
}

type AutoCompleteSvcImpl struct {
	repository CommandRepository
	requester  app.ContextFnInvoker[AutoCompleteRequest, Choices]
}

func (r *AutoCompleteSvcImpl) Invoke(ctx context.Context, request AutocompleteClientRequest) (Choices, error) {
	cmd, err := r.repository.Fetch(ctx, request.Command)
	if err != nil {
		return nil, err
	}

	if !cmd.AutoCompleteFunctionName.Valid {
		return nil, errors.New("autoCompleteFunction does not exist")
	}

	autoCompleteCtx := AutoCompleteRequest{
		Context: request.Context,
		Command: request.Command,
		Params:  request.Params,
	}

	ctxReq := app.Request[AutoCompleteRequest]{
		Token: request.Token,
		FunctionRequest: app.FunctionRequest[AutoCompleteRequest]{
			AppID:        request.Command.AppID,
			FunctionName: cmd.AutoCompleteFunctionName.String,
			Body:         autoCompleteCtx,
		},
	}

	return r.requester.Invoke(ctx, ctxReq)
}

type AutoCompleteRequest struct {
	Context app.ChannelContext
	Command Key
	Params  AutoCompleteArgs
}

func (a AutoCompleteRequest) ChannelContext() app.ChannelContext {
	return a.Context
}
