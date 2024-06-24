package appdisplayfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appdisplay/repo"
	display "github.com/channel-io/ch-app-store/internal/appdisplay/svc"

	"go.uber.org/fx"
)

var AppDisplay = fx.Options(
	AppDisplaySvcs,
	AppDisplayDAOs,
)

var AppDisplaySvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			display.NewDisplayLifecycleSvcImpl,
			fx.As(new(display.DisplayLifecycleSvc)),
		),
		fx.Annotate(
			display.NewAppWithDisplayQuerySvcImpl,
			fx.As(new(display.AppWithDisplayQuerySvc)),
		),
		fx.Annotate(
			display.NewLifeCycleHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)

var AppDisplayDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppDisplayDAO,
			fx.As(new(display.AppDisplayRepository)),
		),
	),
)
