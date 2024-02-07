package brieffx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/internal/brief/repo"
)

var Option = fx.Provide(
	fx.Annotate(repo.NewBriefDao, fx.As(new(domain.BriefRepository))),
)
