package brieffx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/internal/brief/repo"
)

var Brief = fx.Module(
	"brief",
	BriefSvcs,
	BriefDAOs,
)

var BriefSvcs = fx.Module(
	"briefDomain",
	fx.Provide(
		domain.NewInvoker,
		app.NewTypedInvoker[domain.EmptyRequest, domain.BriefResponse],
	),
)

var BriefDAOs = fx.Module(
	"briefDB",
	fx.Provide(
		fx.Annotate(repo.NewBriefDao, fx.As(new(domain.BriefRepository))),
	),
)
