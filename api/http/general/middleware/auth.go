package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
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

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, generalScopePrefix) {
		return
	}

	xAccessToken := ctx.GetHeader(general.Header())
	if len(xAccessToken) <= 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("authorization header is empty"))
		return
	}

	rbac, err := a.parser.Parse(ctx, xAccessToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(fmt.Errorf("parsing token. cause: %w", err)))
		return
	}

	a.logger.Debugw(ctx, "parsed rbac", "id", rbac.ID, "type", rbac.Type)

	ctx.Set(rbacKey, rbac)
}

func RBAC(ctx *gin.Context) general.ParsedRBACToken {
	rawRbac, _ := ctx.Get(rbacKey)
	return rawRbac.(general.ParsedRBACToken)
}
