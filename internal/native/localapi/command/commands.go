package command

import (
	"context"
	"encoding/json"
	"errors"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/localapi/command/action/private"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

type GetCommandChannelActivationsRequest struct {
	AppID     string `json:"appId"`
	ChannelID string `json:"channelId"`
}

type GetCommandChannelActivationsResponse struct {
	Commands    []*model.Command  `json:"commands"`
	Activations model.Activations `json:"activations"`
}

func (r *Handler) GetCommandChannelActivations(
	ctx context.Context,
	token native.Token,
	request native.FunctionRequest,
) native.FunctionResponse {
	var req GetCommandChannelActivationsRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.authorizeGet(ctx, token, req); err != nil {
		return native.WrapCommonErr(err)
	}

	cmds, err := r.commandRepo.FetchAllByAppID(ctx, req.AppID)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	activations, err := r.activationRepo.FetchByChannelIDAndCmdIDs(ctx, req.ChannelID, idsOfCmds(cmds))
	if err != nil {
		return native.WrapCommonErr(err)
	}

	resp, err := json.Marshal(GetCommandChannelActivationsResponse{cmds, activations})
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(resp)
}

func idsOfCmds(cmds []*model.Command) []string {
	ret := make([]string, 0, len(cmds))
	for _, cmd := range cmds {
		ret = append(ret, cmd.ID)
	}
	return ret
}

func (r *Handler) authorizeGet(ctx context.Context, token native.Token, req GetCommandChannelActivationsRequest) error {
	parsedRbac, err := r.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(r.serviceName), private.GetCommandChannelActivations) {
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
