package mockauth

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.uber.org/fx"

	mockgeneral "github.com/channel-io/ch-app-store/generated/mock/auth/general"
	mockprincipal "github.com/channel-io/ch-app-store/generated/mock/auth/principal"
	mockaccount "github.com/channel-io/ch-app-store/generated/mock/auth/principal/account"
	mocksession "github.com/channel-io/ch-app-store/generated/mock/auth/principal/session"
	mocksvc "github.com/channel-io/ch-app-store/generated/mock/remoteapp/development/svc"

	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	"github.com/channel-io/ch-app-store/internal/remoteapp/development/svc"
)

var mockedManager = account.Manager{
	ID:        "1",
	ChannelID: "1",
	AccountID: "1",
	Name:      "fakeManager",
	Email:     "fake@fake.io",
}

var Module = fx.Options(
	fx.Provide(
		func(t *testing.T) svc.RoleClient {
			return mocksvc.NewRoleClient(t)
		},
		func(t *testing.T) account.ManagerFetcher {
			mocked := mockaccount.NewManagerFetcher(t)
			mocked.On("FetchManager", mock.Anything, mock.Anything, mock.Anything).
				Return(account.ManagerPrincipal{Manager: mockedManager, Token: "token"}, nil).Maybe()
			return mocked
		},
		func(t *testing.T) session.UserFetcher {
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
