package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/authfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/dbfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/remoteappfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
)

var AdminHttp = fx.Module(
	"adminHttp",
	gintoolfx.ApiServer,
	adminfx.AdminHandlers,
)

var Admin = fx.Module(
	"appAdmin",
	dbfx.Postgres,
	AdminHttp,
	remoteappfx.RemoteAppDev,
	brieffx.Brief,
	authfx.AdminAuth,
	appfx.App,
	commandfx.Command,
	restyfx.Clients,
)
