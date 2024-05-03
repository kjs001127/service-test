package adminfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/adminfx"
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
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/permissionfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/systemlogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"

	"go.uber.org/fx"
)

var AdminHttp = fx.Options(
	gintoolfx.ApiServer,
	adminfx.AdminHandlers,
)

var Admin = fx.Options(

	AdminHttp,

	authfx.GeneralAuth,

	accountfx.AppAccount,
	appfx.App,
	permissionfx.Permission,
	commandfx.Command,
	approlefx.AppRole,
	appdevfx.AppDev,
	apphttpfx.Function,
	nativefx.Native,
	brieffx.Brief,
	hookfx.Hook,

	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	ddbfx.DynamoDB,
	restyfx.Clients,
	configfx.Values,
	logfx.Logger,
	datadogfx.Datadog,
)
