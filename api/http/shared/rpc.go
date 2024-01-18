package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	appchannel "github.com/channel-io/ch-app-store/internal/appchannel/domain"
	rpc "github.com/channel-io/ch-app-store/internal/rpc/domain"
	"github.com/channel-io/ch-app-store/internal/saga"
)

func ExecuteRpc(invoker *saga.InstallAwareInvoker) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		appID, rpcName := ctx.Param("appId"), ctx.Param("name")

		var body dto.ParamAndContext
		if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
			_ = ctx.Error(err)
			return
		}

		identifier := appchannel.AppChannelIdentifier{
			AppID:     appID,
			ChannelID: body.Context["channelId"].(string),
		}

		rpcRequest := rpc.RpcRequest{
			AppID: appID,
			InvokeRequest: rpc.InvokeRequest{
				Name:    rpcName,
				Context: nil, // auth 후 context 주입 필요
				Params:  body.Params,
			},
		}

		res, err := invoker.Invoke(ctx, saga.InstallAwareInvokeRequest{
			Identifier: identifier,
			Request:    rpcRequest,
		})
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Data(http.StatusOK, "application/octet-stream", res)
	}
}
