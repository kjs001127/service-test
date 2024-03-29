package brieffx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/internal/brief/repo"
)

var Brief = fx.Options(
	BriefSvcs,
	BriefDAOs,
)

var BriefSvcs = fx.Options(
	fx.Provide(
		domain.NewInvoker,
		app.NewTypedInvoker[domain.EmptyRequest, domain.BriefResponse],
		fx.Annotate(
			domain.NewBriefClearHook,
			fx.As(new(app.AppLifeCycleHook)),
			fx.ResultTags(appfx.LifecycleHookGroup),
		),
	),
)

var BriefDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(repo.NewBriefDao, fx.As(new(domain.BriefRepository))),
	),
)
