package dbfx

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

var Postgres = fx.Module(
	"postgres",
	fx.Provide(
		db.BuildDataSource,
	),
	fx.Invoke(func(db *sql.DB) {
		tx.SetDB(db)
	}),

	fx.Provide(
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),
	),
)
