package svc_test

import (
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/hookfx"
	"github.com/channel-io/ch-app-store/lib/datadogfx"
	"github.com/channel-io/ch-app-store/lib/logfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	hookfx.Hook,
	configfx.Values,
	datadogfx.Datadog,
	logfx.Logger,
)
