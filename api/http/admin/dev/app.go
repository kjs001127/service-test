package dev

import (
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"

	_ "github.com/channel-io/ch-app-store/internal/app/model"
	_ "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	_ "github.com/channel-io/ch-app-store/internal/permission/svc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// createApp godoc
//
//	@Summary	create App to app-store
//	@Tags		Public
//
//	@Param		app.AppCreateRequest	body		dto.AppCreateRequest	true	"App title to create"
//
//	@Success	201						{object}	svc.AppWithDisplay
//	@Router		/admin/apps [post]
func (h *Handler) createApp(ctx *gin.Context) {
	var request dto.AppCreateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.modifySvc.Create(ctx, request.ConvertToApp())
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appWithDisplay, err := h.appWithDisplayQuerySvc.AddDisplayToApp(ctx, created)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, appWithDisplay)
}

// deleteApp godoc
//
//	@Summary	delete App to app-store
//	@Tags		Public
//
//	@Param		appId	path	string	true	"appID"
//
//	@Success	204
//	@Router		/admin/apps/{appID}  [delete]
func (h *Handler) deleteApp(ctx *gin.Context) {
	appID := ctx.Param("appID")

	err := h.modifySvc.Delete(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
