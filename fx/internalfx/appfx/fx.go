package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appRepo "github.com/channel-io/ch-app-store/internal/app/repo"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var Option = fx.Module(
	"app",
	fx.Provide(
		fx.Annotate(
			appRepo.NewAppChannelDao,
			fx.As(new(app.AppChannelRepository)),
		),
		fx.Annotate(
			appRepo.NewAppDAO,
			fx.As(new(app.AppRepository)),
		),
		fx.Annotate(
			app.NewAppManagerImpl,
			fx.As(new(app.AppManager)),
		),
		app.NewInvoker[json.RawMessage, json.RawMessage],
		app.NewInvoker[command.ParamInput, command.Action],
		app.NewInvoker[command.AutoCompleteArgs, command.Choices],
		app.NewInvoker[brief.BriefRequest, brief.BriefResponse],
		app.NewAppInstallSvc,
		app.NewQuerySvc,
		app.NewConfigSvc,
		app.NewFileStreamer,
	),
)
