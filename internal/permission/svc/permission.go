package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	ownerType  = "owner"
	permission = "generalSettings"
)

type PermissionUtil struct {
	roleFetcher account.ManagerRoleFetcher
	logger      log.ContextAwareLogger
}

func NewPermissionUtil(roleFetcher account.ManagerRoleFetcher) PermissionUtil {
	return PermissionUtil{roleFetcher: roleFetcher}
}

func (a PermissionUtil) isOwner(ctx context.Context, manager account.Manager) (string, bool) {
	role, err := a.roleFetcher.FetchRole(ctx, manager.ChannelID, manager.RoleID)
	if err != nil {
		a.logger.Error(ctx, "error while fetching role", err)
		return role.RoleType, false
	}
	if role.RoleType == ownerType {
		return role.RoleType, true
	}

	return role.RoleType, false
}

func (a PermissionUtil) hasGeneralSettings(ctx context.Context, manager account.Manager) (string, bool) {
	role, err := a.roleFetcher.FetchRole(ctx, manager.ChannelID, manager.RoleID)
	if err != nil {
		return role.RoleType, false
	}

	if len(role.Permissions) <= 0 {
		return role.RoleType, false
	}

	for _, perm := range role.Permissions {
		if perm.Action == permission {
			return role.RoleType, true
		}
	}

	return role.RoleType, false
}
