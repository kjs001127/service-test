package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type RoleViews map[model.RoleType]RoleView

type RoleView struct {
	AvailableNativeClaims []*model.Claim `json:"availableNativeClaims"`
	AppClaims             []*model.Claim `json:"appClaims"`
	NativeClaims          []*model.Claim `json:"nativeClaims"`
}
