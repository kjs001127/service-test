package controller

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/channel-io/ch-app-store/api/gintool"
)

//go:embed resources/*
var resources embed.FS

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	static, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	router.StaticFS("/public/v1/wam-controller", http.FS(static))
}
