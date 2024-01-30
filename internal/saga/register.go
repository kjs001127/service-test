package saga

import (
	"context"
	"database/sql"

	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type RegisterSaga struct {
	svc *domain.RegisterSvc
}

func NewRegisterSaga(svc *domain.RegisterSvc) *RegisterSaga {
	return &RegisterSaga{svc: svc}
}

func (s *RegisterSaga) Register(ctx context.Context, req domain.RegisterRequest) error {
	return tx.Run(ctx, func(ctx context.Context) error {
		return s.svc.Register(ctx, req)
	}, tx.WithIsolation(sql.LevelSerializable))
}
