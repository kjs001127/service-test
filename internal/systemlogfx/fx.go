package systemlogfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/internal/systemlog/repo"
	systemlog "github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

const (
	ddbTableNameTag = `name:"systemLogTableName"`
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
	fx.Supply(
		fx.Annotate(
			config.Get().SystemLogTableName,
			fx.ResultTags(ddbTableNameTag),
		),
	),
	fx.Provide(
		fx.Annotate(
			repo.NewSystemLogRepository,
			fx.As(new(systemlog.SystemLogRepository)),
			fx.ParamTags(``, ddbTableNameTag),
		),
	),
)
