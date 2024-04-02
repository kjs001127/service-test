package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row

	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Config struct {
	Schema      string `required:"true"`
	DBName      string `required:"true"`
	Host        string `required:"true"`
	Port        string `required:"true"`
	User        string `required:"true"`
	Password    string `required:"true"`
	SSLMode     string `required:"true"`
	MaxOpenConn int    `required:"true"`
}

func BuildDataSource(driverName string, cfg Config) (*sql.DB, error) {
	opened, err := sql.Open(driverName, DataSourceName(cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error while opening psql")
	}
	opened.SetMaxOpenConns(cfg.MaxOpenConn)
	return opened, nil
}

func DataSourceName(cfg Config) string {
	return fmt.Sprintf(
		"host=%s dbname=%s search_path=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.DBName, cfg.Schema, cfg.User, cfg.Password, cfg.SSLMode,
	)
}
