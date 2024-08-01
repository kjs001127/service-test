package i18nfx

import (
	"github.com/channel-io/ch-app-store/lib/i18n"

	"go.uber.org/fx"
)

var I18n = fx.Options(
	fx.Provide(
		fx.Annotate(
			i18n.NewI18nImpl,
			fx.As(new(i18n.I18n)),
		),
	),
)
