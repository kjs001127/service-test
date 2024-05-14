package svc

import (
	"context"

	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	ownerType = "owner"
)

type PermissionUtil struct {
	roleFetcher account.ManagerRoleFetcher
	logger      log.ContextAwareLogger
}

func NewPermissionUtil(roleFetcher account.ManagerRoleFetcher) PermissionUtil {
	return PermissionUtil{roleFetcher: roleFetcher}
}

func (a PermissionUtil) isOwner(ctx context.Context, manager account.Manager) bool {
	role, err := a.roleFetcher.FetchRole(ctx, manager.ChannelID, manager.RoleID)
	if err != nil {
		a.logger.Error(ctx, "error while fetching role", err)
		return false
	}
	if role.RoleType == ownerType {
		return true
	}

	return false
}
