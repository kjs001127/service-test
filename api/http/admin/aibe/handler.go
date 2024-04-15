package aibe

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	brief "github.com/channel-io/ch-app-store/internal/brief/svc"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	systemlog "github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	briefInvoker *brief.Invoker
	systemLogSvc *systemlog.SystemLogSvc
	querySvc     *command.AppCommandQuerySvc
}

func NewHandler(
	briefInvoker *brief.Invoker,
	systemLogSvc *systemlog.SystemLogSvc,
	querySvc *command.AppCommandQuerySvc,
) *Handler {
	return &Handler{briefInvoker: briefInvoker, systemLogSvc: systemLogSvc, querySvc: querySvc}
}

// RegisterRoutes @TODO refactor api spec @Camel
func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/admin/brief", h.invokeBrief)
	router.GET("/admin/ai-be/user-chats/:userChatID/logs", h.queryLog)
	router.POST("/admin/ai-be/user-chats/:userChatID/logs", h.writeLog)

	router.GET("/admin/channels/:channelID/apps", h.query)
}
