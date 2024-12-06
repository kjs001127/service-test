package tx

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/lib/db"
)

func EnableDatabase(DB *sql.DB) {
	transactor = &dbAdapter{DB: DB}
}

type dbAdapter struct {
	*sql.DB
}

func (s *dbAdapter) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	return s.DB.BeginTx(ctx, opts)
}

/*
DB is a Transaction-aware DB implementation
*/
type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

func (s *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return s.txIfExists(ctx).PrepareContext(ctx, query)
}

func (s *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.defaultDB().Exec(query, args...)
}

func (s *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.defaultDB().Query(query, args...)
}

func (s *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.defaultDB().QueryRow(query, args...)
}

func (s *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.txIfExists(ctx).ExecContext(ctx, query, args...)
}

func (s *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.txIfExists(ctx).QueryContext(ctx, query, args...)
}

func (s *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return s.txIfExists(ctx).QueryRowContext(ctx, query, args...)
}

func (s *DB) defaultDB() db.DB {
	return transactor.(*dbAdapter)
}

func (s *DB) txIfExists(ctx context.Context) db.DB {
	if ctx.Value(txKey) == nil {
		return s.defaultDB()
	}

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return tx
	}

	panic(errors.New("invalid tx"))
}
