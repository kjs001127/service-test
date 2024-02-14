package query

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	shareddto "github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

// query godoc
//
//	@Summary	get command, brief, apps of channel
//	@Tags		Admin
//
//	@Param		channelID	path		string	true	"channelID"
//
//	@Success	200			{object}	dto.AppResources
//	@Router		/admin/channels/{channelID}/apps [get]
func (h *Handler) query(ctx *gin.Context) {
	channelID := ctx.Param("channelID")

	apps, err := h.querySvc.QueryAll(ctx, channelID)

	cmds, err := h.cmdRepo.FetchAllByAppIDs(ctx, app.AppIDsOf(apps.AppChannels))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, shareddto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, dto.AppResources{
		Apps:     apps.Apps,
		Commands: cmds,
	})
}
