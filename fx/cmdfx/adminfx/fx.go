package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/configfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/dbfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/authfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/invokelogfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/remoteappfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/logfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
)

var AdminHttp = fx.Module(
	"adminHttp",
	gintoolfx.ApiServer,
	adminfx.AdminHandlers,
)

var Admin = fx.Module(
	"appAdmin",
	AdminHttp,

	authfx.RoleClientOnly,

	appfx.App,
	commandfx.Command,
	remoteappfx.RemoteAppDev,
	nativefx.Native,
	brieffx.Brief,

	invokelogfx.Loggers,

	restyfx.Clients,
	configfx.Values,
	dbfx.Postgres,
	logfx.Logger,
	datadogfx.Datadog,
)
