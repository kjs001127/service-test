package managerfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/manager/svc"
)

const (
	PreInstallHandlerGroup  = `group:"managerPreInstall"`
	PostInstallHandlerGroup = `group:"managerPostInstall"`

	PreToggleHandlerGroup  = `group:"managerPreToggle"`
	PostToggleHandlerGroup = `group:"managerPostToggle"`
)

var Manager = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewManagerAwareInstallSvc,
			fx.ParamTags(``, PreInstallHandlerGroup, PostInstallHandlerGroup),
		),

		fx.Annotate(
			svc.NewManagerAwareToggleSvc,
			fx.ParamTags(``, PreToggleHandlerGroup, PostToggleHandlerGroup),
		),
	),
)
