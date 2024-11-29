package dbfx

import (
	"database/sql"

	"go.uber.org/fx"

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

	fx.Invoke(func(db *sql.DB) {
		tx.EnableDatabase(db)
	}),

	fx.Provide(
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),
		fx.Annotate(
			db.BuildDataSource,
			fx.ParamTags(driverName),
		),
	),
)
