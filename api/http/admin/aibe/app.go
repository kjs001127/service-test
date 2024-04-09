package aibe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	brief "github.com/channel-io/ch-app-store/internal/brief/svc"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
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

	appsInstalled, appChs, err := h.querySvc.QueryAll(ctx, channelID)

	cmds, err := h.cmdRepo.FetchByAppIDsAndScope(ctx, app.AppIDsOf(appChs), cmd.ScopeFront)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppsAndFullCommands{
		Apps:     appsInstalled,
		Commands: cmds,
	})
}

// brief godoc
//
//	@Summaryc	call brief
//	@Tags		Admin

// @Param		dto.BriefRequest	body		dto.BriefRequest	true	"body of Brief"
//
// @Success	200					{object}	brief.BriefResponses
// @Router		/admin/brief  [put]
func (h *Handler) brief(ctx *gin.Context) {
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
