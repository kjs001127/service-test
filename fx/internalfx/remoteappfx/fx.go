package remoteappfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-app-store/internal/remoteapp/infra"
	"github.com/channel-io/ch-app-store/internal/remoteapp/repo"
)

var Option = fx.Provide(
	fx.Annotate(
		infra.NewHttpRequester,
		fx.As(new(domain.HttpRequester)),
	),
	fx.Annotate(
		repo.NewAppDAO,
		fx.As(new(domain.RemoteAppRepository)),
	),
	fx.Annotate(
		domain.NewAppRepositoryAdapter,
		fx.As(new(app.AppRepository)),
	),
)
