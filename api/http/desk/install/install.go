package install

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
)

// install godoc
//
//	@Summary	install an App to Channel
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"id of Channel"
//	@Param		appID		path	string	true	"id of App to install"
//
//	@Success	200
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID} [put]
func (h *Handler) install(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	appID := ctx.Param("appID")

	_, err := h.installer.InstallAppById(ctx, appmodel.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
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
//	@Success		200			{object}	dto.InstalledAppDetailView
//	@Router			/desk/v1/channels/{channelID}/installed-apps/{appID} [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	appFound, err := h.querySvc.Query(ctx, appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdQuerySvc.FetchAllByAppID(ctx, appFound.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmdEnabled, err := h.activateSvc.Check(ctx, appmodel.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.InstalledAppDetailView{
		Commands:       dto.NewCommandViews(cmds),
		CommandEnabled: cmdEnabled,
		App:            dto.NewAppDetailView(appFound),
	})
}

// queryAll godoc
//
//	@Summary		get Apps and AppChannels
//	@Tags			Desk
//	@Description	get All Apps and AppChannels installed to channel.
//
//	@Param			x-account	header	string	true	"access token"
//	@Param			channelID	path	string	true	"id of Channel"
//
//	@Success		200			{array}	dto.AppView
//	@Router			/desk/v1/channels/{channelID}/installed-apps [get]
func (h *Handler) queryAll(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, err := h.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewAppViews(apps))
}

// toggleCmd godoc
//
//	@Summary		toggleCommand
//	@Tags			Desk
//	@Description	get All Apps and AppChannels installed to channel.
//
//	@Param			x-account					header	string						true	"access token"
//	@Param			channelID					path	string						true	"id of Channel"
//	@Param			appID						path	string						true	"id of App"
//	@Param			dto.CommandToggleRequest	body	dto.CommandToggleRequest	true	"toggleCmd body"
//
//	@Success		200
//	@Router			/desk/v1/channels/{channelID}/installed-apps/{appId}/commands [put]
func (h *Handler) toggleCmd(ctx *gin.Context) {
	var body dto.CommandToggleRequest
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	if err := h.activateSvc.Toggle(ctx, appmodel.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	}, body.CommandEnabled); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
