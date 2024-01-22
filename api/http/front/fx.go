package front

import (
	"github.com/channel-io/ch-app-store/api/http/front/appchannel"
	"github.com/channel-io/ch-app-store/api/http/front/command"
)

var HandlerConstructors = []any{
	appchannel.NewHandler,
	command.NewHandler,
}
