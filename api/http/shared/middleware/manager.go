package middleware

import (
	"github.com/channel-io/ch-app-store/api/http/desk/middleware"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

	"github.com/gin-gonic/gin"
)

const (
	ManagerRequesterKey = "ManagerRequester"
)

type ManagerRequest struct{}

func NewManagerRequest() *Request {
	return &Request{}
}

func (r *ManagerRequest) Priority() int {
	return 1
}

func (r *ManagerRequest) Handle(ctx *gin.Context) {
	requester := Requester(ctx)
	manager := middleware.Manager(ctx)

	managerRequester := account.ManagerRequester{
		Requester:        requester,
		ManagerPrincipal: manager,
	}

	ctx.Set(ManagerRequesterKey, managerRequester)
}

func ManagerRequester(ctx *gin.Context) account.ManagerRequester {
	rawRequester, _ := ctx.Get(ManagerRequesterKey)

	return rawRequester.(account.ManagerRequester)
}
