package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appchannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
)

func GetAllWithApp(appRepo app.AppRepository, appChannelRepo appchannel.AppChannelRepository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		channelID := ctx.Param("channelId")
		appID := ctx.Query("appId")

		if appID != "" {
			identifier := appchannel.AppChannelIdentifier{
				AppID:     appID,
				ChannelID: channelID,
			}

			appTarget, err := appRepo.Fetch(ctx, appID)
			if err != nil {
				_ = ctx.Error(err)
				return
			}

			appChannelTarget, err := appChannelRepo.Fetch(ctx, identifier)
			if err != nil {
				_ = ctx.Error(err)
				return
			}

			ctx.JSON(http.StatusOK, &dto.AppAndAppChannel{App: *appTarget, AppChannel: *appChannelTarget})
			return
		}

		appChannels, err := appChannelRepo.FindAllByChannel(ctx, channelID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		var appChannelIDs []string
		for _, appChannelID := range appChannelIDs {
			appChannelIDs = append(appChannelIDs, appChannelID)
		}

		var result []dto.AppAndAppChannel
		for _, appChannelTarget := range appChannels {
			appTarget, err := appRepo.Fetch(ctx, appChannelTarget.AppID)
			if err != nil {
				continue
			}

			result = append(result, dto.AppAndAppChannel{
				App:        *appTarget,
				AppChannel: *appChannelTarget,
			})
		}

		ctx.JSON(http.StatusOK, result)
	}
}

func GetConfig(svc *appchannel.ConfigSvc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		channelID := ctx.Param("channelId")
		appID := ctx.Query("appId")
		identifier := appchannel.AppChannelIdentifier{
			AppID:     appID,
			ChannelID: channelID,
		}

		configMap, err := svc.GetConfig(ctx, identifier)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, configMap)
	}
}
