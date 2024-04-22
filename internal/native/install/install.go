package install

import (
	"context"
	"encoding/json"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
)

type Checker struct {
	serviceName string
	svc         *svc.AppInstallQuerySvc
	rbacParser  authgen.Parser
}

func NewChecker(serviceName string, svc *svc.AppInstallQuerySvc, rbacParser authgen.Parser) *Checker {
	return &Checker{
		serviceName: serviceName,
		svc:         svc,
		rbacParser:  rbacParser,
	}
}

func (c *Checker) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("checkInstall", c.CheckInstall)
}

func (c *Checker) CheckInstall(ctx context.Context, token native.Token, request native.FunctionRequest) native.FunctionResponse {
	var req CheckRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	if err := c.authorize(ctx, token, req); err != nil {
		return native.WrapCommonErr(err)
	}

	installed, err := c.svc.CheckInstall(ctx, model.InstallationID{
		ChannelID: req.ChannelID,
		AppID:     req.AppID,
	})
	if err != nil {
		return native.WrapCommonErr(err)
	}

	marshaled, err := json.Marshal(&CheckResponse{Installed: installed})
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(marshaled)
}

const (
	checkInstallAction = "checkInstall"

	channelScope = "channel"
)

func (c *Checker) authorize(ctx context.Context, token native.Token, req CheckRequest) error {
	parsedRbac, err := c.rbacParser.Parse(ctx, token.Value)
	if err != nil {
		return err
	}

	if !parsedRbac.CheckAction(authgen.Service(c.serviceName), checkInstallAction) {
		return apierr.Unauthorized(errors.New("service, action check fail"))
	}

	if !parsedRbac.CheckScopes(authgen.Scopes{
		channelScope: {req.ChannelID},
	}) {
		return apierr.Unauthorized(errors.New("scope check fail"))
	}

	return nil
}

type CheckRequest struct {
	ChannelID string `json:"channelId"`
	AppID     string `json:"appId"`
}

type CheckResponse struct {
	Installed bool `json:"installed"`
}
