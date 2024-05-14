package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general"
	"github.com/channel-io/ch-app-store/internal/native"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	genauth "github.com/channel-io/ch-app-store/internal/auth/general"
)

// invokeNative godoc
//
//	@Summary	invoke Function
//	@Tags		General
//
//	@Param		x-access-token				header		string						false	"access token"
//	@Param		dto.NativeFunctionRequest	body		dto.NativeFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200							{object}	native.FunctionResponse
//	@Router		/general/v1/native/functions [put]
func (h *Handler) invokeNative(ctx *gin.Context) {
	var req dto.NativeFunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	rbacToken, exists := tokenFrom(ctx)

	resp := h.nativeInvoker.Invoke(ctx,
		native.Token{
			Exists: exists,
			Value:  rbacToken,
		},
		native.FunctionRequest{
			Method: req.Method,
			Params: req.Params,
		},
	)

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

	token, err := h.rbacFrom(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := authFnCall(token, appID, req.Context.Channel.ID, req.Method); err != nil {
		_ = ctx.Error(err)
		return
	}

	res := h.invoker.Invoke(
		ctx,
		appID,
		app.TypedRequest[json.RawMessage]{
			FunctionName: req.Method,
			Context:      fillCaller(token, req.Context),
			Params:       req.Params,
		},
	)

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) rbacFrom(ctx *gin.Context) (genauth.ParsedRBACToken, error) {
	token, exists := tokenFrom(ctx)
	if !exists {
		return genauth.ParsedRBACToken{}, apierr.Unauthorized(errors.New("token not found"))
	}

	parsed, err := h.parser.Parse(ctx, token)
	if err != nil {
		return genauth.ParsedRBACToken{}, apierr.Unauthorized(errors.New("token not valid"))
	}

	return parsed, nil
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
	chCtx.Caller.Type = app.CallerType(rbac.Caller.Type)
	chCtx.Caller.ID = rbac.Caller.ID
	return chCtx
}

const (
	xAccessTokenHeader  = "x-access-token"
	authorizationHeader = "Authorization"
)

func tokenFrom(ctx *gin.Context) (string, bool) {
	xAccessToken := ctx.GetHeader(xAccessTokenHeader)
	if len(xAccessToken) > 0 {
		return xAccessToken, true
	}

	rawAuthHeader := ctx.GetHeader(authorizationHeader)
	if len(rawAuthHeader) > 0 {
		_, token, ok := strings.Cut(rawAuthHeader, " ")
		return token, ok

	}

	return "", false
}
