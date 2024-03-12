package interactionfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/remoteapp/interaction/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/interaction/repo"
	"github.com/channel-io/ch-app-store/internal/remoteapp/interaction/svc"
)

var RemoteAppInteraction = fx.Options(
	RemoteAppInteractionSvcs,
	RemoteAppHttps,
	RemoteAppDAOs,
)

var RemoteAppInteractionSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewInstallHandler,
			fx.As(new(app.InstallHandler)),
		),
		fx.Annotate(
			svc.NewConfigValidator,
			fx.As(new(app.ConfigValidator)),
		),
		fx.Annotate(
			svc.NewInvoker,
			fx.As(new(app.InvokeHandler)),
		),
		fx.Annotate(
			svc.NewFileStreamer,
			fx.ParamTags(``, restyfx.App, ``),
		),
	),
)

var RemoteAppHttps = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(svc.HttpRequester)),
			fx.ParamTags(restyfx.App),
		),
	),
)

var RemoteAppDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppUrlDao,
			fx.As(new(svc.AppUrlRepository)),
		),
	),
)
