package sqlrepo

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type DomainModel interface {
	comparable
}

type BoilModel interface {
	comparable
	Insert(ctx context.Context, executor boil.ContextExecutor, columns boil.Columns) error
	Update(ctx context.Context, executor boil.ContextExecutor, columns boil.Columns) (int64, error)
	Delete(ctx context.Context, executor boil.ContextExecutor) (int64, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
	Reload(ctx context.Context, exec boil.ContextExecutor) error
	Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns boil.Columns, insertColumns boil.Columns) error
}

type BoilModelQuery[M BoilModel, S BoilModelSlice[M]] interface {
	Exec(exec boil.Executor) (sql.Result, error)
	QueryRow(exec boil.Executor) *sql.Row
	ExecContext(ctx context.Context, exec boil.ContextExecutor) (sql.Result, error)
	QueryRowContext(ctx context.Context, exec boil.ContextExecutor) *sql.Row
	QueryContext(ctx context.Context, exec boil.ContextExecutor) (*sql.Rows, error)
	ExecP(exec boil.Executor) sql.Result
	QueryP(exec boil.Executor) *sql.Rows
	BindP(ctx context.Context, exec boil.Executor, obj interface{})
	BindG(ctx context.Context, obj interface{}) error
	Bind(ctx context.Context, exec boil.Executor, obj interface{}) error

	One(ctx context.Context, exec boil.ContextExecutor) (M, error)
	All(ctx context.Context, exec boil.ContextExecutor) (S, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

type BoilModelSlice[M BoilModel] interface {
	~[]M
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

type DTBFunc[D DomainModel, B BoilModel] func(D) (B, error)
type BTDFunc[D DomainModel, B BoilModel] func(B) (D, error)
type QueryFunc[B BoilModel, BS BoilModelSlice[B]] func(...qm.QueryMod) BoilModelQuery[B, BS]
