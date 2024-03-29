package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/app/repo"
)

const (
	FunctionListenersGroup = `group:"functionListeners"`
	InvokeHandlerGroup     = `group:"invokeHandler"`
	RemoteAppName          = `name:"remoteApp"`
	LifecycleHookGroup     = `group:"lifecycle"`
)

var App = fx.Options(
	AppSvcs,
	AppDAOs,
)

var AppSvcs = fx.Options(
	fx.Provide(
		app.NewAppInstallSvc,
		app.NewQuerySvc,
		app.NewConfigSvc,
		fx.Annotate(
			app.NewAppManagerImpl,
			fx.As(new(app.AppManager)),
			fx.ParamTags(``, ``, RemoteAppName, LifecycleHookGroup),
		),
		fx.Annotate(
			app.NewInvoker,
			fx.ParamTags(``, ``, InvokeHandlerGroup, FunctionListenersGroup),
		),
		app.NewTypedInvoker[json.RawMessage, json.RawMessage],
	),
)

var AppDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppChannelDao,
			fx.As(new(app.AppChannelRepository)),
		),
		fx.Annotate(
			repo.NewAppDAO,
			fx.As(new(app.AppRepository)),
		),
	),
)
