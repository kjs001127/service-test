package approlefx

import (
	"fmt"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	devrepo "github.com/channel-io/ch-app-store/internal/approle/repo"
	devsvc "github.com/channel-io/ch-app-store/internal/approle/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	tpa "github.com/channel-io/ch-app-store/internal/native/coreapi/action/thirdparty"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"

	"go.uber.org/fx"
)

const serviceNameGroup = `name:"serviceNames"`

var AppRole = fx.Options(
	RemoteAppDevSvcs,
	AppRoleDaos,
)

var RemoteAppDevSvcs = fx.Options(
	fx.Supply(
		fx.Annotate(
			[]string{config.Get().ServiceName, config.Get().ChannelServiceName},
			fx.ResultTags(serviceNameGroup),
		),
	),
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
				AvailableClaims: []*protomodel.Claim{
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.WriteUserChatMessage,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.WriteGroupMessage,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.GetUser,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.GetUserChat,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.GetManager,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.SearchManagers,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.GetChannel,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.ManageUserChat,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.GetGroup,
						Scope:   []string{"channel-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.BatchGetManagers,
						Scope:   []string{"channel-{id}"},
					},
				},
				DefaultClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: appId,
							Action:  "*",
							Scope:   []string{"channel-{id}"},
						},
					}
				},
			},

			model.RoleTypeUser: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_PRINCIPAL,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				PrincipalTypes: []string{session.XSessionHeader},
				DefaultClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: appId,
							Action:  "*",
							Scope:   []string{"channel-{id}"},
						},
					}
				},
				AvailableClaims: []*protomodel.Claim{
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.WriteUserChatMessageAsUser,
						Scope:   []string{"channel-{id}"},
					},
				},
			},

			model.RoleTypeManager: {
				GrantTypes: []protomodel.GrantType{
					protomodel.GrantType_GRANT_TYPE_PRINCIPAL,
					protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN,
				},
				PrincipalTypes: []string{account.XAccountHeader},
				DefaultClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: appId,
							Action:  "*",
							Scope:   []string{fmt.Sprintf("channel-{id}")},
						},
					}
				},
				AvailableClaims: []*protomodel.Claim{
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.WriteGroupMessageAsManager,
						Scope:   []string{"channel-{id}", "manager-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.WriteUserChatMessageAsManager,
						Scope:   []string{"channel-{id}", "manager-{id}"},
					},
					{
						Service: config.Get().ChannelServiceName,
						Action:  tpa.WriteDirectChatMessageAsManager,
						Scope:   []string{"channel-{id}", "manager-{id}"},
					},
				},
			},
		},
	),
	fx.Provide(
		fx.Annotate(
			devsvc.NewAppRoleSvc,
			fx.ParamTags(``, ``, ``, ``, ``, serviceNameGroup),
		),
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
