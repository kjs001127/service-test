package dto

import (
	"time"

	"github.com/channel-io/ch-app-store/internal/app/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
)

type InstalledApp struct {
	App             *model.App             `json:"app"`
	AppInstallation *model.AppInstallation `json:"appChannel"`
}

type InstalledAppsWithCommands struct {
	Apps             []*model.App             `json:"apps"`
	AppInstallations []*model.AppInstallation `json:"appChannels"`
	Commands         []*CommandDTO            `json:"commands"`
}

type InstalledAppWithCommands struct {
	App             *model.App             `json:"app"`
	AppInstallation *model.AppInstallation `json:"appChannel"`
	Commands        []*CommandDTO          `json:"commands"`
}

type AppsAndCommands struct {
	Apps     []*model.App  `json:"apps"`
	Commands []*CommandDTO `json:"commands"`
}

type CommandDTO struct {
	AppID string    `json:"appId,o"`
	Name  string    `json:"name"`
	Scope cmd.Scope `json:"scope"`

	Description     *string        `json:"description,omitempty"`
	NameDescI18nMap map[string]any `json:"nameDescI18nMap,omitempty"`

	ParamDefinitions cmd.ParamDefinitions `json:"paramDefinitions,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommandDTO(origin *cmd.Command) *CommandDTO {
	return &CommandDTO{
		AppID:            origin.AppID,
		Name:             origin.Name,
		NameDescI18nMap:  origin.NameDescI18NMap,
		Scope:            origin.Scope,
		Description:      origin.Description,
		ParamDefinitions: origin.ParamDefinitions,
		CreatedAt:        origin.CreatedAt,
		UpdatedAt:        origin.UpdatedAt,
	}
}

func NewCommandDTOs(origins []*cmd.Command) []*CommandDTO {
	ret := make([]*CommandDTO, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewCommandDTO(origin))
	}
	return ret
}
