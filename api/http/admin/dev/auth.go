package dev

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	"github.com/channel-io/ch-app-store/internal/approle/svc"
)

// fetchRole godoc
//
//	@Summary	fetch App
//	@Tags		Public
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
	claims, err := h.roleSvc.FetchRole(ctx, appID, roleType)
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
//	@Summary	fetch App
//	@Tags		Public
//
//	@Param		appId		path	string	true	"appId"
//	@Param		roleType	path	string	true	"roleType"
//
//	@Success	200			{array}	model.Claim
//	@Router		/admin/apps/{appId}/auth/roles/{roleType}  [put]
func (h *Handler) modifyClaims(ctx *gin.Context) {
	appID := ctx.Param("appID")
	roleType := model.RoleType(ctx.Param("roleType"))

	var req svc.ClaimsDTO
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.roleSvc.UpdateRole(ctx, appID, roleType, &req)
	if err != nil {
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
//	@Param		appId	path		string	true	"appId"
//
//	@Success	200		{object}	dto.AdminAppSecret
//	@Router		/admin/apps/{appId}/auth/secret [put]
func (h *Handler) refreshSecret(ctx *gin.Context) {
	appID := ctx.Param("appID")

	token, err := h.roleSvc.RefreshAppSecret(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.AdminAppSecret{
		Secret: token,
	})
}
