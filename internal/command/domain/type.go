package domain

import (
	"fmt"
)

type TypeKey string
type Type interface {
	Check(data any) error
}

type TypeValidator struct {
	validTypes map[TypeKey]Type
}

func (v *TypeValidator) Validate(key TypeKey, input any) error {
	t, ok := v.validTypes[key]
	if !ok {
		return fmt.Errorf("invalid type %s", t)
	}

	return t.Check(input)
}
