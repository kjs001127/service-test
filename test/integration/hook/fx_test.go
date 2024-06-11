package integration_test

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/hookfx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"
)

var testOpts = fx.Options(
	hookfx.Hook,
	configfx.Values,
	datadogfx.Datadog,
	logfx.Logger,
)
