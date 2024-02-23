package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/channel-io/go-lib/pkg/log"
	"github.com/gin-gonic/gin"
	wraperr "github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"
)

const (
	PathParamChannelID = "channelID"
	FrontPathPrefix    = "/front"
	UserKey            = "user"
)

type Auth struct {
	userSvc session.UserFetcher
	logger  *log.ChannelLogger
}

func NewAuth(managerSvc session.UserFetcher, logger *log.ChannelLogger) *Auth {
	return &Auth{userSvc: managerSvc, logger: logger}
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, FrontPathPrefix) {
		return
	}

	xSession := ctx.GetHeader(session.XSessionHeader)
	if len(xSession) <= 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("authorization header is empty"))
		return
	}

	user, err := a.userSvc.FetchUser(ctx, xSession)
	if err != nil {
		wrapped := wraperr.Wrap(err, "middleware user fetch fail")
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(wrapped))
		return
	}

	channelID := ctx.Param(PathParamChannelID)
	if len(channelID) >= 0 && channelID != user.ChannelID {
		a.logger.Warnw("channelID doest not match jwt", "path", channelID, "jwt", user.ChannelID)
		err := fmt.Errorf("user auth failed, channelId does not match")
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnauthorizedError(err))
		return
	}

	a.logger.Debugw("injecting user principal", "request", ctx.Request.RequestURI, "user", user)

	ctx.Set(UserKey, user)
}
