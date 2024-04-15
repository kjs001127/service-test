package install

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
)

// install godoc
//
//	@Summary	install an App to Channel
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"id of Channel"
//	@Param		appID		path		string	true	"id of App to install"
//
//	@Success	200			{object}	dto.InstalledApp
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID} [put]
func (h *Handler) install(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	appID := ctx.Param("appID")

	appFound, appInstallation, err := h.installer.InstallAppById(ctx, appmodel.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.InstalledApp{
		App:             appFound,
		AppInstallation: appInstallation,
	})
}

// uninstall godoc
//
//	@Summary	uninstall an App to Channel
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"id of Channel"
//	@Param		appID		path	string	true	"id of App to uninstall"
//
//	@Success	200
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID} [delete]
func (h *Handler) uninstall(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")
	if err := h.installer.UnInstallApp(ctx, appmodel.InstallationID{
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
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"id of Channel"
//	@Param		appID		path		string	true	"id of App"
//	@Param		object		body		object	true	"key-value of Config to set"
//
//	@Success	200			{object}	model.ConfigMap
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID}/configs [put]
func (h *Handler) setConfig(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	var configMap map[string]string
	if err := ctx.ShouldBindBodyWith(&configMap, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	ret, err := h.configSvc.SetConfig(ctx, appmodel.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	}, configMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret.Configs)
}

// getConfig godoc
//
//	@Summary	get App config of a Installation
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		appID		path		string	true	"id of app"
//	@Param		channelID	path		string	true	"id of channel"
//
//	@Success	200			{object}	any		"JSON of configMap"
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID}/configs [get]
func (h *Handler) getConfig(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	cfgs, err := h.configSvc.GetConfig(ctx, appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, cfgs)
}

// query godoc
//
//	@Summary		get App and Installation
//	@Tags			Desk
//	@Description	get App and Installation installed to channel.
//
//	@Param			x-account	header		string	true	"access token"
//	@Param			channelID	path		string	true	"id of Channel"
//	@Param			appID		path		string	false	"id of App"
//
//	@Success		200			{object}	dto.InstalledAppWithCommands
//	@Router			/desk/v1/channels/{channelID}/installed-apps/{appID} [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	appFound, appInstallation, err := h.querySvc.Query(ctx, appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdRepo.FetchAllByAppID(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.InstalledAppWithCommands{
		Commands:        dto.NewCommandDTOs(cmds),
		App:             appFound,
		AppInstallation: appInstallation,
	})
}

// queryAll godoc
//
//	@Summary		get Apps and AppInstallations
//	@Tags			Desk
//	@Description	get All Apps and AppInstallations installed to channel.
//
//	@Param			x-account	header		string	true	"access token"
//	@Param			channelID	path		string	true	"id of Channel"
//
//	@Success		200			{object}	dto.InstalledAppsWithCommands
//	@Router			/desk/v1/channels/{channelID}/installed-apps [get]
func (h *Handler) queryAll(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, appInstallations, err := h.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdRepo.FetchAllByAppIDs(ctx, app.AppIDsOf(appInstallations))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.InstalledAppsWithCommands{
		Commands:         dto.NewCommandDTOs(cmds),
		Apps:             apps,
		AppInstallations: appInstallations,
	})
}
