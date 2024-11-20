package privatefx

import (
	"fmt"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	privatecmd "github.com/channel-io/ch-app-store/internal/native/localapi/command/action/private"
	publiccmd "github.com/channel-io/ch-app-store/internal/native/localapi/command/action/public"
	"github.com/channel-io/ch-app-store/internal/native/localapi/hook/action/private"
	privateinstall "github.com/channel-io/ch-app-store/internal/native/localapi/install/action/private"
	privatesyslog "github.com/channel-io/ch-app-store/internal/native/localapi/systemlog/action/private"
	privatewidget "github.com/channel-io/ch-app-store/internal/native/localapi/widget/action/private"
	privatecore "github.com/channel-io/ch-app-store/internal/native/proxyapi/action/private"
	publiccore "github.com/channel-io/ch-app-store/internal/native/proxyapi/action/public"
	"github.com/channel-io/ch-app-store/internal/role/model"
	devrepo "github.com/channel-io/ch-app-store/internal/role/repo"
	rolesvc "github.com/channel-io/ch-app-store/internal/role/svc"
	"github.com/channel-io/ch-app-store/internal/util"
)

var AppRole = fx.Options(
	RemoteAppDevSvcs,
	AppRoleDAOs,
)

var RemoteAppDevSvcs = fx.Options(
	fx.Provide(
		rolesvc.NewAppRoleSvc,
		rolesvc.NewAppSecretSvc,
		rolesvc.NewChannelAgreementSvc,
		fx.Annotate(
			rolesvc.NewLifecycleListener,
			fx.As(new(appsvc.AppLifeCycleEventListener)),
			fx.ResultTags(appfx.LifecycleListener),
		),
		fx.Annotate(
			rolesvc.NewTokenSvc,
			fx.As(new(rolesvc.TokenSvc)),
		),
		fx.Annotate(
			rolesvc.NewInstallHandler,
			fx.As(new(appsvc.InstallEventListener)),
			fx.ResultTags(appfx.InTrxEventListener),
		),
	),
	fx.Supply(fx.Annotate(PrivateNativeClaims, fx.As(new(rolesvc.ClaimManager)))),
)

var AppRoleDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			devrepo.NewAppRoleDao,
			fx.As(new(rolesvc.AppRoleRepository)),
		),
		fx.Annotate(
			devrepo.NewAppSecretDao,
			fx.As(new(rolesvc.AppSecretRepository)),
		),
		fx.Annotate(
			devrepo.NewChannelRoleAgreementDAO,
			fx.As(new(rolesvc.ChannelRoleAgreementRepo)),
		),
	),
)

var PrivateNativeClaims rolesvc.ClaimManager = rolesvc.StaticClaimManager{
	model.RoleTypeApp: {
		AvailableClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{
				{
					Service: config.Get().ServiceName,
					Action:  private.RegisterInstallHook,
					Scope:   []string{fmt.Sprintf("app-%s", appId)},
				},
				{
					Service: config.Get().ServiceName,
					Action:  private.RegisterToggleHook,
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
		DefaultClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{
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
		AvailableClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{
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
					Action:  privatesyslog.WriteSystemLog,
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
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.DeleteUser,
					Scope:   []string{"channel-{id}"},
				},
			}
		},
		DefaultClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{}
		},
	},

	model.RoleTypeUser: {
		DefaultClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{}
		},
		AvailableClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteUserChatMessageAsUser,
					Scope:   []string{"channel-{id}", "user-{id}"},
				},
			}
		},
	},

	model.RoleTypeManager: {
		DefaultClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{}
		},
		AvailableClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteGroupMessageAsManager,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteUserChatMessageAsManager,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteDirectChatMessageAsManager,
					Scope:   []string{"channel-{id}", "manager-{id}"},
				},
			}
		},
	},
}
