package ratelimiterfx

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/gintoolfx"
	shared "github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/api/httpfx"
	"github.com/channel-io/ch-app-store/configfx"
	"github.com/channel-io/ch-app-store/lib/ratelimiter"

	"go.uber.org/fx"
)

const (
	RateLimiterClient = `name:"rateLimitClient"`
)

var RateLimiter = fx.Options(
	fx.Provide(
		fx.Annotate(
			ratelimiter.NewRateLimiterClient,
			fx.As(new(ratelimiter.Client)),
			fx.ParamTags(httpfx.RateLimiter, configfx.RateLimiterURL),
			fx.ResultTags(RateLimiterClient),
		),
		fx.Annotate(
			shared.NewRateLimiter,
			fx.As(new(gintool.Middleware)),
			fx.ParamTags(RateLimiterClient),
			fx.ResultTags(gintoolfx.MiddlewaresGroup),
		),
		ratelimiter.NewRuleSet,
	),
)
