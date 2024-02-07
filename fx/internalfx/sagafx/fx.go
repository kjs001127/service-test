package sagafx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/saga"
)

var Option = fx.Provide(
	saga.NewRegisterSaga,
)
