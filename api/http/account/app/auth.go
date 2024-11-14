package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/internal/role/svc"
)

// fetchRole godoc
//
//	@Summary	fetch App
//	@Tags		PublicNativeClaims
//
//	@Param		appId		path		string	true	"appId"
//	@Param		roleType	path		string	true	"roleType"
//	@Param		x-account	header		string	true	"token"
//	@Success	200			{object}	dto.RoleView
//	@Router		/desk/account/apps/{appId}/auth/roles/{roleType}  [get]
func (h *Handler) fetchRole(ctx *gin.Context) {
	acc := middleware.Account(ctx)
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	availableClaims, err := h.authPermissionSvc.GetAvailableNativeClaims(ctx, appID, roleType)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	claims, err := h.authPermissionSvc.FetchLatestRole(ctx, appID, roleType, acc.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, h.roleViewWithAvailable(claims, availableClaims))
}

func (h *Handler) roleViewWithAvailable(resp *svc.ClaimsResponse, available model.Claims) *dto.RoleView {

	return &dto.RoleView{
		AvailableNativeClaims: available,
		AppClaims:             resp.AppClaims,
		NativeClaims:          resp.NativeClaims,
		ID:                    resp.ID,
		Type:                  resp.Type,
	}
}

func (h *Handler) roleViewOf(resp *svc.ClaimsResponse) *dto.RoleView {

	return &dto.RoleView{
		AppClaims:    resp.AppClaims,
		NativeClaims: resp.NativeClaims,
		ID:           resp.ID,
		Type:         resp.Type,
	}
}

// modifyClaims godoc
//
//	@Summary	fetch App
//	@Tags		PublicNativeClaims
//
//	@Param		appId			path	string			true	"appId"
//	@Param		roleType		path	string			true	"roleType"
//	@Param		x-account		header	string			true	"token"
//	@Param		svc.ClaimsRequestWithID	body	svc.ClaimsRequestWithID	true	"claims"
//
//	@Success	200
//	@Router		/desk/account/apps/{appId}/auth/roles/{roleType}  [put]
func (h *Handler) modifyClaims(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	var req svc.ClaimsRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	req.AppID = appID

	resp, err := h.authPermissionSvc.CreateRole(ctx, &req, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, h.roleViewOf(resp))
}

// refreshSecret godoc
//
//	@Summary	refresh signing key
//	@Tags		PublicNativeClaims
//
//	@Param		appId		path		string	true	"appId"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	dto.AppSecret
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
//	@Tags		PublicNativeClaims
//
//	@Param		appId		path		string	true	"appId"
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	dto.IssuedBefore
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

// getCallableApps godoc
//
//	@Summary	get callable apps
//	@Tags		Public
//
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	[]dto.AppGeneral
//	@Router		/desk/account/auth/apps  [get]
func (h *Handler) getCallableApps(ctx *gin.Context) {
	account := middleware.Account(ctx)

	apps, err := h.authPermissionSvc.GetAvailableApps(ctx, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromApps(apps))
}
