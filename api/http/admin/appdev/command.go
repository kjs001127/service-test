package appdev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	admindto "github.com/channel-io/ch-app-store/api/http/admin/dto"
)

// registerCommand godoc
//
//	@Summary	register Command to App
//	@Tags		Admin
//
//	@Param		appID			path	string				true	"id of App to register Command"
//	@Param		command.Command	body	dto.RegisterRequest	true	"data of Command to register"
//
//	@Success	201
//	@Router		/admin/apps/{appID}/commands [post]
func (h *Handler) registerCommand(ctx *gin.Context) {
	var request admindto.RegisterRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	appID := ctx.Param("appID")

	if err := h.registerSvc.Register(ctx, appID, request.Commands); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusCreated)
}
