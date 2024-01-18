package desk

import (
	"github.com/channel-io/ch-app-store/api/http/desk/app"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/command"
	"github.com/channel-io/ch-app-store/api/http/desk/wam"
)

var HandlerConstructors = []any{
	app.NewHandler,
	appchannel.NewHandler,
	command.NewHandler,
	wam.NewHandler,
}
