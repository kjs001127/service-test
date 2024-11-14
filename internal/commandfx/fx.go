package commandfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/command/repo"
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

const (
	CommandListenersGroup      = `group:"commandListeners"`
	InTrxToggleListenerGroup   = `group:"managerPreToggle"`
	PostTrxToggleListenerGroup = `group:"managerPostToggle"`
)

var Command = fx.Options(
	CommandDAOs,
	CommandSvcs,
)
var CommandSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewLifecycleListener,
			fx.As(new(app.AppLifeCycleEventListener)),
			fx.ResultTags(appfx.LifecycleListener),
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
			svc.NewPreInstallHandler,
			fx.As(new(app.InstallEventListener)),
			fx.ResultTags(appfx.InTrxEventListener),
		),
		fx.Annotate(
			svc.NewManagerAwareToggleSvc,
			fx.ParamTags(``, InTrxToggleListenerGroup, PostTrxToggleListenerGroup),
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
