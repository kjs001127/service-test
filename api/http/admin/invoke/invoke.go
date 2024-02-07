package invoke

import (
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/auth/appauth"
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
//	@Router		/admin/channels/{channelID}/apps/{appID}/functions/{name} [post]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, name, channelID := ctx.Param("id"), ctx.Param("name"), ctx.Param("channelID")

	var req dto.ParamsAndContext
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
			Scopes:  appauth.Authorities{appauth.Wildcard: {appauth.Wildcard}},
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
