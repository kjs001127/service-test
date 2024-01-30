package main

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http"
	"github.com/channel-io/ch-app-store/internal"
)

func init() {
	// Config Set
}

// HttpModule				   godoc
//
//	@Title		ch-app-store API
//	@Version	1.0
//	@BasePath	/
func main() {
	fx.New(
		internal.Option,
		fx.Provide(
			fx.Annotate(gintool.NewGinEngine, fx.ParamTags(`group:"routes"`)),
			gintool.NewApiServer,
		),
		http.Option,
		fx.Invoke(func(srv *gintool.ApiServer) { panic(srv.Run()) }),
	)
}
