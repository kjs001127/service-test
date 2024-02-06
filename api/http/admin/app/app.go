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
	var target app.RemoteApp
	if err := ctx.ShouldBindBodyWith(&target, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.appRepo.Save(ctx, &target)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// update godoc
//
//	@Summary	update App info
//	@Tags		Admin
//
//	@Param		id				path		string			true	"id of App to update"
//	@Param		app.RemoteApp	body		app.RemoteApp	true	"App to create"
//
//	@Success	200				{object}	app.RemoteApp
//	@Router		/admin/apps/{id} [patch]
func (h *Handler) update(ctx *gin.Context) {
	ID := ctx.Param("id")
	var target app.RemoteApp
	if err := ctx.ShouldBindBodyWith(&target, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	target.ID = ID

	updated, err := h.appRepo.Update(ctx, &target)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, updated)
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

	err := h.appRepo.Delete(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
