package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/app/repo"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
)

const (
	FunctionListenersGroup = `group:"functionListeners"`
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
			repo.NewAppChannelDao,
			fx.As(new(app.AppChannelRepository)),
		),
		fx.Annotate(
			repo.NewAppDAO,
			fx.As(new(app.AppRepository)),
		),
	),
)
