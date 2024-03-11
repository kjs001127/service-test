package svc

import (
	"context"
	"encoding/json"
)

type TypedRequest[REQ any] struct {
	FunctionName string         `json:"functionName"`
	Context      ChannelContext `json:"context"`
	Params       REQ            `json:"params"`
}

type TypedResponse[REQ any] struct {
	Result REQ    `json:"result"`
	Error  *Error `json:"error,omitempty"`
}

type TypedInvoker[REQ any, RES any] struct {
	invoker *Invoker
}

func NewTypedInvoker[REQ any, RES any](
	invoker *Invoker,
) *TypedInvoker[REQ, RES] {
	return &TypedInvoker[REQ, RES]{invoker: invoker}
}

func (i *TypedInvoker[REQ, RES]) Invoke(
	ctx context.Context,
	appID string,
	request TypedRequest[REQ],
) TypedResponse[RES] {
	var ret RES

	marshaled, err := json.Marshal(request.Params)
	if err != nil {
		return TypedResponse[RES]{Error: &Error{Type: "appstore", Message: err.Error()}}
	}

	res := i.invoker.Invoke(ctx, appID, JsonFunctionRequest{
		Method:  request.FunctionName,
		Params:  marshaled,
		Context: request.Context,
	})
	if res.Error != nil {
		return TypedResponse[RES]{Error: res.Error}
	}

	if err := json.Unmarshal(res.Result, &ret); err != nil {
		return TypedResponse[RES]{Error: &Error{Type: "common", Message: err.Error()}}
	}

	return TypedResponse[RES]{Result: ret}
}
