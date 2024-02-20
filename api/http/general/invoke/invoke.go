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
//	@Param		appID				path		string				true	"id of App to invoke Function"
//	@Param		dto.JsonRPCRequest	body		dto.JsonRPCRequest	true	"body of Function to invoke"
//
//	@Success	200					{object}	app.JsonFunctionResponse
//	@Router		/general/v1/apps/{appID}/functions [put]
func (h *Handler) invoke(ctx *gin.Context) {
	appID := ctx.Param("appID")

	var req dto.JsonRPCRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, app.WrapErr(err))
		return
	}

	chCtx, err := authorizeRbac(ctx, appID, req.Method, req.Context)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, app.WrapErr(err))
	}

	res := h.invoker.Invoke(
		ctx,
		app.TypedRequest[json.RawMessage]{
			Endpoint: app.Endpoint{
				ChannelID:    chCtx.Channel.ID,
				AppID:        appID,
				FunctionName: req.Method,
			},
			Body: app.Body[json.RawMessage]{
				Context: chCtx,
				Params:  req.Params,
			},
		})

	ctx.JSON(http.StatusOK, res)
}

func authorizeRbac(ctx *gin.Context, appID string, fn string, chCtx app.ChannelContext) (app.ChannelContext, error) {
	rawRbac, _ := ctx.Get(middleware.RBACKey)
	rbac := rawRbac.(genauth.ParsedRBACToken)

	if ok := rbac.CheckAction(genauth.Service(appID), genauth.Action(fn)); !ok {
		return app.ChannelContext{}, errors.New("function call unauthorized")
	}
	if ok := rbac.CheckScope(general.ChannelScope, chCtx.Channel.ID); !ok {
		return app.ChannelContext{}, errors.New("function call unauthorized")
	}

	return app.ChannelContext{
		Caller: app.Caller{
			Type: rbac.Caller.Type,
			ID:   rbac.Caller.ID,
		},
		Channel: app.Channel{
			ID: chCtx.Channel.ID,
		},
	}, nil
}
