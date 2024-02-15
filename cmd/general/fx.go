package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/apifx/httpfx"
	"github.com/channel-io/ch-app-store/fx/dbfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
	"github.com/channel-io/ch-app-store/fx/testfx/mockappfx"
	"github.com/channel-io/ch-app-store/fx/testfx/mockauthfx"
)

var Option = fx.Module(
	"app",
	dbfx.Option,
	httpfx.Public,
	appfx.Option,
	brieffx.Option,
	commandfx.Option,
	mockappfx.Option,
	mockauthfx.Option,
	restyfx.Option,
)
