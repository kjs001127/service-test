package integration

import (
	"context"

	"github.com/channel-io/ch-app-store/fx/cmdfx/testfx"
	datasource "github.com/channel-io/ch-app-store/lib/db"

	"go.uber.org/fx"
)

type TestApp struct {
	app *fx.App
	DB  *datasource.DB
}

var Populate = fx.Populate

func NewTestApp(testOpts ...fx.Option) *TestApp {
	var dataSource datasource.DB

	fxApp := fx.New(
		append(
			testOpts,
			testfx.Test,
			fx.Populate(&dataSource),
		)...,
	)

	return &TestApp{
		app: fxApp,
		DB:  &dataSource,
	}
}

func Mock[T any](mock any, anns ...fx.Annotation) fx.Option {
	return fx.Decorate(
		fx.Annotate(
			func() T { return mock.(T) },
			anns...,
		),
	)
}

func (t *TestApp) Stop() {
	err := t.app.Stop(context.Background())
	if err != nil {
		return
	}
}

func (t *TestApp) WithPreparedTables(tableNames ...string) *TestApp {
	for _, tableName := range tableNames {
		dataSource := *t.DB
		if _, err := dataSource.Exec("TRUNCATE TABLE " + tableName + " CASCADE"); err != nil {
			panic(err)
		}
	}
	return t
}
