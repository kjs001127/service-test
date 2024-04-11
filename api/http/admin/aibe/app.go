package aibe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	brief "github.com/channel-io/ch-app-store/internal/brief/svc"
	"github.com/channel-io/ch-app-store/internal/systemlog/svc"
)

// invokeBrief godoc
//
//	@Summaryc	call brief
//	@Tags		Admin

//	@Param		dto.BriefRequest	body		dto.BriefRequest	true	"body of Brief"
//
//	@Success	200					{object}	brief.BriefResponses
//	@Router		/admin/brief  [put]
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

//	@Param		svc.QueryRequest	body	svc.QueryRequest	true	"body of Brief"
//
//	@Success	200					{array}	model.SystemLog
//	@Router		/admin/logs [post]
func (h *Handler) queryLog(ctx *gin.Context) {
	var req svc.QueryRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, app.WrapCommonErr(err))
		return
	}

	logs, err := h.systemLogSvc.QueryLog(ctx, &req)
	if err != nil {
		_ = ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, logs)
}
