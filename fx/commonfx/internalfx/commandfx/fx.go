package commandfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/command/repo"
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
		domain.NewInvoker,
		domain.NewAutoCompleteInvoker,
		app.NewTypedInvoker[domain.ParamInput, domain.Action],
		app.NewTypedInvoker[domain.AutoCompleteArgs, domain.Choices],
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
