package adminfx

import (
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/adminfx"
	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/brieffx"
	"github.com/channel-io/ch-app-store/internal/commandfx"
	"github.com/channel-io/ch-app-store/internal/hookfx"
	"github.com/channel-io/ch-app-store/internal/httpfnfx"
	"github.com/channel-io/ch-app-store/internal/invokelogfx"
	"github.com/channel-io/ch-app-store/internal/nativefx"
	"github.com/channel-io/ch-app-store/internal/permissionfx"
	privateapprolefx "github.com/channel-io/ch-app-store/internal/rolefx/privatefx"
	"github.com/channel-io/ch-app-store/internal/sharedfx"
	"github.com/channel-io/ch-app-store/internal/systemlogfx"
	"github.com/channel-io/ch-app-store/internal/widgetfx"
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

	sharedfx.GeneralAuth,
	sharedfx.PrincipalAuth,

	appfx.App,
	permissionfx.Permission,
	commandfx.Command,
	privateapprolefx.AppRole,
	httpfnfx.Function,
	nativefx.Native,
	brieffx.Brief,
	hookfx.Hook,
	widgetfx.AppWidget,

	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	ddbfx.DynamoDB,
	httpfx.Clients,
	configfx.Values,
	logfx.Logger,
	datadogfx.Datadog,
	i18nfx.I18n,
)
