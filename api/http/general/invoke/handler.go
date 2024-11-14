package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	authgen "github.com/channel-io/ch-app-store/internal/shared/general"
	"github.com/channel-io/ch-app-store/internal/native"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker       app.TypedInvoker[json.RawMessage, json.RawMessage]
	nativeInvoker *native.FunctionInvoker
	parser        authgen.Parser
}

func NewHandler(invoker app.TypedInvoker[json.RawMessage, json.RawMessage], nativeInvoker *native.FunctionInvoker, parser authgen.Parser) *Handler {
	return &Handler{invoker: invoker, nativeInvoker: nativeInvoker, parser: parser}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/general/v1/apps/:appID/functions", h.invoke)
	router.PUT("/general/v1/native/functions", h.invokeNative)
}
