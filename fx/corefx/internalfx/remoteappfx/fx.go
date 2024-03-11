package remoteappfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/repo"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

var RemoteAppCommon = fx.Options(
	RemoteAppCommonsSvcs,
	RemoteAppHttps,
	RemoteAppDAOs,
)

var RemoteAppDev = fx.Options(
	RemoteAppCommon,
	RemoteAppDevSvc,
)

var RemoteAppCommonsSvcs = fx.Options(
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
			fx.As(new(app.InvokeHandler)),
		),
		fx.Annotate(
			domain.NewFileStreamer,
			fx.ParamTags(``, restyfx.App, ``),
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

var RemoteAppDevSvc = fx.Options(
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

var RemoteAppHttps = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(domain.HttpRequester)),
			fx.ParamTags(restyfx.App),
		),
	),
)

var RemoteAppDAOs = fx.Options(
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
