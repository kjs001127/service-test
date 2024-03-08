package tx

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Transactor interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

type Tx interface {
	Commit() error
	Rollback() error
}
