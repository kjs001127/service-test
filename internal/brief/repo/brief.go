package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/channel-io/ch-app-store/generated/models"
	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/lib/db"
)

type BriefDao struct {
	db db.DB
}

func (b BriefDao) Fetch(ctx context.Context, appID string) (*domain.Brief, error) {
	one, err := models.Briefs(
		qm.Select("*"),
		qm.Where("app_id = $1", appID),
	).One(ctx, b.db)
	if err != nil {
		return nil, err
	}

	return unmarshal(one), nil
}

func (b BriefDao) FetchAll(ctx context.Context, appIDs []string) ([]*domain.Brief, error) {
	all, err := models.Briefs(
		qm.Select("*"),
		qm.WhereIn("app_id IN $1", appIDs),
	).All(ctx, b.db)

	if err != nil {
		return nil, err
	}

	return unmarshalAll(all), nil
}

func unmarshal(model *models.Brief) *domain.Brief {
	return &domain.Brief{
		AppID:        model.AppID,
		ID:           model.ID,
		FunctionName: model.FunctionID,
	}
}

func unmarshalAll(models models.BriefSlice) []*domain.Brief {
	ret := make([]*domain.Brief, len(models))
	for _, m := range models {
		ret = append(ret, unmarshal(m))
	}
	return ret
}
