package general

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/internal/auth/principal"
)

const (
	issueToken = "/general/auth/" + version + "/v1/token"
)

type RBACExchanger struct {
	cli     *resty.Client
	parser  *ParserImpl
	authURL string
}

func NewRBACExchanger(cli *resty.Client, parser *ParserImpl, authURL string) *RBACExchanger {
	return &RBACExchanger{cli: cli, parser: parser, authURL: authURL}
}

func (e *RBACExchanger) Exchange(
	ctx context.Context,
	token principal.Token,
	scopes Scopes,
	clientID string,
) (IssueResponse, error) {
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
		SetQueryParam("principal_type", token.Type()).
		SetQueryParam("grant_type", "principal").
		SetQueryParam("principal_token", token.Value()).
		SetQueryParam("client_id", clientID)

	resp, err := r.Post(e.authURL + issueToken)
	if err != nil {
		return IssueResponse{}, err
	}

	var unmarshalled IssueResponse
	if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
		return IssueResponse{}, err
	}

	return unmarshalled, nil
}

type IssueResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    string   `json:"expires_in"`
	TokenType    string   `json:"token_type"`
	Scope        []string `json:"scope"`
}
