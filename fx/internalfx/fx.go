package internalfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/remoteappfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/sagafx"
)

var Option = fx.Options(
	appfx.Option,
	brieffx.Option,
	commandfx.Option,
	remoteappfx.Option,
	sagafx.Option,
)
