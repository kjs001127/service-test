package httpfx

import (
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/lib/dial"
)

const (
	DW          = `name:"dw"`
	InternalApp = `name:"internalApp"`
	ExternalApp = `name:"externalApp"`
)
const (
	dwTimeout  = time.Second * 10
	appTimeout = time.Second * 30
)

var dialer = &net.Dialer{
	Timeout:   5 * time.Second,  // Connection timeout
	KeepAlive: 30 * time.Second, // TCP keepalive
}

var filteringDialer = dial.NewIPFilteringWrapper(dialer)

var internalAppTransport = &http.Transport{
	DialContext:           dialer.DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       40 * time.Second,
	MaxConnsPerHost:       runtime.GOMAXPROCS(0) * 4,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var externalAppTransport = &http.Transport{
	DialContext:           filteringDialer.DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	MaxConnsPerHost:       runtime.GOMAXPROCS(0) * 4,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var dwTransport = &http.Transport{
	DialContext:           dialer.DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          runtime.GOMAXPROCS(0) * 5,
	MaxConnsPerHost:       runtime.GOMAXPROCS(0) * 5,
	IdleConnTimeout:       40 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var Clients = fx.Options(
	fx.Supply(
		fx.Annotate(
			dwTransport,
			fx.ResultTags(DW),
			fx.As(new(http.RoundTripper)),
		),
		fx.Annotate(
			internalAppTransport,
			fx.ResultTags(InternalApp),
			fx.As(new(http.RoundTripper)),
		),
		fx.Annotate(
			externalAppTransport,
			fx.ResultTags(ExternalApp),
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
			fx.ParamTags(DW),
			fx.ResultTags(DW),
		),

		fx.Annotate(
			func(tripper http.RoundTripper) *resty.Client {
				ret := resty.New()
				ret.SetTimeout(appTimeout)
				ret.SetTransport(tripper)
				return ret
			},
			fx.ParamTags(ExternalApp),
			fx.ResultTags(ExternalApp),
		),

		fx.Annotate(
			func(tripper http.RoundTripper) *resty.Client {
				ret := resty.New()
				ret.SetTimeout(appTimeout)
				ret.SetTransport(tripper)
				return ret
			},
			fx.ParamTags(InternalApp),
			fx.ResultTags(InternalApp),
		),
	),
)
