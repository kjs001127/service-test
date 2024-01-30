package domain

import (
	"context"
)

type ChannelContextAware interface {
	ChannelContext() ChannelContext
}

type AuthToken string
type AuthInfo struct {
	ChannelID string
	Token     AuthToken
}

type ChannelContext struct {
	Channel Channel
	Caller  Caller
	Trigger Trigger
}

type Channel struct {
	ID     string
	Locale string
}

type Chat struct {
	Type string
	ID   string
}

type Caller struct {
	Type string
	ID   string
}

type Trigger struct {
	Type string
	ID   string
}

type ContextAuthorizer interface {
	Authorize(ctx context.Context, target ChannelContext, info AuthInfo) error
}

type ContextFnInvoker[REQ ChannelContextAware, RES any] interface {
	Invoke(ctx context.Context, req Request[REQ]) (RES, error)
}

type ContextFnInvokerImpl[REQ ChannelContextAware, RES any] struct {
	invoker    *Invoker[REQ, RES]
	authorizer ContextAuthorizer
}

func NewContextFnInvoker[REQ ChannelContextAware, RES any](
	invoker *Invoker[REQ, RES],
	authorizer ContextAuthorizer,
) *ContextFnInvokerImpl[REQ, RES] {
	return &ContextFnInvokerImpl[REQ, RES]{invoker: invoker, authorizer: authorizer}
}

type Request[REQ ChannelContextAware] struct {
	Token AuthToken
	FunctionRequest[REQ]
}

func (i *ContextFnInvokerImpl[REQ, RES]) Invoke(
	ctx context.Context,
	req Request[REQ],
) (RES, error) {
	var empty RES

	auth := AuthInfo{Token: req.Token, ChannelID: req.Body.ChannelContext().Channel.ID}
	if err := i.authorizer.Authorize(ctx, req.Body.ChannelContext(), auth); err != nil {
		return empty, err
	}

	appReq := ChannelFunctionRequest[REQ]{
		ChannelID:       req.Body.ChannelContext().Channel.ID,
		FunctionRequest: req.FunctionRequest,
	}

	return i.invoker.InvokeInChannel(ctx, appReq)
}
