package appstore

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/svc"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo            app.AppRepository
	cmdRepo            command.CommandRepository
	privateAppQuerySvc permission.AccountAppPermissionSvc
	authSvc            *svc.AppRoleSvc
}

func NewHandler(
	appRepo app.AppRepository,
	cmdRepo command.CommandRepository,
	privateAppQuerySvc permission.AccountAppPermissionSvc,
	authSvc *svc.AppRoleSvc,
) *Handler {
	return &Handler{appRepo: appRepo, cmdRepo: cmdRepo, privateAppQuerySvc: privateAppQuerySvc, authSvc: authSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/app-store")

	group.GET("/apps", h.getApps)
	group.GET("/apps/:appID", h.getAppDetail)
	group.GET("/apps/:appID/claims", h.getAppRoles)
}
