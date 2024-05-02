package testfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/generalfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/publicfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/ddbfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/accountfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appdevfx"
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

var Test = fx.Options(
	Http,

	authfx.GeneralAuth,
	authfx.PrincipalAuth,

	accountfx.AppAccount,
	appfx.App,
	permissionfx.Permission,
	brieffx.Brief,
	appdevfx.AppDev,
	commandfx.Command,
	nativefx.Native,
	approlefx.AppRole,
	apphttpfx.Function,
	hookfx.Hook,
	managerfx.Manager,

	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	configfx.Values,
	restyfx.Clients,
	datadogfx.Datadog,
	logfx.Logger,
	ddbfx.DynamoDB,
)

var Http = fx.Options(
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	adminfx.AdminHandlers,
	publicfx.PublicHandlers,
	gintoolfx.ApiServer,
)
