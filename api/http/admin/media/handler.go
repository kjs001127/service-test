package media

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	repo   permission.AppAccountRepo
	parser desk.Parser
}

func NewHandler(repo permission.AppAccountRepo, parser desk.Parser) *Handler {
	return &Handler{repo: repo, parser: parser}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/media/apps/:appID", h.checkOwner)
}
