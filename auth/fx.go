package auth

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/auth/general"
	"github.com/channel-io/ch-app-store/auth/principal/account"
	"github.com/channel-io/ch-app-store/auth/principal/session"
	"github.com/channel-io/ch-app-store/config"
)

var Option = fx.Module(
	"authModule",
	fx.Provide(
		fx.Annotate(
			func(cli *resty.Client, cfg config.Config) *account.ManagerFetcherImpl {
				return account.NewManagerFetcherImpl(cli, cfg.Auth.ManagerFetchURL)
			},
			fx.As(new(account.ManagerFetcher)),
		),

		fx.Annotate(
			func(cli *resty.Client, cfg config.Config) *session.UserFetcherImpl {
				return session.NewUserFetcherImpl(session.JwtServiceKey(cfg.Auth.JWTServiceKey))
			},
			fx.As(new(session.UserFetcher)),
		),
		func(cli *resty.Client, parser *general.Parser, cfg config.Config) *general.PrincipalToRBACExchanger {
			return general.NewPrincipalToRBACExchanger(cli, parser, cfg.Auth.TokenIssueURL)
		},
		general.NewParser,
	),
)
