package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/native/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker       *app.TypedInvoker[json.RawMessage, json.RawMessage]
	nativeInvoker *domain.NativeFunctionInvoker
}

func NewHandler(invoker *app.TypedInvoker[json.RawMessage, json.RawMessage], nativeInvoker *domain.NativeFunctionInvoker) *Handler {
	return &Handler{invoker: invoker, nativeInvoker: nativeInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/general/v1/apps/:appID/functions", h.invoke)
	//router.PUT("/general/v1/native/functions", h.invokeNative)
}
