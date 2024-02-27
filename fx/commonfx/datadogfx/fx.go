package datadogfx

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/datadog"
)

var Datadog = fx.Module("datadog",
	fx.Provide(
		fx.Annotate(
			func() http.RoundTripper {
				return datadog.WrapHttpHandler(http.DefaultTransport)
			},
			fx.ResultTags(restyfx.Dw),
		),
		fx.Annotate(
			func() http.RoundTripper {
				return datadog.WrapHttpHandler(http.DefaultTransport)
			},
			fx.ResultTags(restyfx.App),
		),
	),
	fx.Provide(
		fx.Annotate(
			datadog.NewMethodSpanTagger,
			fx.As(new(app.FunctionRequestListener)),
			fx.ResultTags(`group:"functionListeners"`),
		),
		fx.Annotate(
			datadog.NewDatadog,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"middlewares"`),
		),
	),
)
