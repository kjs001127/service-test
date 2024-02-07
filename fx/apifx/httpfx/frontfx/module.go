package frontfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	_ "github.com/channel-io/ch-app-store/api/http/front/swagger"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/front/app"
	"github.com/channel-io/ch-app-store/api/http/front/invoke"
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/config"
)

const frontPort = `name:"front.port"`

var HttpModule = fx.Module(
	"frontHttpModule",
	fx.Supply(
		fx.Annotate(
			config.Get().Port.Front,
			fx.ResultTags(frontPort),
		),
	),
	fx.Provide(
		fx.Annotate(
			gintool.NewGinEngine,
			fx.ParamTags(`group:"routes"`, `group:"front.auth"`),
			fx.ResultTags(`name:"front.engine"`),
		),
		fx.Annotate(gintool.NewApiServer, fx.ParamTags(`name:"front.engine"`, frontPort), fx.ResultTags(`name:"front.server"`)),

		gintool.AddTag(invoke.NewHandler),
		gintool.AddTag(app.NewHandler),

		fx.Annotate(
			middleware.NewAuth,
			fx.As(new(gintool.Middleware)),
			fx.ResultTags(`group:"front.auth"`),
		),

		fx.Private,
	),
	fx.Invoke(func(server Server) {
		go func() {
			panic(server.Srv.Run())
		}()
	}),
)

type Server struct {
	fx.In
	Srv    *gintool.ApiServer `name:"front.server"`
	Engine *gin.Engine        `name:"front.engine"`
}
