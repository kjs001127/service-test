package fx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/apifx/httpfx"
	"github.com/channel-io/ch-app-store/fx/authfx"
	"github.com/channel-io/ch-app-store/fx/dbfx"
	"github.com/channel-io/ch-app-store/fx/internalfx"
	"github.com/channel-io/ch-app-store/fx/restyfx"
)

var Option = fx.Module(
	"app",
	httpfx.Public,
	internalfx.Option,
	dbfx.Option,
	authfx.Option,
	restyfx.Option,
)
