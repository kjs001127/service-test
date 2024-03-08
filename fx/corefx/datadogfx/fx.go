package datadogfx

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/lib/datadog"
)

var Datadog = fx.Options(
	fx.Decorate(
		fx.Annotate(
			func(client *resty.Client) *resty.Client {
				return datadog.DecorateResty(client)
			},
			fx.ResultTags(restyfx.Dw),
			fx.ParamTags(restyfx.Dw),
		),
		fx.Annotate(
			func(client *resty.Client) *resty.Client {
				return datadog.DecorateResty(client)
			},
			fx.ResultTags(restyfx.App),
			fx.ParamTags(restyfx.App),
		),
		datadog.DecorateLogger,
	),
	fx.Provide(
		fx.Annotate(
			datadog.NewMethodSpanTagger,
			fx.As(new(app.FunctionRequestListener)),
			fx.ResultTags(appfx.FunctionListenersGroup),
		),
		fx.Annotate(
			datadog.NewGinMiddleware,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
	),
)