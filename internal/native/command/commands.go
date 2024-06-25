package command

import (
	"context"
	"encoding/json"
	"errors"

	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
)

type GetCommandRequest struct {
	AppID string `json:"appId"`
}

const (
	getCommands = "getCommands"
)

func (r *Handler) GetCommands(
	ctx context.Context,
	token native.Token,
	request native.FunctionRequest,
) native.FunctionResponse {
	var req GetCommandRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := r.authorizeGet(ctx, token, req); err != nil {
		return native.WrapCommonErr(err)
	}

	res, err := r.commandRepo.FetchAllByAppID(ctx, req.AppID)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	resp, err := json.Marshal(res)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(resp)
}

func (r *Handler) authorizeGet(ctx context.Context, token native.Token, req GetCommandRequest) error {
	parsedRbac, err := r.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(r.serviceName), getCommands) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		appScope: {req.AppID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}
