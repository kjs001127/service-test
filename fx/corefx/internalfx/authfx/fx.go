package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	"github.com/channel-io/ch-app-store/internal/remoteapp/development/svc"
)

var RoleClientOnly = fx.Options(
	fx.Provide(
		fx.Annotate(
			general.NewRoleClientImpl,
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
			fx.As(new(svc.RoleClient)),
		),
	),
)

var PrincipalAuth = fx.Options(
	fx.Provide(
		fx.Annotate(
			account.NewManagerFetcherImpl,
			fx.As(new(account.ManagerFetcher)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
		fx.Annotate(
			session.NewUserFetcherImpl,
			fx.As(new(session.UserFetcher)),
			fx.ParamTags(configfx.JwtServiceKey),
		),
		fx.Annotate(
			principal.NewChatValidator,
			fx.As(new(principal.ChatValidator)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
	),
)

var GeneralAuth = fx.Options(
	fx.Provide(
		fx.Annotate(
			general.NewRoleClientImpl,
			fx.As(new(general.RoleFetcher)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
		fx.Annotate(
			general.NewParser,
			fx.As(new(general.Parser)),
			fx.ParamTags(``, configfx.JwtServiceKey),
		),
	),
)
