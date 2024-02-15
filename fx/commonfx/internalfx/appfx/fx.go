package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/app/repo"
)

var App = fx.Module(
	"app",
	AppDomain,
	AppDB,
)

var AppDomain = fx.Module(
	"appDomain",
	fx.Provide(
		fx.Annotate(
			app.NewAppManagerImpl,
			fx.As(new(app.AppManager)),
		),
		app.NewAppInstallSvc,
		app.NewQuerySvc,
		app.NewConfigSvc,
		app.NewFileStreamer,
		app.NewInvoker[json.RawMessage, json.RawMessage],
	),
)

var AppDB = fx.Module(
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
