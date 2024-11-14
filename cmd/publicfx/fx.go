package publicfx

import (
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	"github.com/channel-io/ch-app-store/api/http/accountfx"
	"github.com/channel-io/ch-app-store/api/http/deskfx"
	"github.com/channel-io/ch-app-store/api/http/frontfx"
	"github.com/channel-io/ch-app-store/api/http/generalfx"
	"github.com/channel-io/ch-app-store/api/http/publicfx"
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
	publicapprolefx "github.com/channel-io/ch-app-store/internal/rolefx/publicfx"
	"github.com/channel-io/ch-app-store/internal/sharedfx"
	"github.com/channel-io/ch-app-store/internal/systemlogfx"
	"github.com/channel-io/ch-app-store/internal/widgetfx"
	"github.com/channel-io/ch-app-store/lib/datadogfx"
	"github.com/channel-io/ch-app-store/lib/ddbfx"
	"github.com/channel-io/ch-app-store/lib/i18nfx"
	"github.com/channel-io/ch-app-store/lib/logfx"
	"github.com/channel-io/ch-app-store/lib/ratelimiterfx"

	"go.uber.org/fx"
)

var Public = fx.Options(

	PublicHttp,

	sharedfx.GeneralAuth,
	sharedfx.PrincipalAuth,

	appfx.App,
	permissionfx.Permission,
	brieffx.Brief,
	commandfx.Command,
	nativefx.Native,
	publicapprolefx.AppRole,
	httpfnfx.Function,
	hookfx.Hook,
	widgetfx.AppWidget,
	invokelogfx.Loggers,
	systemlogfx.SystemLog,

	configfx.Values,
	httpfx.Clients,
	datadogfx.Datadog,
	logfx.Logger,
	ddbfx.DynamoDB,
	i18nfx.I18n,
	ratelimiterfx.RateLimiter,
)

var PublicHttp = fx.Options(
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publicfx.PublicHandlers,
	accountfx.AccountHandlers,
	gintoolfx.ApiServer,
)
