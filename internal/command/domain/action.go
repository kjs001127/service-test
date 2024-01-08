package domain

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ActionType string
type Attributes []string

// Action is a result of Command.
// Must contain Type and Attributes according to that Type.
type Action struct {
	Type       ActionType      `json:"type"`
	Attributes json.RawMessage `json:"attributes"`
}

// ActionValidator checks if certain string is a JSON representation of Action
type ActionValidator struct {
	validActions map[ActionType]reflect.Type
}

func NewActionValidator(actions map[ActionType]reflect.Type) *ActionValidator {
	return &ActionValidator{validActions: actions}
}

func (v *ActionValidator) validate(ret []byte) error {
	var inputAction Action
	if err := json.Unmarshal(ret, &inputAction); err != nil {
		return fmt.Errorf("failed to marshal to action. cause:%w", err)
	}

	requiredAttrType, ok := v.validActions[inputAction.Type]
	if !ok {
		return fmt.Errorf("no action type found. input: %v", inputAction.Type)
	}
	requiredAttr := reflect.Zero(requiredAttrType)
	if err := json.Unmarshal(inputAction.Attributes, &requiredAttr); err != nil {
		return fmt.Errorf("failed to marshal attributes. cause:%w", err)
	}

	return nil
}
