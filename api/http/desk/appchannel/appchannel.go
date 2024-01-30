package appchannel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared"
	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

// install godoc
//
//	@Summary	install an App to Channel
//	@Tags		Desk
//
//	@Param		channelId			path		string				true	"id of Channel"
//	@Param		dto.AppIdRequest	body		dto.AppIDRequest	true	"id of App to install"
//
//	@Success	200					{object}	dto.AppAndAppChannel
//	@Router		/desk/channels/{channelId}/app-channels [post]
func (h *Handler) install(ctx *gin.Context) {
	channelID := ctx.Param("channelId")
	var req dto.AppIDRequest
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	appID := req.AppID

	identifier := domain.AppChannelIdentifier{
		AppID:     appID,
		ChannelID: channelID,
	}

	appChannel, err := h.installSaga.Install(ctx, identifier)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.appRepo.Fetch(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.AppAndAppChannel{
		App:        *app,
		AppChannel: *appChannel,
	})
}

// uninstall godoc
//
//	@Summary	uninstall an App to Channel
//	@Tags		Desk
//
//	@Param		channelId			path	string				true	"id of Channel"
//	@Param		dto.AppIdRequest	body	dto.AppIDRequest	true	"id of App to uninstall"
//
//	@Success	200
//	@Router		/desk/channels/{channelId}/app-channels [delete]
func (h *Handler) uninstall(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelId"), ctx.Param("appId")
	identifier := domain.AppChannelIdentifier{
		AppID:     appID,
		ChannelID: channelID,
	}

	if err := h.installSaga.Uninstall(ctx, identifier); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// setConfig godoc
//
//	@Summary	set config of a Channel
//	@Tags		Desk
//
//	@Param		channelId	path		string	true	"id of Channel"
//	@Param		appId		path		string	true	"id of App"
//	@Param		object		body		object	true	"key-value of Config to set"
//
//	@Success	200			{object}	domain.Configs
//	@Router		/desk/channels/{channelId}/app-channels/configs [put]
func (h *Handler) setConfig(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelId"), ctx.Param("appId")
	identifier := domain.AppChannelIdentifier{
		AppID:     appID,
		ChannelID: channelID,
	}

	var configMap map[string]string
	if err := ctx.ShouldBindBodyWith(configMap, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	ret, err := h.installSaga.SetConfig(ctx, identifier, configMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

// getAllWithApp godoc
//
//	@Summary		get App(s) and AppChannel(s)
//	@Tags			Desk
//	@Description	get App and AppChannel installed to channel. If appId is empty, it will return all Apps and AppChannels
//
//	@Param			channelId	path		string	true	"id of Channel"
//	@Param			appId		query		string	false	"id of App"
//
//	@Success		200			{object}	dto.AppAndAppChannel
//	@Router			/desk/channels/{channelId}/app-channels [get]
func (h *Handler) getAllWithApp() func(*gin.Context) {
	return shared.GetAllWithApp(h.appRepo, h.appChannelRepo)
}

// getConfig godoc
//
//	@Summary	get App config of a AppChannel
//	@Tags		Desk
//
//	@Param		appId		path		string	true	"id of app"
//	@Param		channelId	path		string	true	"id of channel"
//
//	@Success	200			{object}	any		"JSON of configMap"
//	@Router		/desk/channels/{channelId}/app-channels/configs [get]
func (h *Handler) getConfig() func(ctx *gin.Context) {
	return shared.GetConfig(h.appChannelConfigSvc)
}
