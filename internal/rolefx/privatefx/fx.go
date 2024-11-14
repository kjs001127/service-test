package privatefx

import (
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
		AvailableClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: config.Get().ServiceName,
					Action:  private.RegisterInstallHook,
				},
				{
					Service: config.Get().ServiceName,
					Action:  private.RegisterToggleHook,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privateinstall.CheckInstall,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privatewidget.RegisterAppWidgets,
				},
			}
		},
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: config.Get().ServiceName,
					Action:  publiccmd.RegisterCommands,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privatewidget.RegisterAppWidgets,
				},
			}
		},
	},
	model.RoleTypeChannel: {
		AvailableClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteUserChatMessage,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteGroupMessage,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.GetUser,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.GetUserChat,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.GetManager,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.SearchManagers,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.GetChannel,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.ManageUserChat,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.GetGroup,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.BatchGetManagers,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.CreateDirectChat,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.SearchUser,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.FindOrCreateContactAndUser,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.SearchUserChatsByContact,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.UpdateUserChatState,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privatecmd.GetCommandChannelActivations,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privatecmd.ToggleCommand,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privateinstall.CheckInstall,
				},
				{
					Service: config.Get().ServiceName,
					Action:  privatesyslog.WriteSystemLog,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.CreateUserChat,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.SearchGroups,
				},
				{
					Service: config.Get().Services[util.DOCUMENT_API].String(),
					Action:  publiccore.SearchArticles,
				},
				{
					Service: config.Get().Services[util.DOCUMENT_API].String(),
					Action:  publiccore.GetRevision,
				},
				{
					Service: config.Get().Services[util.DOCUMENT_API].String(),
					Action:  publiccore.GetArticle,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.GetPlugin,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.PatchUser,
				},
				{
					Service: config.Get().ChannelServiceName,
					Action:  privatecore.DeleteUser,
				},
			}
		},
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: appId,
					Action:  "*",
				},
			}
		},
	},

	model.RoleTypeUser: {
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: appId,
					Action:  "*",
				},
			}
		},
		AvailableClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.WriteUserChatMessageAsUser,
				},
			}
		},
	},

	model.RoleTypeManager: {
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: appId,
					Action:  "*",
				},
			}
		},
		AvailableClaimsOf: func(appId string) model.Claims {
			return model.Claims{
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
				},
			}
		},
	},
}
