package dbfx

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/pq"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

const (
	postgres = "postgres"
)

var DB = fx.Options(

	Postgres,

	fx.Invoke(func(db *sql.DB, lock tx.LockFn) {
		tx.EnableDatabase(db)
		tx.EnableLock(lock)
	}),

	fx.Provide(
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),
		fx.Annotate(
			db.BuildDataSource,
		),
	),
)

var Postgres = fx.Options(
	fx.Supply(
		db.DriverName(postgres),
		pq.LockFn,
	),
	fx.Provide(
		fx.Annotate(
			pq.NewPsqlErrMapper,
			fx.As(new(db.ErrMapper)),
		),
	),
)
