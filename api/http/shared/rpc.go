package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	appchannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

func ExecuteRpc(invoker *saga.InstallAwareInvokeSaga[any, any]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		appID, rpcName := ctx.Param("appId"), ctx.Param("name")

		var body dto.ArgumentsAndContext
		if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
			_ = ctx.Error(err)
			return
		}

		identifier := appchannel.AppChannelIdentifier{
			AppID:     appID,
			ChannelID: body.Context["channelId"].(string),
		}

		//rpcRequest := rpc.RpcRequest{
		//	AppID: appID,
		//	InvokeRequest: rpc.InvokeRequest{
		//		Name:    rpcName,
		//		Context: nil, // auth 후 context 주입 필요
		//		Arguments:  body.Arguments,
		//	},
		//}
		_ = rpcName

		res, err := invoker.Invoke(ctx, saga.InstallAwareRequest[any]{
			Identifier: identifier,
			Req:        nil,
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
