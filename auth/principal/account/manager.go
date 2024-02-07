package account

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-resty/resty/v2"
)

const XAccountHeader = "x-account"

type Manager struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelId"`
	AccountID string `json:"accountId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type ManagerWithToken struct {
	Manager
	Token string
}

type ManagerFetcher interface {
	FetchManager(ctx context.Context, channelID string, token string) (Manager, error)
}

type ManagerFetcherImpl struct {
	cli        *resty.Client
	managerUrl string
}

func NewManagerFetcherImpl(cli *resty.Client, managerUrl string) *ManagerFetcherImpl {
	return &ManagerFetcherImpl{cli: cli, managerUrl: managerUrl}
}

func (c *ManagerFetcherImpl) FetchManager(ctx context.Context, channelID string, token string) (ManagerWithToken, error) {
	req := c.cli.R()
	req.SetContext(ctx)
	req.Header.Set(XAccountHeader, token)
	req.QueryParam.Set("channelId", channelID)
	resp, err := req.Get(c.managerUrl)
	if err != nil {
		return ManagerWithToken{}, err
	}

	if !isSuccess(resp.StatusCode()) {
		return ManagerWithToken{}, errors.New("auth failed")
	}

	var manager Manager
	if err := json.Unmarshal(resp.Body(), &manager); err != nil {
		return ManagerWithToken{}, err
	}

	return ManagerWithToken{manager, token}, nil
}

func isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

type MockManagerFetcher struct {
}

func NewMockManagerFetcher() *MockManagerFetcher {
	return &MockManagerFetcher{}
}

func (m MockManagerFetcher) FetchManager(ctx context.Context, channelID string, token string) (Manager, error) {
	return Manager{
		ID:        "1",
		ChannelID: "1",
		AccountID: "1",
		Name:      "1",
		Email:     "1",
	}, nil
}
