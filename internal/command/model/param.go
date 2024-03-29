package model

import (
	"github.com/pkg/errors"
)

type ParamName string
type ParamType string

type ParamDefinitions []*ParamDefinition

func (defs ParamDefinitions) toMap() map[ParamName]*ParamDefinition {
	ret := make(map[ParamName]*ParamDefinition)
	for _, def := range defs {
		ret[def.Name] = def
	}
	return ret
}

type ParamDefinition struct {
	Name            ParamName      `json:"name"`
	Type            ParamType      `json:"type"`
	Required        bool           `json:"required"`
	Description     string         `json:"description,omitempty"`
	Choices         Choices        `json:"choices,omitempty"`
	NameDescI18nMap map[string]any `json:"nameDescI18nMap,omitempty"`
	AutoComplete    bool           `json:"autoComplete,omitempty"`
	AlfDescription  string         `json:"alfDescription,omitempty"`
}

type Choices []Choice
type Choice struct {
	Name            string         `json:"name"`
	Value           any            `json:"value"`
	NameDescI18nMap map[string]any `json:"nameDescI18nMap,omitempty"`
}

func (choices Choices) validate() error {
	for _, c := range choices {
		if c.Value == nil || len(c.Name) == 0 {
			return errors.New("name and value of choice must not be empty")
		}
	}
	return nil
}
