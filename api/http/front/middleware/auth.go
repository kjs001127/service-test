package middleware

import (
	"errors"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/auth/principal"
	"github.com/channel-io/ch-app-store/auth/principal/session"
)

const UserKey = "user"

type Auth struct {
	userSvc session.UserFetcher
}

func NewAuth(managerSvc session.UserFetcher) *Auth {
	return &Auth{userSvc: managerSvc}
}

func (a *Auth) Handle(ctx *gin.Context) {
	xSession := ctx.GetHeader(principal.TokenTypeSession.Header())
	if xSession == "" {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("authorization header is empty")))
		return
	}

	user, err := a.userSvc.FetchUser(ctx, xSession)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(err)
		return
	}

	ctx.Set(UserKey, user)
}
