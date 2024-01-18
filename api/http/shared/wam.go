package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/wam/domain"
)

func GetWamUrl(wamSvc *domain.WamSvc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		appID, name := ctx.Param("appId"), ctx.Param("name")
		key := domain.WamKey{
			AppID: appID,
			Name:  name,
		}
		ctx.Redirect(http.StatusTemporaryRedirect, wamSvc.GetWamUrl(ctx, key))
	}
}
