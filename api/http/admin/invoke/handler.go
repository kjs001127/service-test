package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	brief "github.com/channel-io/ch-app-store/internal/brief/domain"
	native "github.com/channel-io/ch-app-store/internal/native/domain"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker       *app.TypedInvoker[json.RawMessage, json.RawMessage]
	briefInvoker  *brief.Invoker
	nativeInvoker *native.NativeFunctionInvoker
}

func NewHandler(
	invoker *app.TypedInvoker[json.RawMessage, json.RawMessage],
	briefInvoker *brief.Invoker,
	nativeInvoker *native.NativeFunctionInvoker,
) *Handler {
	return &Handler{invoker: invoker, briefInvoker: briefInvoker, nativeInvoker: nativeInvoker}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/admin/apps/:id/functions", h.invoke)
	router.PUT("/admin/brief", h.brief)
	router.PUT("/admin/native/functions", h.invokeNative)
}
