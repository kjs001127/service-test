package model

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
)

type ParamName string
type ParamType string

var paramNameRegex = regexp.MustCompile(`^[a-zA-Z_-]{1,20}$`)
var paramNameI18nRegex = regexp.MustCompile(`^[^\s]{1,20}$`)

var maxParamDescriptionLength = 50

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

	if !paramNameRegex.MatchString(string(d.Name)) {
		return apierr.BadRequest(errors.Errorf("param name %s must only have alphabet and numbers with letters below 20", d.Name))
	}

	if utf8.RuneCountInString(d.Description) > maxParamDescriptionLength {
		return apierr.BadRequest(fmt.Errorf("max param description length is %d", maxParamDescriptionLength))
	}

	for _, i18n := range d.NameDescI18nMap {
		if !paramNameI18nRegex.MatchString(i18n.Name) {
			return apierr.BadRequest(errors.Errorf("i18n param name %s must only have letters without space below 20", d.Name))
		}
		if utf8.RuneCountInString(i18n.Description) > maxParamDescriptionLength {
			return apierr.BadRequest(fmt.Errorf("max i18n param description length is %d", maxParamDescriptionLength))
		}
	}

	if utf8.RuneCountInString(d.AlfDescription) > maxAlfDescriptionLength {
		return apierr.BadRequest(fmt.Errorf("alfDescription max length is %d", maxAlfDescriptionLength))
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
