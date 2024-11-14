package dto

import (
	"github.com/channel-io/ch-app-store/internal/role/model"
)

type RoleViews map[model.RoleType]AdminRoleView

type AdminRoleView struct {
	AppClaims    []*model.Claim `json:"appClaims"`
	NativeClaims []*model.Claim `json:"nativeClaims"`
}
