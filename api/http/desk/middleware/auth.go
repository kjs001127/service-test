package middleware

import (
	"errors"
	"regexp"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	channelIDPathParam = "channelID"
	ManagerKey         = "Manager"
)

var regex = regexp.MustCompile(`^/desk/v(\d+)/channels/.*`)

type Auth struct {
	managerSvc account.ManagerFetcher
	logger     log.ContextAwareLogger
}

func NewAuth(managerSvc account.ManagerFetcher, logger log.ContextAwareLogger) *Auth {
	return &Auth{managerSvc: managerSvc, logger: logger}
}

func (a *Auth) Priority() int {
	return 2
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !regex.MatchString(ctx.Request.RequestURI) {
		return
	}

	xAccount := ctx.GetHeader(account.XAccountHeader)
	if len(xAccount) <= 0 {
		ctx.Abort()
		_ = ctx.Error(
			apierr.Unauthorized(errors.New("authorization header is empty")),
		)
		return
	}

	channelID := ctx.Param(channelIDPathParam)
	if len(channelID) <= 0 {
		ctx.Abort()
		_ = ctx.Error(
			apierr.Unauthorized(errors.New("channelID is empty")),
		)
		return
	}

	authenticatedManager, err := a.managerSvc.FetchManager(ctx, channelID, xAccount)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(
			apierr.Unauthorized(errors.New("manager fetch failed")),
		)
		return
	}
	a.logger.Debugw(ctx, "injecting manager principal", "manager", authenticatedManager.ID)
	ctx.Set(ManagerKey, authenticatedManager)
}

func Manager(ctx *gin.Context) account.ManagerPrincipal {
	rawManager, _ := ctx.Get(ManagerKey)
	return rawManager.(account.ManagerPrincipal)
}
