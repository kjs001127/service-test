package adminfx

import (
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/ddbfx"
	"github.com/channel-io/ch-app-store/fx/corefx/httpfx"
	"github.com/channel-io/ch-app-store/fx/corefx/i18nfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appdisplayfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/apphttpfx"
	privateapprolefx "github.com/channel-io/ch-app-store/fx/corefx/internalfx/approlefx/privatefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/authfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/hookfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/invokelogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/permissionfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/systemlogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"

	"go.uber.org/fx"
)

var AdminHttp = fx.Options(
	gintoolfx.ApiServer,
	adminfx.AdminHandlers,
)

var Admin = fx.Options(

	AdminHttp,

	authfx.GeneralAuth,
	authfx.PrincipalAuth,

	appfx.App,
	appdisplayfx.AppDisplay,
	permissionfx.Permission,
	commandfx.Command,
	privateapprolefx.AppRole,
	apphttpfx.Function,
	nativefx.Native,
	brieffx.Brief,
	hookfx.Hook,

	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	ddbfx.DynamoDB,
	httpfx.Clients,
	configfx.Values,
	logfx.Logger,
	datadogfx.Datadog,
	i18nfx.I18n,
)
