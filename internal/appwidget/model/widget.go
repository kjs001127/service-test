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

	Name            string              `json:"name"`
	Description     *string             `json:"description,omitempty"`
	NameDescI18nMap map[string]*I18nMap `json:"nameDescI18nMap,omitempty"`

	DefaultName            *string             `json:"defaultName,omitempty"`
	DefaultDescription     *string             `json:"defaultDescription,omitempty"`
	DefaultNameDescI18nMap map[string]*I18nMap `json:"defaultNameDescI18nMap,omitempty"`

	ActionFunctionName string `json:"actionFunctionName"`
}

type I18nMap struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *AppWidget) Validate() error {
	if len(a.AppID) == 0 {
		return apierr.BadRequest(fmt.Errorf("appId is must not be empty"))
	}

	// check name & description
	if !nameRegex.MatchString(a.Name) {
		return apierr.BadRequest(errors.New("name must be less than 20 letters with only alphabet"))
	}
	if a.Description != nil && utf8.RuneCountInString(*a.Description) > maxDescriptionLength {
		return apierr.BadRequest(errors.New("description length should be less than 40"))
	}

	// check defaultName & description
	if a.DefaultName != nil && !defaultNameRegex.MatchString(*a.DefaultName) {
		return apierr.BadRequest(errors.New("defaultName length should be less than 20"))
	}
	if a.DefaultDescription != nil && utf8.RuneCountInString(*a.DefaultDescription) > maxDefaultDescriptionLength {
		return apierr.BadRequest(errors.New("defaultDescription length should be less than 40"))
	}

	// check nameDescI18nMap
	if a.NameDescI18nMap != nil {
		for _, v := range a.NameDescI18nMap {
			if !i18nNameRegex.MatchString(v.Name) {
				return apierr.BadRequest(errors.New("name length should be less than 20"))
			}

			if utf8.RuneCountInString(v.Description) > maxDescriptionLength {
				return apierr.BadRequest(errors.New("description length should be less than 40"))
			}
		}
	}

	// check defaultNameDescI18nMap
	if a.DefaultNameDescI18nMap != nil {
		for _, v := range a.DefaultNameDescI18nMap {
			if !i18nDefaultNameRegex.MatchString(v.Name) {
				return apierr.BadRequest(errors.New("default name length should be less than 20"))
			}

			if utf8.RuneCountInString(v.Description) > maxDefaultDescriptionLength {
				return apierr.BadRequest(errors.New("default description length should be less than 40"))
			}
		}
	}

	if len(a.ActionFunctionName) == 0 {
		return apierr.BadRequest(fmt.Errorf("actionFunctionName is required"))
	}
	return nil
}
