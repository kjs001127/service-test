package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/channel-io/ch-app-store/internal/resource/domain"
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
	ID           string
	AppID        string
	FunctionName string

	Name  string
	Scope Scope

	Description string

	ParamDefinitions ParamDefinitions

	UpdatedAt time.Time
	CreatedAt time.Time
}

func (c *Command) GetID() string {
	return c.ID
}

func (c *Command) SetID(s string) {
	c.ID = s
}

func (c *Command) GetAppID() string {
	return c.AppID
}

func (c *Command) SetAppID(s string) {
	c.AppID = s
}

func (c *Command) GetName() string {
	return c.Name
}

func (c *Command) SetName(s string) {
	c.Name = s
}

func (c *Command) Validate() error {
	if len(c.AppID) == 0 || len(c.Name) == 0 {
		return fmt.Errorf("appID, name must not be empty")
	}

	if !c.Scope.isDefined() {
		return fmt.Errorf("scope %s is not defined", c.Scope)
	}

	if err := c.ParamDefinitions.validate(); err != nil {
		return err
	}

	return nil
}

type Query struct {
	AppID string
	Scope Scope
	Query string

	Since string
	Limit int
}

type CommandRepository interface {
	FetchByQuery(ctx context.Context, query Query) ([]*Command, error)
	domain.ResourceRepository[*Command]
}
