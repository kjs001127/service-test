package domain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/channel-io/ch-app-store/lib/log"
)

type Token struct {
	Type  string
	Value string
}

type NativeFunctionRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type NativeFunctionResponse struct {
	Error  NativeErr       `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type NativeErr struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func WrapCommonErr(err error) NativeFunctionResponse {
	return NativeFunctionResponse{
		Error: NativeErr{
			Type:    "common",
			Message: err.Error(),
		},
	}
}

type NativeFunctionHandler interface {
	Handle(ctx context.Context, token Token, request NativeFunctionRequest) NativeFunctionResponse
	ListMethods() []string
}

type NativeFunctionInvoker struct {
	router map[string]NativeFunctionHandler
	logger log.ContextAwareLogger
}

func NewNativeFunctionInvoker(handlers []NativeFunctionHandler, logger log.ContextAwareLogger) *NativeFunctionInvoker {
	ret := &NativeFunctionInvoker{router: make(map[string]NativeFunctionHandler), logger: logger}
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
	token Token,
	req NativeFunctionRequest,
) NativeFunctionResponse {
	handler, ok := i.router[req.Method]
	if !ok {
		i.logger.Warnw(ctx, "handler not found", "request", req)
		return NativeFunctionResponse{Error: NativeErr{
			Type:    "common",
			Message: fmt.Sprintf("method not found: %s", req.Method),
		}}
	}

	i.logger.Debugw(ctx, "native function request", "request", req)
	resp := handler.Handle(ctx, token, req)
	i.logger.Debugw(ctx, "native function response", "response", resp)

	if len(resp.Error.Type) >= 0 {
		i.logger.Warnw(ctx, "native function response errored", "request", req, "err", resp.Error)
	}
	return resp
}
