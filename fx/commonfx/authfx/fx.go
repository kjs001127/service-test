package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	"github.com/channel-io/ch-app-store/fx/configfx"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var AdminAuth = fx.Module(
	"adminAuth",
	fx.Provide(
		fx.Annotate(
			general.NewRoleClient,
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
			general.NewRoleClient,
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
		fx.Annotate(
			general.NewParser,
			fx.As(new(general.Parser)),
			fx.ParamTags(restyfx.Dw, ``, configfx.DwAdmin),
		),
		fx.Annotate(
			session.NewContextAuthorizer,
			fx.As(new(session.ContextAuthorizer)),
		),
		fx.Annotate(
			account.NewContextAuthorizer,
			fx.As(new(account.ContextAuthorizer)),
		),
		fx.Annotate(
			principal.NewChatValidator,
			fx.As(new(principal.ChatValidator)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
	),
)
