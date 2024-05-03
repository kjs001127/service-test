package accountfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/internal/permission/infra"
	"github.com/channel-io/ch-app-store/internal/permission/svc"

	"go.uber.org/fx"
)

var AppAccount = fx.Options(
	AppAccountSvc,
	AppAccountInfra,
)

var AppAccountSvc = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAccountChannelSvc,
			fx.As(new(svc.AccountChannelSvc)),
		),
	),
)

var AppAccountInfra = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewChannelFetcherImpl,
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
			fx.As(new(svc.ChannelFetcher)),
		),
	),
)
