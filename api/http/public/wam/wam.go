package wam

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

// downloadWAM godoc
//
//	@Summary	download wam of an app
//	@Tags		Front
//
//	@Param		appID	path		string	true	"id of App"
//	@Param		path	path		string	true	"file path"
//
//	@Success	200		{object}	object
//	@Router		/public/v1/apps/{appID}/wams/{path} [get]
func (h *Handler) downloadWAM(ctx *gin.Context) {
	appID, path := ctx.Param("appID"), ctx.Param("path")

	err := h.wamDownloader.StreamFile(ctx, app.StreamRequest{
		AppID:  appID,
		Path:   path,
		Writer: ctx.Writer,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Writer.Flush()
}
