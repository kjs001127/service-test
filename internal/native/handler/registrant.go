package handler

import (
	"context"
)

type NativeFunctionHandler func(ctx context.Context, token Token, request NativeFunctionRequest) NativeFunctionResponse

type NativeFunctionRegistry interface {
	Register(method string, handler NativeFunctionHandler)
}

type NativeFunctionRegistrant interface {
	RegisterTo(registry NativeFunctionRegistry)
}
