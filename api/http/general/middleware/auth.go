package middleware

import (
	"errors"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/auth/general"
)

const RBACKey = "rbacKey"

type Auth struct {
	parser general.Parser
}

func NewAuth(parser general.Parser) *Auth {
	return &Auth{parser: parser}
}

func (a *Auth) Handle(ctx *gin.Context) {
	xAccessToken := ctx.GetHeader(general.Header())
	if xAccessToken == "" {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("authorization header is empty")))
		return
	}

	rbac, err := a.parser.Parse(ctx, xAccessToken)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(err)
		return
	}

	ctx.Set(RBACKey, rbac)
}
