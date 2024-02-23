package remoteappfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/repo"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

const (
	appTypeRemote = app.AppType("remote")
	remoteAppName = `name:"remoteApp"`
)

var RemoteAppCommon = fx.Module(
	"remoteappCommon",
	RemoteAppCommonsSvcs,
	RemoteAppHttps,
	RemoteAppDAOs,
)

var RemoteAppDev = fx.Module(
	"remoteappDev",
	RemoteAppCommon,
	RemoteAppDevSvc,
)

var RemoteAppCommonsSvcs = fx.Module(
	"remoteappDomain",
	fx.Supply(
		fx.Annotate(
			appTypeRemote,
			fx.ResultTags(remoteAppName),
		),
		fx.Private,
	),
	fx.Provide(
		fx.Annotate(
			app.NewAppManagerImpl,
			fx.As(new(app.AppManager)),
			fx.ParamTags(``, ``, remoteAppName),
		),
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
		fx.Annotate(
			domain.NewFileStreamer,
			fx.ParamTags(``, restyfx.App, ``),
		),
		fx.Annotate(
			app.NewTyped[app.InvokeHandler],
			fx.ParamTags(remoteAppName, `name:"remoteInvoker"`),
			fx.ResultTags(`group:"invokeHandler"`),
		),
	),
)

const (
	roleTypeChannel = "channel"
	roleTypeUser    = "user"
	roleTypeManager = "manager"

	scopeChannel = "channel"
	scopeUser    = "user"
	scopeManager = "manager"
)

var RemoteAppDevSvc = fx.Module(
	"remoteAppDev",
	fx.Supply(
		map[domain.RoleType]domain.TypeRule{
			roleTypeChannel: {
				GrantTypes: []model.GrantType{
					model.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS,
					model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedScopes: []string{
					scopeChannel,
				},
			},

			roleTypeUser: {
				GrantTypes: []model.GrantType{
					model.GrantType_GRANT_TYPE_PRINCIPAL,
					model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedPrincipalTypes: []string{session.XSessionHeader},
				GrantedScopes:         []string{scopeChannel, scopeUser},
			},

			roleTypeManager: {
				GrantTypes: []model.GrantType{
					model.GrantType_GRANT_TYPE_PRINCIPAL,
					model.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedPrincipalTypes: []string{account.XAccountHeader},
				GrantedScopes:         []string{scopeChannel, scopeManager},
			},
		},
	),
	fx.Provide(
		fx.Annotate(
			domain.NewAppDevSvcImpl,
			fx.As(new(domain.AppDevSvc)),
		),
	),
)

var RemoteAppHttps = fx.Module(
	"remoteAppHttp",

	fx.Provide(
		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(domain.HttpRequester)),
			fx.ParamTags(restyfx.App),
		),
	),
)

var RemoteAppDAOs = fx.Module(
	"remoteAppInfra",
	fx.Provide(
		fx.Annotate(
			repo.NewAppUrlDao,
			fx.As(new(domain.AppUrlRepository)),
		),
		fx.Annotate(
			repo.NewAppRoleDao,
			fx.As(new(domain.AppRoleRepository)),
		),
	),
)
