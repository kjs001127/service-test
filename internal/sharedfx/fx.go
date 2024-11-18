package sharedfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/shared/general"
	"github.com/channel-io/ch-app-store/internal/shared/principal"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/internal/shared/principal/front"
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
			desk.NewParserImpl,
			fx.As(new(desk.Parser)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			desk.NewManagerFetcherImpl,
			fx.As(new(desk.ManagerFetcher)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			front.NewUserFetcherImpl,
			fx.As(new(front.UserFetcher)),
			fx.ParamTags(configfx.JwtServiceKey),
		),
		fx.Annotate(
			principal.NewChatValidator,
			fx.As(new(principal.ChatValidator)),
			fx.ParamTags(httpfx.DW, configfx.DWAdmin),
		),
		fx.Annotate(
			desk.NewManagerRoleFetcher,
			fx.As(new(desk.ManagerRoleFetcher)),
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
