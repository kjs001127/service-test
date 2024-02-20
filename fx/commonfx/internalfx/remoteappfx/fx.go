package remoteappfx

import (
	"net/http"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/repo"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

var RemoteApp = fx.Module(
	"remoteApp",
	RemoteAppDomain,
	RemoteAppInfra,
)

var RemoteAppDev = fx.Module(
	"remoteAppDev",
	RemoteAppDevDomain,
	RemoteAppInfra,
)

var RemoteAppDomain = fx.Module(
	"remoteappDomain",
	fx.Supply(
		fx.Annotate(app.AppType("remote"),
			fx.ResultTags(`name:"remoteApp"`),
		),
	),
	fx.Provide(
		fx.Annotate(
			domain.NewInstallHandler,
			fx.As(new(app.InstallHandler)),
		),
		fx.Annotate(
			domain.NewConfigValidator,
			fx.As(new(app.ConfigValidator)),
		),
		fx.Annotate(
			domain.NewInvoker,
			fx.ResultTags(`name:"remoteInvoker"`),
			fx.As(new(app.InvokeHandler)),
		),
		domain.NewFileStreamer,
		fx.Annotate(
			app.NewTyped[app.InvokeHandler],
			fx.ParamTags(`name:"remoteApp"`, `name:"remoteInvoker"`),
			fx.ResultTags(`group:"invokeHandler"`),
		),
	),
)

var RemoteAppDevDomain = fx.Module(
	"remoteAppDev",
	fx.Supply(
		map[domain.RoleType]domain.TypeRule{
			"channel": {
				GrantTypes: []model.GrantType{
					model.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS,
					model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedScopes: []string{
					"channel",
				},
			},

			"user": {
				GrantTypes: []model.GrantType{
					model.GrantType_GRANT_TYPE_PRINCIPAL,
					model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedPrincipalTypes: []string{"x-session"},
				GrantedScopes:         []string{"channel", "user"},
			},

			"manager": {
				GrantTypes: []model.GrantType{
					model.GrantType_GRANT_TYPE_PRINCIPAL,
					model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedPrincipalTypes: []string{"x-account"},
				GrantedScopes:         []string{"channel", "manager"},
			},
		},
	),
	fx.Provide(
		fx.Annotate(
			app.NewAppManagerImpl,
			fx.As(new(app.AppManager)),
			fx.ParamTags(``, ``, `name:"remoteApp"`),
		),
		fx.Annotate(
			domain.NewAppDevSvcImpl,
			fx.As(new(domain.AppDevSvc)),
		),
	),
	RemoteAppDomain,
)

var RemoteAppInfra = fx.Module(
	"remoteAppInfra",
	fx.Supply(
		fx.Annotate(
			http.DefaultTransport,
			fx.As(new(http.RoundTripper)),
		),
	),
	fx.Provide(
		fx.Annotate(
			repo.NewAppUrlDao,
			fx.As(new(domain.AppUrlRepository)),
		),

		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(domain.HttpRequester)),
			fx.ParamTags(`name:"app"`),
		),

		fx.Annotate(
			repo.NewAppRoleDao,
			fx.As(new(domain.AppRoleRepository)),
		),
	),
)
