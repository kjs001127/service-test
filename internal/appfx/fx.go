package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/app/repo"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/hook/svc"
)

const (
	FunctionListener     = `group:"functionListeners"`
	LifecycleListener    = `group:"lifecycle"`
	InTrxEventListener   = `group:"preInstallHandlers"`
	PostTrxEventListener = `group:"postInstallHandlers"`
)

var App = fx.Options(
	AppSvcs,
	AppDAOs,
)

var AppSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			app.NewAppInstallSvc,
			fx.ParamTags(``, ``, ``, InTrxEventListener, PostTrxEventListener),
			fx.As(new(app.AppInstallSvc)),
		),

		app.NewInstallQuerySvc,
		fx.Annotate(
			app.NewAppQuerySvcImpl,
			fx.As(new(app.AppQuerySvc)),
		),
		fx.Annotate(
			app.NewAppLifecycleSvcImpl,
			fx.As(new(app.AppLifecycleSvc)),
			fx.ParamTags(``, ``, ``, ``, LifecycleListener),
		),
		fx.Annotate(
			app.NewInvoker,
			fx.As(new(app.Invoker)),
			fx.ParamTags(``, ``, ``, FunctionListener),
		),
		fx.Annotate(
			app.NewManagerAwareInstallSvc,
			fx.ParamTags(``, InTrxEventListener, PostTrxEventListener),
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
		fx.Annotate(
			repo.NewAppDisplayDAO,
			fx.As(new(app.AppDisplayRepository)),
		),
	),
)
