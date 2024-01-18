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

func main() {
	fx.New(
		fx.Module("server", internal.Option, http.Option()),
		fx.Invoke(func(srv *gintool.ApiServer) { panic(srv.Run()) }),
	)
}
