package dto

import (
	"time"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type ArgumentsAndContext struct {
	Params  cmd.Arguments      `json:"params"`
	Context app.ChannelContext `json:"context"`
}

type AppsAndCommands struct {
	Apps     []*app.AppData `json:"apps"`
	Commands []*CommandDTO  `json:"commands"`
}

type ContextAndAutoCompleteArgs struct {
	Context app.ChannelContext   `json:"context"`
	Params  cmd.AutoCompleteArgs `json:"params"`
}

type CommandDTO struct {
	AppID            string
	Name             string
	Scope            cmd.Scope
	Description      string
	ParamDefinitions cmd.ParamDefinitions

	ActionFunctionName       string
	AutoCompleteFunctionName string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCommandDTO(origin *cmd.Command) *CommandDTO {
	return &CommandDTO{
		AppID:                    origin.AppID,
		Name:                     origin.Name,
		Scope:                    origin.Scope,
		Description:              origin.Description.String,
		ParamDefinitions:         origin.ParamDefinitions,
		ActionFunctionName:       origin.ActionFunctionName,
		AutoCompleteFunctionName: origin.AutoCompleteFunctionName.String,
		CreatedAt:                origin.CreatedAt,
		UpdatedAt:                origin.UpdatedAt,
	}
}

func NewCommandDTOs(origins []*cmd.Command) []*CommandDTO {
	ret := make([]*CommandDTO, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewCommandDTO(origin))
	}
	return ret
}
