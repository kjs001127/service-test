package controller

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/channel-io/ch-app-store/api/gintool"
)

//go:embed resources/*
var resources embed.FS

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	env string
}

func NewHandler(env string) *Handler {
	return &Handler{env: env}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	static, err := fs.Sub(resources, fmt.Sprintf("resources/%s", h.env))
	if err != nil {
		panic(err)
	}
	router.StaticFS("/public/v1/wam-controller", http.FS(static))
	router.StaticFS("/public/wam-controller/v1", http.FS(static))
}
