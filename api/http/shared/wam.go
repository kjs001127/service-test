package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWamUrl() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "")
	}
}
