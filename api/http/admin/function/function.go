package function

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

func (h *Handler) invoke(ctx *gin.Context) {
	appID, name := ctx.Param("appID"), ctx.Param("name")

	var req json.RawMessage
	req, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		_ = ctx.Error(err)
	}

	res, err := h.invoker.Invoke(ctx, app.FunctionRequest[json.RawMessage]{
		AppID:        appID,
		FunctionName: name,
		Body:         req,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
