package dev

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
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
//
//	@Success	200			{object}	dto.AdminRoleView
//	@Router		/admin/apps/{appId}/auth/roles/{roleType}  [get]
func (h *Handler) fetchRole(ctx *gin.Context) {
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))
	view, err := h.roleViewOf(ctx, appID, roleType)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, view)
}

func (h *Handler) roleViewOf(ctx context.Context, appID string, roleType model.RoleType) (*dto.AdminRoleView, error) {
	claims, err := h.roleSvc.FetchLatestRole(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	return &dto.AdminRoleView{
		NativeClaims: claims.NativeClaims,
		AppClaims:    claims.AppClaims,
	}, nil
}

// modifyClaims godoc
//
//	@Summary	overwrite claims of specific role type
//	@Tags		PublicNativeClaims
//
//	@Param		appId		path	string	true	"appId"
//	@Param		roleType	path	string	true	"roleType"
//
//	@Success	200			{array}	model.Claim
//	@Router		/admin/apps/{appId}/auth/roles/{roleType}  [put]
func (h *Handler) modifyClaims(ctx *gin.Context) {
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	var req svc.ClaimsRequest
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	req.AppID = appID
	req.Type = roleType

	resp, err := h.roleSvc.CreateRole(ctx, &req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// addClaims godoc
//
//	@Summary	add claims to specific role type
//	@Tags		PublicNativeClaims
//
//	@Param		appId		path	string	true	"appId"
//	@Param		roleType	path	string	true	"roleType"
//
//	@Success	200			{array}	model.Claim
//	@Router		/admin/apps/{appId}/auth/roles/{roleType}  [patch]
func (h *Handler) addClaims(ctx *gin.Context) {
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	var req svc.ClaimsRequestWithID
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	role, err := h.roleSvc.FetchLatestRole(ctx, appID, roleType)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	req.AppID = appID
	req.ID = role.ID

	if err := h.roleSvc.AppendClaimsToRole(ctx, &req); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// refreshSecret godoc
//
//	@Summary	refresh signing key
//	@Tags		PublicNativeClaims
//
//	@Param		appId	path		string	true	"appId"
//
//	@Success	200		{object}	dto.AdminAppSecret
//	@Router		/admin/apps/{appId}/auth/secret [put]
func (h *Handler) refreshSecret(ctx *gin.Context) {
	appID := ctx.Param("appID")

	token, err := h.secretSvc.RefreshAppSecret(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.AdminAppSecret{
		Secret: token,
	})
}
