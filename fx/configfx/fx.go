package configfx

import (
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
)

const (
	JwtServiceKey = `name:"jwtServiceKey"`
	DwGeneral     = `name:"dwGeneral"`
	DwAdmin       = `name:"dwAdmin"`
)

var Values = fx.Module(
	"configValues",
	fx.Supply(
		fx.Annotate(
			config.Get().Auth.AuthAdminURL,
			fx.ResultTags(DwAdmin),
		),
		fx.Annotate(
			config.Get().Auth.JWTServiceKey,
			fx.ResultTags(JwtServiceKey),
		),
		fx.Annotate(
			config.Get().Auth.AuthGeneralURL,
			fx.ResultTags(DwGeneral),
		),
	),
)
