package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/systemlog/model"
)

type SystemLogRepository interface {
	Save(ctx context.Context, log *model.SystemLog) error
	Query(ctx context.Context, req *QueryRequest) ([]*model.SystemLog, error)
}
