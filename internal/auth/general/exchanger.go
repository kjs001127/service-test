package general

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/internal/auth/principal"
)

const (
	issueTokenVersion = "v1"
	issueToken        = "/api/admin/auth/" + issueTokenVersion + "/token"
)

type RBACExchanger struct {
	cli     *resty.Client
	parser  Parser
	authURL string
}

func NewRBACExchanger(cli *resty.Client, parser Parser, authURL string) *RBACExchanger {
	return &RBACExchanger{cli: cli, parser: parser, authURL: authURL}
}

func (e *RBACExchanger) Refresh(
	ctx context.Context,
	refreshToken string,
) (IssueResponse, error) {
	r := e.cli.R()
	r.SetContext(ctx)

	r.
		SetQueryParam("grant_type", "refresh_token").
		SetQueryParam("refresh_token", refreshToken)

	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	resp, err := r.Post(e.authURL + issueToken)
	if err != nil {
		return IssueResponse{}, err
	}

	var unmarshalled IssueResponse
	if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
		return IssueResponse{}, err
	} else if resp.IsError() {
		return IssueResponse{}, apierr.Unauthorized(fmt.Errorf("token fetch fail, response: %s", resp.Body()))
	}

	return unmarshalled, nil
}

func (e *RBACExchanger) ExchangeWithClientSecret(
	ctx context.Context,
	clientID string,
	clientSecret string,
	scopes Scopes,
) (IssueResponse, error) {
	r := e.cli.R()
	r.SetContext(ctx)

	values := make(url.Values)
	for key, vals := range scopes {
		for _, val := range vals {
			values.Add("scope", fmt.Sprintf("%s-%s", key, val))
		}
	}
	r.
		SetQueryParamsFromValues(values).
		SetQueryParam("grant_type", "client_credentials").
		SetQueryParam("client_secret", clientSecret).
		SetQueryParam("client_id", clientID)

	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	resp, err := r.Post(e.authURL + issueToken)
	if err != nil {
		return IssueResponse{}, err
	} else if resp.IsError() {
		return IssueResponse{}, apierr.Unauthorized(fmt.Errorf("token fetch fail, response: %s", resp.Body()))
	}

	var unmarshalled IssueResponse
	if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
		return IssueResponse{}, err
	}

	return unmarshalled, nil
}

func (e *RBACExchanger) ExchangeWithPrincipal(
	ctx context.Context,
	token principal.Token,
	scopes Scopes,
	clientID string,
) (IssueResponse, error) {
	r := e.cli.R()
	r.SetContext(ctx)

	values := make(url.Values)
	for key, vals := range scopes {
		for _, val := range vals {
			values.Add("scope", fmt.Sprintf("%s-%s", key, val))
		}
	}
	r.
		SetQueryParamsFromValues(values).
		SetQueryParam("grant_type", "principal").
		SetQueryParam("client_id", clientID).
		SetQueryParam("principal_token", token.Value()).
		SetQueryParam("principal_type", token.Type())

	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	resp, err := r.Post(e.authURL + issueToken)
	if err != nil {
		return IssueResponse{}, err
	} else if resp.IsError() {
		return IssueResponse{}, apierr.Unauthorized(fmt.Errorf("token fetch fail, response: %s", resp.Body()))
	}

	var unmarshalled IssueResponse
	if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
		return IssueResponse{}, err
	}

	return unmarshalled, nil
}

type IssueResponse struct {
	AccessToken           string   `json:"access_token"`
	RefreshToken          string   `json:"refresh_token"`
	ExpiresIn             int      `json:"expires_in"`
	ExpiresAt             int64    `json:"expires_at"`
	RefreshTokenExpiresAt int64    `json:"refresh_token_expires_at"`
	RefreshTokenExpiresIn int      `json:"refresh_token_expires_in"`
	TokenType             string   `json:"token_type"`
	Scope                 []string `json:"scope"`
}
