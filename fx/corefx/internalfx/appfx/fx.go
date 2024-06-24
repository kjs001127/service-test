package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/app/repo"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/hook/svc"
)

const (
	FunctionListenersGroup  = `group:"functionListeners"`
	LifecycleHookGroup      = `group:"lifecycle"`
	PreInstallHandlerGroup  = `group:"preInstallHandlers"`
	PostInstallHandlerGroup = `group:"postInstallHandlers"`
)

var App = fx.Options(
	AppSvcs,
	AppDAOs,
)

var AppSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			app.NewAppInstallSvc,
			fx.ParamTags(``, ``, PreInstallHandlerGroup, PostInstallHandlerGroup),
			fx.As(new(app.AppInstallSvc)),
		),

		app.NewInstallQuerySvc,
		fx.Annotate(
			app.NewAppQuerySvcImpl,
			fx.As(new(app.AppQuerySvc)),
		),
		fx.Annotate(
			app.NewAppLifecycleSvc,
			fx.As(new(app.AppLifecycleSvc)),
			fx.ParamTags(``, ``, LifecycleHookGroup),
		),
		fx.Annotate(
			app.NewInvoker,
			fx.As(new(app.Invoker)),
			fx.ParamTags(``, ``, ``, FunctionListenersGroup),
		),
		fx.Annotate(
			app.NewManagerAwareInstallSvc,
			fx.ParamTags(``, PreInstallHandlerGroup, PostInstallHandlerGroup),
		),

		app.NewTypedInvoker[json.RawMessage, json.RawMessage],
		app.NewTypedInvoker[svc.ToggleHookRequest, svc.ToggleHookResponse],
	),
)

var AppDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppInstallationDao,
			fx.As(new(app.AppInstallationRepository)),
		),
		fx.Annotate(
			repo.NewAppDAO,
			fx.As(new(app.AppRepository)),
		),
	),
)
