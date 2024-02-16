package publicfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/doc"
	"github.com/channel-io/ch-app-store/api/http/public/controller"
	"github.com/channel-io/ch-app-store/api/http/public/wam"
)

var PublicHandlers = fx.Module(
	"public",
	fx.Provide(
		gintool.AddTag(wam.NewHandler),
		gintool.AddTag(controller.NewHandler),
	),

	fx.Supply(
		fx.Annotate(
			doc.NewHandler("/swagger/public/*any", "swagger_public"),
			fx.As(new(gintool.RouteRegistrant)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)
