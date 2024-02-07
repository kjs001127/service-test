package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/auth/appauth"
	"github.com/channel-io/ch-app-store/auth/principal"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	cmdRepo       cmd.CommandRepository
	invoker       *app.Invoker[json.RawMessage]
	wamDownloader *app.FileStreamer
	authorizer    appauth.AppAuthorizer[principal.Token]
}

func NewHandler(
	cmdRepo cmd.CommandRepository,
	invoker *app.Invoker[json.RawMessage],
	wamDownloader *app.FileStreamer,
	authorizer appauth.AppAuthorizer[principal.Token],
) *Handler {
	return &Handler{cmdRepo: cmdRepo, invoker: invoker, wamDownloader: wamDownloader, authorizer: authorizer}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v1/channels/:channelID/apps/:appID")
	group.PUT("/commands/:name", h.executeCommand)
	group.PUT("/commands/:name/auto-complete", h.autoComplete)
	group.PUT("/wams/*path", h.downloadWAM)
}
