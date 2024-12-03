package model

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

var (
	nameRegex     = regexp.MustCompile(`^[a-zA-Z_-]{1,20}$`)
	i18nNameRegex = regexp.MustCompile(`^[^@#$%:/\\\\]{1,20}$`)

	defaultNameRegex     = regexp.MustCompile(`^[^@#$%:/\\\\]{0,20}$`)
	i18nDefaultNameRegex = regexp.MustCompile(`^[^@#$%:/\\\\]{0,20}$`)

	maxDescriptionLength        = 40
	maxDefaultDescriptionLength = 40
)

type AppWidget struct {
	ID    string `json:"id"`
	AppID string `json:"appId"`
	Scope Scope  `json:"scope"`

	Name            string              `json:"name"`
	Description     *string             `json:"description,omitempty"`
	NameDescI18nMap map[string]*I18nMap `json:"nameDescI18nMap,omitempty"`

	DefaultName            *string             `json:"defaultName,omitempty"`
	DefaultDescription     *string             `json:"defaultDescription,omitempty"`
	DefaultNameDescI18nMap map[string]*I18nMap `json:"defaultNameDescI18nMap,omitempty"`

	ActionFunctionName string `json:"actionFunctionName"`
}

type Scope string

const (
	Front      = "front"
	Desk       = "desk"
	ScopeFront = Scope(Front)
	ScopeDesk  = Scope(Desk)
)

var validScopes = []Scope{ScopeFront, ScopeDesk}

func (s Scope) IsDefined() bool {
	for _, validScope := range validScopes {
		if validScope == s {
			return true
		}
	}
	return false
}

type I18nMap struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *AppWidget) Validate() error {
	if len(a.AppID) == 0 {
		return apierr.UnprocessableEntity(fmt.Errorf("appId is must not be empty"))
	}

	if !a.Scope.IsDefined() {
		return apierr.UnprocessableEntity(fmt.Errorf("scope %s is not defined", a.Scope))
	}

	// check name & description
	if !nameRegex.MatchString(a.Name) {
		return apierr.UnprocessableEntity(errors.New("name must be less than 20 letters with only alphabet"))
	}
	if a.Description != nil && utf8.RuneCountInString(*a.Description) > maxDescriptionLength {
		return apierr.UnprocessableEntity(errors.New("description length should be less than 40"))
	}

	// check defaultName & description
	if a.DefaultName != nil && !defaultNameRegex.MatchString(*a.DefaultName) {
		return apierr.UnprocessableEntity(errors.New("defaultName length should be less than 20"))
	}
	if a.DefaultDescription != nil && utf8.RuneCountInString(*a.DefaultDescription) > maxDefaultDescriptionLength {
		return apierr.UnprocessableEntity(errors.New("defaultDescription length should be less than 40"))
	}

	// check nameDescI18nMap
	if a.NameDescI18nMap != nil {
		for _, v := range a.NameDescI18nMap {
			if !i18nNameRegex.MatchString(v.Name) {
				return apierr.UnprocessableEntity(errors.New("name length should be less than 20"))
			}

			if utf8.RuneCountInString(v.Description) > maxDescriptionLength {
				return apierr.UnprocessableEntity(errors.New("description length should be less than 40"))
			}
		}
	}

	// check defaultNameDescI18nMap
	if a.DefaultNameDescI18nMap != nil {
		for _, v := range a.DefaultNameDescI18nMap {
			if !i18nDefaultNameRegex.MatchString(v.Name) {
				return apierr.UnprocessableEntity(errors.New("default name length should be less than 20"))
			}

			if utf8.RuneCountInString(v.Description) > maxDefaultDescriptionLength {
				return apierr.UnprocessableEntity(errors.New("default description length should be less than 40"))
			}
		}
	}

	if len(a.ActionFunctionName) == 0 {
		return apierr.UnprocessableEntity(fmt.Errorf("actionFunctionName is required"))
	}
	return nil
}
