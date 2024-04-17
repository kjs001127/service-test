package command

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	command "github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/native"
)

type RegisterHandler struct {
	rbacParser  authgen.Parser
	registerSvc *command.RegisterSvc
}

func NewRegisterHandler(rbacParser authgen.Parser, registerSvc *command.RegisterSvc) *RegisterHandler {
	return &RegisterHandler{rbacParser: rbacParser, registerSvc: registerSvc}
}

func (r RegisterHandler) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("registerCommands", r.RegisterCommand)
}

func (r *RegisterHandler) RegisterCommand(
	ctx context.Context,
	token native.Token,
	request native.FunctionRequest,
) native.FunctionResponse {
	var req command.CommandRegisterRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.authorize(ctx, token, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.registerSvc.Register(ctx, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	return native.Empty()
}

const (
	appStoreService  = "api.app-store.channel.io"
	registerCommands = "registerCommands"

	appScope = "app"
)

func (r *RegisterHandler) authorize(ctx context.Context, token native.Token, req *command.CommandRegisterRequest) error {
	parsedRbac, err := r.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(appStoreService, registerCommands) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope: {req.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
