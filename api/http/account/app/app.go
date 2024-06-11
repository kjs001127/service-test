package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/account/dto"
	"github.com/channel-io/ch-app-store/api/http/account/middleware"

	_ "github.com/channel-io/ch-app-store/internal/app/model"
	_ "github.com/channel-io/ch-app-store/internal/permission/svc"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

// createApp godoc
//
//	@Summary	create App to app-store
//	@Tags		Public
//
//	@Param		app.AppRequest	body		dto.AppCreateRequest	true	"App title to create"
//	@Param		x-account		header		string					true	"token"
//
//	@Success	201				{object}	dto.AppGeneral
//	@Router		/desk/account/apps [post]
func (h *Handler) createApp(ctx *gin.Context) {
	account := middleware.Account(ctx)
	var request dto.AppCreateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := request.Validate()
	if err != nil {
		_ = ctx.Error(apierr.BadRequest(err))
		return
	}
	created, err := h.appPermissionSvc.CreateApp(ctx, request.Title, account.ID)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appWithDisplay, err := h.appWithDisplayQuerySvc.AddDisplayToApp(ctx, created)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.FromAppWithDisplay(appWithDisplay))
}

// deleteApp godoc
//
//	@Summary	delete App to app-store
//	@Tags		Public
//
//	@Param		x-account	header	string	true	"token"
//	@Param		appID		path	string	true	"appID"
//
//	@Success	204
//	@Router		/desk/account/apps/{appID}  [delete]
func (h *Handler) deleteApp(ctx *gin.Context) {
	account := middleware.Account(ctx)
	appID := ctx.Param("appID")

	err := h.appPermissionSvc.DeleteApp(ctx, appID, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// listApps godoc
//
//	@Summary	list apps
//	@Tags		Public
//
//	@Param		x-account	header		string	true	"token"
//
//	@Success	200			{object}	[]dto.AppView
//	@Router		/desk/account/apps  [get]
func (h *Handler) listApps(ctx *gin.Context) {
	account := middleware.Account(ctx)

	appsWithDisplay, err := h.displayPermissionSvc.GetAppsWithDisplayByAccount(ctx, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppWithDisplayViewsFrom(appsWithDisplay))
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

	appsWithDisplay, err := h.displayPermissionSvc.GetCallableAppsWithDisplay(ctx, account.ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.FromAppsWithDisplay(appsWithDisplay))
}
