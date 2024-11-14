package store

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
	"github.com/channel-io/ch-app-store/internal/role/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appQuerySvc        app.AppQuerySvc
	cmdRepo            command.CommandRepository
	privateAppQuerySvc permission.AccountAppPermissionSvc
	authSvc            *svc.AppRoleSvc
}

func NewHandler(
	appQuerySvc app.AppQuerySvc,
	cmdRepo command.CommandRepository,
	privateAppQuerySvc permission.AccountAppPermissionSvc,
	authSvc *svc.AppRoleSvc,
) *Handler {
	return &Handler{
		appQuerySvc:        appQuerySvc,
		cmdRepo:            cmdRepo,
		privateAppQuerySvc: privateAppQuerySvc,
		authSvc:            authSvc,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/v1/channels/:channelID/app-store")

	group.GET("/apps", h.getApps)
	group.GET("/apps/:appID", h.getAppDetail)
	group.GET("/apps/:appID/roles", h.getAppRoles)
}
