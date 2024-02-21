package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/auth/general"
)

const RBACKey = "rbacKey"

type Auth struct {
	parser general.Parser
}

func NewAuth(parser general.Parser) *Auth {
	return &Auth{parser: parser}
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, "/general") {
		return
	}

	xAccessToken := ctx.GetHeader(general.Header())
	if xAccessToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("authorization header is empty"))
		return
	}

	rbac, err := a.parser.Parse(ctx, xAccessToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Set(RBACKey, rbac)
}
