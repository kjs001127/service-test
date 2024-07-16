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
		return body(ctx)
	}

	tx, txCtx, err := beginTx(ctx, opts...)
	if err != nil {
		var empty R
		return empty, err
	}

	defer func() {
		if err := recover(); err != nil {
			_ = rollbackTx(ctx, tx, nil, opts...)
			panic(err)
		}

		if retErr != nil {
			retErr = rollbackTx(ctx, tx, retErr, opts...)
			return
		}

		retErr = commitTx(ctx, tx, opts...)
	}()

	return body(txCtx)
}

func hasTx(ctx context.Context) bool {
	return ctx.Value(txKey) != nil
}

func commitTx(ctx context.Context, tx Tx, options ...Option) error {
	for _, o := range options {
		if err := o.onCommit(ctx); err != nil {
			return rollbackTx(ctx, tx, err, options...)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed. cause: %v", err)
	}

	return nil
}

func rollbackTx(ctx context.Context, tx Tx, cause error, options ...Option) error {
	for _, o := range options {
		o.onRollback(ctx)
	}

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

	for _, option := range options {
		if err := option.onBegin(ctx); err != nil {
			return nil, nil, rollbackTx(ctx, tx, err, options...)
		}
	}

	return tx, context.WithValue(ctx, txKey, tx), nil
}

// Option is options for transaction
type Option interface {
	apply(options *sql.TxOptions)
	onBegin(ctx context.Context) error
	onCommit(ctx context.Context) error
	onRollback(ctx context.Context)
}

type IsolationOption sql.IsolationLevel

func (i IsolationOption) onBegin(ctx context.Context) error {
	return nil
}

func (i IsolationOption) onCommit(ctx context.Context) error {
	return nil
}

func (i IsolationOption) onRollback(ctx context.Context) {
}

func (i IsolationOption) apply(options *sql.TxOptions) {
	options.Isolation = sql.IsolationLevel(i)
}

type ReadOnlyOption bool

func (i ReadOnlyOption) onBegin(ctx context.Context) error {
	return nil
}

func (i ReadOnlyOption) onCommit(ctx context.Context) error {
	return nil
}

func (i ReadOnlyOption) onRollback(ctx context.Context) {
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
