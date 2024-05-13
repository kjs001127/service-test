package dto

import (
	"github.com/channel-io/ch-app-store/internal/approle/model"
)

type RoleViews map[model.RoleType]RoleView

type RoleView struct {
	AvailableClaims []*model.Claim `json:"availableClaims"`
	Claims          []*model.Claim `json:"claims"`
}
