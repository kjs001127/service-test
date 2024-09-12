package privatefx

import (
	"fmt"

	"github.com/channel-io/ch-app-store/config"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	devrepo "github.com/channel-io/ch-app-store/internal/approle/repo"
	devsvc "github.com/channel-io/ch-app-store/internal/approle/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
	privatecmd "github.com/channel-io/ch-app-store/internal/native/localapi/command/action/private"
	publiccmd "github.com/channel-io/ch-app-store/internal/native/localapi/command/action/public"
	privatehook "github.com/channel-io/ch-app-store/internal/native/localapi/hook/action/private"
	privateinstall "github.com/channel-io/ch-app-store/internal/native/localapi/install/action/private"
	privatesystemlog "github.com/channel-io/ch-app-store/internal/native/localapi/systemlog/action/private"
	privatewidget "github.com/channel-io/ch-app-store/internal/native/localapi/widget/action/private"
	privatecore "github.com/channel-io/ch-app-store/internal/native/proxyapi/action/private"
	publiccore "github.com/channel-io/ch-app-store/internal/native/proxyapi/action/public"
	"github.com/channel-io/ch-app-store/internal/util"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"

	"go.uber.org/fx"
)

const serviceNameGroup = `name:"serviceNames"`

var AppRole = fx.Options(
	RemoteAppDevSvcs,
	AppRoleDaos,
)

func services() []string {
	var ret []string
	for _, service := range config.Get().Services {
		ret = append(ret, service.String())
	}
	return ret
}

var RemoteAppDevSvcs = fx.Options(
	fx.Supply(
		fx.Annotate(
			services(),
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
				AvailableClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: config.Get().ServiceName,
							Action:  privatehook.RegisterInstallHook,
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privatehook.RegisterToggleHook,
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privateinstall.CheckInstall,
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privatewidget.RegisterAppWidgets,
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
					}
				},
				DefaultClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: config.Get().ServiceName,
							Action:  publiccmd.RegisterCommands,
							Scope:   []string{fmt.Sprintf("app-%s", appId)},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privatewidget.RegisterAppWidgets,
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
				AvailableClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.WriteUserChatMessage,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.WriteGroupMessage,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.GetUser,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.GetUserChat,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.GetManager,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.SearchManagers,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.GetChannel,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.ManageUserChat,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.GetGroup,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.BatchGetManagers,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.CreateDirectChat,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.SearchUser,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.FindOrCreateContactAndUser,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.SearchUserChatsByContact,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.UpdateUserChatState,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privatecmd.GetCommandChannelActivations,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privatecmd.ToggleCommand,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privateinstall.CheckInstall,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ServiceName,
							Action:  privatesystemlog.WriteSystemLog,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.CreateUserChat,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.SearchGroups,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().Services[util.DOCUMENT_API].String(),
							Action:  publiccore.SearchArticles,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().Services[util.DOCUMENT_API].String(),
							Action:  publiccore.GetRevision,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().Services[util.DOCUMENT_API].String(),
							Action:  publiccore.GetArticle,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.GetPlugin,
							Scope:   []string{"channel-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  privatecore.PatchUser,
							Scope:   []string{"channel-{id}"},
						},
					}
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
				AvailableClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.WriteUserChatMessageAsUser,
							Scope:   []string{"channel-{id}", "user-{id}"},
						},
					}
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
				AvailableClaimsOf: func(appId string) []*protomodel.Claim {
					return []*protomodel.Claim{
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.WriteGroupMessageAsManager,
							Scope:   []string{"channel-{id}", "manager-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.WriteUserChatMessageAsManager,
							Scope:   []string{"channel-{id}", "manager-{id}"},
						},
						{
							Service: config.Get().ChannelServiceName,
							Action:  publiccore.WriteDirectChatMessageAsManager,
							Scope:   []string{"channel-{id}", "manager-{id}"},
						},
					}
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
