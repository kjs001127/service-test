package svc_test

import (
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	"go.uber.org/fx"
)

var testOpts = fx.Options(
	configfx.Values,
	datadogfx.Datadog,
	commandfx.Command,
)
