package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	app "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

// create godoc
//
//	@Summary	create App to app-store
//	@Tags		Admin
//
//	@Param		app.RemoteApp	body		app.RemoteApp	true	"App to create"
//
//	@Success	201				{object}	app.RemoteApp
//	@Router		/admin/apps [post]
func (h *Handler) create(ctx *gin.Context) {
	var target app.AppRequest
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
//	@Param		id	path	string	true	"id of App to delete"
//
//	@Success	204
//	@Router		/admin/apps/{id} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := h.appDevSvc.DeleteApp(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
