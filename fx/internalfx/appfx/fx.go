package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appRepo "github.com/channel-io/ch-app-store/internal/app/repo"
)

var Option = fx.Provide(
	app.NewFileStreamer,
	fx.Annotate(
		appRepo.NewAppChannelDao,
		fx.As(new(app.AppChannelRepository)),
	),
	fx.Annotate(
		appRepo.NewAppDAO,
		fx.As(new(app.AppRepository)),
	),
	app.NewInvoker[json.RawMessage],
	app.NewAppInstallSvc,
	app.NewQuerySvc,
	app.NewConfigSvc,
	app.NewAppManagerImpl,
)
