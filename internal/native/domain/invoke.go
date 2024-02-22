package domain

import (
	"context"
	"encoding/json"
	"fmt"
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

func WrapCommonErr(err error) NativeFunctionResponse {
	return NativeFunctionResponse{
		Error: Error{
			Type:    "common",
			Message: err.Error(),
		},
	}
}

type NativeFunctionHandler interface {
	Handle(ctx context.Context, request NativeFunctionRequest) NativeFunctionResponse
	ListMethods() []string
}

type NativeFunctionInvoker struct {
	router map[string]NativeFunctionHandler
}

func NewNativeFunctionInvoker(handlers []NativeFunctionHandler) *NativeFunctionInvoker {
	ret := &NativeFunctionInvoker{router: make(map[string]NativeFunctionHandler)}
	for _, r := range handlers {
		ret.registerHandler(r)
	}
	return ret
}

func (i *NativeFunctionInvoker) registerHandler(r NativeFunctionHandler) {
	for _, m := range r.ListMethods() {
		if _, alreadyExists := i.router[m]; alreadyExists {
			panic(fmt.Errorf("method %s already has handler registered", m))
		}
		i.router[m] = r
	}
}

func (i *NativeFunctionInvoker) Invoke(
	ctx context.Context,
	req NativeFunctionRequest,
) NativeFunctionResponse {
	handler, ok := i.router[req.Method]
	if !ok {
		return NativeFunctionResponse{Error: Error{
			Type:    "common",
			Message: fmt.Sprintf("method not found: %s", req.Method),
		}}
	}
	return handler.Handle(ctx, req)
}
