package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/configfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/auth/principal"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var AdminAuth = fx.Module(
	"adminAuth",
	fx.Provide(
		fx.Annotate(
			general.NewRoleClientImpl,
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
			fx.As(new(remoteapp.RoleClient)),
		),
	),
)

var Auth = fx.Module(
	"authModule",
	fx.Provide(
		fx.Annotate(
			general.NewRBACExchanger,
			fx.ParamTags(restyfx.Dw, "", configfx.DwGeneral),
		),
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
			general.NewRoleClientImpl,
			fx.As(new(general.RoleFetcher)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
		fx.Annotate(
			general.NewParser,
			fx.As(new(general.Parser)),
			fx.ParamTags(``, configfx.JwtServiceKey),
		),
		fx.Annotate(
			principal.NewCommandCtxAuthorizer,
			fx.As(new(principal.CommandCtxAuthorizer)),
		),
		fx.Annotate(
			principal.NewChatValidator,
			fx.As(new(principal.ChatValidator)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
	),
)
