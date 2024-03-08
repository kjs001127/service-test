package brieffx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/internal/brief/repo"
)

var Brief = fx.Options(
	BriefSvcs,
	BriefDAOs,
)

var BriefSvcs = fx.Options(
	fx.Provide(
		domain.NewInvoker,
		app.NewTypedInvoker[domain.EmptyRequest, domain.BriefResponse],
	),
)

var BriefDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(repo.NewBriefDao, fx.As(new(domain.BriefRepository))),
	),
)
