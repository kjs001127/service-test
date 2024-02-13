package dto

import (
	"encoding/json"
	"time"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type ParamsAndContext struct {
	Params  cmd.ParamInput     `json:"params"`
	Context app.ChannelContext `json:"context"`
}

type AppsAndCommands struct {
	Apps     []*app.App    `json:"apps"`
	Commands []*CommandDTO `json:"commands"`
}

type ContextAndAutoCompleteArgs struct {
	Context app.ChannelContext `json:"context"`
	Params  json.RawMessage    `json:"params"`
}

type CommandDTO struct {
	AppID            string               `json:"appId"`
	Name             string               `json:"name"`
	Scope            cmd.Scope            `json:"scope"`
	Description      *string              `json:"description"`
	ParamDefinitions cmd.ParamDefinitions `json:"paramDefinitions"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommandDTO(origin *cmd.Command) *CommandDTO {
	return &CommandDTO{
		AppID:            origin.AppID,
		Name:             origin.Name,
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
