package mockauthfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	"github.com/channel-io/ch-app-store/test/mockauth"
)

var AuthMocked = fx.Options(
	fx.Supply(
		fx.Annotate(
			new(mockauth.UserFetcher),
			fx.As(new(session.UserFetcher)),
		),
		fx.Annotate(
			new(mockauth.ManagerFetcher),
			fx.As(new(account.ManagerFetcher)),
		),
	),
)

var GeneralAuthMocked = fx.Options(
	fx.Supply(
		fx.Annotate(
			new(mockauth.Parser),
			fx.As(new(general.Parser)),
		),
	),
)
