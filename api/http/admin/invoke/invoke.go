package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	localdto "github.com/channel-io/ch-app-store/api/http/admin/dto"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
)

var callerAdmin = app.Caller{
	Type: "admin",
	ID:   "-",
}

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		appID					path		string					true	"id of App to invoke Function"
//	@Param		name					path		string					true	"name of Function to invoke"
//	@Param		dto.JsonFunctionRequest	body		dto.JsonFunctionRequest	true	"body of Function to invoke"
//
//	@Success	200						{object}	json.RawMessage
//	@Router		/admin/apps/{appID}/functions [put]
func (h *Handler) invoke(ctx *gin.Context) {
	appID := ctx.Param("id")

	var req dto.JsonFunctionRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, app.WrapErr(err))
		return
	}

	res := h.invoker.Invoke(ctx, app.TypedRequest[json.RawMessage]{
		Endpoint: app.Endpoint{
			ChannelID:    req.Context.Channel.ID,
			AppID:        appID,
			FunctionName: req.Method,
		},
		Body: app.Body[json.RawMessage]{
			Context: app.ChannelContext{
				Channel: app.Channel{ID: req.Context.Channel.ID},
				Caller:  callerAdmin,
			},
			Params: req.Params,
		},
	})

	ctx.JSON(http.StatusOK, res)
}

// brief godoc
//
//	@Summaryc	call brief
//	@Tags		Admin

// @Param		dto.BriefRequest	body		dto.BriefRequest	true	"body of Brief"
//
// @Success	200					{object}	brief.BriefResponses
// @Router		/admin/brief  [put]
func (h *Handler) brief(ctx *gin.Context) {
	var req localdto.BriefRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.JSON(http.StatusOK, app.WrapErr(err))
		return
	}

	var ret brief.BriefResponses
	ret, err := h.briefInvoker.Invoke(ctx, req.Context)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
