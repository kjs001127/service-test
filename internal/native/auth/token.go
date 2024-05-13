package auth

import (
	"context"
	"encoding/json"

	"github.com/channel-io/ch-app-store/internal/approle/svc"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
)

type Request struct {
	Secret    string  `json:"secret"`
	ChannelID *string `json:"channelId,omitempty"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type TokenIssueHandler struct {
	svc *svc.TokenSvc
}

func NewTokenIssueHandler(svc *svc.TokenSvc) *TokenIssueHandler {
	return &TokenIssueHandler{
		svc: svc,
	}
}

func (c *TokenIssueHandler) RegisterTo(registry native.FunctionRegistry) {
	registry.Register("issueToken", c.IssueToken)
	registry.Register("refreshToken", c.refreshToken)
}

func (c *TokenIssueHandler) refreshToken(ctx context.Context, token native.Token, request native.FunctionRequest) native.FunctionResponse {
	var req RefreshRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	resp, err := c.svc.RefreshToken(ctx, req.RefreshToken)

	marshaled, err := json.Marshal(&resp)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(marshaled)
}

func (c *TokenIssueHandler) IssueToken(ctx context.Context, token native.Token, request native.FunctionRequest) native.FunctionResponse {
	var req Request
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	resp, err := c.doIssueToken(ctx, req)

	marshaled, err := json.Marshal(&resp)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	return native.ResultSuccess(marshaled)
}

func (c *TokenIssueHandler) doIssueToken(ctx context.Context, req Request) (authgen.IssueResponse, error) {
	if req.ChannelID != nil {
		return c.svc.IssueChannelToken(ctx, *req.ChannelID, req.Secret)
	} else {
		return c.svc.IssueAppToken(ctx, req.Secret)
	}
}
