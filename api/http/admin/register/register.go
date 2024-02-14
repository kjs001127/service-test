package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

// registerCommand godoc
//
//	@Summary	register Command to App
//	@Tags		Admin
//
//	@Param		id						path	string					true	"id of App to register Command"
//	@Param		command.RegisterRequest	body	command.RegisterRequest	true	"data of Command to register"
//
//	@Success	204
//	@Router		/admin/apps/{id}/commands [post]
func (h *Handler) registerCommand(ctx *gin.Context) {
	var request command.RegisterRequest
	if err := ctx.ShouldBindBodyWith(request, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	if err := h.registerSaga.Register(ctx, request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
