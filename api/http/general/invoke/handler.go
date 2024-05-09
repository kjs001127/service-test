package invoke

import (
	"encoding/json"

	"github.com/channel-io/ch-app-store/api/gintool"
	"github.com/channel-io/ch-app-store/api/http/general/middleware"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/native"
)

var _ gintool.RouteRegistrant = (*Handler)(nil)

type Handler struct {
	invoker       *app.TypedInvoker[json.RawMessage, json.RawMessage]
	nativeInvoker *native.FunctionInvoker
	rbacAuth      middleware.Auth
}

func NewHandler(invoker *app.TypedInvoker[json.RawMessage, json.RawMessage], nativeInvoker *native.FunctionInvoker, rbacAuth middleware.Auth) *Handler {
	return &Handler{invoker: invoker, nativeInvoker: nativeInvoker, rbacAuth: rbacAuth}
}

func (h *Handler) RegisterRoutes(router gintool.Router) {
	router.PUT("/general/v1/apps/:appID/functions", h.rbacAuth.Handle, h.invoke)
	router.PUT("/general/v1/native/functions", h.rbacAuth.Handle, h.invokeNative)
}
