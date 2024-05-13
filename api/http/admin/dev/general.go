package dev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
)

// readGeneral godoc
//
//	@Summary	fetch App general
//	@Tags		Public
//
//	@Param		appID	path		string	true	"appID"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200		{object}	dto.AppResponse
//	@Router		/desk/account/apps/{appID}/general  [get]
func (h *Handler) readGeneral(ctx *gin.Context) {
	appID := ctx.Param("appID")

	app, err := h.querySvc.Read(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromApp(app))
}

// modifyGeneral godoc
//
//	@Summary	modify App general
//	@Tags		Public
//
//	@Param		appID					path		string					true	"appID"
//	@Param		x-account				header		string					true	"token"
//	@Param		svc.AppModifyRequest	body		dto.AppModifyRequest	true	"dto"
//
//	@Success	200						{object}	dto.AppResponse
//	@Router		/desk/account/apps/{appID}/general  [put]
func (h *Handler) modifyGeneral(ctx *gin.Context) {
	appID := ctx.Param("appID")
	var request dto.AppModifyRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.modifySvc.Update(ctx, request.ConvertToApp(appID))

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}
