package tx

import (
	"context"
	"database/sql"
	"hash/fnv"
)

type lockOption struct {
	isShared bool
	name     string
	key      int64
}

func (l lockOption) apply(options *sql.TxOptions) {
}

func (l lockOption) onBegin(ctx context.Context) error {
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

func (l lockOption) onCommit(ctx context.Context) error {
	return nil
}

func (l lockOption) onRollback(ctx context.Context) {
}

func Xlock(name string) Option {
	return lockOption{
		name:     name,
		isShared: false,
		key:      hash(name),
	}
}

func SLock(name string) Option {
	return lockOption{
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
