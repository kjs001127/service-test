package dto

import (
	"github.com/channel-io/ch-app-store/api/http/desk/dto"
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

type SimpleInstalledAppView struct {
	*dto.SimpleAppView
	ShouldUpdateAgreement bool `json:"shouldUpdateAgreement"`
}

type DetailedInstalledAppPageView struct {
	App           *dto.DetailedAppView    `json:"app"`
	Commands      []*InstalledCommandView `json:"commands"`
	UnAgreedRoles DeskRoleViews           `json:"unAgreedRoles"`
}

type InstalledCommandView struct {
	*dto.SimpleCommandView
	Enabled bool `json:"enabled"`
}

func NewInstalledCommandView(cmd *svc.CommandWithActivation) *InstalledCommandView {
	return &InstalledCommandView{
		SimpleCommandView: dto.NewCommandView(cmd.Command),
		Enabled:           cmd.Enabled,
	}
}

func NewInstalledCommandViews(origins []*svc.CommandWithActivation) []*InstalledCommandView {
	ret := make([]*InstalledCommandView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewInstalledCommandView(origin))
	}
	return ret
}
