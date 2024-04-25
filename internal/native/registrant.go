package native

import (
	"context"
)

type FunctionHandler func(ctx context.Context, token Token, request FunctionRequest) FunctionResponse

type FunctionRegistry interface {
	Register(method string, handler FunctionHandler)
}

type FunctionRegistrant interface {
	RegisterTo(registry FunctionRegistry)
}
