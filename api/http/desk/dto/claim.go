package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type DeskRoleView struct {
	Type         model.RoleType `json:"type"`
	AppClaims    model.Claims   `json:"appClaims"`
	NativeClaims model.Claims   `json:"nativeClaims"`
}

type DeskRoleViews []DeskRoleView
