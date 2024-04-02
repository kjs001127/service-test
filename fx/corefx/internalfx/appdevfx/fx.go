package appdevfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/appdev/svc"
)

var AppDev = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAppDevSvcImpl,
			fx.As(new(svc.AppDevSvc)),
		),
	),
)
