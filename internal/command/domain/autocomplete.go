package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/friendsofgo/errors"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
)

type AutoCompleteRequest struct {
	ChannelID string           `json:"channelId"`
	Command   CommandKey       `json:"command"`
	Body      AutoCompleteBody `json:"body"`
	Caller    Caller           `json:"caller"`
}

type AutoCompleteBody struct {
	Chat  Chat             `json:"chat"`
	Input AutoCompleteArgs `json:"input"`
}

type AutoCompleteArgs []*AutoCompleteArg
type AutoCompleteArg struct {
	Focused bool   `json:"focused"`
	Name    string `json:"name"`
	Value   any    `json:"value"`
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
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func (choices Choices) validate() error {
	for _, c := range choices {
		if c.Value == nil || len(c.Name) == 0 {
			return errors.New("name and value of choice must not be empty")
		}
	}
	return nil
}

type AutoCompleteInvoker struct {
	invoker *app.TypedInvoker[AutoCompleteBody, Choices]
	repo    CommandRepository
}

func NewAutoCompleteInvoker(
	invoker *app.TypedInvoker[AutoCompleteBody, Choices],
	repo CommandRepository,
) *AutoCompleteInvoker {
	return &AutoCompleteInvoker{invoker: invoker, repo: repo}
}

func (i *AutoCompleteInvoker) Invoke(ctx context.Context, request AutoCompleteRequest) (Choices, error) {
	cmd, err := i.repo.Fetch(ctx, request.Command)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if cmd.AutoCompleteFunctionName == nil {
		return nil, apierr.NotFound(errors.New("autocomplete function not found"))
	}

	res := i.invoker.Invoke(ctx, request.Command.AppID, app.TypedRequest[AutoCompleteBody]{
		FunctionName: *cmd.AutoCompleteFunctionName,
		Context: app.ChannelContext{
			Channel: app.Channel{
				ID: request.ChannelID,
			}, Caller: app.Caller{
				Type: request.Caller.Type,
				ID:   request.Caller.ID,
			},
		},
		Params: request.Body,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Result, nil
}
