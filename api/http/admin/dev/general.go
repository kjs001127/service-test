package dev

import (
	"net/http"

	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// readGeneral godoc
//
//	@Summary	fetch App general
//	@Tags		Public
//
//	@Param		appID	path		string	true	"appID"
//
//	@Success	200		{object}	displaysvc.AppWithDisplay
//	@Router		/admin/apps/{appID}/general  [get]
func (h *Handler) readGeneral(ctx *gin.Context) {
	appID := ctx.Param("appID")

	app, err := h.querySvc.Read(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appWithDisplay, err := h.appWithDisplayQuerySvc.AddDisplayToApp(ctx, app)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, appWithDisplay)
}

// modifyGeneral godoc
//
//	@Summary	modify App general
//	@Tags		Public
//
//	@Param		appID						path		string						true	"appID"
//	@Param		displaysvc.AppWithDisplay	body		displaysvc.AppWithDisplay	true	"dto"
//
//	@Success	200							{object}	displaysvc.AppWithDisplay
//	@Router		/admin/apps/{appID}/general  [put]
func (h *Handler) modifyGeneral(ctx *gin.Context) {
	appID := ctx.Param("appID")
	var request displaysvc.AppWithDisplay
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	request.ID = appID

	app, err := h.modifySvc.Update(ctx, request.ConvertToApp())
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	display, err := h.displayModifySvc.Update(ctx, request.ConvertToDisplay())
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appWithDisplay := displaysvc.MergeAppAndDisplay(app, display)

	ctx.JSON(http.StatusOK, appWithDisplay)
}
