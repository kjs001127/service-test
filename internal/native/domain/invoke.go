package domain

import (
	"context"
	"encoding/json"
)

type Token struct {
	Type  string
	Value string
}

type NativeFunctionRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Token  Token           `json:"token"`
}

type NativeFunctionResponse struct {
	Error  Error           `json:"error"`
	Result json.RawMessage `json:"result"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func WrapError(err error) NativeFunctionResponse {
	return NativeFunctionResponse{
		Error: Error{
			Type:    "common",
			Message: err.Error(),
		},
	}
}

type NativeFunctions map[string]NativeFunction
type NativeFunction func(ctx context.Context, params json.RawMessage, token Token) NativeFunctionResponse

type NativeFunctionProvider interface {
	Provide() NativeFunctions
}

type NativeFunctionInvoker struct {
	functions map[string]NativeFunction
}

func NewNativeFunctionInvoker(provider NativeFunctionProvider) *NativeFunctionInvoker {
	return &NativeFunctionInvoker{functions: provider.Provide()}
}

func (i *NativeFunctionInvoker) Invoke(
	ctx context.Context,
	req NativeFunctionRequest,
) NativeFunctionResponse {
	fn, ok := i.functions[req.Method]
	if !ok {
		return NativeFunctionResponse{Error: Error{
			Type:    "common",
			Message: "method not found",
		}}
	}
	return fn(ctx, req.Params, req.Token)
}
