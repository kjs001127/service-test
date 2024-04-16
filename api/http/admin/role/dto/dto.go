package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type AppRoles struct {
	AppID   string   `json:"appId"`
	RoleIDs []string `json:"roleIds"`
}

func AppRoleFrom(appId string, models []*model.AppRole) *AppRoles {
	roleIds := make([]string, 0, len(models))
	for _, appRole := range models {
		roleIds = append(roleIds, appRole.RoleID)
	}
	return &AppRoles{
		AppID:   appId,
		RoleIDs: roleIds,
	}
}
