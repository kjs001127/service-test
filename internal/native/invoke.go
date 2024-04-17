package native

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

type FunctionRequest struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type FunctionResponse struct {
	Error  Err             `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type Err struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func WrapCommonErr(err error) FunctionResponse {
	return FunctionResponse{
		Error: Err{
			Type:    "common",
			Message: err.Error(),
		},
	}
}

func ResultSuccess(resp json.RawMessage) FunctionResponse {
	return FunctionResponse{
		Result: resp,
	}
}

func Empty() FunctionResponse {
	return FunctionResponse{}
}

type FunctionInvoker struct {
	router map[string]FunctionHandler
	logger log.ContextAwareLogger
}

func NewNativeFunctionInvoker(handlers []FunctionRegistrant, logger log.ContextAwareLogger) *FunctionInvoker {
	ret := &FunctionInvoker{router: make(map[string]FunctionHandler), logger: logger}
	for _, r := range handlers {
		r.RegisterTo(ret)
	}
	return ret
}

func (i *FunctionInvoker) Register(method string, handler FunctionHandler) {
	i.router[method] = handler
}

func (i *FunctionInvoker) Invoke(
	ctx context.Context,
	token Token,
	req FunctionRequest,
) FunctionResponse {
	handler, ok := i.router[req.Method]
	if !ok {
		i.logger.Warnw(ctx, "handler not found", "request", req)
		return FunctionResponse{Error: Err{
			Type:    "common",
			Message: fmt.Sprintf("method not found: %s", req.Method),
		}}
	}

	i.logger.Debugw(ctx, "native function request", "request", req)
	resp := handler(ctx, token, req)
	i.logger.Debugw(ctx, "native function response", "response", resp)

	if len(resp.Error.Type) >= 0 {
		i.logger.Warnw(ctx, "native function response errored", "request", req, "err", resp.Error)
	}
	return resp
}
