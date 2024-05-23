package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/httpfx"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

var RoleClientOnly = fx.Options(
	fx.Provide(
		fx.Annotate(
			general.NewRoleClientImpl,
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
			fx.As(new(general.RoleFetcher)),
		),
	),
)

var PrincipalAuth = fx.Options(
	fx.Provide(
		fx.Annotate(
			account.NewParserImpl,
			fx.As(new(account.Parser)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			account.NewManagerFetcherImpl,
			fx.As(new(account.ManagerFetcher)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			session.NewUserFetcherImpl,
			fx.As(new(session.UserFetcher)),
			fx.ParamTags(configfx.JwtServiceKey),
		),
		fx.Annotate(
			principal.NewChatValidator,
			fx.As(new(principal.ChatValidator)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			account.NewManagerRoleFetcher,
			fx.As(new(account.ManagerRoleFetcher)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
	),
)

var GeneralAuth = fx.Options(
	fx.Provide(
		fx.Annotate(
			general.NewRoleClientImpl,
			fx.As(new(general.RoleFetcher)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			general.NewParser,
			fx.As(new(general.Parser)),
			fx.ParamTags(``, configfx.JwtServiceKey),
		),
		fx.Annotate(
			general.NewRBACExchanger,
			fx.ParamTags(httpfx.DW, ``, configfx.DWAdmin),
		),
	),
)
