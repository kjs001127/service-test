package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/apifx/httpfx"
	"github.com/channel-io/ch-app-store/fx/authfx"
	"github.com/channel-io/ch-app-store/fx/dbfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/remoteappfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
)

var Option = fx.Module(
	"app",
	dbfx.Option,
	httpfx.Admin,
	appfx.Option,
	brieffx.Option,
	commandfx.Option,
	restyfx.Option,

	remoteappfx.Option,
	authfx.Option,
)
