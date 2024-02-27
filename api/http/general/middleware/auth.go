package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/channel-io/go-lib/pkg/log"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/auth/general"
)

const (
	generalScopePrefix = "/general"
	RBACKey            = "rbacKey"
)

type Auth struct {
	parser general.Parser
	logger *log.ChannelLogger
}

func NewAuth(parser general.Parser, logger *log.ChannelLogger) *Auth {
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

	a.logger.Debugw("parsed rbac", "token", rbac, "request", ctx.Request.RequestURI) // TODO: remove

	ctx.Set(RBACKey, rbac)
}
