package restyfx

import (
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

const timeout = time.Second * 10

var Option = fx.Module(
	"resty",
	fx.Provide(func() *resty.Client {
		ret := resty.New()
		ret.SetTimeout(timeout)
		return ret
	}),
)
