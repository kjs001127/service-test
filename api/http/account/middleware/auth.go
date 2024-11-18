package middleware

import (
	"errors"
	"strings"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	accountScope = "/desk/account"
	AccountKey   = "Account"
)

type Auth struct {
	parser desk.Parser
	logger log.ContextAwareLogger
}

func NewAuth(parser desk.Parser, logger log.ContextAwareLogger) *Auth {
	return &Auth{parser: parser, logger: logger}
}

func (a *Auth) Priority() int {
	return 2
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, accountScope) {
		return
	}

	xAccount := ctx.GetHeader(desk.XAccountHeader)
	if len(xAccount) <= 0 {
		ctx.Abort()
		_ = ctx.Error(
			apierr.Unauthorized(errors.New("authorization header is empty")),
		)
		return
	}

	acc, err := a.parser.ParseAccount(ctx, xAccount)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(
			apierr.Unauthorized(errors.New("account fetch failed")),
		)
		return
	}

	a.logger.Debugw(ctx, "injecting account principal", "account", acc.Account)
	ctx.Set(AccountKey, acc)
}

func Account(ctx *gin.Context) desk.Principal {
	rawAccount, _ := ctx.Get(AccountKey)
	return rawAccount.(desk.Principal)
}
