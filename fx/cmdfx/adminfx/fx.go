package adminfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/authfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/invokelogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/remoteappfx/developmentfx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
)

var AdminHttp = fx.Options(
	gintoolfx.ApiServer,
	adminfx.AdminHandlers,
)

var Admin = fx.Options(

	AdminHttp,

	authfx.RoleClientOnly,

	appfx.App,
	commandfx.Command,
	developmentfx.RemoteAppDevelopment,
	nativefx.Native,
	brieffx.Brief,

	invokelogfx.Loggers,

	restyfx.Clients,
	configfx.Values,
	logfx.Logger,
	datadogfx.Datadog,
)
