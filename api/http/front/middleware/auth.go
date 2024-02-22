package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/channel-io/go-lib/pkg/log"
	"github.com/gin-gonic/gin"
	errors2 "github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

const UserKey = "user"

type Auth struct {
	userSvc session.UserFetcher
	logger  *log.ChannelLogger
}

func NewAuth(managerSvc session.UserFetcher, logger *log.ChannelLogger) *Auth {
	return &Auth{userSvc: managerSvc, logger: logger}
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
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(errors2.Wrap(err, "middleware user fetch fail")))
		return
	}

	channelID := ctx.Param("channelID")
	if channelID != "" && channelID != user.ChannelID {
		a.logger.Warnw("channelID doest not match jwt", "path", channelID, "jwt", user.ChannelID)
	}

	a.logger.Debugw("injecting user principal", "request", ctx.Request.RequestURI, "user", user)

	ctx.Set(UserKey, user)
}
