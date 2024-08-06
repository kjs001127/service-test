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
	"github.com/channel-io/ch-app-store/internal/appdisplayfx"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/apphttpfx"
	publicapprolefx "github.com/channel-io/ch-app-store/internal/approlefx/publicfx"
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
	"github.com/channel-io/ch-app-store/lib/ratelimiterfx"

	"go.uber.org/fx"
)

var Public = fx.Options(

	PublicHttp,

	authfx.GeneralAuth,
	authfx.PrincipalAuth,

	appfx.App,
	appdisplayfx.AppDisplay,
	permissionfx.Permission,
	brieffx.Brief,
	commandfx.Command,
	nativefx.Native,
	publicapprolefx.AppRole,
	apphttpfx.Function,
	hookfx.Hook,

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
