package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	appChannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	command "github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

func QueryCommands(
	svc *command.QuerySvc,
	appRepo app.AppRepository,
	appChannelRepo appChannel.AppChannelRepository,
	scope command.Scope,
) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		channelID := ctx.Param("channelId")
		appChannels, err := appChannelRepo.FindAllByChannel(ctx, channelID)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		var appIDs []string
		for _, appChannelTarget := range appChannels {
			appIDs = append(appIDs, appChannelTarget.AppID)
		}

		apps, err := appRepo.FindAll(ctx, appIDs)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		commands, err := svc.QueryCommands(ctx, command.Query{ // paging?
			AppIDs: appIDs,
			Scope:  scope,
		})

		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, &dto.AppsAndCommands{
			Apps:     apps,
			Commands: commands,
		})
	}
}

func AutoComplete(invokeSaga *saga.InstallAwareInvokeSaga[any, any], scope command.Scope) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		channelID, appID, name := ctx.Param("channelId"), ctx.Param("appId"), ctx.Param("name")
		var body dto.ContextAndAutoCompleteArgs
		if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
			_ = ctx.Error(err)
			return
		}

		identifier := appChannel.AppChannelIdentifier{
			AppID:     appID,
			ChannelID: channelID,
		}

		res, err := invokeSaga.Invoke(ctx, saga.InstallAwareRequest[any]{
			Identifier: identifier,
			Req: command.AutoCompleteRequest{
				Key: command.Key{
					AppID: appID,
					Scope: string(scope),
					Name:  name,
				},
				Context:   body.Context,
				Arguments: body.Params,
			},
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
