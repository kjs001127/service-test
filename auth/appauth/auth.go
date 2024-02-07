package appauth

import (
	"context"

	"github.com/channel-io/ch-app-store/auth"
)

type Authorities map[string][]string

const Wildcard = "*"

type ChannelContext struct {
	Channel struct {
		ID string `json:"id"`
	}
	Caller struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
	Chat struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}
}

type AppUseRequest[T auth.Token] struct {
	AppID        string
	FunctionName string
	ChCtx        ChannelContext
	Token        T
}

type AppAuthorizer[T auth.Token] interface {
	Handle(ctx context.Context, req AppUseRequest[T]) (Authorities, error)
}

func (c ChannelContext) RequiredAuthorities() Authorities {
	return map[string][]string{
		"channel":     {c.Channel.ID},
		c.Caller.Type: {c.Caller.ID},
		c.Chat.Type:   {c.Chat.ID},
	}
}

type MockAuthorizer[T auth.Token] struct {
}

func NewMockAuthorizer[T auth.Token]() *MockAuthorizer[T] {
	return &MockAuthorizer[T]{}
}

func (m MockAuthorizer[T]) Handle(ctx context.Context, req AppUseRequest[T]) (Authorities, error) {
	return Authorities{Wildcard: {Wildcard}}, nil
}
