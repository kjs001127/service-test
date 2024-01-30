package app

import (
	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	appRepo app.AppRepository
	cmdRepo command.CommandRepository

	wamDownloader       *app.FileStreamer
	cmdInvoker          *command.InvokeSvc
	autoCompleteInvoker *command.AutoCompleteSvc
}

func NewHandler(
	appRepo app.AppRepository,
	cmdRepo command.CommandRepository,
	wamDownloader *app.FileStreamer,
	cmdInvoker *command.InvokeSvc,
	autoCompleteInvoker *command.AutoCompleteSvc,
) *Handler {
	return &Handler{appRepo: appRepo, cmdRepo: cmdRepo, wamDownloader: wamDownloader, cmdInvoker: cmdInvoker, autoCompleteInvoker: autoCompleteInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/desk/apps")

	group.GET("/", h.getApps)
	group.GET("/:appID/commands", h.getCommands)
	group.PUT("/:appID/commands/:name", h.executeCommand)
	group.PUT("/:appID/commands/:name/auto-complete", h.autoComplete)
	group.GET("/:appID/wams/*path", h.downloadWAM)
}
