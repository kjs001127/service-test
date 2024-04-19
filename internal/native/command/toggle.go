package command

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
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

	if err := r.activationSvc.Toggle(ctx, model.InstallationID{
		AppID:     req.AppID,
		ChannelID: req.ChannelID,
	}, req.CommandEnabled); err != nil {
		return native.WrapCommonErr(err)
	}

	return native.Empty()
}

type ToggleCommandRequest struct {
	ChannelID      string `json:"channelId"`
	AppID          string `json:"appID"`
	CommandEnabled bool   `json:"commandEnabled"`
}

const (
	toggleCommand = "toggleCommand"
	channelScope  = "channel"
)

func (r *Handler) authorizeToggle(ctx context.Context, token native.Token, req ToggleCommandRequest) error {
	parsedRbac, err := r.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(r.serviceName), toggleCommand) {
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
