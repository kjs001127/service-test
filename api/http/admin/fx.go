package admin

import (
	"github.com/channel-io/ch-app-store/api/http/admin/app"
	"github.com/channel-io/ch-app-store/api/http/admin/register"
)

var HandlerConstructors = []any{
	app.NewHandler,
	register.NewHandler,
}
