package httpfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/generalfx"
)

var Option = fx.Module("http",
	generalfx.HttpModule,
	frontfx.HttpModule,
	adminfx.HttpModule,
	deskfx.HttpModule,
)
