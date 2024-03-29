package repo

import (
	"context"
	"database/sql"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type BriefDao struct {
	db db.DB
}

func (b BriefDao) DeleteAllByAppID(ctx context.Context, appID string) error {
	//TODO implement me
	panic("implement me")
}

func NewBriefDao(db db.DB) *BriefDao {
	return &BriefDao{db: db}
}

func (b BriefDao) Fetch(ctx context.Context, appID string) (*domain.Brief, error) {
	one, err := models.Briefs(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).One(ctx, b.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, apierr.NotFound(err)
	} else if err != nil {
		return nil, errors.Wrap(err, "error while querying brief")
	}

	return unmarshal(one), nil
}

func (b BriefDao) FetchAll(ctx context.Context, appIDs []string) ([]*domain.Brief, error) {
	slice := make([]interface{}, len(appIDs))
	for i, v := range appIDs {
		slice[i] = v
	}

	all, err := models.Briefs(
		qm.Select("*"),
		qm.WhereIn("app_id IN ?", slice...),
	).All(ctx, b.db)

	if err != nil {
		return nil, errors.Wrap(err, "error while querying brief")
	}

	return unmarshalAll(all), nil
}

func (b *BriefDao) DeleteByAppID(ctx context.Context, appID string) error {
	_, err := models.Briefs(qm.Where("app_id = $1", appID)).DeleteAll(ctx, b.db)
	return errors.WithStack(err)
}

func unmarshal(model *models.Brief) *domain.Brief {
	return &domain.Brief{
		AppID:             model.AppID,
		ID:                model.ID,
		BriefFunctionName: model.BriefFunctionName,
	}
}

func unmarshalAll(models models.BriefSlice) []*domain.Brief {
	ret := make([]*domain.Brief, 0, len(models))
	for _, m := range models {
		ret = append(ret, unmarshal(m))
	}
	return ret
}
