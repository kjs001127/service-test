package mockauth

import (
	"context"

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
		Caller: general.Caller{
			Type: "user",
			ID:   "userID",
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
