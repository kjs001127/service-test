package auth

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/role/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	tokenSvc svc.TokenSvc
}

func NewHandler(tokenSvc svc.TokenSvc) *Handler {
	return &Handler{tokenSvc: tokenSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/desk/v1/channels/:channelID/apps/:appID/token", h.issueToken)
}
