package media

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/shared/principal/account"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	repo   permission.AppAccountRepo
	parser account.Parser
}

func NewHandler(repo permission.AppAccountRepo, parser account.Parser) *Handler {
	return &Handler{repo: repo, parser: parser}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/media/apps/:appID", h.checkOwner)
}
