package hook

import (
	"context"
	"encoding/json"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/hook/model"
	"github.com/channel-io/ch-app-store/internal/hook/svc"
	"github.com/channel-io/ch-app-store/internal/native"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/pkg/errors"
)

type Hook struct {
	serviceName string
	svc         *svc.InstallHookSvc
	toggleSvc   *svc.ToggleHookSvc
	rbacParser  authgen.Parser
}

func NewHook(serviceName string, svc *svc.InstallHookSvc, toggleSvc *svc.ToggleHookSvc, rbacParser authgen.Parser) *Hook {
	return &Hook{
		serviceName: serviceName,
		svc:         svc,
		toggleSvc:   toggleSvc,
		rbacParser:  rbacParser,
	}
}

func (h *Hook) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("registerInstallHook", h.RegisterInstallHook)
	registry.Register("registerToggleHook", h.RegisterToggleHook)
}

func (h *Hook) RegisterInstallHook(ctx context.Context, token native.Token, req native.FunctionRequest) native.FunctionResponse {
	var installHook model.AppInstallHooks
	if err := json.Unmarshal(req.Params, &installHook); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := h.authorizeInstallHook(ctx, token, &installHook); err != nil {
		return native.WrapCommonErr(err)
	}

	hooks, err := h.svc.RegisterHook(ctx, installHook.AppID, &installHook)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	res, err := json.Marshal(hooks)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(res)
}

func (h *Hook) RegisterToggleHook(ctx context.Context, token native.Token, req native.FunctionRequest) native.FunctionResponse {
	var hook model.CommandToggleHooks
	if err := json.Unmarshal(req.Params, &hook); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := h.authorizeToggleHook(ctx, token, &hook); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := h.toggleSvc.RegisterHook(ctx, &hook); err != nil {
		return native.WrapCommonErr(err)
	}

	res, err := json.Marshal(hook)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(res)
}

const (
	installHookAction = "registerInstallHook"
	toggleHookAction  = "registerToggleHook"
	appScope          = "app"
)

func (h *Hook) authorizeInstallHook(ctx context.Context, token native.Token, installHook *model.AppInstallHooks) error {
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

func (h *Hook) authorizeToggleHook(ctx context.Context, token native.Token, toggleHook *model.CommandToggleHooks) error {
	parsedRbac, err := h.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(h.serviceName), toggleHookAction) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope: {toggleHook.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
