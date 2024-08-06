package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/approle/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
)

// fetchRole godoc
//
//	@Summary	fetch App
//	@Tags		Public
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
	view, err := h.roleViewOf(ctx, appID, acc.Account, roleType)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, view)
}

func (h *Handler) roleViewOf(ctx context.Context, appID string, account account.Account, roleType model.RoleType) (*dto.RoleView, error) {
	claims, err := h.authPermissionSvc.FetchRole(ctx, appID, roleType, account.ID)
	if err != nil {
		return nil, err
	}
	availableClaims, err := h.authPermissionSvc.GetAvailableNativeClaims(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}
	return &dto.RoleView{
		AvailableNativeClaims: availableClaims,
		AppClaims:             claims.AppClaims,
		NativeClaims:          claims.NativeClaims,
	}, nil
}

// modifyClaims godoc
//
//	@Summary	fetch App
//	@Tags		Public
//
//	@Param		appId			path	string			true	"appId"
//	@Param		roleType		path	string			true	"roleType"
//	@Param		x-account		header	string			true	"token"
//	@Param		model.Claims	body	model.Claims	true	"claims"
//
//	@Success	200				{array}	model.Claim
//	@Router		/desk/account/apps/{appId}/auth/roles/{roleType}  [put]
func (h *Handler) modifyClaims(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	var req svc.ClaimsDTO
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.authPermissionSvc.UpdateRole(ctx, appID, roleType, &req, account.ID); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// refreshSecret godoc
//
//	@Summary	refresh signing key
//	@Tags		Public
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
//	@Tags		Public
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
