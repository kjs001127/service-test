package admin

import (
	"github.com/channel-io/ch-app-store/api/http/admin/app"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
	"github.com/channel-io/ch-app-store/api/http/admin/wam"
)

var HandlerConstructors = []any{
	app.NewHandler,
	register.NewHandler,
	wam.NewHandler,
}
