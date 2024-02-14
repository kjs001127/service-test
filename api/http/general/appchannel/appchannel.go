package appchannel

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/general"
	"github.com/channel-io/ch-app-store/api/http/general/middleware"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	authgen "github.com/channel-io/ch-app-store/auth/general"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const (
	fetchConfig = authgen.Action("fetchConfig")
)

// getConfig godoc
//
//	@Summary	get App config of a Channel
//	@Tags		General
//
//	@Param		x-access-token	header		string	true	"access token"
//	@Param		appID			path		string	true	"id of app"
//	@Param		channelID		path		string	true	"id of channel"
//
//	@Success	200				{object}	any		"JSON of configMap"
//	@Router		/general/v1/channels/{channelID}/app-channels/{appID}/configs [get]
func (h *Handler) getConfig(ctx *gin.Context) {

	appID, channelID := ctx.Param("appID"), ctx.Param("channelID")

	rawRbac, _ := ctx.Get(middleware.RBACKey)
	rbac := rawRbac.(authgen.ParsedRBACToken)
	if ok := rbac.CheckAction(general.AppStoreService, fetchConfig); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("function call unauthorized"))
		return
	}

	if ok := rbac.CheckScopes(authgen.Scopes{
		general.ChannelScope: {channelID},
		general.AppScope:     {appID},
	}); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("function call unauthorized"))
		return
	}

	installedApp, err := h.querySvc.Query(ctx, app.Install{AppID: appID, ChannelID: channelID})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, installedApp.AppChannel.Configs)
}
