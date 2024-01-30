package domain

import (
	"fmt"

	"github.com/friendsofgo/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type ParamName string
type ParamDefinitions map[ParamName]*ParamDefinition

type ParamDefinition struct {
	Name           ParamName      `json:"name"`
	Type           TypeKey        `json:"type"`
	Required       bool           `json:"required"`
	Attributes     map[string]any `json:"attributes"`
	AlfDescription string         `json:"alfDescription"`
}

type ParamValidator struct {
	typeValidator *TypeValidator
}

func (v *ParamValidator) ValidateDefs(defs ParamDefinitions) error {
	for _, def := range defs {
		if err := v.typeValidator.Validate(def.Type, def.Attributes); err != nil {
			return errors.Wrap(err, fmt.Sprintf("param type %s is not valid", def.Type))
		}

		if len(def.Name) <= 0 {
			return fmt.Errorf("param name must not be empty")
		}
	}

	dupCheck := make(map[ParamName]bool)

	for _, def := range defs {
		if dupCheck[def.Name] {
			return fmt.Errorf("duplicate param name %s", def.Name)
		}
		dupCheck[def.Name] = true
	}
	return nil

}

type Arguments map[ParamName]any

func (r Arguments) ChannelContext() app.ChannelContext {
	return r[contextParamName].(app.ChannelContext)
}

type ArgsValidator struct {
	typeValidator *TypeValidator
}

func (v *ArgsValidator) ValidateArgs(defs ParamDefinitions, input Arguments) error {
	if err := v.validateExistence(defs, input); err != nil {
		return err
	}
	if err := v.validateTypes(defs, input); err != nil {
		return err
	}
	return nil
}

func (v *ArgsValidator) validateExistence(defs ParamDefinitions, params Arguments) error {
	if err := v.validateRequiredParams(defs, params); err != nil {
		return err
	}

	if err := v.validateOptionalParams(defs, params); err != nil {
		return err
	}

	return nil
}

func (v *ArgsValidator) validateRequiredParams(defs ParamDefinitions, params Arguments) error {
	for _, def := range defs {
		if !def.Required {
			continue
		}

		if _, ok := params[def.Name]; !ok {
			return fmt.Errorf("required param does not exists, required: %v", def)
		}
	}
	return nil
}

func (v *ArgsValidator) validateOptionalParams(defs ParamDefinitions, params Arguments) error {
	for name, _ := range params {
		if _, ok := defs[name]; !ok {
			return fmt.Errorf("param does not exist in paramDefinition, key: %v", name)
		}
	}

	return nil
}

func (v *ArgsValidator) validateTypes(defs ParamDefinitions, input Arguments) error {
	for name, val := range input {
		def := defs[name]
		if err := v.typeValidator.Validate(def.Type, val); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to validate tyep of param %s", name))
		}
	}
	return nil
}
