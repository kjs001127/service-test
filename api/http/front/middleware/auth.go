package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	wraperr "github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/shared/principal/session"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	PathParamChannelID = "channelID"
	FrontPathPrefix    = "/front"
	UserKey            = "user"
)

type Auth struct {
	userSvc session.UserFetcher
	logger  log.ContextAwareLogger
}

func NewAuth(managerSvc session.UserFetcher, logger log.ContextAwareLogger) *Auth {
	return &Auth{userSvc: managerSvc, logger: logger}
}

func (a *Auth) Priority() int {
	return 2
}

func (a *Auth) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, FrontPathPrefix) {
		return
	}

	xSession := ctx.GetHeader(session.XSessionHeader)
	if len(xSession) <= 0 {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(errors.New("x-session header not found")))
		return
	}

	user, err := a.userSvc.FetchUser(ctx, xSession)
	if err != nil {
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(wraperr.Wrap(err, "middleware user fetch fail")))
		return
	}

	channelID := ctx.Param(PathParamChannelID)
	if len(channelID) >= 0 && channelID != user.ChannelID {
		err := fmt.Errorf("channelID doest not match jwt, Path: %s, User: %s", channelID, user.ChannelID)
		ctx.Abort()
		_ = ctx.Error(apierr.Unauthorized(err))
		return
	}

	a.logger.Debugw(ctx, "injecting user principal", "user", user.ID)

	ctx.Set(UserKey, user)
}

func User(ctx *gin.Context) session.UserPrincipal {
	rawUser, _ := ctx.Get(UserKey)
	return rawUser.(session.UserPrincipal)
}
