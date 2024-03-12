package developmentfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/remoteappfx/interactionfx"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	"github.com/channel-io/ch-app-store/internal/remoteapp/development/model"
	devrepo "github.com/channel-io/ch-app-store/internal/remoteapp/development/repo"
	devsvc "github.com/channel-io/ch-app-store/internal/remoteapp/development/svc"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
)

const (
	roleTypeChannel = "channel"
	roleTypeUser    = "user"
	roleTypeManager = "manager"

	scopeChannel = "channel-{id}"
	scopeUser    = "user-{id}"
	scopeManager = "manager-{id}"
)

var RemoteAppDevelopment = fx.Options(
	interactionfx.RemoteAppInteraction,
	RemoteAppDevSvcs,
	RemoteAppDevDaos,
)

var RemoteAppDevSvcs = fx.Options(
	fx.Supply(
		map[model.RoleType]devsvc.TypeRule{
			roleTypeChannel: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedScopes: []string{
					scopeChannel,
				},
			},

			roleTypeUser: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_PRINCIPAL,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedPrincipalTypes: []string{session.XSessionHeader},
				GrantedScopes:         []string{scopeChannel, scopeUser},
			},

			roleTypeManager: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_PRINCIPAL,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				GrantedPrincipalTypes: []string{account.XAccountHeader},
				GrantedScopes:         []string{scopeChannel, scopeManager},
			},
		},
	),
	fx.Provide(
		fx.Annotate(
			devsvc.NewAppDevSvcImpl,
			fx.As(new(devsvc.AppDevSvc)),
		),
	),
)

var RemoteAppDevDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			devrepo.NewAppRoleDao,
			fx.As(new(devsvc.AppRoleRepository)),
		),
	),
)
