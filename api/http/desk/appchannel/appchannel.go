package appchannel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

func (h *Handler) install(ctx *gin.Context) {
	channelID := ctx.Param("channelID")
	var req dto.AppIDRequest
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}
	appID := req.AppID

	identifier := domain.AppChannelIdentifier{
		AppID:     appID,
		ChannelID: channelID,
	}

	appChannel, err := h.installSaga.Install(ctx, identifier)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	app, err := h.appRepo.Fetch(ctx, appID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.AppAndAppChannel{
		App:        *app,
		AppChannel: *appChannel,
	})
}

func (h *Handler) uninstall(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelID"), ctx.Param("appID")
	identifier := domain.AppChannelIdentifier{
		AppID:     appID,
		ChannelID: channelID,
	}

	if err := h.installSaga.Uninstall(ctx, identifier); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) setConfig(ctx *gin.Context) {
	channelID, appID := ctx.Param("channelId"), ctx.Param("appId")
	identifier := domain.AppChannelIdentifier{
		AppID:     appID,
		ChannelID: channelID,
	}

	var configMap map[string]string
	if err := ctx.ShouldBindBodyWith(configMap, binding.JSON); err != nil {
		_ = ctx.Error(err)
		return
	}

	ret, err := h.installSaga.SetConfig(ctx, identifier, configMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
