package dto

import (
	"time"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	cmd "github.com/channel-io/ch-app-store/internal/command/model"
)

type SimpleAppView struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	AvatarURL   *string                `json:"avatarUrl,omitempty"`
	Description *string                `json:"description,omitempty"`
	IsBuiltIn   bool                   `json:"isBuiltIn"`
	I18nMap     map[string]AppViewI18n `json:"i18nMap"`
	IsPrivate   bool                   `json:"isPrivate"`
}

type AppViewI18n struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func convertAppViewI18n(app *appmodel.App) map[string]AppViewI18n {
	ret := make(map[string]AppViewI18n)
	for lang, i18n := range app.I18nMap {
		ret[lang] = AppViewI18n{
			Title:       i18n.Title,
			Description: i18n.Description,
		}
	}
	return ret
}

func NewAppView(origin *appmodel.App) *SimpleAppView {
	return &SimpleAppView{
		ID:          origin.ID,
		Title:       origin.Title,
		AvatarURL:   origin.AvatarURL,
		Description: origin.Description,
		IsBuiltIn:   origin.IsBuiltIn,
		IsPrivate:   origin.IsPrivate,
		I18nMap:     convertAppViewI18n(origin),
	}
}

func NewAppViews(origins []*appmodel.App) []*SimpleAppView {
	ret := make([]*SimpleAppView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewAppView(origin))
	}
	return ret
}

type SimpleCommandView struct {
	AppID string    `json:"appId,o"`
	Name  string    `json:"name"`
	Scope cmd.Scope `json:"scope"`

	Description     *string                        `json:"description,omitempty"`
	NameDescI18nMap map[string]cmd.NameDescI18nMap `json:"nameDescI18nMap,omitempty"`

	ParamDefinitions cmd.ParamDefinitions `json:"paramDefinitions,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommandView(origin *cmd.Command) *SimpleCommandView {
	return &SimpleCommandView{
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

func NewCommandViews(origins []*cmd.Command) []*SimpleCommandView {
	ret := make([]*SimpleCommandView, 0, len(origins))
	for _, origin := range origins {
		ret = append(ret, NewCommandView(origin))
	}
	return ret
}
