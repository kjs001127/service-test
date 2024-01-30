package tx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/db"
)

var FxOption = fx.Provide(
	fx.Annotate(
		NewDB,
		fx.As(new(db.DB)),
	),
)
