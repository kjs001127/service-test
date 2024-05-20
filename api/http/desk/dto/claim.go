package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type DeskRoleView struct {
	Type         model.RoleType `json:"type"`
	AppClaims    model.Claims   `json:"appClaims,omitempty"`
	NativeClaims model.Claims   `json:"nativeClaims,omitempty"`
}

type DeskRoleViews []DeskRoleView
