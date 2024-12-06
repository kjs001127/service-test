package pq

import (
	"context"
	"database/sql"
	"errors"
	"hash/fnv"
	"reflect"
	"unsafe"

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
	return wrapRowsErr(func() (sql.Result, error) {
		return p.delegate.Exec(query, args...)
	})
}

func (p PsqlErrMapper) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return wrapRowsErr(func() (*sql.Rows, error) {
		return p.delegate.Query(query, args...)
	})
}

func (p PsqlErrMapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return wrapRowsErr(func() (*sql.Rows, error) {
		return p.delegate.QueryContext(ctx, query, args...)
	})
}

func (p PsqlErrMapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return wrapRowsErr(func() (sql.Result, error) {
		return p.delegate.ExecContext(ctx, query, args...)
	})
}

func (p PsqlErrMapper) QueryRow(query string, args ...interface{}) *sql.Row {
	return wrapRowErr(p.delegate.QueryRow(query, args...))
}

func (p PsqlErrMapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return wrapRowErr(p.delegate.QueryRowContext(ctx, query, args...))
}

const errFieldName = "err"

func wrapRowErr(row *sql.Row) *sql.Row {
	if row.Err() != nil {
		rowValue := reflect.ValueOf(row).Elem()
		errField := rowValue.FieldByName(errFieldName)

		unsafeErrField := reflect.NewAt(errField.Type(), unsafe.Pointer(errField.UnsafeAddr())).Elem()
		unsafeErrField.Set(reflect.ValueOf(mapErr(row.Err())))
	}

	return row
}

func wrapRowsErr[RET any](f func() (RET, error)) (RET, error) {
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

var LockFn tx.LockFn = func(ctx context.Context, tx *sql.Tx, lock tx.Lock) error {
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
