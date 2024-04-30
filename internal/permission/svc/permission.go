package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
)

const (
	ownerType  = "owner"
	permission = "generalSettings"
)

type PermissionUtil struct {
	roleFetcher account.ManagerRoleFetcher
}

func NewPermissionUtil(roleFetcher account.ManagerRoleFetcher) PermissionUtil {
	return PermissionUtil{roleFetcher: roleFetcher}
}

func (a PermissionUtil) isOwner(ctx context.Context, manager account.Manager) bool {
	role, err := a.roleFetcher.FetchRole(ctx, manager.RoleID)
	if err != nil {
		return false
	}
	if role.RoleType == ownerType {
		return true
	}

	return false
}

func (a PermissionUtil) hasPermission(ctx context.Context, manager account.Manager) bool {
	role, err := a.roleFetcher.FetchRole(ctx, manager.RoleID)
	if err != nil {
		return false
	}

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
