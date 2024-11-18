package dao_test

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/rolefx/privatefx"
	"github.com/channel-io/ch-app-store/lib/datadogfx"
)

var testOpts = fx.Options(
	datadogfx.Datadog,
	configfx.Values,
	privatefx.AppRole,
)
