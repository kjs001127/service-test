package doc

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/channel-io/ch-app-store/api/gintool"
	_ "github.com/channel-io/ch-app-store/api/http/swagger"
)

type Handler struct {
	path string
	name string
}

func NewHandler(path string, name string) *Handler {
	return &Handler{path: path, name: name}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.GET(h.path, ginSwagger.WrapHandler(
		swaggerFiles.NewHandler(),
		ginSwagger.InstanceName(h.name),
	))
}
