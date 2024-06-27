package command

import (
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/native"
)

type Handler struct {
	serviceName    string
	rbacParser     authgen.Parser
	registerSvc    *command.RegisterSvc
	commandRepo    command.CommandRepository
	activationRepo command.ActivationRepository
	activationSvc  command.ActivationSvc
}

func NewHandler(
	serviceName string,
	rbacParser authgen.Parser,
	registerSvc *command.RegisterSvc,
	commandRepo command.CommandRepository,
	activationRepo command.ActivationRepository,
	activationSvc command.ActivationSvc,
) *Handler {
	return &Handler{
		serviceName:    serviceName,
		rbacParser:     rbacParser,
		registerSvc:    registerSvc,
		commandRepo:    commandRepo,
		activationRepo: activationRepo,
		activationSvc:  activationSvc,
	}
}

func (r Handler) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("registerCommands", r.RegisterCommand)
	registry.Register("getCommandChannelActivations", r.GetCommandChannelActivations)
	registry.Register("toggleCommand", r.ToggleCommand)
}
