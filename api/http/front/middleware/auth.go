package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

const UserKey = "user"

type Auth struct {
	userSvc session.UserFetcher
}

func NewAuth(managerSvc session.UserFetcher) *Auth {
	return &Auth{userSvc: managerSvc}
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, "/front") {
		return
	}

	xSession := ctx.GetHeader(session.XSessionHeader)
	if xSession == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("authorization header is empty"))
		return
	}

	user, err := a.userSvc.FetchUser(ctx, xSession)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.Set(UserKey, user)
}
