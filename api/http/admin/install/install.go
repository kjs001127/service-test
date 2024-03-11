package install

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/model"
)

// install godoc
//
//	@Summary	install an App to Channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"id of Channel"
//	@Param		appID		path		string	true	"id of App to install"
//
//	@Success	200			{object}	app.InstalledApp
//	@Router		/admin/channels/{channelID}/app-channels/{appID} [put]
func (h *Handler) install(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	appID := ctx.Param("appID")

	appInstalled, appCh, err := h.installer.InstallApp(ctx, app.AppChannelID{
		AppID:     appID,
		ChannelID: channelID,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.InstalledApp{
		App:        appInstalled,
		AppChannel: appCh,
	})
}

// uninstall godoc
//
//	@Summary	uninstall an App to Channel
//	@Tags		Admin
//
//	@Param		channelID	path	string	true	"id of Channel"
//	@Param		appID		path	string	true	"id of App to uninstall"
//
//	@Success	200
//	@Router		/admin/channels/{channelID}/app-channels/{appID} [delete]
func (h *Handler) uninstall(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")
	if err := h.installer.UnInstallApp(ctx, app.AppChannelID{
		AppID:     appID,
		ChannelID: channelID,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// setConfig godoc
//
//	@Summary	set config of a Channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"id of Channel"
//	@Param		appID		path		string	true	"id of App"
//	@Param		object		body		object	true	"key-value of Config to set"
//
//	@Success	200			{object}	app.ConfigMap
//	@Router		/admin/channels/{channelID}/app-channels/{appID}/configs [put]
func (h *Handler) setConfig(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	var configMap map[string]string
	if err := ctx.ShouldBindBodyWith(&configMap, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	ret, err := h.configSvc.SetConfig(ctx, app.AppChannelID{
		AppID:     appID,
		ChannelID: channelID,
	}, configMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

// getConfig godoc
//
//	@Summary	get App config of a AppChannel
//	@Tags		Admin
//
//	@Param		appID		path		string	true	"id of app"
//	@Param		channelID	path		string	true	"id of channel"
//
//	@Success	200			{object}	any		"JSON of configMap"
//	@Router		/admin/channels/{channelID}/app-channels/{appID}/configs [get]
func (h *Handler) getConfig(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	cfgs, err := h.configSvc.GetConfig(ctx, app.AppChannelID{
		ChannelID: channelID,
		AppID:     appID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, cfgs)
}
