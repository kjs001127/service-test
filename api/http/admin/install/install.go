package install

import (
	"net/http"

	app "github.com/channel-io/ch-app-store/internal/app/model"

	_ "github.com/channel-io/ch-app-store/internal/appdisplay/svc"

	"github.com/gin-gonic/gin"
)

// checkInstall godoc
//
//	@Summary	checkInstall an App to Channel
//	@Tags		Admin
//
//	@Param		channelID	path	string	true	"id of Channel"
//	@Param		appID		path	string	true	"id of App to install"
//
//	@Success	200
//	@Router		/admin/channels/{channelID}/installed-apps/{appID} [get]
func (h *Handler) checkInstall(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	appID := ctx.Param("appID")

	_, err := h.querySvc.Query(ctx, app.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// install godoc
//
//	@Summary	install an App to Channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"id of Channel"
//	@Param		appID		path		string	true	"id of App to install"
//
//	@Success	200			{object}	svc.AppWithDisplay
//	@Router		/admin/channels/{channelID}/installed-apps/{appID} [put]
func (h *Handler) install(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	appID := ctx.Param("appID")

	appInstalled, err := h.installer.InstallAppById(ctx, app.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appWithDisplay, err := h.appWithDisplayQuerySvc.AddDisplayToApp(ctx, appInstalled)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, appWithDisplay)
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
//	@Router		/admin/channels/{channelID}/installed-apps/{appID} [delete]
func (h *Handler) uninstall(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")
	if err := h.installer.UnInstallApp(ctx, app.InstallationID{
		AppID:     appID,
		ChannelID: channelID,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
