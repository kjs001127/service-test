package brieffx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/internal/brief/repo"
)

var Brief = fx.Module(
	"brief",
	BriefDomain,
	BriefDB,
)

var BriefDomain = fx.Module(
	"briefDomain",
	fx.Provide(
		domain.NewInvoker,
		app.NewTypedInvoker[domain.EmptyRequest, domain.BriefResponse],
	),
)

var BriefDB = fx.Module(
	"briefDB",
	fx.Provide(
		fx.Annotate(repo.NewBriefDao, fx.As(new(domain.BriefRepository))),
	),
)
