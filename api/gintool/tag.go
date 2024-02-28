package gintool

import (
	"go.uber.org/fx"
)

func AddTag(handlerConstructor any) any {
	return fx.Annotate(
		handlerConstructor,
		fx.As(new(RouteRegistrant)),
		fx.ResultTags(`group:"routes"`),
	)
}
