package appstore

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	display "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/approle/svc"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo                app.AppRepository
	displayRepo            display.AppDisplayRepository
	cmdRepo                command.CommandRepository
	privateAppQuerySvc     permission.AccountAppPermissionSvc
	privateDisplayQuerySvc permission.AccountDisplayPermissionSvc
	appWithDisplayQuerySvc display.AppWithDisplayQuerySvc
	authSvc                *svc.AppRoleSvc
}

func NewHandler(
	appRepo app.AppRepository,
	displayRepo display.AppDisplayRepository,
	cmdRepo command.CommandRepository,
	privateAppQuerySvc permission.AccountAppPermissionSvc,
	privateDisplayQuerySvc permission.AccountDisplayPermissionSvc,
	appWithDisplayQuerySvc display.AppWithDisplayQuerySvc,
	authSvc *svc.AppRoleSvc,
) *Handler {
	return &Handler{
		appRepo:                appRepo,
		displayRepo:            displayRepo,
		cmdRepo:                cmdRepo,
		privateAppQuerySvc:     privateAppQuerySvc,
		privateDisplayQuerySvc: privateDisplayQuerySvc,
		appWithDisplayQuerySvc: appWithDisplayQuerySvc,
		authSvc:                authSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/app-store")

	group.GET("/apps", h.getApps)
	group.GET("/apps/:appID", h.getAppDetail)
	group.GET("/apps/:appID/roles", h.getAppRoles)
}
