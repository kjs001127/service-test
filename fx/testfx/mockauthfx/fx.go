package mockauthfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	"github.com/channel-io/ch-app-store/test/mockauth"
)

var AuthMocked = fx.Module("mockauth",
	fx.Supply(
		fx.Annotate(
			new(mockauth.Parser),
			fx.As(new(general.Parser)),
		),
		fx.Annotate(
			new(mockauth.UserFetcher),
			fx.As(new(session.UserFetcher)),
		),
		fx.Annotate(
			new(mockauth.ManagerFetcher),
			fx.As(new(account.ManagerFetcher)),
		),
		fx.Annotate(
			new(mockauth.SessionCtxAuthorizer),
			fx.As(new(session.ContextAuthorizer)),
		),
		fx.Annotate(
			new(mockauth.ManagerCtxAuthorizer),
			fx.As(new(account.ContextAuthorizer)),
		),
	),
)
