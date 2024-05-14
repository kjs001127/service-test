package approlefx

import (
	"fmt"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	devrepo "github.com/channel-io/ch-app-store/internal/approle/repo"
	devsvc "github.com/channel-io/ch-app-store/internal/approle/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
)

var AppRole = fx.Options(
	RemoteAppDevSvcs,
	AppRoleDaos,
)

var RemoteAppDevSvcs = fx.Options(
	fx.Supply(
		map[model.RoleType]devsvc.TypeRule{
			model.RoleTypeApp: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				DefaultClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: appId,
							Action:  "*",
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
						{
							Service: config.Get().ServiceName,
							Action:  "registerCommands",
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
					}
				},
			},
			model.RoleTypeChannel: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				DefaultClaimsOf: func(s string) []*protomodel.Claim {
					return nil
				},
			},

			model.RoleTypeUser: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_PRINCIPAL,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				PrincipalTypes: []string{session.XSessionHeader},
				DefaultClaimsOf: func(s string) []*protomodel.Claim {
					return nil
				},
			},

			model.RoleTypeManager: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_PRINCIPAL,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				PrincipalTypes: []string{account.XAccountHeader},
				DefaultClaimsOf: func(appID string) []*protomodel.Claim {
					return nil
				},
			},
		},
	),
	fx.Provide(
		devsvc.NewAppRoleSvc,
		fx.Annotate(
			devsvc.NewTokenSvc,
			fx.As(new(devsvc.TokenSvc)),
		),
		fx.Annotate(
			devsvc.NewRoleClearHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)

var AppRoleDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			devrepo.NewAppRoleDao,
			fx.As(new(devsvc.AppRoleRepository)),
		),
		fx.Annotate(
			devrepo.NewAppSecretDao,
			fx.As(new(devsvc.AppSecretRepository)),
		),
	),
)
