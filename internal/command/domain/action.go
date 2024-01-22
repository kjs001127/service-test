package domain

import (
	"reflect"

	"github.com/friendsofgo/errors"
)

type ActionType reflect.Type

// Action is a result of Command.
// Must contain Type and Attributes according to that Type.
type Action struct {
	Type       TypeKey
	Attributes map[string]any
}

type ActionValidator struct {
	typeValidator *TypeValidator
}

func NewActionValidator(typeValidator *TypeValidator) *ActionValidator {
	return &ActionValidator{typeValidator: typeValidator}
}

func (v *ActionValidator) ValidateAction(input Action) error {

	if len(input.Type) <= 0 {
		return errors.New("type must not be empty")
	}

	if err := v.typeValidator.Validate(input.Type, input.Attributes); err != nil {
		return errors.Wrap(err, "action attribute parse fail")
	}

	return nil
}
