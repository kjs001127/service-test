package middleware

import (
	"strings"

	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/internal/shared/principal/front"

	"github.com/gin-gonic/gin"
)

const (
	UserRequesterKey = "UserRequester"
)

type UserRequest struct{}

func NewUserRequest() *UserRequest {
	return &UserRequest{}
}

func (r *UserRequest) Priority() int {
	return 3
}

func (r *UserRequest) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, FrontPathPrefix) {
		return
	}
	requester := middleware.Requester(ctx)
	user := User(ctx)

	userRequester := front.UserRequester{
		Requester:     requester,
		UserPrincipal: user,
	}

	ctx.Set(UserRequesterKey, userRequester)
}

func UserRequester(ctx *gin.Context) front.UserRequester {
	rawRequester, _ := ctx.Get(UserRequesterKey)

	return rawRequester.(front.UserRequester)
}
