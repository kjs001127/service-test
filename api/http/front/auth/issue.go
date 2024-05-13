package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/front/middleware"

	_ "github.com/channel-io/ch-app-store/internal/auth/general"
)

// issueToken godoc
//
//	@Summary	issue desk token
//	@Tags		Desk
//
//	@Param		x-account			header		string				true	"access token"
//	@Param		channelID			path		string				true	"id of Channel"
//	@Param		appID				path		string				true	"id of App"
//	@Param		dto.SessionToken	body		dto.SessionToken	true	"account jwt"
//	@Success	200					{object}	general.IssueResponse
//	@Router		/front/v1/channels/{channelID}/apps/{appID}/token [put]
func (h *Handler) issueToken(ctx *gin.Context) {
	appID := ctx.Param("appID")
	user := middleware.User(ctx)

	/*
		var body dto.SessionToken
		if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			_ = ctx.Error(err)
			return
		}


		if body.SessionToken != user.Token.Value() {
			_ = ctx.Error(apierr.Unauthorized(errors.New("token does not match with x-account")))
		}
	*/

	res, err := h.tokenSvc.IssueUserToken(ctx, appID, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
