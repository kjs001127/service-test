package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	uploadDetailImages godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		app.AppRequest	body		dto.AppCreateRequest	true	"App title to create"
//
// @Success	201				{object}	model.App
// @Router		/desk/account/apps/{appId}/detail-images [post]
func (h *Handler) uploadDetailImages(ctx *gin.Context) {
	var accountID string

	file, err := ctx.FormFile("image")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

//	uploadAvatar godoc
//
// @Summary	create App to app-store
// @Tags		Public
//
// @Param		app.AppRequest	body		dto.AppCreateRequest	true	"App title to create"
//
// @Success	201				{object}	model.App
// @Router		/desk/account/apps/{appId}/avatar [post]
func (h *Handler) uploadAvatar(ctx *gin.Context) {
	var accountID string

	file, err := ctx.FormFile("image")
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	}
	ctx.Status(http.StatusOK)
}
