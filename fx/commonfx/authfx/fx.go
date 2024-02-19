package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	"github.com/channel-io/ch-app-store/config"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

const (
	jwtServiceKey = `name:"jwtServiceKey"`
	authGeneral   = `name:"authGeneral"`
	authAdmin     = `name:"authAdmin"`
)

var AdminAuth = fx.Module(
	"authAdmin",
	fx.Supply(
		fx.Annotate(
			config.Get().Auth.AuthAdminURL,
			fx.ResultTags(authAdmin),
		),
		map[remoteapp.RoleType][]model.GrantType{
			remoteapp.RoleTypeChannel: {
				model.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS,
				model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
			},
			remoteapp.RoleTypeFront: {
				model.GrantType_GRANT_TYPE_PRINCIPAL,
				model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
			},
			remoteapp.RoleTypeDesk: {
				model.GrantType_GRANT_TYPE_PRINCIPAL,
				model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
			},
		},
		map[remoteapp.RoleType][]string{
			remoteapp.RoleTypeFront: {"x-session"},
			remoteapp.RoleTypeDesk:  {"x-account"},
		},
	),
	fx.Provide(
		fx.Annotate(
			general.NewRoleClient,
			fx.ParamTags(`name:"dw"`, authAdmin),
		),
		fx.Annotate(
			general.NewRoleClientAdapter,
			fx.As(new(remoteapp.RoleClient)),
		),
	),
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
	),
)
