package appstore

import (
	"context"
	"net/http"
	"strconv"

	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	displaysvc "github.com/channel-io/ch-app-store/internal/appdisplay/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/gin-gonic/gin"
)

// getAppRole godoc
//
//	@Summary	get claims of an app
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		channelID	path	string	true	"channelID"
//	@Param		appID		path	string	true	"appID"
//
//	@Success	200		{array} 	dto.DeskRoleView
//	@Router		/desk/v1/channels/{channelID}/app-store/apps/{appID}/roles  [get]
func (h *Handler) getAppRoles(ctx *gin.Context) {
	appID := ctx.Param("appID")

	views, err := h.roleViewsOf(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, views)
}

func (h *Handler) roleViewsOf(ctx context.Context, appID string) (dto.DeskRoleViews, error) {
	roleTypes := []model.RoleType{model.RoleTypeManager, model.RoleTypeUser, model.RoleTypeChannel}
	ret := make(dto.DeskRoleViews, 0, len(roleTypes))
	for _, roleType := range roleTypes {
		role, err := h.authSvc.FetchRole(ctx, appID, roleType)
		if apierr.IsNotFound(err) {
			continue
		} else if err != nil {
			return nil, err
		}
		ret = append(ret, dto.DeskRoleView{
			Type:         roleType,
			NativeClaims: role.NativeClaims,
			AppClaims:    role.AppClaims,
		})
	}
	return ret, nil
}

// getApps godoc
//
//	@Summary	get list of Apps
//	@Tags		Desk
//
//	@Param		x-account	header	string	true	"access token"
//	@Param		since		query	string	false	"get App after this id"
//	@Param		limit		query	string	true	"max count of return data"
//	@Param		channelID	path	string	true	"channelID"
//
//	@Success	200			{array}	dto.AppStoreAppView
//	@Router		/desk/v1/channels/{channelID}/app-store/apps  [get]
func (h *Handler) getApps(ctx *gin.Context) {
	appsWithDisplay, err := h.findAppsWithDisplay(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewAppWithDisplayViews(appsWithDisplay))
}

func (h *Handler) findAppsWithDisplay(ctx *gin.Context) ([]*displaysvc.AppWithDisplay, error) {
	since, limit := ctx.Query("since"), ctx.Query("limit")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	if isPrivate(ctx) {
		return h.privateDisplayQuerySvc.GetPrivateAppsWithDisplayByAccount(ctx, middleware.Manager(ctx).AccountID)
	} else {
		return h.appWithDisplayQuerySvc.ReadPublicAppsWithDisplay(ctx, since, limitNumber)
	}
}

func isPrivate(ctx *gin.Context) bool {
	isPrivateQuery := ctx.Query("isPrivate")
	if len(isPrivateQuery) <= 0 {
		return false
	}
	isPrivate, err := strconv.ParseBool(isPrivateQuery)
	if err != nil {
		return false
	}

	return isPrivate
}

// getAppDetail godoc
//
//	@Summary	get list of Apps
//	@Tags		Desk
//
//	@Param		x-account	header		string	true	"access token"
//	@Param		channelID	path		string	true	"channelID"
//	@Param		appID		path		string	true	"appID"
//
//	@Success	200			{object}	dto.AppStoreDetailView
//	@Router		/desk/v1/channels/{channelID}/app-store/apps/{appID}  [get]
func (h *Handler) getAppDetail(ctx *gin.Context) {
	appID := ctx.Param("appID")

	appWithDisplay, err := h.appWithDisplayQuerySvc.Read(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdRepo.FetchAllByAppID(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppStoreDetailView{
		App:      dto.NewAppWithDisplayDetailView(appWithDisplay),
		Commands: dto.NewCommandViews(cmds),
	})
}
