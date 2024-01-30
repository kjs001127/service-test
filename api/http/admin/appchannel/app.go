package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

// query godoc
//
//	@Summary	get command, brief, apps of channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"channelID"
//
//	@Success	200			{object}	object
//	@Router		/admin/v1/channels/{channelID}/apps [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, err := h.appRepo.QueryAll(ctx, channelID)

	briefs, err := h.briefRepo.FetchAll(ctx, app.AppIDsOf(apps.AppChannels))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	cmds, err := h.cmdRepo.FetchAllByAppIDs(ctx, app.AppIDsOf(apps.AppChannels))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AppResources{
		Apps:     apps.Apps,
		Commands: cmds,
		Briefs:   briefs,
	})
}
