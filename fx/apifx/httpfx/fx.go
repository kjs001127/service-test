package httpfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/apifx/httpfx/generalfx"
)

const port = `name:"port"`

var Public = fx.Module(
	"httpPublic",
	generalfx.HttpModule,
	frontfx.HttpModule,
	deskfx.HttpModule,
	gintoolfx.Option,
)

var Admin = fx.Module(
	"httpAdmin",
	adminfx.HttpModule,
	gintoolfx.Option,
)
