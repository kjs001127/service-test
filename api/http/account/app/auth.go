package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

// fetchClaims godoc
//
//	@Summary	fetch App
//	@Tags		Public
//
//	@Param		appId		path	string	true	"appId"
//	@Param		roleType	path	string	true	"roleType"
//
//	@Success	200			{array}	model.Claim
//	@Router		/desk/account/apps/{appId}/auth/roles/{roleType}  [get]
func (h *Handler) fetchClaims(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	claims, err := h.authPermissionSvc.FetchRole(ctx, appID, roleType, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, claims)
}

// modifyClaims godoc
//
//	@Summary	fetch App
//	@Tags		Public
//
//	@Param		appId		path	string	true	"appId"
//	@Param		roleType	path	string	true	"roleType"
//
//	@Success	200			{array}	model.Claim
//	@Router		/desk/account/apps/{appId}/auth/roles/{roleType}  [put]
func (h *Handler) modifyClaims(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	var claims []*model.Claim
	if err := ctx.ShouldBindBodyWith(&claims, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	claims, err := h.authPermissionSvc.UpdateRole(ctx, appID, roleType, claims, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, claims)
}

// refreshSecret godoc
//
//	@Summary	refresh signing key
//	@Tags		Public
//
//	@Param		appId	path		string	true	"appId"
//
//	@Success	200		{object}	dto.AppSecret
//	@Router		/desk/account/apps/{appId}/auth/secret [put]
func (h *Handler) refreshSecret(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	token, err := h.authPermissionSvc.RefreshToken(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.AppSecret{
		Secret: token,
	})
}

// checkSecret godoc
//
//	@Summary	check token issued before
//	@Tags		Public
//
//	@Param		appId	path		string	true	"appId"
//
//	@Success	200		{object}	dto.IssuedBefore
//	@Router		/desk/account/apps/{appId}/auth/secret [get]
func (h *Handler) checkSecret(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	issued, err := h.authPermissionSvc.HasTokenIssuedBefore(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.IssuedBefore{
		IssuedBefore: issued,
	})
}
