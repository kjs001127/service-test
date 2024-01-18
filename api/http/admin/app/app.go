package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

func (h *Handler) create(ctx *gin.Context) {
	var target app.App
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

func (h *Handler) update(ctx *gin.Context) {
	ID := ctx.Param("id")
	var target app.App
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

func (h *Handler) delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := h.appRepo.Delete(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
