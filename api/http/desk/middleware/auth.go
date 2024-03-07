package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/lib/log"
)

const ManagerKey = "Manager"

type Auth struct {
	managerSvc account.ManagerFetcher
	logger     log.ContextAwareLogger
}

func NewAuth(managerSvc account.ManagerFetcher, logger log.ContextAwareLogger) *Auth {
	return &Auth{managerSvc: managerSvc, logger: logger}
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
	a.logger.Debugw(ctx, "injecting manager principal", "request", ctx.Request.RequestURI, "manager", authenticatedManager.ID)
	ctx.Set(ManagerKey, authenticatedManager)
}
