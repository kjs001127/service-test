package tx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/channel-io/ch-app-store/lib/db"
)

const txKey = "tx"

var defaultDB *sql.DB

func SetDB(newDB *sql.DB) {
	defaultDB = newDB
}

func Run(
	ctx context.Context,
	body func(context.Context) error,
	sqlOptions ...Option,
) error {
	_, err := RunWith(
		ctx,
		func(ctx context.Context) (interface{}, error) {
			return nil, body(ctx)
		},
		sqlOptions...,
	)
	return err
}

func RunWith[R any](
	ctx context.Context,
	body func(context.Context) (R, error),
	sqlOptions ...Option,
) (ret R, retErr error) {
	var empty R
	if defaultDB == nil {
		return empty, fmt.Errorf("defaultDB does not exist")
	}

	txOptions := sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}
	for _, opt := range sqlOptions {
		opt.apply(&txOptions)
	}

	if ctx.Value(txKey) == nil {
		tx, err := defaultDB.BeginTx(ctx, &txOptions)
		if err != nil {
			return empty, err
		}
		ctx = context.WithValue(ctx, txKey, tx)

		defer func() {
			if err := recover(); err != nil {
				_ = tx.Rollback()
				panic(err)
			}

			if retErr != nil {
				if err := tx.Rollback(); err != nil {
					retErr = fmt.Errorf("rollback fail, origin: %w", err)
				}
			} else {
				if err := tx.Commit(); err != nil {
					retErr = err
				}
			}
		}()
	} else if _, ok := ctx.Value(txKey).(*sql.Tx); !ok {
		return empty, fmt.Errorf("found conn in context, but is not db.Conn")
	}

	return body(ctx)
}

type Option interface {
	apply(options *sql.TxOptions)
}

type IsolationOption sql.IsolationLevel

func (i IsolationOption) apply(options *sql.TxOptions) {
	options.Isolation = sql.IsolationLevel(i)
}

type ReadOnlyOption bool

func (i ReadOnlyOption) apply(options *sql.TxOptions) {
	options.ReadOnly = bool(i)
}

func WithIsolation(level sql.IsolationLevel) Option {
	return IsolationOption(level)
}

func WithReadOnly() Option {
	return ReadOnlyOption(true)
}

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

func (s *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return defaultDB.Exec(query, args...)
}

func (s *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return defaultDB.Query(query, args...)
}

func (s *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return defaultDB.QueryRow(query, args...)
}

func (s *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return withTx(ctx).ExecContext(ctx, query, args...)
}

func (s *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return withTx(ctx).QueryContext(ctx, query, args...)
}

func (s *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return withTx(ctx).QueryRowContext(ctx, query, args...)
}

func withTx(ctx context.Context) db.DB {
	if ctx.Value(txKey) == nil {
		return defaultDB
	}

	if tx, ok := ctx.Value(txKey).(db.DB); ok {
		return tx
	}

	panic(errors.New("invalid tx"))

}
