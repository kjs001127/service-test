package tx

import (
	"context"
	"database/sql"
)

type IdleTransactor struct {
}

func (i IdleTransactor) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	return &IdleTx{}, nil
}

type IdleTx struct {
}

func (i IdleTx) Commit() error {
	return nil
}

func (i IdleTx) Rollback() error {
	return nil
}
