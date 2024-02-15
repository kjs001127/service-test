package restyfx

import (
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

const (
	dwTimeout  = time.Second * 10
	appTimeout = time.Second * 10
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
			fx.ResultTags(`name:"dw"`),
		),
	),

	fx.Provide(
		fx.Annotate(
			func() *resty.Client {
				ret := resty.New()
				ret.SetTimeout(appTimeout)
				return ret
			},
			fx.ResultTags(`name:"app"`),
		),
	),
)
