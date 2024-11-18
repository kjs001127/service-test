package desk

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	roleFetch = "/api/admin/channels/{channelID}/roles/{roleID}"
)

type Permission struct {
	Action string `json:"action"`
}

type ManagerRole struct {
	Permissions []Permission `json:"permissions"`
	Name        string       `json:"name"`
	RoleType    string       `json:"type"`
	ChannelID   string       `json:"channelId"`
	ID          string       `json:"id"`
}

type ManagerRoleResponse struct {
	Role ManagerRole `json:"role"`
}

type ManagerRoleFetcher interface {
	FetchRole(ctx context.Context, channelID, roleID string) (ManagerRole, error)
}

type ManagerRoleFetcherImpl struct {
	cli     *resty.Client
	authURL string
}

func NewManagerRoleFetcher(cli *resty.Client, roleURL string) *ManagerRoleFetcherImpl {
	return &ManagerRoleFetcherImpl{cli: cli, authURL: roleURL}
}

func (c *ManagerRoleFetcherImpl) FetchRole(ctx context.Context, channelID, roleID string) (ManagerRole, error) {
	req := c.cli.R()
	req.SetContext(ctx)
	req.SetPathParam("channelID", channelID)
	req.SetPathParam("roleID", roleID)

	resp, err := req.Get(c.authURL + roleFetch)
	if err != nil {
		return ManagerRole{}, err
	}
	if !resp.IsSuccess() {
		return ManagerRole{}, errors.New("failed to fetch role")
	}

	body := resp.Body()
	var roleResp ManagerRoleResponse
	if err := json.Unmarshal(body, &roleResp); err != nil {
		return ManagerRole{}, err
	}

	return roleResp.Role, nil
}

const (
	ownerType  = "owner"
	permission = "generalSettings"
)

func (role ManagerRole) IsOwner() bool {
	return role.RoleType == ownerType
}

func (role ManagerRole) HasGeneralSettings() bool {
	if len(role.Permissions) <= 0 {
		return false
	}

	for _, perm := range role.Permissions {
		if perm.Action == permission {
			return true
		}
	}

	return false
}
