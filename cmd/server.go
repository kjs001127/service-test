package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/admin"
	"github.com/channel-io/ch-app-store/api/http/desk"
	"github.com/channel-io/ch-app-store/api/http/front"
	"github.com/channel-io/ch-app-store/api/http/general"
	"github.com/channel-io/ch-app-store/internal"
)

func init() {
	// Config Set
}

func main() {
	fx.New(
		internal.Option,
		fx.Provide(
			fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`)),
			gintool.NewApiServer,
		),
		admin.HttpModule(),
		desk.HttpModule(),
		front.HttpModule(),
		general.HttpModule(),
		fx.Invoke(func(srv *gintool.ApiServer) { panic(srv.Run()) }),
	)
}
