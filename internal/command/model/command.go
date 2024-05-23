package model

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
)

type AlfMode string

const (
	AlfModeDisable   = AlfMode("disable")
	AlfModeRecommend = AlfMode("recommend")
)

var validAlfModes = []AlfMode{AlfModeRecommend, AlfModeDisable}
var maxAlfDescriptionLength = 500

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

var nameRegex = regexp.MustCompile(`^[a-zA-Z]{1,20}$`)
var i18nNameRegex = regexp.MustCompile(`^[^\s_]{1,20}$`)
var maxDescriptionLength = 100

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
	if len(c.AppID) == 0 {
		return apierr.BadRequest(fmt.Errorf("appID must not be empty"))
	}

	if !c.Scope.isDefined() {
		return apierr.BadRequest(fmt.Errorf("scope %s is not defined", c.Scope))
	}

	if err := c.ParamDefinitions.validate(); err != nil {
		return err
	}

	if !nameRegex.MatchString(c.Name) {
		return apierr.BadRequest(errors.New("name must be less than 20 letters with only alphabet and numbers"))
	}

	if c.Description != nil && utf8.RuneCountInString(*c.Description) > maxDescriptionLength {
		return apierr.BadRequest(fmt.Errorf("max description length is %d", maxDescriptionLength))
	}

	for _, i18n := range c.NameDescI18NMap {
		if !i18nNameRegex.MatchString(i18n.Name) {
			return apierr.BadRequest(errors.New("i18n name must be less than 20 letters without space"))
		}
		if utf8.RuneCountInString(i18n.Description) > maxDescriptionLength {
			return apierr.BadRequest(fmt.Errorf("max i18n description length is %d", maxDescriptionLength))
		}
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

	if c.AlfDescription != nil && utf8.RuneCountInString(*c.AlfDescription) > maxAlfDescriptionLength {
		return apierr.BadRequest(fmt.Errorf("alfDescription max length is %d", maxAlfDescriptionLength))
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
