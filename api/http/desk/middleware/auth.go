package middleware

import (
	"errors"
	"strings"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/auth/principal"
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

	xAccount := ctx.GetHeader(principal.TokenTypeAccount.Header())
	if xAccount == "" {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("authorization header is empty")))
		return
	}

	channelID := ctx.Param("channelID")
	if channelID == "" {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("channelID is empty")))
		return
	}

	authenticatedManager, err := a.managerSvc.FetchManager(ctx, channelID, xAccount)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(err)
		return
	}

	ctx.Set(ManagerKey, authenticatedManager)
}
