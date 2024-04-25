package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/app/repo"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
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
		),
		app.NewQuerySvc,
		fx.Annotate(
			app.NewAppCrudSvcImpl,
			fx.As(new(app.AppCrudSvc)),
			fx.ParamTags(``, ``, LifecycleHookGroup),
		),
		fx.Annotate(
			app.NewInvoker,
			fx.ParamTags(``, ``, ``, FunctionListenersGroup),
		),
		app.NewTypedInvoker[json.RawMessage, json.RawMessage],
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
