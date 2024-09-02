package widget

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/appwidget/svc"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/localapi/widget/action/public"
)

func (r *Handler) RegisterAppWidgets(
	ctx context.Context,
	token native.Token,
	request native.FunctionRequest,
) native.FunctionResponse {
	var req svc.AppWidgetRegisterRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.authorizeReg(ctx, token, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.registerSvc.Register(ctx, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	return native.Empty()
}

const (
	appScope = "app"
)

func (r *Handler) authorizeReg(ctx context.Context, token native.Token, req *svc.AppWidgetRegisterRequest) error {
	parsedRbac, err := r.parser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(r.serviceName), public.RegisterAppWidgets) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope: {req.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
