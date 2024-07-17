package tx

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

const txKey = "tx"

var transactor Transactor = &IdleTransactor{}

func Do(
	ctx context.Context,
	body func(context.Context) error,
	sqlOptions ...Option,
) error {
	_, err := DoReturn(
		ctx,
		func(ctx context.Context) (interface{}, error) {
			return nil, body(ctx)
		},
		sqlOptions...,
	)
	return err
}

func DoReturn[R any](
	ctx context.Context,
	body func(context.Context) (R, error),
	opts ...Option,
) (ret R, retErr error) {
	if transactor == nil {
		panic(errors.New("transactor is not configured"))
	}

	if hasTx(ctx) {
		for _, opt := range opts {
			if err := opt.onBegin(ctx); err != nil {
				retErr = err
				return
			}
		}
		return body(ctx)
	}

	tx, txCtx, err := beginTx(ctx, opts...)
	if err != nil {
		var empty R
		return empty, err
	}

	defer func() {
		if err := recover(); err != nil {
			_ = rollbackTx(tx, nil)
			panic(err)
		}

		if retErr != nil {
			retErr = rollbackTx(tx, retErr)
		} else {
			retErr = commitTx(tx)
		}
	}()

	for _, opt := range opts {
		if err := opt.onBegin(ctx); err != nil {
			retErr = err
			return
		}
	}

	return body(txCtx)
}

func hasTx(ctx context.Context) bool {
	return ctx.Value(txKey) != nil
}

func commitTx(tx Tx) error {
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed. cause: %v", err)
	}

	return nil
}

func rollbackTx(tx Tx, cause error) error {
	if err := tx.Rollback(); err != nil {
		return fmt.Errorf("rollback fail, err:%v rollback cause: %w", err, cause)
	}
	return cause
}

func beginTx(ctx context.Context, options ...Option) (Tx, context.Context, error) {
	txOptions := sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}
	for _, opt := range options {
		opt.apply(&txOptions)
	}
	tx, err := transactor.BeginTx(ctx, &txOptions)
	if err != nil {
		return nil, nil, err
	}

	return tx, context.WithValue(ctx, txKey, tx), nil
}

// Option is options for transaction
type Option interface {
	apply(options *sql.TxOptions)
	onBegin(ctx context.Context) error
}

type IsolationOption sql.IsolationLevel

func (i IsolationOption) onBegin(ctx context.Context) error {
	return nil
}

func (i IsolationOption) apply(options *sql.TxOptions) {
	options.Isolation = sql.IsolationLevel(i)
}

type ReadOnlyOption bool

func (i ReadOnlyOption) onBegin(ctx context.Context) error {
	return nil
}

func (i ReadOnlyOption) apply(options *sql.TxOptions) {
	options.ReadOnly = bool(i)
}

func Isolation(level sql.IsolationLevel) Option {
	return IsolationOption(level)
}

func ReadOnly() Option {
	return ReadOnlyOption(true)
}
