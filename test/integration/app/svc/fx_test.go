package integration_test

import (
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appdisplayfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	datadogfx.Datadog,
	configfx.Values,
	appfx.App,
	appdisplayfx.AppDisplay,
)
