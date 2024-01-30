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
	app.NewInvoker[json.RawMessage, json.RawMessage],
	app.NewInvoker[app.ChannelContextAware, any],
	func() app.ContextAuthorizer { return nil }, // ContextAuthorizer 등록 필요
	app.NewInvoker[command.Arguments, command.Action],
	fx.Annotate(
		app.NewContextFnInvoker[command.Arguments, command.Action],
		fx.As(new(app.ContextFnInvoker[command.Arguments, command.Action])),
	),
	app.NewInvoker[brief.BriefRequest, brief.BriefResponse],
	fx.Annotate(
		app.NewContextFnInvoker[brief.BriefRequest, brief.BriefResponse],
		fx.As(new(app.ContextFnInvoker[brief.BriefRequest, brief.BriefResponse])),
	),
	app.NewInvoker[command.AutoCompleteRequest, command.Choices],
	fx.Annotate(
		app.NewContextFnInvoker[command.AutoCompleteRequest, command.Choices],
		fx.As(new(app.ContextFnInvoker[command.AutoCompleteRequest, command.Choices])),
	),
	app.NewAppInstallSvc,
	app.NewQuerySvc,
	app.NewConfigSvc,
)
