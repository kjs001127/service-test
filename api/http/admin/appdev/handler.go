package appdev

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/appdev/svc"
	rolesvc "github.com/channel-io/ch-app-store/internal/approle/svc"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appDevSvc   svc.AppDevSvc
	appRoleSvc  *rolesvc.AppRoleSvc
	appManager  app.AppCrudSvc
	registerSvc *command.RegisterSvc
}

func NewHandler(appDevSvc svc.AppDevSvc, appRoleSvc *rolesvc.AppRoleSvc, appManager app.AppCrudSvc, registerSvc *command.RegisterSvc) *Handler {
	return &Handler{appDevSvc: appDevSvc, appRoleSvc: appRoleSvc, appManager: appManager, registerSvc: registerSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/admin/apps")

	group.POST("/", h.create)
	group.GET("/:appID", h.queryDetail)
	group.DELETE("/:appID", h.delete)
	group.POST("/:appID/commands", h.registerCommand)

	group.GET("/", h.queryLegacy)
}
