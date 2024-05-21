package model

import (
	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
)

type ParamName string
type ParamType string

var validParamTypes = []ParamType{
	ParamTypeString, ParamTypeFloat, ParamTypeInt, ParamTypeBool,
}

const (
	ParamTypeString = ParamType("string")
	ParamTypeFloat  = ParamType("float")
	ParamTypeInt    = ParamType("int")
	ParamTypeBool   = ParamType("bool")
)

type ParamDefinitions []*ParamDefinition

func (defs ParamDefinitions) toMap() map[ParamName]*ParamDefinition {
	ret := make(map[ParamName]*ParamDefinition)
	for _, def := range defs {
		ret[def.Name] = def
	}
	return ret
}

func (defs ParamDefinitions) validate() error {
	for _, def := range defs {
		if err := def.validate(); err != nil {
			return err
		}
	}
	return nil
}

type ParamDefinition struct {
	Name            ParamName                `json:"name"`
	Type            ParamType                `json:"type"`
	Required        bool                     `json:"required"`
	Description     string                   `json:"description,omitempty"`
	Choices         Choices                  `json:"choices,omitempty"`
	NameDescI18nMap map[string]ParamDefI18ns `json:"nameDescI18nMap,omitempty"`
	AutoComplete    bool                     `json:"autoComplete,omitempty"`
	AlfDescription  string                   `json:"alfDescription,omitempty"`
}

type ParamDefI18ns struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d *ParamDefinition) validate() error {
	if !d.isValidType() {
		return apierr.BadRequest(errors.Errorf("param name %s has invalid type %s", d.Name, d.Type))
	}
	if err := d.Choices.validate(); err != nil {
		return err
	}

	return nil
}

func (d *ParamDefinition) isValidType() bool {
	for _, t := range validParamTypes {
		if t == d.Type {
			return true
		}
	}
	return false
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
