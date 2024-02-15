package remoteappfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/repo"
)

var RemoteApp = fx.Module(
	"remoteApp",
	RemoteAppDomain,
	RemoteAppInfra,
)

var RemoteAppDev = fx.Module(
	"remoteAppDev",
	RemoteAppDevDomain,
	RemoteAppInfra,
)

var RemoteAppDomain = fx.Module(
	"remoteappDomain",
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
			domain.NewFileStreamHandler,
			fx.As(new(app.FileStreamHandler)),
		),
	),
)

var RemoteAppDevDomain = fx.Module(
	"remoteAppDev",
	fx.Provide(
		fx.Annotate(
			domain.NewAppDevSvcImpl,
			fx.As(new(domain.AppDevSvc)),
		),
	),
	RemoteAppDomain,
)

var RemoteAppInfra = fx.Module(
	"remoteAppInfra",
	fx.Provide(
		fx.Annotate(
			repo.NewAppUrlDao,
			fx.As(new(domain.AppUrlRepository)),
		),

		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(domain.HttpRequester)),
			fx.ParamTags(`name:"app"`),
		),

		fx.Annotate(
			repo.NewAppRoleDao,
			fx.As(new(domain.AppRoleRepository)),
		),
	),
)
