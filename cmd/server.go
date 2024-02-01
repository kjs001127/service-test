package main

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http"
	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/internal"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

func init() {
	// Config Set
}

// HttpModule				   godoc
//
//	@Title		ch-app-store API
//	@Version	1.0
//	@BasePath	/
func main() {
	fx.New(
		internal.Option,
		fx.Provide(
			fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`)),
			gintool.NewApiServer,
		),
		http.Option,
		fx.Invoke(func() {
			conf := config.Get()
			db, err := sql.Open(
				"postgres",
				tx.BuildDataSourceName(
					conf.Psql.Host,
					conf.Psql.DBName,
					conf.Psql.Schema,
					conf.Psql.User,
					conf.Psql.Password,
					conf.Psql.SSLMode,
				),
			)
			if err != nil {
				panic(err)
			}
			db.SetMaxOpenConns(50)

			tx.SetDB(db)
		}),
		fx.Invoke(func(srv *gintool.ApiServer) { panic(srv.Run()) }),
	)
}
