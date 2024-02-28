package dto

import (
	"encoding/json"
)

type NativeFunctionRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}
