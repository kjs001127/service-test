package saga

import (
	"go.uber.org/fx"
)

var Option = fx.Provide(
	NewRegisterSaga,
)
