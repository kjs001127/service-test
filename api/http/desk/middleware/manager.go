package middleware

import (
	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"

	"github.com/gin-gonic/gin"
)

const (
	ManagerRequesterKey = "ManagerRequester"
)

type ManagerRequest struct{}

func NewManagerRequest() *ManagerRequest {
	return &ManagerRequest{}
}

func (r *ManagerRequest) Priority() int {
	return 3
}

func (r *ManagerRequest) Handle(ctx *gin.Context) {
	if !regex.MatchString(ctx.Request.RequestURI) {
		return
	}
	requester := middleware.Requester(ctx)
	manager := Manager(ctx)

	managerRequester := desk.ManagerRequester{
		Requester:        requester,
		ManagerPrincipal: manager,
	}

	ctx.Set(ManagerRequesterKey, managerRequester)
}

func ManagerRequester(ctx *gin.Context) desk.ManagerRequester {
	rawRequester, _ := ctx.Get(ManagerRequesterKey)

	return rawRequester.(desk.ManagerRequester)
}
