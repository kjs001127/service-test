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

	xAccessTokenHeader  = "x-access-token"
	authorizationHeader = "Authorization"
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

func tokenFrom(ctx *gin.Context) (string, bool) {
	xAccessToken := ctx.GetHeader(xAccessTokenHeader)
	if len(xAccessToken) <= 0 {
		return "", false
	}

	rawAuthHeader := ctx.GetHeader(authorizationHeader)
	if len(rawAuthHeader) <= 0 {
		return "", false
	}

	_, token, ok := strings.Cut(rawAuthHeader, " ")
	if !ok {
		return "", false
	}

	return token, true
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, generalScopePrefix) {
		return
	}

	token, exists := tokenFrom(ctx)
	if !exists {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("authorization header is empty")))
		return
	}

	rbac, err := a.parser.Parse(ctx, token)
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
