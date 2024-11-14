package widget

import (
	authgen "github.com/channel-io/ch-app-store/internal/shared/general"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/localapi/widget/action/private"
	"github.com/channel-io/ch-app-store/internal/widget/svc"
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
	registry.Register(private.RegisterAppWidgets, r.RegisterAppWidgets)
}
