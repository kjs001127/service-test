package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

func (h *Handler) registerCommand(ctx *gin.Context) {
	var request command.RegisterRequest
	if err := ctx.ShouldBindBodyWith(request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.appRegisterSvc.Register(ctx, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
