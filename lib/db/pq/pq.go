package pq

import (
	"context"
	"database/sql"
	"errors"
	"hash/fnv"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/lib/pq"

	"github.com/channel-io/ch-app-store/lib/db"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type PsqlErrMapper struct {
	delegate db.DB
}

func Wrap(delegate db.DB) db.DB {
	return &PsqlErrMapper{delegate: delegate}
}

func (p PsqlErrMapper) Exec(query string, args ...interface{}) (sql.Result, error) {
	return withErrMap(func() (sql.Result, error) {
		return p.delegate.Exec(query, args...)
	})
}

func (p PsqlErrMapper) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return withErrMap(func() (*sql.Rows, error) {
		return p.delegate.Query(query, args...)
	})
}

func (p PsqlErrMapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return withErrMap(func() (*sql.Rows, error) {
		return p.delegate.QueryContext(ctx, query, args...)
	})
}

func (p PsqlErrMapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return withErrMap(func() (sql.Result, error) {
		return p.delegate.ExecContext(ctx, query, args...)
	})
}

func (p PsqlErrMapper) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.delegate.QueryRow(query, args...)
}

func (p PsqlErrMapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.delegate.QueryRowContext(ctx, query, args...)
}

func withErrMap[RET any](f func() (RET, error)) (RET, error) {
	ret, err := f()
	return ret, mapErr(err)
}

func mapErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(err)
	}

	var pqErr *pq.Error
	ok := errors.As(err, &pqErr)
	if !ok {
		return err
	}

	switch pqErr.Code {
	case "23305":
		return apierr.Conflict(err)
	case "23503":
		return apierr.UnprocessableEntity(err)
	}

	return err
}

func Lock(ctx context.Context, tx *sql.Tx, lock tx.Lock) error {
	if lock.IsShared {
		_, err := tx.ExecContext(ctx, "SELECT pg_advisory_xact_lock_shared($1, $2)", hash(lock.Namespace), hash(lock.Id))
		return err
	} else {
		_, err := tx.ExecContext(ctx, "SELECT pg_advisory_xact_lock($1, $2)", hash(lock.Namespace), hash(lock.Id))
		return err
	}
}

func hash(s string) int32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return int32(h.Sum32())
}
