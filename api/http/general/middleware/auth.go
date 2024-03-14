package middleware

import (
	"errors"
	"strings"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	wraperr "github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	generalScopePrefix = "/general"
	rbacKey            = "rbacKey"
)

type Auth struct {
	parser general.Parser
	logger log.ContextAwareLogger
}

func NewAuth(parser general.Parser, logger log.ContextAwareLogger) *Auth {
	return &Auth{parser: parser, logger: logger}
}

func (a *Auth) Priority() int {
	return 2
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, generalScopePrefix) {
		return
	}

	xAccessToken := ctx.GetHeader(general.Header())
	if len(xAccessToken) <= 0 {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("authorization header is empty")))
		return
	}

	rbac, err := a.parser.Parse(ctx, xAccessToken)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(wraperr.Wrap(err, "parsing x-access-token fail")))
		return
	}

	a.logger.Debugw(ctx, "parsed rbac", "id", rbac.ID, "type", rbac.Type)

	ctx.Set(rbacKey, rbac)
}

func RBAC(ctx *gin.Context) general.ParsedRBACToken {
	rawRbac, _ := ctx.Get(rbacKey)
	return rawRbac.(general.ParsedRBACToken)
}
