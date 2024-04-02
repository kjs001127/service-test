package functionfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/function/infra"
	"github.com/channel-io/ch-app-store/internal/function/repo"
	"github.com/channel-io/ch-app-store/internal/function/svc"
)

var Function = fx.Options(
	FunctionSvcs,
	FunctionHttps,
	FunctionDaos,
)

var FunctionSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewInvoker,
			fx.As(new(app.InvokeHandler)),
		),
		fx.Annotate(
			svc.NewAppHttpProxy,
			fx.ParamTags(``, restyfx.App),
		),
	),
)

var FunctionHttps = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(svc.HttpRequester)),
			fx.ParamTags(restyfx.App),
		),
	),
)

var FunctionDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppUrlDao,
			fx.As(new(svc.AppUrlRepository)),
		),
	),
)
