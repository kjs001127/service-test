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
	"github.com/channel-io/ch-app-store/fx/corefx/httpfx"
	"github.com/channel-io/ch-app-store/fx/corefx/i18nfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appdisplayfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/apphttpfx"
	publicapprolefx "github.com/channel-io/ch-app-store/fx/corefx/internalfx/approlefx/publicfx"
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
)

var PublicHttp = fx.Options(
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publicfx.PublicHandlers,
	accountfx2.AccountHandlers,
	gintoolfx.ApiServer,
)
