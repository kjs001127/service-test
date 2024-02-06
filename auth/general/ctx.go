package general

import (
	"context"
	"errors"

	"github.com/channel-io/ch-app-store/auth"
	"github.com/channel-io/ch-app-store/auth/chctx"
)

type Request[T auth.Token] struct {
	AppID string
	ChCtx chctx.ChannelContext
	Token T
}

type AppAuthorizer[T auth.Token] interface {
	Handle(ctx context.Context, req Request[T]) (Scopes, error)
}

type AuthorizerImpl struct {
	parser *Parser
}

func NewAuthorizerImpl(parser *Parser) *AuthorizerImpl {
	return &AuthorizerImpl{parser: parser}
}

func (g *AuthorizerImpl) Handle(ctx context.Context, req Request[Token]) (Scopes, error) {
	rbac, err := g.parser.Parse(ctx, req.Token.Value())
	if err != nil {
		return nil, err
	}

	if ok := rbac.CheckScopes(ExtractScopes(req.ChCtx)); !ok {
		return nil, errors.New("scope check fail")
	}

	return rbac.Scopes, nil
}

func ExtractScopes(c chctx.ChannelContext) Scopes {
	return map[ScopeKey][]ScopeValue{
		"channel":               {ScopeValue(c.Channel.ID)},
		ScopeKey(c.Caller.Type): {ScopeValue(c.Caller.ID)},
		ScopeKey(c.Chat.Type):   {ScopeValue(c.Chat.ID)},
	}
}
