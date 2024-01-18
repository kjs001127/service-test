package http

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin"
	"github.com/channel-io/ch-app-store/api/http/desk"
	"github.com/channel-io/ch-app-store/api/http/front"
	"github.com/channel-io/ch-app-store/api/http/general"
)

func Option() fx.Option {
	var constructors []any
	constructors = append(constructors, fx.Annotate(gintool.NewGinRouter, fx.ParamTags(`group:"routes"`)))
	constructors = append(constructors, gintool.NewApiServer)

	constructors = append(constructors, admin.HandlerConstructors...)
	constructors = append(constructors, desk.HandlerConstructors...)
	constructors = append(constructors, front.HandlerConstructors...)
	constructors = append(constructors, general.NewHandler)

	return fx.Provide(constructors...)
}
