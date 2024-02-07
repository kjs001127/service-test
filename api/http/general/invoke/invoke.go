package invoke

import (
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general/dto"
	"github.com/channel-io/ch-app-store/auth/appauth"
	"github.com/channel-io/ch-app-store/auth/general"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const tokenHeader = "x-access-token"

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
//	@Router		/general/v1/channels/{channelID}/apps/{appID}/functions/{name} [post]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, name, channelID := ctx.Param("id"), ctx.Param("name"), ctx.Param("channelID")

	var req dto.JsonRPCRequest
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	scopes, err := h.authorizer.Handle(ctx, appauth.AppUseRequest[general.Token]{
		Token: general.Token(ctx.GetHeader(general.Header())),
		AppID: appID,
		ChCtx: req.Context,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: name,
		},
		Body: app.Body{
			Scopes:  scopes,
			Context: req.Context,
			Params:  req.Params,
		},
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
