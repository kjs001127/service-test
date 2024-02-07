package fx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/apifx/httpfx"
	"github.com/channel-io/ch-app-store/fx/authfx"
	"github.com/channel-io/ch-app-store/fx/internalfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
	"github.com/channel-io/ch-app-store/lib/dbfx"
)

var Option = fx.Module(
	"app",
	httpfx.Option,
	internalfx.Option,
	dbfx.Option,
	authfx.Option,
	restyfx.Option,
)
