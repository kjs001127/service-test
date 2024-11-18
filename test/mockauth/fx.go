package mockauth

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.uber.org/fx"

	mockaccount "github.com/channel-io/ch-app-store/generated/mock/shared/principal/desk"

	mocksession "github.com/channel-io/ch-app-store/generated/mock/shared/principal/front"

	mockgeneral "github.com/channel-io/ch-app-store/generated/mock/shared/general"
	mockprincipal "github.com/channel-io/ch-app-store/generated/mock/shared/principal"

	"github.com/channel-io/ch-app-store/internal/shared/general"
	"github.com/channel-io/ch-app-store/internal/shared/principal"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/internal/shared/principal/front"
)

var mockedManager = desk.Manager{
	ID:        "1",
	ChannelID: "1",
	AccountID: "1",
	Name:      "fakeManager",
	Email:     "fake@fake.io",
}

var Module = fx.Options(
	fx.Provide(
		func(t *testing.T) desk.ManagerFetcher {
			mocked := mockaccount.NewManagerFetcher(t)
			mocked.On("FetchManager", mock.Anything, mock.Anything, mock.Anything).
				Return(desk.ManagerPrincipal{Manager: mockedManager, Token: "token"}, nil).Maybe()
			return mocked
		},
		func(t *testing.T) front.UserFetcher {
			return mocksession.NewUserFetcher(t)
		},
		func(t *testing.T) general.RoleFetcher {
			return mockgeneral.NewRoleFetcher(t)
		},
		func(t *testing.T) general.Parser {
			return mockgeneral.NewParser(t)
		},
		func(t *testing.T) principal.ChatValidator {
			return mockprincipal.NewChatValidator(t)
		},
	),
)
