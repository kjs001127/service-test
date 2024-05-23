package integration_test

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
)

var testOpts = fx.Options(
	datadogfx.Datadog,
	configfx.Values,
	appfx.App,
)
