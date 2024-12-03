package dbfx

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/pq"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

const (
	postgres   = "postgres"
	DriverName = `name:"postgres"`
)

var DB = fx.Options(

	Postgres,

	fx.Invoke(func(db *sql.DB) {
		tx.EnableDatabase(db)
	}),

	fx.Provide(
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),
		fx.Annotate(
			db.BuildDataSource,
			fx.ParamTags(DriverName),
		),
	),
)

var Postgres = fx.Options(
	fx.Supply(
		fx.Annotate(
			postgres,
			fx.ResultTags(DriverName),
		),
	),

	fx.Decorate(pq.Wrap),

	fx.Invoke(func() {
		tx.EnableLock(pq.Lock)
	}),
)
