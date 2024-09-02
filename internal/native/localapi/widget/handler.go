package widget

import (
	"github.com/channel-io/ch-app-store/internal/appwidget/svc"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/localapi/widget/action/public"
)

type Handler struct {
	serviceName string
	registerSvc svc.RegisterSvc
	parser      authgen.Parser
}

func NewHandler(serviceName string, registerSvc svc.RegisterSvc, parser authgen.Parser) *Handler {
	return &Handler{serviceName: serviceName, registerSvc: registerSvc, parser: parser}
}

func (r *Handler) RegisterTo(registry native.FunctionRegistry) {
	registry.Register(public.RegisterAppWidgets, r.RegisterAppWidgets)
}
