package commandfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/command/repo"
)

const (
	CommandListenersGroup = `group:"commandListeners"`
)

var Command = fx.Module(
	"command",
	CommandDAOs,
	CommandSvcs,
)
var CommandSvcs = fx.Module(
	"commandDomain",
	fx.Provide(
		domain.NewParamValidator,
		domain.NewRegisterService,
		domain.NewAutoCompleteInvoker,
		app.NewTypedInvoker[domain.CommandBody, domain.Action],
		app.NewTypedInvoker[domain.AutoCompleteBody, domain.Choices],
	),

	fx.Provide(
		fx.Annotate(
			domain.NewInvoker,
			fx.ParamTags(``, ``, ``, CommandListenersGroup),
		),
	),
)

var CommandDAOs = fx.Module(
	"commandDB",
	fx.Provide(
		fx.Annotate(
			repo.NewCommandDao,
			fx.As(new(domain.CommandRepository)),
		),
	),
)
