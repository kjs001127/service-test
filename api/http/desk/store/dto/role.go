package dto

import (
	"github.com/channel-io/ch-app-store/internal/role/model"
)

type DeskRoleView struct {
	Type         model.RoleType `json:"type"`
	ID           string         `json:"id"`
	AppClaims    model.Claims   `json:"appClaims,omitempty"`
	NativeClaims model.Claims   `json:"nativeClaims,omitempty"`
}

type DeskRoleViews []DeskRoleView
