package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"errors"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/general"
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
//	@Param		name				path		string				true	"name of Function to invoke"
//	@Param		dto.JsonRPCRequest	body		dto.JsonRPCRequest	true	"body of Function to invoke"
//
//	@Success	200					{object}	json.RawMessage
//	@Router		/general/v1/channels/{channelID}/apps/{appID}/functions/{name} [put]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, name, channelID := ctx.Param("id"), ctx.Param("name"), ctx.Param("channelID")

	var req dto.JsonRPCRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	rawRbac, _ := ctx.Get(middleware.RBACKey)
	rbac := rawRbac.(genauth.ParsedRBACToken)
	if ok := rbac.CheckAction(genauth.Service(appID), genauth.Action(name)); !ok {
		_ = ctx.Error(apierr.Unauthorized(errors.New("function call unauthorized")))
		return
	}
	if ok := rbac.CheckScope(general.ChannelScope, channelID); !ok {
		_ = ctx.Error(apierr.Unauthorized(errors.New("function call unauthorized")))
		return
	}

	res, err := h.invoker.InvokeChannelFunction(
		ctx,
		channelID,
		app.FunctionRequest[json.RawMessage]{
			Endpoint: app.Endpoint{
				AppID:        appID,
				FunctionName: name,
			},
			Body: app.Body[json.RawMessage]{
				Caller: app.Caller{
					Type: rbac.Type,
					ID:   rbac.ID,
				},
				Params: req.Params,
			},
		})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}