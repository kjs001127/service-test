package util

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/channel-io/ch-app-store/api/http/swagger"

	"github.com/channel-io/ch-app-store/api/gintool"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET("/ping", h.ping)
}

// ping godoc
//
//	@Summary	send ping to server
//	@Tags		Util
//	@Success	200	{string}	string
//	@Router		/ping [get]
func (h *Handler) ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
