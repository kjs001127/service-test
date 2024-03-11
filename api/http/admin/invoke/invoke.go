package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	localdto "github.com/channel-io/ch-app-store/api/http/admin/dto"
	"github.com/channel-io/ch-app-store/api/http/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
	"github.com/channel-io/ch-app-store/internal/native/domain"
)

// invokeNative godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		domain.NativeFunctionRequest	body		domain.NativeFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200								{object}	domain.NativeFunctionResponse
//	@Router		/admin/native/functions [put]
func (h *Handler) invokeNative(ctx *gin.Context) {
	var req domain.NativeFunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	resp := h.nativeInvoker.Invoke(ctx, domain.Token{}, domain.NativeFunctionRequest{
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

// brief godoc
//
//	@Summaryc	call brief
//	@Tags		Admin

//	@Param		dto.BriefRequest	body		dto.BriefRequest	true	"body of Brief"
//
//	@Success	200					{object}	brief.BriefResponses
//	@Router		/admin/brief  [put]
func (h *Handler) brief(ctx *gin.Context) {
	var req localdto.BriefRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, app.WrapCommonErr(err))
		return
	}

	var ret brief.BriefResponses
	ret, err := h.briefInvoker.Invoke(ctx, req.Context)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
