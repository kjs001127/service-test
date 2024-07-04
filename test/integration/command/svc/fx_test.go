package svc_test

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
)

var testOpts = fx.Options(
	configfx.Values,
	datadogfx.Datadog,
	commandfx.Command,
)
