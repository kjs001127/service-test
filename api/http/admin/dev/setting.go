package dev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	settingsvc "github.com/channel-io/ch-app-store/internal/apphttp/svc"
)

// fetchEndpoints godoc
//
//	@Summary	fetch App
//	@Tags		Public
//
//	@Param		appID		path		string	true	"appID"
//
//	@Success	200		{object}	settingsvc.Urls
//	@Router		/admin/apps/{appID}/server-settings/endpoints  [get]
func (h *Handler) fetchEndpoints(ctx *gin.Context) {
	appID := ctx.Param("appID")

	settings, err := h.settingSvc.FetchUrls(ctx, appID)
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
//
//	@Success	200
//	@Router		/admin/apps/{appID}/server-settings/endpoints  [put]
func (h *Handler) modifyEndpoints(ctx *gin.Context) {
	appID := ctx.Param("appID")
	var request settingsvc.Urls
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.settingSvc.UpsertUrls(ctx, appID, request); err != nil {
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
//
//	@Success	200		{object}	dto.SigningKey
//	@Router		/admin/apps/{appID}/server-settings/signing-key  [put]
func (h *Handler) refreshSigningKey(ctx *gin.Context) {
	appID := ctx.Param("appID")

	key, err := h.settingSvc.RefreshSigningKey(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.SigningKey{
		SigningKey: key,
	})
}
