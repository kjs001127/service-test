package configfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
)

const (
	JwtServiceKey      = `name:"jwtServiceKey"`
	DWAdmin            = `name:"dwAdmin"`
	Stage              = `name:"stage"`
	ServiceName        = `name:"serviceName"`
	ChannelServiceName = `name:"channelServiceName"`
)

var Values = fx.Options(
	fx.Supply(
		fx.Annotate(
			config.Get().Auth.AuthAdminURL,
			fx.ResultTags(DWAdmin),
		),
		fx.Annotate(
			config.Get().Auth.JWTServiceKey,
			fx.ResultTags(JwtServiceKey),
		),
		fx.Annotate(
			config.Get().Stage,
			fx.ResultTags(Stage),
		),
		fx.Annotate(
			config.Get().ServiceName,
			fx.ResultTags(ServiceName),
		),
		fx.Annotate(
			config.Get().ChannelServiceName,
			fx.ResultTags(ChannelServiceName),
		),
		config.Get().Log,
		config.Get().DDB,
		config.Get().Psql,
	),
)