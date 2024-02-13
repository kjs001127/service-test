package domain

import (
	"fmt"
	"reflect"

	"github.com/friendsofgo/errors"
)

type ParamInput map[ParamName]any

type ParamName string
type ParamType string

const (
	ParamTypeInt    = ParamType("int")
	ParamTypeString = ParamType("string")
	ParamTypeFloat  = ParamType("float")
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

type ParamDefinition struct {
	Name           ParamName `json:"name"`
	Type           ParamType `json:"type"`
	Required       bool      `json:"required"`
	Description    string    `json:"description"`
	Choices        Choices   `json:"choices"`
	AutoComplete   bool      `json:"autoComplete"`
	AlfDescription string    `json:"alfDescription"`
}

type Validator func(param any) error
type TypeValidator map[ParamType]Validator

type ParamValidator struct {
	typeValidator TypeValidator
}

func NewParamValidator() *ParamValidator {
	ret := &ParamValidator{typeValidator: make(TypeValidator)}
	ret.typeValidator[ParamTypeInt] = validateInt
	ret.typeValidator[ParamTypeString] = validateString
	ret.typeValidator[ParamTypeFloat] = validateFloat
	ret.typeValidator[ParamTypeBool] = validateBool

	return ret
}

func (v *ParamValidator) ValidateDefs(defs ParamDefinitions) error {
	for _, def := range defs {
		if _, ok := v.typeValidator[def.Type]; !ok {
			return fmt.Errorf("param type %s is not valid", def.Type)
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

func (v *ParamValidator) ValidateParamInput(defs ParamDefinitions, input ParamInput) error {
	if err := v.validateExistence(defs, input); err != nil {
		return err
	}
	if err := v.validateTypes(defs, input); err != nil {
		return err
	}
	return nil
}

func (v *ParamValidator) validateExistence(defs ParamDefinitions, params ParamInput) error {
	if err := v.validateRequiredParams(defs, params); err != nil {
		return err
	}

	if err := v.validateOptionalParams(defs, params); err != nil {
		return err
	}

	return nil
}

func (v *ParamValidator) validateRequiredParams(defs ParamDefinitions, params ParamInput) error {
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

func (v *ParamValidator) validateOptionalParams(defs ParamDefinitions, params ParamInput) error {
	defMap := defs.toMap()
	for name, _ := range params {
		if _, ok := defMap[name]; !ok {
			return fmt.Errorf("param does not exist in paramDefinition, key: %v", name)
		}
	}

	return nil
}

func (v *ParamValidator) validateTypes(defs ParamDefinitions, input ParamInput) error {
	defMap := defs.toMap()
	for name, val := range input {
		def := defMap[name]
		validator, ok := v.typeValidator[def.Type]
		if !ok {
			return errors.New("invalid type")
		}
		if err := validator(val); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to validate type of param %s", name))
		}
	}
	return nil
}

func validateInt(param any) error {
	switch reflect.TypeOf(param).Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return nil
	default:
		return errors.New("not a int")
	}
}

func validateFloat(param any) error {
	switch reflect.TypeOf(param).Kind() {
	case reflect.Float64, reflect.Float32:
		return nil
	default:
		return errors.New("not a float")
	}
}

func validateString(param any) error {
	switch reflect.TypeOf(param).Kind() {
	case reflect.String:
		return nil
	default:
		return errors.New("not a string")
	}
}

func validateBool(param any) error {
	switch reflect.TypeOf(param).Kind() {
	case reflect.Bool:
		return nil
	default:
		return errors.New("not a bool")
	}
}
