package remoteappfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/repo"
)

var Option = fx.Provide(
	fx.Annotate(
		infra.NewHttpRequester,
		fx.As(new(domain.HttpRequester)),
		fx.ParamTags(`name:"app"`),
	),
	fx.Annotate(
		repo.NewAppUrlDao,
		fx.As(new(domain.AppUrlRepository)),
	),
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
	fx.Annotate(
		domain.NewAppDevSvcImpl,
		fx.As(new(domain.AppDevSvc)),
	),
	fx.Annotate(
		repo.NewAppRoleDao,
		fx.As(new(domain.AppRoleRepository)),
	),
)
