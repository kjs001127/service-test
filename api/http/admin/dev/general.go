package dev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	appsvc "github.com/channel-io/ch-app-store/internal/app/svc"
)

// readGeneral godoc
//
//	@Summary	fetch App general
//	@Tags		Public
//
//	@Param		appID	path		string	true	"appID"
//
//	@Success	200		{object}	svc.AppDetail
//	@Router		/admin/apps/{appID}/general  [get]
func (h *Handler) readGeneral(ctx *gin.Context) {
	appID := ctx.Param("appID")

	app, err := h.querySvc.ReadDetail(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

// modifyGeneral godoc
//
//	@Summary	modify App general
//	@Tags		Public
//
//	@Param		appID				path		string				true	"appID"
//	@Param		appsvc.AppDetail	body		appsvc.AppDetail	true	"dto"
//
//	@Success	200					{object}	svc.AppDetail
//	@Router		/admin/apps/{appID}/general  [put]
func (h *Handler) modifyGeneral(ctx *gin.Context) {
	appID := ctx.Param("appID")
	var request *appsvc.AppDetail
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	request.ID = appID

	app, err := h.modifySvc.UpdateDetail(ctx, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}
