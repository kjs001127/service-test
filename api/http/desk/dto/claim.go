package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type DeskRoleView struct {
	Type   model.RoleType
	Claims []*model.Claim
}

type DeskRoleViews []DeskRoleView
