package commandfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/repo"
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

const (
	CommandListenersGroup  = `group:"commandListeners"`
	PreToggleHandlerGroup  = `group:"managerPreToggle"`
	PostToggleHandlerGroup = `group:"managerPostToggle"`
)

var Command = fx.Options(
	CommandDAOs,
	CommandSvcs,
)
var CommandSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAppLifecycleHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
		svc.NewRegisterSvc,
		svc.NewAutoCompleteInvoker,
		svc.NewWysiwygQuerySvc,
		app.NewTypedInvoker[svc.CommandBody, svc.Action],
		app.NewTypedInvoker[svc.AutoCompleteBody, svc.AutoCompleteResponse],
		fx.Annotate(
			svc.NewActivationSvc,
			fx.As(new(svc.ActivationSvc)),
		),
		fx.Annotate(
			svc.NewManagerAwareToggleSvc,
			fx.ParamTags(``, PreToggleHandlerGroup, PostToggleHandlerGroup),
		),
		svc.NewInstalledCommandQuerySvc,
	),

	fx.Provide(
		fx.Annotate(
			svc.NewInvoker,
			fx.ParamTags(``, ``, ``, CommandListenersGroup),
		),
	),
)

var CommandDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewCommandDao,
			fx.As(new(svc.CommandRepository)),
		),
		fx.Annotate(
			repo.NewActivationRepository,
			fx.As(new(svc.ActivationRepository)),
		),
	),
)
