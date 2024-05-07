package account

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	repo permission.AppAccountRepo
}

func NewHandler(repo permission.AppAccountRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/admin/accounts/:accountID/apps/:appID", h.checkOwner)
}
