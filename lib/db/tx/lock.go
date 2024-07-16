package tx

import (
	"context"
	"database/sql"
	"hash/fnv"
)

type LockOption struct {
	isShared bool
	name     string
	key      int64
}

func (l LockOption) apply(options *sql.TxOptions) {
}

func (l LockOption) onBegin(ctx context.Context) error {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		if l.isShared {
			_, err := tx.QueryContext(ctx, "SELECT pg_advisory_xact_lock_shared($1)", l.key)
			return err
		} else {
			_, err := tx.QueryContext(ctx, "SELECT pg_advisory_xact_lock($1)", l.key)
			return err
		}
	}
	return nil
}

func (l LockOption) onCommit(ctx context.Context) error {
	return nil
}

func (l LockOption) onRollback(ctx context.Context) {
}

func Xlock(name string) Option {
	return LockOption{
		name:     name,
		isShared: false,
		key:      hash(name),
	}
}

func SLock(name string) Option {
	return LockOption{
		name:     name,
		isShared: true,
		key:      hash(name),
	}
}

func hash(s string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return int64(h.Sum64())
}
