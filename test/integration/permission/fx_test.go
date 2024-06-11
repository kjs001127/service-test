package integration_test

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/permissionfx"
)

var testOpts = fx.Options(
	configfx.Values,
	datadogfx.Datadog,
	permissionfx.Permission,
	appfx.App,
)
