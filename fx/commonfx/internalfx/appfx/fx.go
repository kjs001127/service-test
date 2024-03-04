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
)

var App = fx.Module(
	"app",
	AppSvcs,
	AppDAOs,
)

var AppSvcs = fx.Module(
	"appDomain",

	fx.Provide(
		app.NewAppInstallSvc,
		app.NewQuerySvc,
		app.NewConfigSvc,
		fx.Annotate(
			app.NewInvoker,
			fx.ParamTags(``, ``, InvokeHandlerGroup, FunctionListenersGroup),
		),
		app.NewTypedInvoker[json.RawMessage, json.RawMessage],
	),
)

var AppDAOs = fx.Module(
	"appDB",
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
