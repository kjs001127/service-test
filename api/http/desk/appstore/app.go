package appstore

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
)

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
//	@Success	200			{array}	dto.AppView
//	@Router		/desk/v1/channels/{channelID}/app-store/apps  [get]
func (h *Handler) getApps(ctx *gin.Context) {
	apps, err := h.findApps(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewAppViews(apps))
}

func (h *Handler) findApps(ctx *gin.Context) ([]*appmodel.App, error) {
	since, limit := ctx.Query("since"), ctx.Query("limit")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	if isPrivate(ctx) {
		return h.privateAppQuerySvc.GetAppsByAccount(ctx, middleware.Manager(ctx).AccountID)
	} else {
		return h.appRepo.FindPublicApps(ctx, since, limitNumber)
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
//	@Success	200			{object}	dto.AppDetailView
//	@Router		/desk/v1/channels/{channelID}/app-store/apps/{appID}  [get]
func (h *Handler) getAppDetail(ctx *gin.Context) {
	appID := ctx.Param("appID")

	app, err := h.appRepo.FindApp(ctx, appID)
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
		App:      dto.NewAppDetailView(app),
		Commands: dto.NewCommandViews(cmds),
	})
}
