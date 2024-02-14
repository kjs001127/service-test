package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/friendsofgo/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type AutoCompleteRequest struct {
	Command   Key    `json:"command"`
	ChannelID string `json:"channelId"`
	app.Body[AutoCompleteArgs]
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

type AutoCompleteInvoker struct {
	invoker app.Invoker[AutoCompleteArgs, Choices]
	repo    CommandRepository
}

func NewAutoCompleteInvoker(
	invoker app.Invoker[AutoCompleteArgs, Choices],
	repo CommandRepository,
) *AutoCompleteInvoker {
	return &AutoCompleteInvoker{invoker: invoker, repo: repo}
}

func (i *AutoCompleteInvoker) Invoke(ctx context.Context, request AutoCompleteRequest) (Choices, error) {
	cmd, err := i.repo.Fetch(ctx, request.Command)
	if err != nil {
		return nil, err
	}
	if cmd.AutoCompleteFunctionName == nil {
		return nil, apierr.NotFound(errors.New("autocomplete function not found"))
	}

	return i.invoker.InvokeChannelFunction(ctx, request.ChannelID, app.FunctionRequest[AutoCompleteArgs]{
		Endpoint: app.Endpoint{
			AppID:        request.Command.AppID,
			FunctionName: *cmd.AutoCompleteFunctionName,
		},
		Body: request.Body,
	})
}