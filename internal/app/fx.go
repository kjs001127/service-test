package app

import (
	"encoding/json"

	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appRepo "github.com/channel-io/ch-app-store/internal/app/repo"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var Option = fx.Provide(
	app.NewFileStreamer,
	fx.Annotate(
		appRepo.NewAppChannelDao,
		fx.As(new(app.AppChannelRepository)),
	),

	app.NewInvoker[json.RawMessage],
	app.NewInvoker[command.Action],
	app.NewInvoker[brief.BriefResponse],
	app.NewInvoker[command.Choices],

	app.NewAppInstallSvc,
	app.NewQuerySvc,
	app.NewConfigSvc,
)
