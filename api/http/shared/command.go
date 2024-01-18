package shared

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	command "github.com/channel-io/ch-app-store/internal/command/domain"
)

func QueryCommands(svc *command.QuerySvc, scope command.Scope) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		appID := ctx.Param("appId")
		since, limit, query := ctx.Query("since"), ctx.Query("limit"), ctx.Query("query")
		limitNumber, err := strconv.Atoi(limit)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		commands, err := svc.QueryCommands(ctx, command.Query{
			AppID: appID,
			Scope: scope,
			Query: query,
			Since: since,
			Limit: limitNumber,
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, commands)
	}
}
