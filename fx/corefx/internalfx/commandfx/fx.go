package commandfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
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
		fx.Annotate(
			svc.NewCommandClearHook,
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
			fx.As(new(app.InstallHandler)),
			fx.ResultTags(appfx.PreInstallHandlerGroup),
		),
		svc.NewActivationSvc,
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
		fx.Annotate(
			repo.NewActivationSettingRepository,
			fx.As(new(svc.ActivationSettingRepository)),
		),
	),
)
