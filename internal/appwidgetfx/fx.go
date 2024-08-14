package appwidgetfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/appwidget/infra"
	"github.com/channel-io/ch-app-store/internal/appwidget/repo"
	"github.com/channel-io/ch-app-store/internal/appwidget/svc"
)

var AppWidget = fx.Options(
	AppWidgetDaos,
	AppWidgetSvcs,
	AppWidgetInfra,
)

var AppWidgetSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAppLifeCycleHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
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
			fx.As(new(app.InstallHandler)),
			fx.ResultTags(appfx.PostInstallHandlerGroup),
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
