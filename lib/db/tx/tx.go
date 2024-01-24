package tx

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/channel-io/ch-app-store/lib/db"
)

const txKey = "tx"

var txDB *sql.DB

func SetDB(newDB *sql.DB) {
	txDB = newDB
}

func Run(
	ctx context.Context,
	body func(context.Context) error,
	sqlOptions ...Option,
) error {
	_, err := RunWithReturn(
		ctx,
		func(ctx context.Context) (interface{}, error) {
			return nil, body(ctx)
		},
		sqlOptions...,
	)
	return err
}

func RunWithReturn[R any](
	ctx context.Context,
	body func(context.Context) (R, error),
	sqlOptions ...Option,
) (ret R, retErr error) {
	var empty R
	if txDB == nil {
		return empty, fmt.Errorf("txDB does not exist")
	}

	txOptions := sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}
	for _, opt := range sqlOptions {
		opt.apply(&txOptions)
	}

	tx, err := txDB.BeginTx(ctx, &txOptions)
	if err != nil {
		return empty, err
	}

	defer func() {
		if err := recover().(error); err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				retErr = fmt.Errorf("rollback fail. err: %v. cause: %w", txErr, err)
			} else {
				retErr = err
			}
		}
	}()

	if ctx.Value(txKey) == nil {
		ctx = context.WithValue(ctx, txKey, tx)
	} else if _, ok := ctx.Value(txKey).(db.Conn); !ok {
		return empty, fmt.Errorf("found conn in context, but is not db.Conn")
	}

	result, err := body(ctx)

	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return empty, fmt.Errorf("rollback fail. err: %v. cause: %w", txErr, err)
		}
		return empty, err
	}

	if err := tx.Commit(); err != nil {
		return empty, fmt.Errorf("tx commit fail. cause: %w", err)
	}

	return result, nil
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

func WithReadOnly(readOnly bool) Option {
	return ReadOnlyOption(readOnly)
}

type Source struct {
}

func NewSource() *Source {
	return &Source{}
}

func (s *Source) New(ctx context.Context) (db.Conn, error) {
	if val, ok := ctx.Value(txKey).(db.Conn); ok {
		return val, nil
	}

	if txDB == nil {
		return nil, fmt.Errorf("txDB does not exist")
	}

	return txDB, nil
}
