package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type RoleViews map[model.RoleType]AdminRoleView

type AdminRoleView struct {
	AppClaims    []*model.Claim `json:"appClaims"`
	NativeClaims []*model.Claim `json:"nativeClaims"`
}
