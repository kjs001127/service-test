package model

import (
	"fmt"
	"time"
)

type AlfMode string

const (
	AlfModeDisable   = AlfMode("disable")
	AlfModeRecommend = AlfMode("recommend")
)

var validAlfModes = []AlfMode{AlfModeRecommend, AlfModeDisable}

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

	Description     *string            `json:"description"`
	NameDescI18NMap map[string]I18nMap `json:"nameDescI18nMap"`

	AlfDescription *string `json:"alfDescription"`
	AlfMode        AlfMode `json:"alfMode"`

	ActionFunctionName       string  `json:"actionFunctionName"`
	AutoCompleteFunctionName *string `json:"autoCompleteFunctionName"`

	ParamDefinitions ParamDefinitions `json:"paramDefinitions"`

	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type I18nMap struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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

	if err := c.validateAlf(); err != nil {
		return err
	}

	return nil
}

func (c *Command) validateAlf() error {
	for _, mode := range validAlfModes {
		if mode == c.AlfMode {
			return nil
		}
	}
	return fmt.Errorf("alfMode %s is not valid mode", c.AlfMode)
}

type Query struct {
	AppIDs []string
	Scope  Scope
}

type CommandKey struct {
	AppID string
	Scope Scope
	Name  string
}
