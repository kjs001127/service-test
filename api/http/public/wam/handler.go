package wam

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	wamDownloader *app.FileStreamer
}

func NewHandler(wamDownloader *app.FileStreamer) *Handler {
	return &Handler{wamDownloader: wamDownloader}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/public/v1/apps/:appID/wams/*path", h.downloadWAM)
}
