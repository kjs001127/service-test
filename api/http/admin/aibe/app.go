package aibe

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	brief "github.com/channel-io/ch-app-store/internal/brief/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

// query godoc
//
//	@Summary	get command, brief, apps of channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"channelID"
//
//	@Success	200			{object}	dto.AppsAndFullCommands
//	@Router		/admin/channels/{channelID}/apps [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	appsInstalled, cmds, err := h.querySvc.Query(ctx, channelID, cmd.ScopeFront)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndFullCommands{
		Apps:     appsInstalled,
		Commands: cmds,
	})
}

// invokeBrief godoc
//
//	@Summaryc	call brief
//	@Tags		Admin

// @Param		dto.BriefRequest	body		dto.BriefRequest	true	"body of Brief"
//
// @Success	200					{object}	brief.BriefResponses
// @Router		/admin/brief  [put]
func (h *Handler) invokeBrief(ctx *gin.Context) {
	var req dto.BriefRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, app.WrapCommonErr(err))
		return
	}

	var ret brief.BriefResponses
	ret, err := h.briefInvoker.Invoke(ctx, req.Context)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

// queryLog godoc
//
//	@Summary	query log
//	@Tags		Admin

// @Param		userChatID	path	string	true	"userChatID"
// @Param		sortOrder	query	string	true	"sortOrder"
// @Param		since		query	string	true	"since"
// @Param		limit		query	int		true	"limit"
//
// @Success	200			{array}	object
// @Router		/admin/ai-be/user-chats/{userChatID}/logs [get]
func (h *Handler) queryLog(ctx *gin.Context) {
	userChatId, sortOrder, since, limit :=
		ctx.Param("userChatID"),
		ctx.Query("sortOrder"),
		ctx.Query("since"),
		ctx.Query("limit")

	logs, err := h.systemLogSvc.QueryLog(ctx, &svc.QueryRequest{
		CursorID: since,
		Limit:    limitFrom(limit),
		Order:    svc.Order(sortOrder),
		ChatType: svc.ChatTypeUserChat,
		ChatId:   userChatId,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, logs)
}

func limitFrom(limitStr string) int32 {
	val, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0
	}
	return int32(val)
}
