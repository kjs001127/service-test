package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"net/http"

	"github.com/channel-io/ch-app-store/api/http/general"
	"github.com/channel-io/ch-app-store/internal/native/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	genauth "github.com/channel-io/ch-app-store/internal/auth/general"
)

// invokeNative godoc
//
//	@Summary	invoke Function
//	@Tags		General
//
//	@Param		x-access-token				header		string						true	"access token"
//	@Param		dto.NativeFunctionRequest	body		dto.NativeFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200							{object}	domain.NativeFunctionResponse
//	@Router		/general/v1/native/functions [put]
func (h *Handler) invokeNative(ctx *gin.Context) {
	var req dto.NativeFunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	rbacToken := middleware.RBAC(ctx)

	resp := h.nativeInvoker.Invoke(ctx, domain.NativeFunctionRequest{
		Token: domain.Token{
			Type:  rbacToken.Token.Header(),
			Value: rbacToken.Token.Value(),
		},
		Method: req.Method,
		Params: req.Params,
	})

	ctx.JSON(http.StatusOK, resp)
}

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		General
//
//	@Param		x-access-token			header		string					true	"access token"
//	@Param		appID					path		string					true	"id of App to invoke Function"
//	@Param		dto.JsonFunctionRequest	body		dto.JsonFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200						{object}	object
//	@Router		/general/v1/apps/{appID}/functions [put]
func (h *Handler) invoke(ctx *gin.Context) {
	appID := ctx.Param("appID")

	var req dto.JsonFunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	rbacToken := middleware.RBAC(ctx)

	if err := authFnCall(rbacToken, appID, req.Context.Channel.ID, req.Method); err != nil {
		_ = ctx.Error(err)
		return
	}

	res := h.invoker.Invoke(
		ctx,
		app.TypedRequest[json.RawMessage]{
			AppID:        appID,
			FunctionName: req.Method,
			Context:      fillCaller(rbacToken, req.Context),
			Params:       req.Params,
		})

	ctx.JSON(http.StatusOK, res)
}

func authFnCall(rbac genauth.ParsedRBACToken, appID string, channelID string, fn string) error {
	if ok := rbac.CheckAction(genauth.Service(appID), genauth.Action(fn)); !ok {
		return errors.New("function call unauthorized")
	}
	if ok := rbac.CheckScope(general.ChannelScope, channelID); !ok {
		return errors.New("function call unauthorized")
	}
	return nil
}

func fillCaller(rbac genauth.ParsedRBACToken, chCtx app.ChannelContext) app.ChannelContext {
	chCtx.Caller.Type = rbac.Caller.Type
	chCtx.Caller.ID = rbac.Caller.ID
	return chCtx
}
