package svc_test

import (
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/appdisplayfx"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/lib/datadogfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	datadogfx.Datadog,
	configfx.Values,
	appfx.App,
	appdisplayfx.AppDisplay,
)
