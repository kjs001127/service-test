package accountfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/internal/account/svc"

	"go.uber.org/fx"
)

var AppAccount = fx.Options(
	AppAccountSvc,
)

var AppAccountSvc = fx.Options(
	fx.Provide(
		fx.Annotate(
			svc.NewAccountChannelSvc,
			fx.As(new(svc.AccountChannelSvc)),
		),
	),
	fx.Provide(
		fx.Annotate(
			svc.NewChannelFetcherImpl,
			fx.As(new(svc.ChannelFetcher)),
			fx.ParamTags(restyfx.Dw, configfx.DwAdmin),
		),
	),
)
