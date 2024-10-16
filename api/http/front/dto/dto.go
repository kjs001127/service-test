package dto

import (
	"time"

	"github.com/channel-io/ch-app-store/internal/app/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
)

type AppsAndCommands struct {
	Apps     []*AppDTO     `json:"apps"`
	Commands []*CommandDTO `json:"commands"`
}

type AppDTO struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	AvatarURL   *string               `json:"avatarUrl,omitempty"`
	Description *string               `json:"description,omitempty"`
	I18nMap     map[string]AppDTOI18n `json:"i18NMap,omitempty"`
	IsBuiltIn   bool                  `json:"isBuiltIn"`
}

type AppDTOI18n struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func NewAppDTO(origin *model.App) *AppDTO {
	return &AppDTO{
		ID:          origin.ID,
		Title:       origin.Title,
		AvatarURL:   origin.AvatarURL,
		Description: origin.Description,
		IsBuiltIn:   origin.IsBuiltIn,
		I18nMap:     convertAppDTOI18n(origin),
	}
}

func convertAppDTOI18n(app *model.App) map[string]AppDTOI18n {
	ret := make(map[string]AppDTOI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppDTOI18n{
			Title:       i18n.Title,
			Description: i18n.Description,
		}
	}
	return ret
}

func NewAppDTOs(origins []*model.App) []*AppDTO {
	ret := make([]*AppDTO, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewAppDTO(origin))
	}
	return ret
}

type CommandDTO struct {
	AppID string    `json:"appId,o"`
	Name  string    `json:"name"`
	Scope cmd.Scope `json:"scope"`

	Description     *string                        `json:"description,omitempty"`
	NameDescI18nMap map[string]cmd.NameDescI18nMap `json:"nameDescI18nMap,omitempty"`

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
