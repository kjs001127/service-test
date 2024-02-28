package restyfx

import (
	"net/http"
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
	fx.Supply(
		fx.Annotate(
			http.DefaultTransport,
			fx.ResultTags(Dw),
			fx.As(new(http.RoundTripper)),
		),
		fx.Annotate(
			http.DefaultTransport,
			fx.ResultTags(App),
			fx.As(new(http.RoundTripper)),
		),
	),
	fx.Provide(
		fx.Annotate(
			func(tripper http.RoundTripper) *resty.Client {
				ret := resty.New()
				ret.SetTimeout(dwTimeout)
				ret.SetTransport(tripper)
				return ret
			},
			fx.ParamTags(Dw),
			fx.ResultTags(Dw),
		),
	),

	fx.Provide(
		fx.Annotate(
			func(tripper http.RoundTripper) *resty.Client {
				ret := resty.New()
				ret.SetTimeout(appTimeout)
				ret.SetTransport(tripper)
				return ret
			},
			fx.ParamTags(App),
			fx.ResultTags(App),
		),
	),
)
