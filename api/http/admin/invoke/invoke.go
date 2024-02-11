package invoke

import (
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const (
	callerAdmin = "admin"
	idInferred  = "-"
)

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		appID				path		string				true	"id of App to invoke Function"
//	@Param		name				path		string				true	"name of Function to invoke"
//	@Param		dto.JsonRPCRequest	body		dto.JsonRPCRequest	true	"body of Function to invoke"
//
//	@Success	200					{object}	json.RawMessage
//	@Router		/admin/channels/{channelID}/apps/{appID}/functions/{name} [post]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, name, channelID := ctx.Param("id"), ctx.Param("name"), ctx.Param("channelID")

	var req dto.JsonRPCRequest
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: name,
		},
		Body: app.Body{
			Caller: app.Caller{
				Type: callerAdmin,
				ID:   idInferred,
			},
			Params: req.Params,
		},
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
