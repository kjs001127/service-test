package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/general"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	genauth "github.com/channel-io/ch-app-store/auth/general"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const callerTypeApp = "app"

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

	caller, err := authorizeRbac(ctx, channelID, appID, req.Method)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
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
					Caller:  caller,
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

func authorizeRbac(ctx *gin.Context, channelID, appID, functionName string) (app.Caller, error) {
	rawRbac, _ := ctx.Get(middleware.RBACKey)
	rbac := rawRbac.(genauth.ParsedRBACToken)
	if rbac.Caller.Type == callerTypeApp && rbac.Caller.ID == appID {
		return app.Caller{
			Type: rbac.Caller.Type,
			ID:   rbac.Caller.ID,
		}, nil
	}

	if ok := rbac.CheckAction(genauth.Service(appID), genauth.Action(functionName)); !ok {
		return app.Caller{}, errors.New("function call unauthorized")
	}
	if ok := rbac.CheckScope(general.ChannelScope, channelID); !ok {
		return app.Caller{}, errors.New("function call unauthorized")
	}

	return app.Caller{
		Type: rbac.Caller.Type,
		ID:   rbac.Caller.ID,
	}, nil
}