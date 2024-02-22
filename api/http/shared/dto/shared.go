package dto

import (
	"time"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type CommandInput struct {
	Params cmd.CommandBody `json:"params"`
}

type AppsAndCommands struct {
	Apps     []*app.App    `json:"apps"`
	Commands []*CommandDTO `json:"commands"`
}

type ContextAndAutoCompleteArgs struct {
	Params cmd.AutoCompleteBody `json:"params"`
}

type CommandDTO struct {
	AppID       string            `json:"appId"`
	Name        string            `json:"name"`
	NameI18nMap map[string]string `json:"nameI18NMap"`

	Scope              cmd.Scope            `json:"scope"`
	Description        *string              `json:"description"`
	DescriptionI18nMap map[string]string    `json:"descriptionI18NMap"`
	ParamDefinitions   cmd.ParamDefinitions `json:"paramDefinitions"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommandDTO(origin *cmd.Command) *CommandDTO {
	return &CommandDTO{
		AppID:              origin.AppID,
		Name:               origin.Name,
		NameI18nMap:        origin.NameI18nMap,
		Scope:              origin.Scope,
		Description:        origin.Description,
		ParamDefinitions:   origin.ParamDefinitions,
		DescriptionI18nMap: origin.DescriptionI18nMap,
		CreatedAt:          origin.CreatedAt,
		UpdatedAt:          origin.UpdatedAt,
	}
}

func NewCommandDTOs(origins []*cmd.Command) []*CommandDTO {
	ret := make([]*CommandDTO, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewCommandDTO(origin))
	}
	return ret
}
