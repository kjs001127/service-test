package front

import (
	"github.com/channel-io/ch-app-store/api/http/front/appchannel"
	"github.com/channel-io/ch-app-store/api/http/front/command"
	"github.com/channel-io/ch-app-store/api/http/front/wam"
)

var HandlerConstructors = []any{
	appchannel.NewHandler,
	command.NewHandler,
	wam.NewHandler,
}
