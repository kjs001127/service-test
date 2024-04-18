package appdev

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/internal/appdev/svc"
)

// create godoc
//
//	@Summary	create App to app-store
//	@Tags		Admin
//
//	@Param		app.AppRequest	body		svc.AppRequest	true	"App to create"
//
//	@Success	201				{object}	svc.AppResponse
//	@Router		/admin/apps [post]
func (h *Handler) create(ctx *gin.Context) {
	var target svc.AppRequest
	if err := ctx.ShouldBindBodyWith(&target, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	created, err := h.appDevSvc.CreateApp(ctx, target)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// delete godoc
//
//	@Summary	delete App from app-store
//	@Tags		Admin
//
//	@Param		ID	path	string	true	"id of App to delete"
//
//	@Success	204
//	@Router		/admin/apps/{ID} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	ID := ctx.Param("appID")

	err := h.appDevSvc.DeleteApp(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// queryDetail godoc
//
//	@Summary	query App from app-store
//	@Tags		Admin
//
//	@Param		appID	path	string	true "appId"
//
//	@Success	200  	{object} svc.AppResponse
//	@Router		/admin/apps/{appID} [get]
func (h *Handler) queryDetail(ctx *gin.Context) {
	ID := ctx.Param("appID")
	appFound, err := h.appDevSvc.FetchApp(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, appFound)
}

// queryLegacy godoc
//
//	@Summary	query App from app-store
//	@Tags		Admin
//
//	@Param		roleId	query	string	true "roleId of App to query"
//
//	@Success	200  	{object} model.App
//	@Router		/admin/apps [get]
func (h *Handler) queryLegacy(ctx *gin.Context) {
	ID := ctx.Query("roleId")

	role, err := h.appRoleSvc.FetchRole(ctx, ID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	appFound, err := h.appManager.Read(ctx, role.AppID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, AppLegacy{
		ID:                 appFound.ID,
		State:              string(appFound.State),
		Title:              appFound.Title,
		DetailDescriptions: appFound.DetailDescriptions,
		AppType:            "",
		DetailImageURLs:    appFound.DetailImageURLs,
		ConfigSchemas:      make([]map[string]any, 0),
	})
}

type AppLegacy struct {
	ID                 string           `json:"id"`
	State              string           `json:"state"`
	Title              string           `json:"title"`
	Description        string           `json:"description"`
	IsPrivate          bool             `json:"isPrivate"`
	ManualURL          string           `json:"manualUrl"`
	DetailDescriptions []map[string]any `json:"detailDescriptions"`
	DetailImageURLs    []string         `json:"detailImageUrls"`
	ConfigSchemas      []map[string]any `json:"configSchemas"`
	AppType            string           `json:"appType"`
}
