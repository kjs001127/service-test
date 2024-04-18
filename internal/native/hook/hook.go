package hook

import (
	"context"
	"encoding/json"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/installhook/model"
	"github.com/channel-io/ch-app-store/internal/installhook/svc"
	"github.com/channel-io/ch-app-store/internal/native"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type Hook struct {
	serviceName string
	svc         *svc.HookSvc
	rbacParser  authgen.Parser
}

func NewHook(serviceName string, svc *svc.HookSvc, rbacParser authgen.Parser) *Hook {
	return &Hook{
		serviceName: serviceName,
		svc:         svc,
		rbacParser:  rbacParser,
	}
}

func (h *Hook) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("registerHook", h.RegisterHook)
}

func (h *Hook) RegisterHook(ctx context.Context, token native.Token, req native.FunctionRequest) native.FunctionResponse {
	var installHook model.AppInstallHooks
	if err := json.Unmarshal(req.Params, &installHook); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := h.authorize(ctx, token, &installHook); err != nil {
		return native.WrapCommonErr(err)
	}

	hooks, err := h.svc.Upsert(ctx, installHook.AppID, &installHook)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	res, err := json.Marshal(hooks)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(res)
}

const (
	installHookAction = "installHook"

	appScope = "app"
)

func (h *Hook) authorize(ctx context.Context, token native.Token, installHook *model.AppInstallHooks) error {
	parsedRbac, err := h.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(h.serviceName), installHookAction) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope: {installHook.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
