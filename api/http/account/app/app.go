package app

import (
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/internal/permission/svc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//	createApp godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		app.AppRequest	body		dto.AppCreateRequest	true	"App title to create"
//
// @Success	201				{object}	model.App
// @Router		/desk/account/apps [post]
func (h *Handler) createApp(ctx *gin.Context) {
	var accountID string
	var request dto.AppCreateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.appPermissionSvc.CreateApp(ctx, request.Title, accountID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

func (h *Handler) modifyApp(ctx *gin.Context) {

	var accountID string
	appID := ctx.Param("appId")
	var request svc.AppModifyRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.appPermissionSvc.ModifyApp(ctx, request, appID, accountID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *Handler) deleteApp(ctx *gin.Context) {
	var accountID string
	appID := ctx.Param("appID")

	err := h.appPermissionSvc.DeleteApp(ctx, accountID, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
