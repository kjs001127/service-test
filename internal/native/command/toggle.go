package command

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/command/svc"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/command/action/private"
)

func (r *Handler) ToggleCommand(
	ctx context.Context,
	token native.Token,
	request native.FunctionRequest,
) native.FunctionResponse {
	var req ToggleCommandRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.authorizeToggle(ctx, token, req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.activationSvc.ToggleByKey(ctx, svc.ToggleRequest{
		Enabled:   req.CommandEnabled,
		ChannelID: req.ChannelID,
		Command:   req.CommandKey,
	}); err != nil {
		return native.WrapCommonErr(err)
	}

	return native.Empty()
}

type ToggleCommandRequest struct {
	cmd.CommandKey
	ChannelID      string `json:"channelId"`
	CommandEnabled bool   `json:"commandEnabled"`
}

const (
	channelScope = "channel"
)

func (r *Handler) authorizeToggle(ctx context.Context, token native.Token, req ToggleCommandRequest) error {
	parsedRbac, err := r.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(r.serviceName), private.ToggleCommand) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope:     {req.AppID},
		channelScope: {req.ChannelID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
