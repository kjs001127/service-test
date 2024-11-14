package widgetfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/widget/infra"
	"github.com/channel-io/ch-app-store/internal/widget/repo"
	"github.com/channel-io/ch-app-store/internal/widget/svc"
)

var AppWidget = fx.Options(
	AppWidgetDaos,
	AppWidgetSvcs,
	AppWidgetInfra,
)

var AppWidgetSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewLifecycleListener,
			fx.As(new(app.AppLifeCycleEventListener)),
			fx.ResultTags(appfx.LifecycleListener),
		),
		fx.Annotate(
			svc.NewRegisterSvc,
			fx.As(new(svc.RegisterSvc)),
		),
		fx.Annotate(
			svc.NewAppWidgetInvokerImpl,
			fx.As(new(svc.AppWidgetInvoker)),
		),
		fx.Annotate(
			svc.NewAppWidgetFetcherImpl,
			fx.As(new(svc.AppWidgetFetcher)),
		),
		fx.Annotate(
			svc.NewAppInstallListener,
			fx.As(new(app.InstallEventListener)),
			fx.ResultTags(appfx.PostTrxEventListener),
		),
	),
)

var AppWidgetInfra = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewDropwizardEventPublisher,
			fx.As(new(svc.EventPublisher)),
			fx.ParamTags(configfx.DWAdmin, httpfx.DW),
		),
	),
)

var AppWidgetDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewApWidgetDao,
			fx.As(new(svc.AppWidgetRepository)),
		),
	),
)
