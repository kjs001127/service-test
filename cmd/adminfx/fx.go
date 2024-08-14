package adminfx

import (
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/adminfx"
	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/appdisplayfx"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/apphttpfx"
	privateapprolefx "github.com/channel-io/ch-app-store/internal/approlefx/privatefx"
	"github.com/channel-io/ch-app-store/internal/appwidgetfx"
	"github.com/channel-io/ch-app-store/internal/authfx"
	"github.com/channel-io/ch-app-store/internal/brieffx"
	"github.com/channel-io/ch-app-store/internal/commandfx"
	"github.com/channel-io/ch-app-store/internal/hookfx"
	"github.com/channel-io/ch-app-store/internal/invokelogfx"
	"github.com/channel-io/ch-app-store/internal/nativefx"
	"github.com/channel-io/ch-app-store/internal/permissionfx"
	"github.com/channel-io/ch-app-store/internal/systemlogfx"
	"github.com/channel-io/ch-app-store/lib/datadogfx"
	"github.com/channel-io/ch-app-store/lib/ddbfx"
	"github.com/channel-io/ch-app-store/lib/i18nfx"
	"github.com/channel-io/ch-app-store/lib/logfx"

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
	appwidgetfx.AppWidget,

	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	ddbfx.DynamoDB,
	httpfx.Clients,
	configfx.Values,
	logfx.Logger,
	datadogfx.Datadog,
	i18nfx.I18n,
)
