package domain

import (
	"context"
	"fmt"
	"time"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type Scope string

const (
	ScopeFront = Scope("front")
	ScopeDesk  = Scope("desk")
)

var validScopes = []Scope{ScopeFront, ScopeDesk}

func (s Scope) isDefined() bool {
	for _, validScope := range validScopes {
		if validScope == s {
			return true
		}
	}
	return false
}

type Command struct {
	ID string `json:"id"`

	AppID string `json:"appId"`
	Name  string `json:"name"`
	Scope Scope  `json:"scope"`

	NameI18nMap map[string]string `json:"nameI18nMap"`
	Description *string           `json:"description"`

	AlfDescription *string `json:"alfDescription"`
	AlfMode        string  `json:"alfMode"`

	ActionFunctionName       string  `json:"actionFunctionName"`
	AutoCompleteFunctionName *string `json:"autoCompleteFunctionName"`

	ParamDefinitions ParamDefinitions `json:"paramDefinitions"`

	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (c *Command) Validate() error {
	if len(c.AppID) == 0 || len(c.Name) == 0 {
		return fmt.Errorf("appID, name must not be empty")
	}

	if !c.Scope.isDefined() {
		return fmt.Errorf("scope %s is not defined", c.Scope)
	}

	return nil
}

type Query struct {
	AppIDs []string
	Scope  Scope
}

type Key struct {
	AppID string
	Scope Scope
	Name  string
}

type CommandRepository interface {
	FetchByQuery(ctx context.Context, query Query) ([]*Command, error)
	Fetch(ctx context.Context, key Key) (*Command, error)

	FetchAllByAppIDs(ctx context.Context, appIDs []string) ([]*Command, error)
	FetchAllByAppID(ctx context.Context, appID string) ([]*Command, error)

	Delete(ctx context.Context, key Key) error
	Save(ctx context.Context, resource *Command) (*Command, error)
}

type Invoker struct {
	repository CommandRepository

	requester *app.InvokeTyper[ParamInput, Action]
	validator *ParamValidator
}

func NewInvoker(
	repository CommandRepository,
	requester *app.InvokeTyper[ParamInput, Action],
	validator *ParamValidator,
) *Invoker {
	return &Invoker{repository: repository, requester: requester, validator: validator}
}

type CommandRequest struct {
	Key
	ChannelID string
	app.Body[ParamInput]
}

func (r *Invoker) Invoke(ctx context.Context, request CommandRequest) (Action, error) {
	cmd, err := r.repository.Fetch(ctx, request.Key)
	if err != nil {
		return Action{}, err
	}

	ctxReq := app.FunctionRequest[ParamInput]{
		Endpoint: app.Endpoint{
			AppID:        cmd.AppID,
			ChannelID:    request.ChannelID,
			FunctionName: cmd.ActionFunctionName,
		},
		Body: request.Body,
	}

	ret := r.requester.InvokeChannelFunction(ctx, ctxReq)
	return ret.Result, ret.Error
}
