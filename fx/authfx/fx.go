package authfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/auth/appauth"
	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	"github.com/channel-io/ch-app-store/config"
)

const (
	jwtServiceKey   = `name:"jwtServiceKey"`
	managerFetchUrl = `name:"managerFetchUrl"`
	authGeneral     = `name:authGeneral`
	authAdmin       = `name:authAdmin`
)

var Option = fx.Module(
	"authModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Auth.ManagerFetchURL,
			fx.ResultTags(managerFetchUrl),
		),
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
			fx.ParamTags("", "", authGeneral),
		),
		fx.Annotate(
			account.NewManagerFetcherImpl,
			fx.As(new(account.ManagerFetcher)),
			fx.ParamTags("", managerFetchUrl),
		),
		fx.Annotate(
			session.NewUserFetcherImpl,
			fx.As(new(session.UserFetcher)),
			fx.ParamTags(jwtServiceKey),
		),
		fx.Annotate(
			general.NewParser,
			fx.ParamTags("", authAdmin, jwtServiceKey),
		),

		fx.Private,
	),
	fx.Provide(
		fx.Annotate(
			general.NewAuthorizer,
			fx.As(new(appauth.AppAuthorizer[general.Token])),
		),
		fx.Annotate(
			general.NewPrincipalAuthorizer,
			fx.As(new(appauth.AppAuthorizer[principal.Token])),
		),
	),
)

var MockOption = fx.Module("mockauth",
	fx.Provide(
		fx.Annotate(
			appauth.NewMockAuthorizer[principal.Token],
			fx.As(new(appauth.AppAuthorizer[principal.Token])),
		),
		fx.Annotate(
			appauth.NewMockAuthorizer[general.Token],
			fx.As(new(appauth.AppAuthorizer[general.Token])),
		),
		fx.Annotate(
			account.NewMockManagerFetcher,
			fx.As(new(account.ManagerFetcher)),
		),
		fx.Annotate(
			session.NewMockUserFetcher,
			fx.As(new(session.UserFetcher)),
		),
	),
)
