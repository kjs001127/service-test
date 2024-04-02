package svc

import (
	"context"

	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type RoleClient interface {
	GetRole(ctx context.Context, roleID string) (*service.GetRoleResult, error)
	CreateRole(ctx context.Context, request *service.CreateRoleRequest) (*service.CreateRoleResult, error)
	UpdateRole(ctx context.Context, roleID string, request *service.ReplaceRoleClaimsRequest) (*service.ReplaceRoleClaimsResult, error)
	DeleteRole(ctx context.Context, roleID string) (*service.DeleteRoleResult, error)
}
