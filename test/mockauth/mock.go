package mockauth

import (
	"context"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

type Parser struct {
}

func (p *Parser) Parse(ctx context.Context, token string) (general.ParsedRBACToken, error) {
	return general.ParsedRBACToken{
		Actions: map[general.Service][]general.Action{
			"*": {"*"},
		},
		Scopes: map[string][]string{
			"*": {"*"},
		},
	}, nil
}

type ManagerFetcher struct {
}

func (m *ManagerFetcher) FetchManager(ctx context.Context, channelID string, token string) (account.ManagerPrincipal, error) {
	return account.ManagerPrincipal{
		Token: "",
		Manager: account.Manager{
			ID:        "1",
			ChannelID: "1",
			AccountID: "1",
			Name:      "testManager",
			Email:     "1@channel.io",
		},
	}, nil
}

type UserFetcher struct {
}

func (u *UserFetcher) FetchUser(ctx context.Context, token string) (session.UserPrincipal, error) {
	return session.UserPrincipal{
		Token: "",
		User: session.User{
			ID:        "1",
			ChannelID: "1",
		},
	}, nil
}

type ManagerCtxAuthorizer struct {
}

func (m ManagerCtxAuthorizer) Authorize(ctx context.Context, channelContext app.ChannelContext, invoker account.ManagerPrincipal) error {
	return nil
}

type SessionCtxAuthorizer struct {
}

func (s SessionCtxAuthorizer) Authorize(ctx context.Context, channelContext app.ChannelContext, invoker session.UserPrincipal) error {
	return nil
}
