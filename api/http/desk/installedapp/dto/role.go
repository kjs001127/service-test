package dto

import (
	"github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/internal/role/svc"
)

type DeskRoleView struct {
	ID           string         `json:"string"`
	Type         model.RoleType `json:"type"`
	AppClaims    model.Claims   `json:"appClaims,omitempty"`
	NativeClaims model.Claims   `json:"nativeClaims,omitempty"`
}

type DeskRoleViews []*DeskRoleView

func FromRoles(roles []*svc.ClaimsResponse) DeskRoleViews {
	ret := make(DeskRoleViews, 0, len(roles))
	for _, role := range roles {

		ret = append(ret, &DeskRoleView{
			Type:         role.Type,
			NativeClaims: role.NativeClaims,
			AppClaims:    role.AppClaims,
		})
	}

	return ret
}
