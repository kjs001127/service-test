package domain

import (
	"fmt"
	"reflect"

	"github.com/channel-io/ch-app-store/internal/rpc/domain"
)

type Params map[string]any
type ParamType string

const (
	ParamTypeNumber = ParamType("number")
	ParamTypeString = ParamType("string")
	ParamTypeBool   = ParamType("bool")
)

var validParamTypes = []ParamType{ParamTypeNumber, ParamTypeString, ParamTypeBool}

func (p ParamType) isDefined() bool {
	for _, validParamType := range validParamTypes {
		if p == validParamType {
			return true
		}
	}
	return false
}

func (p ParamType) isAssignable(param any) bool {
	kind := reflect.TypeOf(param).Kind()
	switch p {
	case ParamTypeNumber:
		return kind == reflect.Int ||
			kind == reflect.Int32 ||
			kind == reflect.Int64 ||
			kind == reflect.Float64 ||
			kind == reflect.Float32
	case ParamTypeString:
		return kind == reflect.String
	case ParamTypeBool:
		return kind == reflect.Bool
	default:
		return false
	}
}

type ParamDefinition struct {
	Name     string
	Type     ParamType
	Required bool
}

func (d ParamDefinition) validate() error {
	if !d.Type.isDefined() {
		return fmt.Errorf("paramType %v is not valid", d.Type)
	}
	if len(d.Name) <= 0 {
		return fmt.Errorf("param name must not be empty")
	}
	return nil
}

type ParamDefinitions map[string]*ParamDefinition

func (d ParamDefinitions) validate() error {
	for _, def := range d {
		if err := def.validate(); err != nil {
			return err
		}
	}

	dupCheck := make(map[string]bool)
	for _, def := range d {
		if dupCheck[def.Name] {
			return fmt.Errorf("duplicate param name %s", def.Name)
		}
		dupCheck[def.Name] = true
	}

	return nil
}

func (d ParamDefinitions) validateParamInput(params domain.Params) error {
	if err := d.validateRequiredParams(params); err != nil {
		return err
	}

	if err := d.validateOptionalParams(params); err != nil {
		return err
	}

	return nil
}

func (d ParamDefinitions) validateRequiredParams(params domain.Params) error {
	for _, def := range d {
		if !def.Required {
			continue
		}

		if _, ok := params[def.Name]; !ok {
			return fmt.Errorf("required param does not exists, required: %v", def)
		}

		if !def.Type.isAssignable(params[def.Name]) {
			return fmt.Errorf("param type does not matches, required: %v, offered: %v", def, params[def.Name])
		}
	}
	return nil
}

func (d ParamDefinitions) validateOptionalParams(params domain.Params) error {
	for name, value := range params {
		if _, ok := d[name]; !ok {
			return fmt.Errorf("param does not exist in paramDefinition, key: %v", name)
		}
		if !d[name].Type.isAssignable(value) {
			return fmt.Errorf("param type does not matches, required: %v, offered: %v", d[name], value)
		}
	}

	return nil
}
