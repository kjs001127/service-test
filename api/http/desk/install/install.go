package install

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	cmdmodel "github.com/channel-io/ch-app-store/internal/command/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/svc"
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

	manager := middleware.Manager(ctx)

	_, err := h.installer.Install(ctx, manager.Manager, appmodel.InstallationID{
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
	manager := middleware.Manager(ctx)

	if err := h.installer.UnInstall(ctx, manager.Manager, appmodel.InstallationID{
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

	installID := appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	}

	appFound, err := h.appQuerySvc.Query(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdQuerySvc.FetchAllWithActivation(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.InstalledAppDetailView{
		Commands: dto.NewInstalledCommandViews(cmds),
		App:      dto.NewAppDetailView(appFound),
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

	apps, err := h.appQuerySvc.QueryAll(ctx, channelID)
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
//	@Router			/desk/v1/channels/{channelID}/installed-apps/{appID}/commands [put]
func (h *Handler) toggleCmd(ctx *gin.Context) {
	var body dto.CommandToggleRequest
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	manager := middleware.Manager(ctx)
	manager.Manager.Language = body.Language

	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	if err := h.activateSvc.ToggleByKey(ctx, manager.Manager, cmd.ToggleRequest{
		Command: cmdmodel.CommandKey{
			Name:  body.Name,
			Scope: body.Scope,
			AppID: appID,
		},
		Enabled:   body.Enabled,
		ChannelID: channelID,
	}); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// queryCommands godoc
//
//	@Summary		get App and Installation
//	@Tags			Desk
//	@Description	get App and Installation installed to channel.
//
//	@Param			x-account	header		string	true	"access token"
//	@Param			channelID	path		string	true	"id of Channel"
//	@Param			appID		path		string	false	"id of App"
//
//	@Success		200			{array}	dto.InstalledCommandView
//	@Router			/desk/v1/channels/{channelID}/installed-apps/{appID}/commands [get]
func (h *Handler) queryCommands(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	installID := appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	}

	cmds, err := h.cmdQuerySvc.FetchAllWithActivation(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if scope := cmdmodel.Scope(ctx.Query("scope")); len(scope) > 0 {
		cmds = h.filterWithScope(cmds, scope)
	}

	ctx.JSON(http.StatusOK, dto.NewInstalledCommandViews(cmds))
}

func (h *Handler) filterWithScope(cmds []*cmd.CommandWithActivation, scope cmdmodel.Scope) []*cmd.CommandWithActivation {
	ret := make([]*cmd.CommandWithActivation, 0)
	for _, command := range cmds {
		if command.Scope == scope {
			ret = append(ret, command)
		}
	}
	return ret
}