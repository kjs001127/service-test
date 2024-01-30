package function

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		appID			path	string			true	"id of App to invoke Function"
//	@Param		name			path	string			true	"name of Function to invoke"
//	@Param		json.RawMessage	body	json.RawMessage	true	"body of Function to invoke"
//
//	@Success	204
//	@Router		/admin/apps/{appID}/functions/{name} [post]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, name := ctx.Param("ID"), ctx.Param("name")

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
