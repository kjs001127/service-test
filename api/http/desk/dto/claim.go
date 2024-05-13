package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type RoleView struct {
	Type   model.RoleType
	Claims []*model.Claim
}

type RoleViews []RoleView
