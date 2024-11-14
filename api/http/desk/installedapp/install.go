package installedapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	commondto "github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/api/http/desk/installedapp/dto"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
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

	managerRequester := middleware.ManagerRequester(ctx)

	_, err := h.installer.Install(ctx, managerRequester.Manager, appmodel.InstallationID{
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

	managerRequester := middleware.ManagerRequester(ctx)

	if err := h.installer.UnInstall(ctx, managerRequester.Manager, appmodel.InstallationID{
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
//	@Success		200			{object}	dto.DetailedInstalledAppPageView
//	@Router			/desk/v1/channels/{channelID}/installed-apps/{appID} [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	installID := appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	}

	appFound, err := h.installQuerySvc.QueryInstalledApp(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appDetail, err := h.appQuerySvc.ReadDetail(ctx, appFound.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdQuerySvc.FetchAllWithActivation(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	roles, err := h.agreementSvc.FetchUnAgreedRoles(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.DetailedInstalledAppPageView{
		Commands:      dto.NewInstalledCommandViews(cmds),
		App:           commondto.FromDetail(appDetail),
		UnAgreedRoles: dto.FromRoles(roles),
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
//	@Success		200			{array}	dto.SimpleInstalledAppView
//	@Router			/desk/v1/channels/{channelID}/installed-apps [get]
func (h *Handler) queryAll(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, err := h.installQuerySvc.QueryInstalledAppsByChannelID(ctx, channelID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	rolesToAgree, err := h.agreementSvc.BulkFetchUnAgreedRoles(ctx, channelID, app.AppIDsOf(apps))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	agreementMap := make(map[string]bool)
	for _, role := range rolesToAgree {
		agreementMap[role.AppID] = true
	}

	var ret []*dto.SimpleInstalledAppView
	for _, installed := range apps {
		ret = append(ret, &dto.SimpleInstalledAppView{
			ShouldUpdateAgreement: agreementMap[installed.ID],
			SimpleAppView:         commondto.NewAppView(installed),
		})
	}

	ctx.JSON(http.StatusOK, ret)
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

	managerRequester := middleware.ManagerRequester(ctx)

	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	if err := h.activateSvc.ToggleByKey(ctx, managerRequester, cmd.ToggleRequest{
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
//	@Param			x-account	header	string	true	"access token"
//	@Param			channelID	path	string	true	"id of Channel"
//	@Param			appID		path	string	false	"id of App"
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

// getAppRole godoc
//
//	@Summary	get claims of an app
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"channelID"
//	@Param		appID		path	string	true	"appID"
//
//	@Success	200			{array}	dto.DeskRoleView
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID}/roles [get]
func (h *Handler) getAppRoles(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	installID := appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	}

	ret, err := h.agreementSvc.FetchUnAgreedRoles(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}

// agreeToRoles godoc
//
//	@Summary	get claims of an app
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"channelID"
//	@Param		appID		path	string	true	"appID"
//
//	@Success	200			{array}	dto.DeskRoleView
//	@Router		/desk/v1/channels/{channelID}/installed-apps/{appID}/roles [post]
func (h *Handler) agreeToRoles(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")

	var appRoleIDs []string
	if err := ctx.ShouldBindBodyWith(&appRoleIDs, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.agreementSvc.Agree(ctx, channelID, appRoleIDs); err != nil {
		_ = ctx.Error(err)
		return
	}

	installID := appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	}

	ret, err := h.agreementSvc.FetchUnAgreedRoles(ctx, installID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, dto.FromRoles(ret))
}
