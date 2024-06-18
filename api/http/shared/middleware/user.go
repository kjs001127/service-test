package middleware

import (
	"github.com/channel-io/ch-app-store/api/http/front/middleware"
	"github.com/channel-io/ch-app-store/internal/auth/principal/session"

	"github.com/gin-gonic/gin"
)

const (
	UserRequesterKey = "UserRequester"
)

type UserRequest struct{}

func NewUserRequest() *Request {
	return &Request{}
}

func (r *UserRequest) Priority() int {
	return 1
}

func (r *UserRequest) Handle(ctx *gin.Context) {
	requester := Requester(ctx)
	user := middleware.User(ctx)

	userRequester := session.UserRequester{
		Requester:     requester,
		UserPrincipal: user,
	}

	ctx.Set(UserRequesterKey, userRequester)
}

func UserRequester(ctx *gin.Context) session.UserRequester {
	rawRequester, _ := ctx.Get(RequesterKey)

	return rawRequester.(session.UserRequester)
}
