package commandfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/command/repo"
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

const (
	CommandListenersGroup = `group:"commandListeners"`
)

var Command = fx.Options(
	CommandDAOs,
	CommandSvcs,
)
var CommandSvcs = fx.Options(
	fx.Provide(
		svc.NewRegisterService,
		svc.NewAutoCompleteInvoker,
		app.NewTypedInvoker[svc.CommandBody, svc.Action],
		app.NewTypedInvoker[svc.AutoCompleteBody, model.Choices],
	),

	fx.Provide(
		fx.Annotate(
			svc.NewInvoker,
			fx.ParamTags(``, ``, CommandListenersGroup),
		),
	),
)

var CommandDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewCommandDao,
			fx.As(new(svc.CommandRepository)),
		),
	),
)
