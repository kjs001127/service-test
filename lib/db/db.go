package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/channel-io/ch-app-store/config"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row

	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

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
