package publicfx

import (
	"fmt"

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
		AvailableClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{}
		},
		DefaultClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{
				{
					Service: config.Get().ServiceName,
					Action:  publiccmd.RegisterCommands,
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
					Action:  publiccore.SearchGroups,
					Scope:   []string{"channel-{id}"},
				},
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
			}
		},
		DefaultClaimsOf: func(appId string) rolesvc.AvailableClaims {
			return rolesvc.AvailableClaims{	}
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
}
