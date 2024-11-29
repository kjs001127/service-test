package svc_test

import (
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/lib/dbfx"
	"github.com/channel-io/ch-app-store/lib/logfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	dbfx.Postgres,
	configfx.Values,
	appfx.App,
	logfx.Logger,
)
