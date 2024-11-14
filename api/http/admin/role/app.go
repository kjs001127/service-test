package role

import (
	"fmt"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/admin/role/dto"
)

// queryByRoleID godoc
//
//	@Summary	query App from app-store
//	@Tags		Admin
//
//	@Param		roleId	path	string	true "roleId of App to query"
//
//	@Success	200  	{object} dto.AppRoles
//	@Router		/admin/app-roles [get]
func (h *Handler) queryByRoleID(ctx *gin.Context) {
	roleID := ctx.Query("roleId")

	role, err := h.roleRepo.FindByRoleID(ctx, roleID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	fullRoles, err := h.roleRepo.FindAllByAppID(ctx, role.AppID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppRoleFrom(role.AppID, fullRoles))
}

// queryByAppID godoc
//
//	@Summary	query App from app-store
//	@Tags		Admin
//
//	@Param		appId	query	string	true "appId of App to query"
//
//	@Success	200  	{object} dto.AppRoles
//	@Router		/admin/app-roles/{appID} [get]
func (h *Handler) queryByAppID(ctx *gin.Context) {
	appID := ctx.Param("appID")

	appRoles, err := h.roleRepo.FindAllByAppID(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if len(appRoles) <= 0 {
		_ = ctx.Error(apierr.NotFound(fmt.Errorf("no app role found with appID: %s", appID)))
	}

	ctx.JSON(http.StatusOK, dto.AppRoleFrom(appID, appRoles))
}
