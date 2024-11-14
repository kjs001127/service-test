package httpfnfx

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/api/httpfx"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appfx"
	"github.com/channel-io/ch-app-store/internal/httpfn/infra"
	"github.com/channel-io/ch-app-store/internal/httpfn/model"
	"github.com/channel-io/ch-app-store/internal/httpfn/repo"
	"github.com/channel-io/ch-app-store/internal/httpfn/svc"
)

var Function = fx.Options(
	FunctionSvcs,
	FunctionHttps,
	FunctionDaos,
)

var FunctionSvcs = fx.Options(
	fx.Provide(
		svc.NewAppHttpProxy,
		fx.Annotate(
			svc.NewInvoker,
			fx.As(new(app.InvokeHandler)),
		),
		fx.Annotate(
			svc.NewServerSettingSvcImpl,
			fx.As(new(svc.ServerSettingSvc)),
		),
		fx.Annotate(
			svc.NewLifecycleListener,
			fx.As(new(app.AppLifeCycleEventListener)),
			fx.ResultTags(appfx.LifecycleListener),
		),
	),
)

var FunctionHttps = fx.Options(
	fx.Provide(
		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(svc.HttpRequester)),
			fx.ParamTags(httpfx.InternalApp),
			fx.ResultTags(httpfx.InternalApp),
		),

		fx.Annotate(
			infra.NewHttpRequester,
			fx.As(new(svc.HttpRequester)),
			fx.ParamTags(httpfx.ExternalApp),
			fx.ResultTags(httpfx.ExternalApp),
		),
	),
	fx.Provide(
		fx.Annotate(
			func(internal svc.HttpRequester, external svc.HttpRequester) svc.RequesterMap {
				return svc.RequesterMap{
					model.AccessType_Internal: internal,
					model.AccessType_External: external,
				}
			},
			fx.ParamTags(httpfx.InternalApp, httpfx.ExternalApp),
		),
		fx.Annotate(
			func(internal http.RoundTripper, external http.RoundTripper) svc.RoundTripperMap {
				return svc.RoundTripperMap{
					model.AccessType_Internal: internal,
					model.AccessType_External: external,
				}
			},
			fx.ParamTags(httpfx.InternalApp, httpfx.ExternalApp),
		),
	),
)

var FunctionDaos = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppServerSettingDao,
			fx.As(new(svc.AppServerSettingRepository)),
		),
	),
)
