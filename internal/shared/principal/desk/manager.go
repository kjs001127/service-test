package desk

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/channel-io/ch-app-store/internal/shared/principal"

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

type ManagerResponse struct {
	Manager Manager `json:"manager"`
}

type Manager struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelId"`
	AccountID string `json:"accountId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	RoleID    string `json:"roleId"`

	RoleFetcher ManagerRoleFetcher
	role        *ManagerRole
}

// Role lazy fetch manager role
func (m *Manager) Role(ctx context.Context) (*ManagerRole, error) {
	if m.role == nil {
		ret, err := m.RoleFetcher.FetchRole(ctx, m.ChannelID, m.RoleID)
		if err != nil {
			return nil, err
		}

		m.role = &ret
	}

	return m.role, nil
}

type ManagerPrincipal struct {
	Manager
	Token Token
}

type ManagerRequester struct {
	principal.Requester
	ManagerPrincipal
}

type ManagerFetcher interface {
	FetchManager(ctx context.Context, channelID string, token string) (ManagerPrincipal, error)
}

type ManagerFetcherImpl struct {
	cli         *resty.Client
	roleFetcher ManagerRoleFetcher
	authUrl     string
}

func NewManagerFetcherImpl(cli *resty.Client, managerUrl string, roleFetcher ManagerRoleFetcher) *ManagerFetcherImpl {
	return &ManagerFetcherImpl{cli: cli, authUrl: managerUrl, roleFetcher: roleFetcher}
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

	body := resp.Body()
	var managerResp ManagerResponse
	if err := json.Unmarshal(body, &managerResp); err != nil {
		return ManagerPrincipal{}, err
	}

	manager := managerResp.Manager
	manager.RoleFetcher = c.roleFetcher

	return ManagerPrincipal{manager, Token(token)}, nil
}

func isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
