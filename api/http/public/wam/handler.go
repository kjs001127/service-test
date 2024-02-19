package wam

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	remoteapp "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamDownloader *remoteapp.FileStreamer
}

func NewHandler(wamDownloader *remoteapp.FileStreamer) *Handler {
	return &Handler{wamDownloader: wamDownloader}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/public/v1/apps/:appID/wams/*path", h.downloadWAM)
}
