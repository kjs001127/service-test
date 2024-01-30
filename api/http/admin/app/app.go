package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/remoteapp/domain"
)

// create godoc
//
//	@Summary	create App to app-store
//	@Tags		Admin
//
//	@Param		app.App	body		app.App	true	"App to create"
//
//	@Success	201		{object}	app.App
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
//	@Param		id		path		string	true	"id of App to update"
//	@Param		app.App	body		app.App	true	"App to create"
//
//	@Success	200		{object}	app.App
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

func (h *Handler) query(ctx *gin.Context) {
	appID := ctx.Param("appID")

	found, err := h.appRepo.Fetch(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	brief, err := h.briefRepo.Fetch(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.commandRepo.FetchAllByAppID(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppResource{
		App:      found.Data(),
		Commands: cmds,
		Brief:    brief,
	})
}
