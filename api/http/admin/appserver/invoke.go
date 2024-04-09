package appserver

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/native"
)

// invokeNative godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		handler.NativeFunctionRequest	body		handler.NativeFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200								{object}	handler.NativeFunctionResponse
//	@Router		/admin/native/functions [put]
func (h *Handler) invokeNative(ctx *gin.Context) {
	var req native.FunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	token := ctx.GetHeader("x-access-token")
	resp := h.nativeInvoker.Invoke(ctx, native.Token{Type: "x-access-token", Value: token}, native.FunctionRequest{
		Method: req.Method,
		Params: req.Params,
	})

	ctx.JSON(http.StatusOK, resp)
}

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		appID					path		string					true	"id of App to invoke Function"
//	@Param		dto.JsonFunctionRequest	body		dto.JsonFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200						{object}	json.RawMessage
//	@Router		/admin/apps/{appID}/functions [put]
func (h *Handler) invoke(ctx *gin.Context) {
	appID := ctx.Param("id")

	var req dto.JsonFunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, app.WrapCommonErr(err))
		return
	}

	res := h.invoker.Invoke(ctx, appID, app.TypedRequest[json.RawMessage]{
		FunctionName: req.Method,
		Context:      req.Context,
		Params:       req.Params,
	})

	ctx.JSON(http.StatusOK, res)
}
