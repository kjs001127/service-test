package mockauth

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/shared/general"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/internal/shared/principal/front"
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

func (m *ManagerFetcher) FetchManager(ctx context.Context, channelID string, token string) (desk.ManagerPrincipal, error) {
	return desk.ManagerPrincipal{
		Token: "",
		Manager: desk.Manager{
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

func (u *UserFetcher) FetchUser(ctx context.Context, token string) (front.UserPrincipal, error) {
	return front.UserPrincipal{
		Token: "",
		User: front.User{
			ID:        "1",
			ChannelID: "1",
		},
	}, nil
}
