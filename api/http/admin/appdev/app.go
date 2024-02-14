package appdev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	created, err := h.appDevSvc.CreateApp(ctx, target)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
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
	ID := ctx.Param("ID")

	err := h.appDevSvc.DeleteApp(ctx, ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) query(ctx *gin.Context) {
	ID := ctx.Query("roleId")

	appFound, err := h.appDevSvc.FetchAppByRoleID(ctx, ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, appFound)
}

func (h *Handler) queryDetail(ctx *gin.Context) {
	ID := ctx.Param("appID")
	appFound, err := h.appDevSvc.FetchApp(ctx, ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, appFound)
}
