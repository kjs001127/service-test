package commandfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/command/repo"
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
		domain.NewParamValidator,
		domain.NewRegisterService,
		domain.NewAutoCompleteInvoker,
		app.NewTypedInvoker[domain.CommandBody, domain.Action],
		app.NewTypedInvoker[domain.AutoCompleteBody, domain.AutoCompleteResponse],
		fx.Annotate(
			domain.NewCommandClearHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),

	fx.Provide(
		fx.Annotate(
			domain.NewInvoker,
			fx.ParamTags(``, ``, ``, CommandListenersGroup),
		),
	),
)

var CommandDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewCommandDao,
			fx.As(new(domain.CommandRepository)),
		),
	),
)
