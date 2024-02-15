package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/generalfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/publicfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/dbfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/remoteappfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	"github.com/channel-io/ch-app-store/fx/testfx/mockauthfx"
)

var Public = fx.Module(
	"appAdmin",
	restyfx.Clients,
	dbfx.Postgres,
	PublicHttp,
	remoteappfx.RemoteApp,
	brieffx.Brief,
	mockauthfx.AuthMocked,
	appfx.App,
	commandfx.Command,
)

var PublicHttp = fx.Module(
	"httpPublic",
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publicfx.PublicHandlers,
	gintoolfx.ApiServer,
)
