package desk

import (
	"github.com/channel-io/ch-app-store/api/http/desk/app"
	"github.com/channel-io/ch-app-store/api/http/desk/appchannel"
	"github.com/channel-io/ch-app-store/api/http/desk/command"
)

var HandlerConstructors = []any{
	app.NewHandler,
	appchannel.NewHandler,
	command.NewHandler,
}
