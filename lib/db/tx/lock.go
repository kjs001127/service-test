package tx

import (
	"context"
	"database/sql"
	"hash/fnv"
)

type lockOption struct {
	isShared  bool
	namespace int32
	id        int32
}

func (l lockOption) apply(options *sql.TxOptions) {
}

func (l lockOption) onBegin(ctx context.Context) error {

	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		if l.isShared {
			_, err := tx.ExecContext(ctx, "SELECT pg_advisory_xact_lock_shared($1, $2)", l.namespace, l.id)
			return err
		} else {
			_, err := tx.ExecContext(ctx, "SELECT pg_advisory_xact_lock($1, $2)", l.namespace, l.id)
			return err
		}
	}
	return nil
}

func XLock(namespace string, id string) Option {
	return lockOption{
		namespace: hash(namespace),
		id:        hash(id),
		isShared:  false,
	}
}

func SLock(namespace string, id string) Option {
	return lockOption{
		namespace: hash(namespace),
		id:        hash(id),
		isShared:  true,
	}
}

func hash(s string) int32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return int32(h.Sum32())
}
