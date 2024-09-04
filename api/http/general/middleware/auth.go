package middleware

import (
	"strings"

	genauth "github.com/channel-io/ch-app-store/internal/auth/general"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	generalPathPrefix = "/general/v1"

	RbacTokenKey = "RbacToken"
)

type Auth struct {
	parser genauth.Parser
}

func NewAuth(parser genauth.Parser) *Auth {
	return &Auth{parser: parser}
}

func (r *Auth) Priority() int {
	return 3
}

func (r *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, generalPathPrefix) {
		return
	}

	token, err := r.rbacFrom(ctx)
	if err != nil {
		return
	}

	ctx.Set(RbacTokenKey, token)
}

func (r *Auth) rbacFrom(ctx *gin.Context) (genauth.ParsedRBACToken, error) {
	token, exists := tokenFrom(ctx)

	if !exists {
		return genauth.ParsedRBACToken{}, apierr.Unauthorized(errors.New("token not found"))
	}

	parsed, err := r.parser.Parse(ctx, token)
	if err != nil {
		return genauth.ParsedRBACToken{}, apierr.Unauthorized(errors.New("token not valid"))
	}

	return parsed, nil
}

const (
	xAccessTokenHeader  = "x-access-token"
	authorizationHeader = "Authorization"
)

func tokenFrom(ctx *gin.Context) (string, bool) {
	xAccessToken := ctx.GetHeader(xAccessTokenHeader)
	if len(xAccessToken) > 0 {
		return xAccessToken, true
	}

	rawAuthHeader := ctx.GetHeader(authorizationHeader)
	if len(rawAuthHeader) > 0 {
		_, token, ok := strings.Cut(rawAuthHeader, " ")
		return token, ok

	}
	return "", false
}

func Rbac(ctx *gin.Context) (genauth.ParsedRBACToken, bool) {
	rawRbac, exists := ctx.Get(RbacTokenKey)
	if !exists {
		return genauth.ParsedRBACToken{}, false
	}

	return rawRbac.(genauth.ParsedRBACToken), true
}
