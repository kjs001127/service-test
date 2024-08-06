package permission_test

import (
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/appdisplayfx"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/permissionfx"
	"github.com/channel-io/ch-app-store/lib/datadogfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	configfx.Values,
	datadogfx.Datadog,
	permissionfx.Permission,
	appfx.App,
	appdisplayfx.AppDisplay,
)
