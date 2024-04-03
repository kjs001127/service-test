package datadog

import (
	"database/sql"

	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"

	"github.com/channel-io/ch-app-store/lib/db"
)

func NewDataSource(driverName string, cfg db.Config) (*sql.DB, error) {
	open, err := sqltrace.Open(driverName, db.DataSourceName(cfg))
	if err != nil {
		return nil, err
	}
	open.SetMaxOpenConns(cfg.MaxOpenConn)
	return open, nil
}
