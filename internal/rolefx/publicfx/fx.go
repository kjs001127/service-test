package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	publiccmd "github.com/channel-io/ch-app-store/internal/native/localapi/command/action/public"
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
	fx.Supply(fx.Annotate(PublicNativeClaims, fx.As(new(rolesvc.ClaimManager)))),
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

var PublicNativeClaims rolesvc.ClaimManager = rolesvc.StaticClaimManager{
	model.RoleTypeApp: {
		AvailableClaimsOf: func(appId string) model.Claims {
			return model.Claims{}
		},
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: config.Get().ServiceName,
					Action:  publiccmd.RegisterCommands,
				},
			}
		},
	},
	model.RoleTypeChannel: {
		AvailableClaimsOf: func(appId string) model.Claims {
			return model.Claims{
				{
					Service: config.Get().ChannelServiceName,
					Action:  publiccore.SearchGroups,
				},
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
			}
		},
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{}
		},
	},

	model.RoleTypeUser: {
		DefaultClaimsOf: func(appId string) model.Claims {
			return model.Claims{}
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
			return model.Claims{}
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
