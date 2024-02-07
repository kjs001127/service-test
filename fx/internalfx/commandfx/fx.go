package commandfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/command/repo"
)

var Option = fx.Provide(
	fx.Annotate(
		repo.NewCommandDao,
		fx.As(new(domain.CommandRepository)),
	),
	domain.NewParamValidator,
	domain.NewAutoCompleteSvc,
	domain.NewInvokeSvc,
	domain.NewQueryService,
	domain.NewRegisterService,
)
