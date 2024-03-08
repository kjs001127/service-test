package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/generalfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/publicfx"
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

var Public = fx.Module(
	"appPublic",
	PublicHttp,

	authfx.GeneralAuth,
	authfx.PrincipalAuth,

	appfx.App,
	brieffx.Brief,
	commandfx.Command,
	nativefx.Native,
	remoteappfx.RemoteAppCommon,

	invokelogfx.Loggers,

	configfx.Values,
	restyfx.Clients,
	dbfx.Postgres,
	datadogfx.Datadog,
	logfx.Logger,
)

var PublicHttp = fx.Module(
	"httpPublic",
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publicfx.PublicHandlers,
	gintoolfx.ApiServer,
)
