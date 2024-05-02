package command

import (
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/native"
)

type Handler struct {
	serviceName   string
	rbacParser    authgen.Parser
	registerSvc   *command.RegisterSvc
	activationSvc *command.ToggleSvc
}

// TODO: fx 주입시, serviceName ParamTag 추가
func NewHandler(
	serviceName string,
	rbacParser authgen.Parser,
	registerSvc *command.RegisterSvc,
	activationSvc *command.ToggleSvc,
) *Handler {
	return &Handler{
		serviceName:   serviceName,
		rbacParser:    rbacParser,
		registerSvc:   registerSvc,
		activationSvc: activationSvc,
	}
}

func (r Handler) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("registerCommands", r.RegisterCommand)
	registry.Register("toggleCommand", r.ToggleCommand)
}
