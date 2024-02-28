package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/app/repo"
)

var App = fx.Module(
	"app",
	AppSvcs,
	AppDAOs,
	AppListeners,
)

var AppSvcs = fx.Module(
	"appDomain",

	fx.Provide(
		app.NewAppInstallSvc,
		app.NewQuerySvc,
		app.NewConfigSvc,
		fx.Annotate(
			app.NewInvoker,
			fx.ParamTags(``, ``, `group:"invokeHandler"`, `group:"functionListeners"`),
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

var AppListeners = fx.Module(
	"appListeners",
	fx.Provide(
		fx.Annotate(
			app.NewFunctionDBLogger,
			fx.As(new(app.FunctionRequestListener)),
			fx.ResultTags(`group:"functionListeners"`),
		),
	),
)
