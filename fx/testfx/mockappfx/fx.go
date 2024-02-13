package mockappfx

import (
	"go.uber.org/fx"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/test/mockapp"
)

var Option = fx.Module(
	"mockapp",
	fx.Supply(
		fx.Annotate(
			new(mockapp.ConfigValidator),
			fx.As(new(app.ConfigValidator)),
		),
		fx.Annotate(
			new(mockapp.InvokeHandler),
			fx.As(new(app.InvokeHandler)),
		),
		fx.Annotate(
			new(mockapp.FileStreamer),
			fx.As(new(app.FileStreamHandler)),
		),
		fx.Annotate(
			new(mockapp.InstallHandler),
			fx.As(new(app.InstallHandler)),
		),
	),

	fx.Invoke(mockapp.SetUpMockApp),
)
