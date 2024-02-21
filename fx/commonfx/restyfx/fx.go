package restyfx

import (
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

const (
	Dw  = `name:"dw"`
	App = `name:"app"`
)
const (
	dwTimeout  = time.Second * 10
	appTimeout = time.Second * 30
)

var Clients = fx.Module(
	"resty",
	fx.Provide(
		fx.Annotate(
			func() *resty.Client {
				ret := resty.New()
				ret.SetTimeout(dwTimeout)
				return ret
			},
			fx.ResultTags(Dw),
		),
	),

	fx.Provide(
		fx.Annotate(
			func() *resty.Client {
				ret := resty.New()
				ret.SetTimeout(appTimeout)
				return ret
			},
			fx.ResultTags(App),
		),
	),
)
