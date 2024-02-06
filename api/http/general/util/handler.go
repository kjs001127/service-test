package util

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

func (h *Handler) ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
