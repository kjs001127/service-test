package publicfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	accountfx2 "github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/accountfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/generalfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/publicfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/ddbfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/accountfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/apphttpfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/approlefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/authfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/hookfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/invokelogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/managerfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/permissionfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/systemlogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"

	"go.uber.org/fx"
)

var Public = fx.Options(

	PublicHttp,

	authfx.GeneralAuth,
	authfx.PrincipalAuth,

	accountfx.AppAccount,
	appfx.App,
	permissionfx.Permission,
	brieffx.Brief,
	commandfx.Command,
	nativefx.Native,
	approlefx.AppRole,
	apphttpfx.Function,
	managerfx.Manager,
	hookfx.Hook,

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
	accountfx2.AccountHandlers,
	gintoolfx.ApiServer,
)
