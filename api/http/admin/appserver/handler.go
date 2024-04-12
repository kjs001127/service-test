package appserver

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	native "github.com/channel-io/ch-app-store/internal/native"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker       *app.TypedInvoker[json.RawMessage, json.RawMessage]
	nativeInvoker *native.FunctionInvoker
}

func NewHandler(
	invoker *app.TypedInvoker[json.RawMessage, json.RawMessage],
	nativeInvoker *native.FunctionInvoker,
) *Handler {
	return &Handler{invoker: invoker, nativeInvoker: nativeInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/admin/apps/:id/functions", h.invoke)
	router.PUT("/admin/native/functions", h.invokeNative)
}
