package datadog

import (
	"database/sql"
	"time"

	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"

	_ "github.com/lib/pq"

	"github.com/channel-io/ch-app-store/lib/db"
)

func NewDataSource(driverName string, cfg db.Config) (*sql.DB, error) {
	open, err := sqltrace.Open(driverName, db.DataSourceName(cfg))
	if err != nil {
		return nil, err
	}
	open.SetMaxOpenConns(cfg.MaxOpenConn)
	open.SetMaxIdleConns(cfg.MaxOpenConn)
	open.SetConnMaxLifetime(1 * time.Hour)
	open.SetConnMaxIdleTime(15 * time.Minute)
	return open, nil
}
