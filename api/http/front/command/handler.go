package command

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/shared"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	commandQuerySvc *command.QuerySvc
	invoker         *saga.InstallAwareInvokeSaga[any, any]

	appRepo        app.AppRepository
	appChannelRepo appChannel.AppChannelRepository
}

func NewHandler(
	commandQuerySvc *command.QuerySvc,
	invoker *saga.InstallAwareInvokeSaga[any, any],
) *Handler {
	return &Handler{
		commandQuerySvc: commandQuerySvc,
		invoker:         invoker,
	}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	group := router.Group("/front/v6/channels/:channelId")

	group.GET("/commands", h.queryCommands())
	group.PUT("/apps/:appId/commands/:name", h.executeCommand())
	group.PUT("/apps/{appId}/commands/:name/auto-complete", h.autoComplete())
}

// queryCommands godoc
//
//	@Summary	get Commands of Channel
//	@Tags		Front
//
//	@Param		channelId	path		string	true	"id of Channel"
//
//	@Success	200			{object}	dto.AppsAndCommands
//	@Router		/front/v6/channels/{channelId} [get]
func (h *Handler) queryCommands() func(*gin.Context) {
	return shared.QueryCommands(
		h.commandQuerySvc,
		h.appRepo,
		h.appChannelRepo,
		command.ScopeFront,
	)
}

// executeCommand godoc
//
//	@Summary	execute selected Command
//	@Tags		Front
//
//	@Param		channelId	path		string	true	"id of Channel"
//	@Param		appId		path		string	true	"id of App"
//	@Param		name		path		string	true	"name of Command to execute"
//
//	@Success	200			{object}	object
//	@Router		/front/v6/channels/{channelId}/apps/{appId}/commands/{name} [put]
func (h *Handler) executeCommand() func(*gin.Context) {
	return shared.ExecuteRpc(h.invoker)
}

// autoComplete godoc
//
//	@Summary	execute selected AutoComplete of Command
//	@Tags		Front
//
//	@Param		channelId	path		string	true	"id of Channel"
//	@Param		appId		path		string	true	"id of App"
//	@Param		name		path		string	true	"name of Command to execute autoComplete"
//
//	@Success	200			{object}	object
//	@Router		/front/v6/channels/{channelId}/apps/{appId}/commands/{name}/auto-complete [put]
func (h *Handler) autoComplete() func(*gin.Context) {
	return shared.AutoComplete(h.invoker, command.ScopeDesk)
}
