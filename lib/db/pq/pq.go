package pq

import (
	"context"
	"database/sql"
	"errors"
	"hash/fnv"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/lib/pq"

	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type PsqlErrMapper struct {
}

func NewPsqlErrMapper() *PsqlErrMapper {
	return &PsqlErrMapper{}
}

func (p *PsqlErrMapper) Map(err error) error {
	if err == nil {
		return nil
	}

	var pqErr *pq.Error
	ok := errors.As(err, &pqErr)
	if !ok {
		return err
	}

	// https://www.postgresql.org/docs/8.4/errcodes-appendix.html
	class := pqErr.Code[:2]
	switch class {
	case "20":
		return apierr.NotFound(err)
	case "22":
		return apierr.BadRequest(err)
	case "23":
		return apierr.Conflict(err)
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
