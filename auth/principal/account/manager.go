package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

const (
	managerFetch   = "/api/admin/manager"
	XAccountHeader = "x-account"
)

type Token string

func (t Token) Type() string {
	return XAccountHeader
}

func (t Token) Value() string {
	return string(t)
}

type Manager struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelId"`
	AccountID string `json:"accountId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type ManagerPrincipal struct {
	Manager
	Token Token
}

type ManagerFetcher interface {
	FetchManager(ctx context.Context, channelID string, token string) (ManagerPrincipal, error)
}

type ManagerFetcherImpl struct {
	cli     *resty.Client
	authUrl string
}

func NewManagerFetcherImpl(cli *resty.Client, managerUrl string) *ManagerFetcherImpl {
	return &ManagerFetcherImpl{cli: cli, authUrl: managerUrl}
}

func (c *ManagerFetcherImpl) FetchManager(ctx context.Context, channelID string, token string) (ManagerPrincipal, error) {
	req := c.cli.R()
	req.SetContext(ctx)
	req.Header.Set(XAccountHeader, token)
	req.QueryParam.Set("channelId", channelID)
	resp, err := req.Get(c.authUrl + managerFetch)
	if err != nil {
		return ManagerPrincipal{}, err
	}

	if !isSuccess(resp.StatusCode()) {
		return ManagerPrincipal{}, errors.New("auth failed")
	}

	var manager Manager
	if err := json.Unmarshal(resp.Body(), &manager); err != nil {
		return ManagerPrincipal{}, err
	}

	return ManagerPrincipal{manager, Token(token)}, nil
}

func isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
