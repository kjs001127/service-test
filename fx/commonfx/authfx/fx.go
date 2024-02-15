package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	"github.com/channel-io/ch-app-store/config"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

const (
	jwtServiceKey = `name:"jwtServiceKey"`
	authGeneral   = `name:"authGeneral"`
	authAdmin     = `name:"authAdmin"`
)

var Auth = fx.Module(
	"authModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Auth.JWTServiceKey,
			fx.ResultTags(jwtServiceKey),
		),
		fx.Annotate(
			config.Get().Auth.AuthGeneralURL,
			fx.ResultTags(authGeneral),
		),
		fx.Annotate(
			config.Get().Auth.AuthAdminURL,
			fx.ResultTags(authAdmin),
		),
	),
	fx.Provide(
		fx.Annotate(
			general.NewRBACExchanger,
			fx.ParamTags(`name:"dw"`, "", authGeneral),
		),
		fx.Annotate(
			account.NewManagerFetcherImpl,
			fx.As(new(account.ManagerFetcher)),
			fx.ParamTags(`name:"dw"`, authAdmin),
		),
		fx.Annotate(
			session.NewUserFetcherImpl,
			fx.As(new(session.UserFetcher)),
			fx.ParamTags(jwtServiceKey),
		),
		fx.Annotate(
			general.NewRoleClient,
			fx.ParamTags(`name:"dw"`, authAdmin),
		),
		fx.Annotate(
			general.NewParser,
			fx.As(new(general.Parser)),
			fx.ParamTags(`name:"dw"`, ``, jwtServiceKey),
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
			fx.ParamTags(`name:"dw"`, authAdmin),
		),
		fx.Annotate(
			general.NewRoleClientAdapter,
			fx.As(new(remoteapp.RoleClient)),
		),
	),
)
