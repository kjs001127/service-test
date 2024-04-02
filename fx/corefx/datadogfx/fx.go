package datadogfx

import (
	"database/sql"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/tx"

	"github.com/channel-io/ch-app-store/lib/datadog"
)

const (
	postgres   = "postgres"
	driverName = `name:"postgres"`
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
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),

		fx.Annotate(
			datadog.NewDataSource,
			fx.ParamTags(driverName, ``),
		),

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

	fx.Supply(
		fx.Annotate(
			postgres,
			fx.ResultTags(driverName),
		),
	),

	fx.Invoke(func(db *sql.DB) {
		tx.EnableDatabase(db)
	}),
)
