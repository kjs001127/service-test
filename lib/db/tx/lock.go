package tx

import (
	"context"
	"database/sql"
)

type LockFn func(ctx context.Context, tx *sql.Tx, lock Lock) error
type Lock struct {
	Namespace string
	Id        string
	IsShared  bool
}

var lockFn LockFn = func(ctx context.Context, tx *sql.Tx, lock Lock) error {
	return nil
}

func EnableLock(fn LockFn) {
	lockFn = fn
}

type lockOption struct {
	lock Lock
}

func (l lockOption) apply(options *sql.TxOptions) {
}

func (l lockOption) onBegin(ctx context.Context) error {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {

		if err := lockFn(ctx, tx, l.lock); err != nil {
			return err
		}
	}
	return nil
}

func XLock(namespace string, id string) Option {
	return lockOption{
		lock: Lock{
			Namespace: namespace,
			Id:        id,
			IsShared:  false,
		},
	}
}

func SLock(namespace string, id string) Option {
	return lockOption{
		lock: Lock{
			Namespace: namespace,
			Id:        id,
			IsShared:  true,
		},
	}
}
