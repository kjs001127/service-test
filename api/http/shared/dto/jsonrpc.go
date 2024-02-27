package dto

import (
	"encoding/json"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type JsonFunctionRequest struct {
	Method  string             `json:"method"`
	Params  json.RawMessage    `json:"params"`
	Context app.ChannelContext `json:"context"`
}

type NativeFunctionRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}
