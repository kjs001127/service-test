package internal

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/internal/app"
	"github.com/channel-io/ch-app-store/internal/brief"
	"github.com/channel-io/ch-app-store/internal/command"
	"github.com/channel-io/ch-app-store/internal/remoteapp"
	"github.com/channel-io/ch-app-store/internal/saga"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

var Option = fx.Options(
	tx.FxOption,
	app.Option,
	brief.Option,
	command.Option,
	remoteapp.Option,
	saga.Option,
)
