package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	settingsvc "github.com/channel-io/ch-app-store/internal/apphttp/svc"
)

// fetchEndpoints godoc
//
//	@Summary	fetch App
//	@Tags		Public
//
//	@Param		appID		path		string	true	"appID"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	settingsvc.Urls
//	@Router		/desk/account/apps/{appID}/server-settings/endpoints  [get]
func (h *Handler) fetchEndpoints(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	settings, err := h.settingPermissionSvc.FetchURLs(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, settings)
}

// modifyEndpoints godoc
//
//	@Summary	modify App
//	@Tags		Public
//
//	@Param		appID			path	string			true	"appID"
//	@Param		settingsvc.Urls	body	settingsvc.Urls	true	"dto"
//	@Param		x-account		header	string			true	"token"
//
//	@Success	200
//	@Router		/desk/account/apps/{appID}/server-settings/endpoints  [put]
func (h *Handler) modifyEndpoints(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")
	var request settingsvc.Urls
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.settingPermissionSvc.UpsertURLs(ctx, appID, request, account.ID); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// refreshSigningKey godoc
//
//	@Summary	refresh signing key
//	@Tags		Public
//
//	@Param		appID		path		string	true	"appID"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	dto.SigningKey
//	@Router		/desk/account/apps/{appID}/server-settings/signing-key  [put]
func (h *Handler) refreshSigningKey(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	key, err := h.settingPermissionSvc.RefreshSigningKey(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.SigningKey{
		SigningKey: key,
	})
}

// checkSigningKey godoc
//
//	@Summary	check signing key has issued before
//	@Tags		Public
//
//	@Param		appID		path		string	true	"appID"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	dto.IssuedBefore
//	@Router		/desk/account/apps/{appID}/server-settings/signing-key  [get]
func (h *Handler) checkSigningKey(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	issued, err := h.settingPermissionSvc.HasIssuedBefore(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.IssuedBefore{
		IssuedBefore: issued,
	})
}
