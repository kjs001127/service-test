package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/generalfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/publicfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/ddbfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/apphttpfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/approlefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/authfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/installhookfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/invokelogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/systemlogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
)

var Public = fx.Options(

	PublicHttp,

	authfx.GeneralAuth,
	authfx.PrincipalAuth,

	appfx.App,
	brieffx.Brief,
	commandfx.Command,
	nativefx.Native,
	approlefx.AppRole,
	apphttpfx.Function,
	installhookfx.InstallHooks,

	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	configfx.Values,
	restyfx.Clients,
	datadogfx.Datadog,
	logfx.Logger,
	ddbfx.DynamoDB,
)

var PublicHttp = fx.Options(
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publicfx.PublicHandlers,
	gintoolfx.ApiServer,
)
