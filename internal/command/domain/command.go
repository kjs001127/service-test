package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
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

	AppID string `json:"appID"`
	Name  string `json:"name"`
	Scope Scope  `json:"scope"`

	Description    null.String `json:"description"`
	AlfDescription null.String `json:"alfDescription"`

	ActionFunctionName       string           `json:"actionFunctionName"`
	AutoCompleteFunctionName null.String      `json:"autoCompleteFunctionName"`
	ParamDefinitions         ParamDefinitions `json:"paramDefinitions"`

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
