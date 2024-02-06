package general

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/auth/principal"
)

type PrincipalToRBACExchanger struct {
	cli     *resty.Client
	parser  *Parser
	authURL string
}

func NewPrincipalToRBACExchanger(cli *resty.Client, parser *Parser, authURL string) *PrincipalToRBACExchanger {
	return &PrincipalToRBACExchanger{cli: cli, parser: parser, authURL: authURL}
}

func (e *PrincipalToRBACExchanger) Exchange(ctx context.Context, token principal.Token, scopes Scopes, clientID string) (TokenResponse, error) {
	r := e.cli.R()
	r.SetContext(ctx)

	var values url.Values
	for key, vals := range scopes {
		for _, val := range vals {
			values.Add("scope", fmt.Sprintf("%s-%s", key, val))
		}
	}
	r.
		SetQueryParamsFromValues(values).
		SetQueryParam("principal_type", token.Type().Header()).
		SetQueryParam("grant_type", "principal").
		SetQueryParam("principal_token", token.Value()).
		SetQueryParam("client_id", clientID)

	resp, err := r.Post(e.authURL)
	if err != nil {
		return TokenResponse{}, err
	}

	var unmarshalled TokenResponse
	if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
		return TokenResponse{}, err
	}

	return unmarshalled, nil
}

type TokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    string   `json:"expires_in"`
	TokenType    string   `json:"token_type"`
	Scope        []string `json:"scope"`
}
