package wam

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/apphttp/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamDownloader *svc.AppHttpProxy
}

func NewHandler(wamDownloader *svc.AppHttpProxy) *Handler {
	return &Handler{wamDownloader: wamDownloader}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/public/v1/apps/:appID/wams/*path", h.downloadWAM)
}
