package dto

import (
	"github.com/channel-io/ch-app-store/internal/role/model"
)

type RoleViews map[model.RoleType]RoleView

type RoleView struct {
	ID                    string         `json:"id"`
	Type                  model.RoleType `json:"type"`
	AvailableNativeClaims []*model.Claim `json:"availableNativeClaims,omitempty"`
	AppClaims             []*model.Claim `json:"appClaims"`
	NativeClaims          []*model.Claim `json:"nativeClaims"`
}
