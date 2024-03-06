package dbfx

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/datadog"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

const (
	postgres   = "postgres"
	driverName = `name:"postgres"`
)

var Postgres = fx.Module(
	"postgres",

	fx.Supply(
		fx.Annotate(
			postgres,
			fx.ResultTags(driverName),
		),
	),
	fx.Provide(
		fx.Annotate(
			datadog.NewDataSource,
			fx.ParamTags(driverName, ``),
		),
		fx.Private,
	),

	fx.Invoke(func(db *sql.DB) {
		tx.EnableDatabase(db)
	}),

	fx.Provide(
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),
	),
)
