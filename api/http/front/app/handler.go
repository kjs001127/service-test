package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	cmdInvokeSvc    cmd.InvokeSvc
	autoCompleteSvc cmd.AutoCompleteSvc
	wamDownloader   *app.FileStreamer
}

func NewHandler(
	cmdInvokeSvc cmd.InvokeSvc,
	autoCompleteSvc cmd.AutoCompleteSvc,
	wamDownloader *app.FileStreamer,
) *Handler {
	return &Handler{cmdInvokeSvc: cmdInvokeSvc, autoCompleteSvc: autoCompleteSvc, wamDownloader: wamDownloader}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v6/apps/:appID")
	group.PUT("/commands/:name", h.executeCommand)
	group.PUT("/commands/:name/auto-complete", h.autoComplete)
	group.PUT("/wams/*path", h.downloadWAM)
}
