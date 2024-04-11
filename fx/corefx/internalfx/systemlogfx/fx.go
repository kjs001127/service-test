package systemlogfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/systemlog/repo"
	systemlog "github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

var SystemLog = fx.Options(
	SystemLogSvc,
	SystemLogRepo,
)

var SystemLogSvc = fx.Options(
	fx.Provide(
		systemlog.NewSystemLogSvc,
	),
)

var SystemLogRepo = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewSystemLogRepository,
			fx.As(new(systemlog.SystemLogRepository)),
		),
	),
)
