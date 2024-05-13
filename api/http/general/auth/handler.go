package auth

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/approle/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	tokenSvc *svc.TokenSvc
}

func NewHandler(tokenSvc *svc.TokenSvc) *Handler {
	return &Handler{tokenSvc: tokenSvc}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/general/v1/token/refresh", h.refreshToken)
}
