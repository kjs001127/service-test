package brieffx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/brief/repo"
	"github.com/channel-io/ch-app-store/internal/brief/svc"
)

var Brief = fx.Options(
	BriefSvcs,
	BriefDAOs,
)

var BriefSvcs = fx.Options(
	fx.Provide(
		svc.NewInvoker,
		app.NewTypedInvoker[svc.BriefRequest, svc.BriefResponse],
		fx.Annotate(
			svc.NewBriefClearHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)

var BriefDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(repo.NewBriefDao, fx.As(new(svc.BriefRepository))),
	),
)
