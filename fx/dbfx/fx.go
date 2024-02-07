package dbfx

import (
	"database/sql"
	"fmt"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

var Option = fx.Module(
	"postgres",
	fx.Provide(
		fx.Annotate(tx.NewDB, fx.As(new(db.DB))),
	),
	fx.Provide(
		BuildDataSource,
		fx.Private,
	),
	fx.Invoke(func(db *sql.DB) {
		tx.SetDB(db)
	}),
)

func BuildDataSource() (*sql.DB, error) {
	cfg := config.Get()
	psql := cfg.Psql
	name := fmt.Sprintf(
		"host=%s dbname=%s search_path=%s user=%s password=%s sslmode=%s",
		psql.Host, psql.DBName, psql.Schema, psql.User, psql.Password, psql.SSLMode,
	)

	opened, err := sql.Open("postgres", name)
	if err != nil {
		return nil, err
	}
	opened.SetMaxOpenConns(50)
	return opened, nil
}
