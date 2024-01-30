package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

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
//	@Router		/admin/app-store/v1/apps/{id}/commands [post]
func (h *Handler) registerCommand(ctx *gin.Context) {
	var request command.RegisterRequest
	if err := ctx.ShouldBindBodyWith(request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.registerSaga.Register(ctx, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
