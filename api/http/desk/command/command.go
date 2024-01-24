package command

import (
	"net/http"

	"github.com/gin-gonic/gin"

	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

func (h *Handler) getCommands(ctx *gin.Context) {
	appID := ctx.Param("appId")

	query := command.Query{
		AppIDs: []string{appID},
		Scope:  command.ScopeDesk,
	}
	commands, err := h.commandQuerySvc.QueryCommands(ctx, query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, commands)
}
