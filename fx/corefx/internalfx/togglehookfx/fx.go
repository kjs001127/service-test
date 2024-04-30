package togglehookfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/togglehook/repo"
	"github.com/channel-io/ch-app-store/internal/togglehook/svc"
)

var ToggleHook = fx.Options(
	ToggleHookSvcs,
	ToggleHookDaos,
)

var ToggleHookSvcs = fx.Options(
	fx.Provide(
		svc.NewHookSendingActivationSvc,
		app.NewTypedInvoker[svc.ToggleHookRequest, svc.ToggleHookResponse],
	),
)

var ToggleHookDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewCommandToggleHookDao,
			fx.As(new(svc.HookRepository)),
		),
	),
)
