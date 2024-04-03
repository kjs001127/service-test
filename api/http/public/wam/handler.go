package wam

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/internal/remoteapp/interaction/svc"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamDownloader *svc.FileStreamer
}

func NewHandler(wamDownloader *svc.FileStreamer) *Handler {
	return &Handler{wamDownloader: wamDownloader}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/public/v1/apps/:appID/wams/*path", h.downloadWAM)
}
