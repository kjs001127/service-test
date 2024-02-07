package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// getApps godoc
//
//	@Summary	get list of Apps
//	@Tags		Desk
//
//	@Param		since	query	string	false	"get App after this id"
//	@Param		limit	query	string	true	"max count of return data"
//
//	@Success	200		dto.AppsAndCommands
//	@Router		/desk/v1/channels/{channelId}/apps [get]
func (h *Handler) getApps(ctx *gin.Context) {
	since, limit := ctx.Query("since"), ctx.Query("limit")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	apps, err := h.appRepo.Index(ctx, since, limitNumber)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ids := make([]string, 0, len(apps))
	for _, a := range apps {
		ids = append(ids, a.Attributes().ID)
	}

	cmds, err := h.cmdRepo.FetchByQuery(ctx, command.Query{Scope: command.ScopeDesk, AppIDs: ids})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndCommands{Apps: app.AppDatasOf(apps), Commands: dto.NewCommandDTOs(cmds)})
}

// getCommands godoc
//
//	@Summary	get Commands of specific app
//	@Tags		Desk
//
//	@Param		appID	path		string	true	"id of App"
//
//	@Success	200		{array}		dto.CommandDTO
//	@Router		/desk/v1/channels/{channelID}/apps/{appID}/commands [get]
func (h *Handler) getCommands(ctx *gin.Context) {
	appID := ctx.Param("appID")
	commands, err := h.cmdRepo.FetchByQuery(ctx, command.Query{Scope: command.ScopeDesk, AppIDs: []string{appID}})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCommandDTOs(commands))
}
