package wam

import (
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/remoteapp/interaction/svc"
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

	reqCloned := *ctx.Request
	reqCloned.URL.Path = path
	err := h.wamDownloader.StreamFile(ctx, svc.WamProxyRequest{
		AppID:  appID,
		Writer: ctx.Writer,
		Req:    ctx.Request,
	})

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Writer.Flush()
}
