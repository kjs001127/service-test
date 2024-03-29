package wysiwygfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/wysiwyg/svc"
)

var WysiwygQuery = fx.Options(
	fx.Provide(
		svc.NewAppCommandQuerySvc,
	),
)
