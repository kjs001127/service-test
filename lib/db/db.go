package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row

	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func BuildDataSourceName(
	host string,
	dbName string,
	schemaName string,
	user string,
	password string,
	sslMode string,
) string {
	return fmt.Sprintf(
		"host=%s dbname=%s search_path=%s user=%s password=%s sslmode=%s",
		host, dbName, schemaName, user, password, sslMode,
	)
}
