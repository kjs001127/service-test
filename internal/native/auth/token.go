package auth

import (
	"context"
	"encoding/json"

	"github.com/channel-io/ch-app-store/internal/approle/svc"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	"github.com/channel-io/ch-app-store/internal/native"
	"github.com/channel-io/ch-app-store/internal/native/auth/action/public"
)

type Request struct {
	Secret    string  `json:"secret"`
	ChannelID *string `json:"channelId,omitempty"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

func fromIssueResponse(response authgen.IssueResponse) *Response {
	return &Response{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		ExpiresIn:    response.ExpiresIn,
	}
}

type TokenIssueHandler struct {
	svc svc.TokenSvc
}

func NewTokenIssueHandler(svc svc.TokenSvc) *TokenIssueHandler {
	return &TokenIssueHandler{
		svc: svc,
	}
}

func (c *TokenIssueHandler) RegisterTo(registry native.FunctionRegistry) {
	registry.Register(public.IssueToken, c.IssueToken)
	registry.Register(public.RefreshToken, c.refreshToken)
}

func (c *TokenIssueHandler) refreshToken(ctx context.Context, token native.Token, request native.FunctionRequest) native.FunctionResponse {
	var req RefreshRequest
	if err := json.Unmarshal(request.Params, &req); err != nil {
		return native.WrapCommonErr(err)
	}

	resp, err := c.svc.RefreshToken(ctx, req.RefreshToken)

	marshaled, err := json.Marshal(fromIssueResponse(resp))
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
	if err != nil {
		return native.WrapCommonErr(err)
	}

	marshaled, err := json.Marshal(fromIssueResponse(resp))
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
