package appdev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/internal/appdev/svc"
)

// create godoc
//
//	@Summary	create App to app-store
//	@Tags		Admin
//
//	@Param		app.AppRequest	body		svc.AppRequest	true	"App to create"
//
//	@Success	201				{object}	svc.AppResponse
//	@Router		/admin/apps [post]
func (h *Handler) create(ctx *gin.Context) {
	var target svc.AppRequest
	if err := ctx.ShouldBindBodyWith(&target, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.appDevSvc.CreateApp(ctx, target)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// delete godoc
//
//	@Summary	delete App from app-store
//	@Tags		Admin
//
//	@Param		ID	path	string	true	"id of App to delete"
//
//	@Success	204
//	@Router		/admin/apps/{ID} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	ID := ctx.Param("appID")

	err := h.appDevSvc.DeleteApp(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// queryDetail godoc
//
//	@Summary	query App from app-store
//	@Tags		Admin
//
//	@Param		appID	path	string	true "appId"
//
//	@Success	200  	{object} svc.AppResponse
//	@Router		/admin/apps/{appID} [get]
func (h *Handler) queryDetail(ctx *gin.Context) {
	ID := ctx.Param("appID")
	appFound, err := h.appDevSvc.FetchApp(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, appFound)
}
