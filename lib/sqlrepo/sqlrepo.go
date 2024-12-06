package sqlrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"

	"github.com/channel-io/go-lib/pkg/errors"
	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/channel-io/ch-app-store/lib/db"
)

const columnID = "id"

type ReadSQLRepo[D DomainModel] interface {
	Fetch(ctx context.Context, id string) (D, error)
	FetchBy(ctx context.Context, mods ...qm.QueryMod) (D, error)
	Find(ctx context.Context, id string) (D, error)
	FindBy(ctx context.Context, mods ...qm.QueryMod) (D, error)
	FindAll(ctx context.Context, ids ...string) ([]D, error)
	FindAllBy(ctx context.Context, mods ...qm.QueryMod) ([]D, error)
	CountBy(ctx context.Context, mods ...qm.QueryMod) (int64, error)
}

type SQLRepo[D DomainModel] interface {
	ReadSQLRepo[D]

	Create(ctx context.Context, d D) (D, error)
	Update(ctx context.Context, d D) (D, error)
	Upsert(ctx context.Context, d D, conflictColumns ...string) (D, error)
	Delete(ctx context.Context, id string) error
	DeleteBy(ctx context.Context, mods ...qm.QueryMod) error
	DeleteAllBy(ctx context.Context, mods ...qm.QueryMod) error
}

func New[D DomainModel, B BoilModel, BS BoilModelSlice[B]](
	db db.DB,
	dtbFunc DTBFunc[D, B],
	btdFunc BTDFunc[D, B],
	modelsFunc QueryFunc[B, BS],
	errMapper db.ErrMapper,
) SQLRepo[D] {
	return &sqlRepo[D, B, BS]{
		db:         db,
		dtbFunc:    dtbFunc,
		btdFunc:    btdFunc,
		modelsFunc: modelsFunc,
		errMapper:  errMapper,
	}
}

type sqlRepo[D DomainModel, B BoilModel, BS BoilModelSlice[B]] struct {
	db         db.DB
	errMapper  db.ErrMapper
	dtbFunc    DTBFunc[D, B]
	btdFunc    BTDFunc[D, B]
	modelsFunc QueryFunc[B, BS]
}

func (repo *sqlRepo[D, B, BS]) Fetch(ctx context.Context, id string) (D, error) {
	domainModel, err := repo.FetchBy(ctx, qmhelper.Where(columnID, qmhelper.EQ, id))
	return domainModel, repo.errMapper.Map(err)
}

func (repo *sqlRepo[D, B, BS]) FetchBy(ctx context.Context, mods ...qm.QueryMod) (D, error) {
	domainModel, err := repo.FindBy(ctx, mods...)
	if err != nil {
		return lo.Empty[D](), err
	}
	if lo.IsEmpty(domainModel) {
		return lo.Empty[D](), errors.Wrap(apierr.NotFound(
			errors.New(fmt.Sprintf("%T not found", domainModel)),
		))
	}
	return domainModel, nil
}

func (repo *sqlRepo[D, B, BS]) Find(ctx context.Context, id string) (D, error) {
	boilModel, err := repo.FindBy(ctx, qmhelper.Where(columnID, qmhelper.EQ, id))
	return boilModel, repo.errMapper.Map(err)
}

func (repo *sqlRepo[D, B, BS]) FindBy(ctx context.Context, mods ...qm.QueryMod) (D, error) {
	boilModel, err := repo.modelsFunc(mods...).One(ctx, repo.db)
	if errors.Is(err, sql.ErrNoRows) {
		return lo.Empty[D](), nil
	}
	if err != nil {
		return lo.Empty[D](), repo.errMapper.Map(err)
	}
	return repo.btdFunc(boilModel)
}

func (repo *sqlRepo[D, B, BS]) FindAll(ctx context.Context, ids ...string) ([]D, error) {
	ret, err := repo.FindAllBy(ctx, qm.WhereIn(`"id" IN ?`, lo.Map(ids, func(id string, _ int) interface{} {
		return id
	})...))
	return ret, repo.errMapper.Map(err)
}

func (repo *sqlRepo[D, B, BS]) FindAllBy(ctx context.Context, mods ...qm.QueryMod) ([]D, error) {
	boilModels, err := repo.modelsFunc(mods...).All(ctx, repo.db)
	if err != nil {
		return lo.Empty[[]D](), repo.errMapper.Map(err)
	}

	var domainModels []D
	for _, boilModel := range boilModels {
		domainModel, err := repo.btdFunc(boilModel)
		if err != nil {
			return lo.Empty[[]D](), errors.Wrap(err)
		}
		domainModels = append(domainModels, domainModel)
	}
	return domainModels, nil
}

func (repo *sqlRepo[D, B, BS]) CountBy(ctx context.Context, mods ...qm.QueryMod) (int64, error) {
	ret, err := repo.modelsFunc(mods...).Count(ctx, repo.db)
	return ret, repo.errMapper.Map(err)
}

func (repo *sqlRepo[D, B, BS]) Create(ctx context.Context, domainModel D) (D, error) {
	boilModel, err := repo.dtbFunc(domainModel)
	if err != nil {
		return lo.Empty[D](), errors.Wrap(err)
	}
	if err := boilModel.Insert(ctx, repo.db, boil.Infer()); err != nil {
		return lo.Empty[D](), repo.errMapper.Map(err)
	}
	return repo.btdFunc(boilModel)
}

func (repo *sqlRepo[D, B, BS]) Update(ctx context.Context, domainModel D) (D, error) {
	boilModel, err := repo.dtbFunc(domainModel)
	if err != nil {
		return lo.Empty[D](), errors.Wrap(err)
	}

	if _, err := boilModel.Update(ctx, repo.db, boil.Infer()); err != nil {
		return lo.Empty[D](), repo.errMapper.Map(err)
	}

	return repo.btdFunc(boilModel)
}

func (repo *sqlRepo[D, B, BS]) Upsert(ctx context.Context, domainModel D, conflictColumns ...string) (D, error) {
	boilModel, err := repo.dtbFunc(domainModel)
	if err != nil {
		return lo.Empty[D](), errors.Wrap(err)
	}

	if err = boilModel.Upsert(
		ctx,
		repo.db,
		true,
		conflictColumns,
		boil.Blacklist(conflictColumns...),
		boil.Infer(),
	); err != nil {
		return lo.Empty[D](), repo.errMapper.Map(err)
	}

	return repo.btdFunc(boilModel)
}

func (repo *sqlRepo[D, B, BS]) Delete(ctx context.Context, id string) error {
	return repo.DeleteBy(ctx, qmhelper.Where(columnID, qmhelper.EQ, id))
}

func (repo *sqlRepo[D, B, BS]) DeleteBy(ctx context.Context, mods ...qm.QueryMod) error {
	boilModel, err := repo.modelsFunc(mods...).One(ctx, repo.db)
	if errors.Is(err, sql.ErrNoRows) {
		return apierr.NotFound(err)
	} else if err != nil {
		return repo.errMapper.Map(err)
	}

	_, err = boilModel.Delete(ctx, repo.db)
	if err != nil {
		return repo.errMapper.Map(err)
	}
	return errors.Wrap(err)
}

func (repo *sqlRepo[D, B, BS]) DeleteAllBy(ctx context.Context, mods ...qm.QueryMod) error {
	_, err := repo.modelsFunc(mods...).DeleteAll(ctx, repo.db)
	return repo.errMapper.Map(err)
}
