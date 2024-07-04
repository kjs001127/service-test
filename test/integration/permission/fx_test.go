package permission_test

import (
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appdisplayfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/permissionfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	configfx.Values,
	datadogfx.Datadog,
	permissionfx.Permission,
	appfx.App,
	appdisplayfx.AppDisplay,
)
