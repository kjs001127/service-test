package main

import (
	"database/sql"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/http/admin"
	"github.com/channel-io/ch-app-store/api/http/desk"
	"github.com/channel-io/ch-app-store/api/http/front"
	"github.com/channel-io/ch-app-store/api/http/general"
	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/internal"
	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type Module int

const (
	DeskModule = iota
	FrontModule
	GeneralModule
	AdminModule
)

func init() {
	// Config Set
}

func main() {
	startModule(DeskModule, FrontModule, GeneralModule, AdminModule)
}

func startModule(modules ...Module) {
	if len(modules) == 0 {
		panic("no module to start")
	}

	for _, module := range modules {
		var option fx.Option
		switch module {
		case DeskModule:
			option = desk.HttpModule()
		case FrontModule:
			option = front.HttpModule()
		case GeneralModule:
			option = general.HttpModule()
		case AdminModule:
			option = admin.HttpModule()
		}

		go fx.New(
			internal.Option,
			option,
			fx.Invoke(func() {
				conf := config.Get()
				database, err := sql.Open(
					"postgres",
					db.BuildDataSourceName(
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
				database.SetMaxOpenConns(50)

				tx.SetDB(database)
			}),
		)
	}

	select {}
}
