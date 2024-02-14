package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/auth/principal/account"
)

const ManagerKey = "Manager"

type Auth struct {
	managerSvc account.ManagerFetcher
}

func NewAuth(managerSvc account.ManagerFetcher) *Auth {
	return &Auth{managerSvc: managerSvc}
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, "/desk") {
		return
	}

	xAccount := ctx.GetHeader(account.XAccountHeader)
	if xAccount == "" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			dto.HttpUnauthorizedError(errors.New("authorization header is empty")),
		)
		return
	}

	channelID := ctx.Param("channelID")
	if channelID == "" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			dto.HttpUnauthorizedError(errors.New("channelID is empty")),
		)
		return
	}

	authenticatedManager, err := a.managerSvc.FetchManager(ctx, channelID, xAccount)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			dto.HttpUnprocessableEntityError(err),
		)
		return
	}

	ctx.Set(ManagerKey, authenticatedManager)
}
