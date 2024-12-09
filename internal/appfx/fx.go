package appfx

import (
	"encoding/json"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/generated/models"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/repo"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/hook/svc"
	"github.com/channel-io/ch-app-store/lib/sqlrepo"
)

const (
	FunctionListener     = `group:"functionListeners"`
	LifecycleListener    = `group:"lifecycle"`
	InTrxEventListener   = `group:"preInstallHandlers"`
	PostTrxEventListener = `group:"postInstallHandlers"`
)

var App = fx.Options(
	AppSvcs,
	AppDAOs,
)

var AppSvcs = fx.Options(
	fx.Provide(
		fx.Annotate(
			app.NewAppInstallSvc,
			fx.ParamTags(``, ``, ``, InTrxEventListener, PostTrxEventListener),
			fx.As(new(app.AppInstallSvc)),
		),

		app.NewInstallQuerySvc,
		fx.Annotate(
			app.NewAppQuerySvcImpl,
			fx.As(new(app.AppQuerySvc)),
		),
		fx.Annotate(
			app.NewAppLifecycleSvcImpl,
			fx.As(new(app.AppLifecycleSvc)),
			fx.ParamTags(``, ``, ``, ``, LifecycleListener),
		),
		fx.Annotate(
			app.NewInvoker,
			fx.As(new(app.Invoker)),
			fx.ParamTags(``, ``, ``, FunctionListener),
		),
		fx.Annotate(
			app.NewManagerAwareInstallSvc,
			fx.ParamTags(``, InTrxEventListener, PostTrxEventListener),
		),

		app.NewTypedInvoker[json.RawMessage, json.RawMessage],
		app.NewTypedInvoker[svc.ToggleHookRequest, svc.ToggleHookResponse],
	),
)

var AppDAOs = fx.Options(
	fx.Provide(
		fx.Annotate(
			repo.NewAppInstallationDao,
			fx.As(new(app.AppInstallationRepository)),
		),
		fx.Annotate(
			repo.NewAppDAO,
			fx.As(new(app.AppRepository)),
		),
		fx.Annotate(
			repo.NewAppDisplayDAO,
			fx.As(new(app.AppDisplayRepository)),
		),
	),

	fx.Provide(
		fx.Annotate(
			sqlrepo.New[*appmodel.App, *models.App, models.AppSlice],
			fx.As(new(sqlrepo.SQLRepo[*appmodel.App])),
		),
		fx.Annotate(
			sqlrepo.New[*appmodel.AppInstallation, *models.AppInstallation, models.AppInstallationSlice],
			fx.As(new(sqlrepo.SQLRepo[*appmodel.AppInstallation])),
		),

		fx.Annotate(
			sqlrepo.New[*appmodel.AppDisplay, *models.AppDisplay, models.AppDisplaySlice],
			fx.As(new(sqlrepo.SQLRepo[*appmodel.AppDisplay])),
		),
	),

	fx.Supply(
		repo.MarshalInstallation,
		repo.UnmarshalInstallation,
		repo.QueryInstallation,
	),

	fx.Supply(
		repo.MarshalApp,
		repo.UnmarshalApp,
		repo.QueryApp,
	),

	fx.Supply(
		repo.MarshalDisplay,
		repo.UnmarshalDisplay,
		repo.QueryDisplay,
	),
)
