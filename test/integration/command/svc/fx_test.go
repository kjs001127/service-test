package svc_test

import (
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/internal/commandfx"
	"github.com/channel-io/ch-app-store/lib/dbfx"

	"go.uber.org/fx"
)

var testOpts = fx.Options(
	configfx.Values,
	dbfx.DB,
	commandfx.Command,
)
