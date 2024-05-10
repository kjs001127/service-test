package integration

import (
	"context"

	datasource "github.com/channel-io/ch-app-store/lib/db"

	"go.uber.org/fx"
)

type TestHelper struct {
	app *fx.App
	DB  *datasource.DB
}

func NewTestHelper(testOpts ...fx.Option) *TestHelper {
	var dataSource datasource.DB

	fxApp := fx.New(
		append(
			testOpts,
			fx.Populate(&dataSource),
		)...,
	)

	return &TestHelper{
		app: fxApp,
		DB:  &dataSource,
	}
}

func Mock[T any](mock any, anns ...fx.Annotation) fx.Option {
	return fx.Provide(
		fx.Annotate(
			func() T { return mock.(T) },
			anns...,
		),
	)
}

func (t *TestHelper) Stop() {
	err := t.app.Stop(context.Background())
	if err != nil {
		panic(err)
	}
}

func (t *TestHelper) WithPreparedTables(tableNames ...string) *TestHelper {
	for _, tableName := range tableNames {
		dataSource := *t.DB
		if _, err := dataSource.Exec("TRUNCATE TABLE " + tableName + " CASCADE"); err != nil {
			panic(err)
		}
	}
	return t
}
