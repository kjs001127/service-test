package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"github.com/channel-io/ch-app-store/api/http/general"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	genauth "github.com/channel-io/ch-app-store/auth/general"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const tokenHeader = "x-access-token"

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		General
//
//	@Param		x-access-token		header		string				true	"access token"
//	@Param		channelID			path		string				true	"id of Channel to invoke Function"
//	@Param		appID				path		string				true	"id of App to invoke Function"
//	@Param		dto.JsonRPCRequest	body		dto.JsonRPCRequest	true	"body of Function to invoke"
//
//	@Success	200					{object}	json.RawMessage
//	@Router		/general/v1/channels/{channelID}/apps/{appID}/functions [put]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, channelID := ctx.Param("id"), ctx.Param("channelID")

	var req dto.JsonRPCRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	rawRbac, _ := ctx.Get(middleware.RBACKey)
	rbac := rawRbac.(genauth.ParsedRBACToken)
	if ok := rbac.CheckAction(genauth.Service(appID), genauth.Action(req.Method)); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("function call unauthorized"))
		return
	}
	if ok := rbac.CheckScope(general.ChannelScope, channelID); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("function call unauthorized"))
		return
	}

	res, err := h.invoker.InvokeChannelFunction(
		ctx,
		app.FunctionRequest[json.RawMessage]{
			Endpoint: app.Endpoint{
				ChannelID:    channelID,
				AppID:        appID,
				FunctionName: req.Method,
			},
			Body: app.Body[json.RawMessage]{
				Context: app.ChannelContext{
					Channel: app.Channel{ID: channelID},
					Caller: app.Caller{
						Type: rbac.Type,
						ID:   rbac.ID,
					},
				},
				Params: req.Params,
			},
		})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}
